package web
import (
	//"net/http"
	"appengine"
//	"appengine/blobstore"
	"fmt"
	"io"
	"os"
//	"errors"
	"regexp"
	"mime"
	"mime/multipart"
	"archive/zip"

	"store"
	"zpack"


)


type _callback struct {
	c appengine.Context
	rutil *store.ResourceUtil
	vutil *store.VpathUtil
}
// implements zpack.ZCallback interface
func (cb *_callback)Run(zr io.Reader,fi os.FileInfo){
	if !fi.IsDir(){
		fn := fi.Name()
		re,err := regexp.Compile(`\..*$`)
		if err != nil {
			cb.c.Errorf("regexp:%v",err)
		}
		n := re.FindStringIndex(fn)
		ext := fn[n[0]:]
		mediaType := mime.TypeByExtension(ext)
		if mediaType == "" {
			mediaType = "application/octet-stream"
		}
		res,err :=cb.rutil.SaveUnique(zr,mediaType)
		if err != nil {
			cb.c.Errorf("%v",err)
			return
		}
		vp,err := cb.vutil.SaveOrUpdate(res.Key,fn)
		if err != nil {
			cb.c.Errorf("%v",err)
		}
		if vp == nil { return }
	}
}
func tarupload(c appengine.Context,f multipart.File) error {

	err := zpack.TarForEach(f,&_callback{c,store.NewResourceUtil(c),store.NewVpathUtil(c)})
	if err != nil {
		c.Errorf("%v",err)
		return err
	}
	return nil
}
func zipupload(c appengine.Context,f multipart.File) error {
	// reach end of file's size
	size,err := f.Seek(0,2)
	if err != nil {
		return err
	}
	mr,err := zip.NewReader(f,size)
	if err != nil {
		return err
	}
	err = zpack.ZipForEach(mr,&_callback{c,store.NewResourceUtil(c),store.NewVpathUtil(c)})
	if err != nil {
		return err
	}
	return nil
}
func upload(c appengine.Context,f multipart.File, fh *multipart.FileHeader, fp string) error {
	util :=store.NewResourceUtil(c)
	res,err := util.SaveUnique(f,fh.Header.Get("Content-type"))
	if err != nil {
		return err
	}
	if fp == "" {
		return fmt.Errorf("file name must be provided")
	}
	u1 := store.NewVpathUtil(c)
	u1.SaveOrUpdate(res.Key,fp)
	return nil
}
