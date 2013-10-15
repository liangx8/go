package web
import (
	"net/http"
	"appengine"
	"appengine/datastore"
	"appengine/blobstore"
	"fmt"
	"io"
	"os"
	"errors"
	"regexp"
	"mime"

	"store"
	"zpack"
)
func byFileName(w http.ResponseWriter,r *http.Request) {
	uri := r.URL.RequestURI()[1:]
	if uri == "" {uri = "index.html"}
	c := appengine.NewContext(r)
	util := store.NewVpathUtil(c)
	var vp store.Vpath
	key,err := util.FindOne(uri,&vp)
	if(err !=  nil){
		c.Errorf("%v",err)
		return
	}
	if key == nil {
		c.Infof("file(%s) is not found!",uri)
		http.NotFound(w,r)
	}
	blobstore.Send(w,vp.Key)
	return
}
func upload(w http.ResponseWriter, r *http.Request,name string) error{
	c := appengine.NewContext(r)
	f,fh,err:=r.FormFile(name)
	if err != nil {
		return err
	}
	mimeType := fh.Header.Get("Content-type")
	util := store.NewResourceUtil(c)
	util.SaveUnique(f,mimeType)
	if err != nil {
		return err
	}
	return nil
}
func vupload(w http.ResponseWriter, r *http.Request,name,pa string) error {
	c := appengine.NewContext(r)
	f,fh,err := r.FormFile(name)
	if err != nil {
		return nil
	}
	mimeType := fh.Header.Get("Content-type")
	util := store.NewResourceUtil(c)
	res,err := util.SaveUnique(f,mimeType)
	if err != nil {
		return err
	}
	filepath := r.FormValue(pa)
	if filepath == "" {
		return errors.New("file name must be provided")
	}
	util1 := store.NewVpathUtil(c)
	util1.SaveOrUpdate(res.Key,filepath)
	return nil
}
func list(w http.ResponseWriter, r *http.Request, page int) error{
	c := appengine.NewContext(r)
	util := store.NewResourceUtil(c)
	return util.All(&rCallback{w})
}

type rCallback struct{w io.Writer}
func (cb *rCallback)Head(isEmpty bool){
	if isEmpty {return}
	fmt.Fprint(cb.w,`<TABLE bgcolor="black"><TR bgcolor="gray"><TD>Seq</TD><TD>MD5</TD><TD>OPEN</TD></TR>`)
}
func (cb *rCallback)Each(k *datastore.Key, e interface{},num int) error{
	fmt.Fprintf(cb.w,`<TR BGCOLOR="white"><TD>%d</TD><TD>%s</TD><TD><a href="bykey?bk=%s">OPEN</a></TD></TR>`,num,e.(*store.Resource).Md5str,e.(*store.Resource).Key)
	return nil
}
func (cb *rCallback)Tail(count int){
	if count == 0 {return}
	fmt.Fprint(cb.w,`</TABLE>`)
}
func vlist(w http.ResponseWriter, r *http.Request, page int) error{
	c := appengine.NewContext(r)
	util := store.NewVpathUtil(c)
	return util.All(&fCallback{w})
}
type fCallback struct{
	w io.Writer
}
func (cb *fCallback)Head(isEmpty bool){
	if isEmpty {
		fmt.Fprint(cb.w, "zero record")
		return
	}
	fmt.Fprint(cb.w,`<TABLE bgcolor="black"><TR bgcolor="gray"><TD>Seq</TD><TD>path</TD><TD>OPEN</TD></TR>`)
}
func (cb *fCallback)Each(k *datastore.Key, e interface{}, num int) error {
	fmt.Fprintf(cb.w,`<TR BGCOLOR="white"><TD>%d</TD><TD>%s</TD><TD><a href="/%s">OPEN</a></TD></TR>`,num,e.(*store.Vpath).Name,e.(*store.Vpath).Name)
	return nil
}
func (cb *fCallback)Tail(count int){
	fmt.Fprint(cb.w,"</TABLE>")
}
func tarupload(w http.ResponseWriter, r *http.Request, name string) error{
	f,_,err := r.FormFile(name)
	c := appengine.NewContext(r)
	rutil := store.NewResourceUtil(c)
	vutil := store.NewVpathUtil(c)

	if err != nil {
		c.Errorf("formfile:%v",err)
		return err
	}
	re,err := regexp.Compile(`\..*$`)
	if err != nil {
		c.Errorf("regexp:%v",err)
		return err
	}
	err = zpack.TarForEach(f,func(zr io.Reader,fi os.FileInfo){
		if !fi.IsDir(){
			fn := fi.Name()
			n := re.FindStringIndex(fn)
			ext := fn[n[0]:]
			mediaType := mime.TypeByExtension(ext)
			if mediaType == "" {
				mediaType = "application/octet-stream"
			}
			res,err :=rutil.SaveUnique(zr,mediaType)
			if err != nil {
				c.Errorf("%v",err)
				return
			}
			vp,err := vutil.SaveOrUpdate(res.Key,fn)
			if err != nil {
				c.Errorf("%v",err)
			}
			if vp == nil { return }
		}
	})
	if err != nil {
		c.Errorf("%v",err)
		return err
	}
	return nil
}
