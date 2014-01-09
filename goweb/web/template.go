package web

import (
	"net/http"
	"appengine"
)
func tmpladmin(w http.ResponseWriter, r *http.Request){
        c:=appengine.NewContext(r)
	if r.Method == "POST" {
		err := saveTmpl(c,r)
	        if err != nil {
			c.Errorf("%v",err)
			return
		}
	}
	m := map[string]interface{} {
        "view":"tmpl_admin.tmpl",
	"tlist":map[string]string{"tmpl_admin.tmpl":"模板管理","filelist.tmpl":"文件列表","bloblist.tmpl":"blob列表",},
	}
        err := render(w,m)
        if err != nil {
                c.Errorf("%v",err)
                return
        }

}
func saveTmpl(c appengine.Context, r *http.Request) error {
	return nil
}

