package web

import (
	"net/http"
	"fmt"
	"regexp"
	"strconv"

	"appengine"
	"appengine/blobstore"

	"click"
)
const (
	ADMIN_URI = "/__admin__/"
	FILE_FIELD="file"
)
func init(){
	http.HandleFunc("/",handler)
	http.HandleFunc(ADMIN_URI,adminHandler)
}
func handler(w http.ResponseWriter, r *http.Request){
	click.Counter(r)
	byFileName(w,r)
	//uri := r.URL.RequestURI()[1:]
	//fmt.Fprint(w, uri)
}
func adminHandler(w http.ResponseWriter, r *http.Request){
	uri := r.URL.RequestURI()[11:]
	relist,err := regexp.Compile(`^list\.*(\d*)$`)
	c := appengine.NewContext(r)
	if err != nil {
		c.Errorf("list:Impossible error %v",err)
		return
	}
	revlist,err := regexp.Compile(`^vlist\.*(\d*)$`)
	if err != nil {
		c.Errorf("vlist:Impossible error %v",err)
		return
	}
	reclick,err := regexp.Compile(`^click\.*(\d*)$`)
	if err != nil {
		c.Errorf("click:impossible error %v",err)
		return
	}
	switch{
	case relist.MatchString(uri):
		w.Header().Add("Content-type","text/html;charset=UTF-8")
		var page int
		pstr := relist.FindStringSubmatch(uri)
		page,err = strconv.Atoi(pstr[1])
		if err != nil {page = 0}
		fmt.Fprintf(w,HTML_HEAD,"List all resources")
		list(w,r,page)
		fmt.Fprintf(w,HTML_UPLOADFORM,"upload.do",FILE_FIELD)
		fmt.Fprintf(w,HTML_TAIL)
	case revlist.MatchString(uri):
		w.Header().Add("Content-type","text/html;charset=UTF-8")
		fmt.Fprintf(w,HTML_HEAD,"List all resources")
		pstr := revlist.FindStringSubmatch(uri)[1]
//		fmt.Fprint(w,pstr)
		var page int
		page,err = strconv.Atoi(pstr)
		if err != nil {page = 0}
		vlist(w,r,page)
		fmt.Fprintf(w,HTML_UPLOADFORM,"vupload.do",FILE_FIELD)
		fmt.Fprintf(w,HTML_TAIL)
	case reclick.MatchString(uri):
		w.Header().Add("Content-type","text/html;charset=UTF-8")
		var page int
		pstr := reclick.FindStringSubmatch(uri)
		page,err = strconv.Atoi(pstr[1])
		if err != nil {page = 0}
		fmt.Fprintf(w,HTML_HEAD,"statistices")
		click.ListCounter(w,r,page)
		fmt.Fprintf(w,HTML_TAIL)
	case uri == "tarupload":
		w.Header().Add("Content-type","text/html;charset=UTF-8")
		fmt.Fprintf(w,HTML_HEAD,"upload tar archive")
		fmt.Fprint(w,"<h2>Upload batch files in tgz archive</h2><hr />")
		fmt.Fprintf(w,HTML_UPLOADFORM,"tarupload.do",FILE_FIELD)
		fmt.Fprintf(w,HTML_TAIL)
	case uri == "tarupload.do":
		err := tarupload(w,r,FILE_FIELD)
		w.Header().Add("Content-type","text/html;charset=UTF-8")
		fmt.Fprintf(w,HTML_HEAD,"tar file upload result")
		if err == nil{
			fmt.Fprint(w,"Successful!");
		} else {
			fmt.Fprint(w,"Failure!");
		}
		fmt.Fprintf(w,HTML_TAIL)
	case uri == "zipupload":
		w.Header().Add("Content-type","text/html;charset=UTF-8")
		fmt.Fprintf(w,HTML_HEAD,"upload tar archive")
		fmt.Fprint(w,"<h2>Upload batch files in zip archive</h2><hr />")
		fmt.Fprintf(w,HTML_UPLOADFORM,"zipupload.do",FILE_FIELD)
		fmt.Fprintf(w,HTML_TAIL)
	case uri == "zipupload.do":
		err := zipupload(w,r,FILE_FIELD)
		w.Header().Add("Content-type","text/html;charset=UTF-8")
		fmt.Fprintf(w,HTML_HEAD,"zip file upload result")
		if err == nil{
			fmt.Fprint(w,"Successful!");
		} else {
			fmt.Fprint(w,"Failure!");
			fmt.Fprint(w,err)
		}
		fmt.Fprintf(w,HTML_TAIL)
	case uri == "upload.do":
		upload(w,r,FILE_FIELD)
	case uri == "vupload.do":
		vupload(w,r,FILE_FIELD,"filepath")
	case uri[:5] == "bykey":
		// show file by blobkey
		bk := r.FormValue("bk")
		blobstore.Send(w,appengine.BlobKey(bk))
	default:
		w.Header().Add("Content-type","text/html;charset=UTF-8")
		fmt.Fprintf(w,HTML_HEAD,"function is not implements")
		fmt.Fprintf(w,"(%s)Not implemented!还没实现的功能!",uri)
		fmt.Fprintf(w,HTML_TAIL)
	}
}
const (
	HTML_HEAD="<HTML><HEAD><TITLE>%s</TITLE></HEAD><BODY>"
	HTML_UPLOADFORM=`<form action="%s" method="POST" enctype="multipart/form-data"> filename:<input type="file" name="%s" value="" /> path:<input name="filepath" /><input type="submit" /></form>`
	HTML_TAIL="</BODY></HTML>"
)
