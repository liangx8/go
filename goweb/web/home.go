package web

import (
	"net/http"
	"appengine"
	"appengine/blobstore"

	"store"
	"session"
	"click"

)

func root(w http.ResponseWriter,r *http.Request,s *session.Session){
	click.Counter(r,s)
	byFileName(w,r)
}
func byFileName(w http.ResponseWriter,r *http.Request) {
	c := appengine.NewContext(r)
	uri := r.URL.RequestURI()
	if uri == "/" {uri = "/index.html"}
	util := store.NewVpathUtil(c)
	var vp store.Vpath
	key,err := util.FindOne(uri[1:],&vp)
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
