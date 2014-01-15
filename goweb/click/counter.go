package click
import (
	"net/http"
	"time"
	"appengine"
	"appengine/datastore"

	"session"
)
type Click struct{
	Ip string
	Host string
	When int64
	Request string
	Agent string
}
func Counter(r *http.Request,s *session.Session){
	c:=appengine.NewContext(r)
	if s.IsUsed(c) { return }
	uri := r.URL.RequestURI()
	clk := Click{r.RemoteAddr,"",time.Now().UnixNano(),uri,r.UserAgent()}
	_,err:=datastore.Put(c,datastore.NewIncompleteKey(c,"Click",nil),&clk)
	if err != nil {
		c.Errorf("%v",err)
	}
}
