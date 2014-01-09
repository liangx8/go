package web

import (
	"net/http"
	"html/template"
	"fmt"
	"regexp"
	"strconv"

	"appengine"
	"appengine/blobstore"

	"store"
	"click"
	"session"
)
const (
	ADMIN_URI = "/__admin__/"
)
var orig_t *template.Template
func init(){
	handleFunc(ADMIN_URI,adminHandler)
	handleFunc(ADMIN_URI+"list",list)
	handleFunc(ADMIN_URI+"template",tmpladmin)
	handleFunc(ADMIN_URI+"bykey",bykey)
	handleFunc(ADMIN_URI+"bloblist",bloblist)
	handleFunc(ADMIN_URI+"listsession",listsession)
	handleFunc("/",root)
	orig_t = template.Must(template.New("admin.tmpl").Parse(ADMIN_HOME))
	template.Must(orig_t.New("tmpl_admin.tmpl").Parse(TMPL_ADMIN))
	template.Must(orig_t.New("list.tmpl").Parse(LIST))
	template.Must(orig_t.New("bloblist.tmpl").Parse(BLOBLIST))
	template.Must(orig_t.New("click.tmpl").Parse(CLICKLIST))
	template.Must(orig_t.New("session.tmpl").Parse(SESSIONLIST))
}
func bykey(w http.ResponseWriter, r *http.Request){
	bk := r.FormValue("bk")
	blobstore.Send(w,appengine.BlobKey(bk))
}
func bloblist(w http.ResponseWriter, r *http.Request){
	c:=appengine.NewContext(r)
	res := store.NewResourceUtil(c)
	ch := make(chan *store.Resource)
	go func() {
		res.All(func(rs *store.Resource){ ch <- rs})
		close(ch)
	}()
	m:=map[string]interface{}{"view":"bloblist.tmpl","data":ch}
	err := render(w,m)
	if err != nil {
		c.Errorf("%v",err)
		return
	}
}
func listsession(w http.ResponseWriter, r *http.Request){
	c:=appengine.NewContext(r)
	m:=map[string]interface{}{"view":"session.tmpl","data":session.List()}
	err := render(w,m)
	if err != nil {
		c.Errorf("%v",err)
        }

}
func list(w http.ResponseWriter, r *http.Request){
	c:=appengine.NewContext(r)
	vpu := store.NewVpathUtil(c)
	ch := make(chan *store.Vpath)
	go func() {
		vpu.All(func(vp *store.Vpath){ ch <- vp })
		close(ch)
	}()
	m:=map[string]interface{}{"view":"list.tmpl","data":ch}
	err := render(w,m)
	if err != nil {
		c.Errorf("%v",err)
		return
	}
//	session.Log(c)
}
func adminHandler(w http.ResponseWriter, r *http.Request){
	c:=appengine.NewContext(r)
	uri := r.URL.RequestURI()[11:]
	reclick,err := regexp.Compile(`^click\.*(\d*)$`)
	if err==nil && reclick.MatchString(uri){
		pstr := reclick.FindStringSubmatch(uri)
		page,err := strconv.Atoi(pstr[1])
		if err != nil {page = 0}

		m := click.ListCounter(w,c,page)
		err = render(w,m)
		if err != nil {
			c.Errorf("%v",err)
		}
		return
	}
	action := r.FormValue("action")
	switch action {
		case "tarupload":
		case "zipupload":
			// it should be a zip, so file header is not need
			f,_,err := r.FormFile("filename")
			if err != nil {
				c.Errorf("%v",err)
				return
			}
			c.Infof("%s",action)
			err = zipupload(c,f)
			if err !=nil {
				c.Errorf("%v",err)
				return
			}
		case "upload":
                        f,fh,err := r.FormFile("filename")
                        if err != nil {
                                c.Errorf("%v",err)
                                return
                        }
			fp := r.FormValue("filepath")
                        c.Infof("%s",action)
			err = upload(c,f,fh,fp)
                        if err !=nil {
                                c.Errorf("%v",err)
                                return
                        }
		default:
			err := render(w,map[string]interface{}{"view":"admin.tmpl"})
			if err != nil {
				c.Errorf("%v",err)
			}
			return
	}

}
func render(w http.ResponseWriter,model map[string]interface{})error{
	tn,ok := model["view"].(string)
	if !ok {
		return fmt.Errorf("not view name")
	}
	t:=orig_t.Lookup(tn)
	if t== nil {
		return fmt.Errorf("template %s not define",tn)
	}
	return t.Execute(w,model)
}

