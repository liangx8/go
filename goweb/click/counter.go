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
type send struct {
	clk *Click
	c appengine.Context
}
var ch chan *send
func Counter(r *http.Request,s session.Session){
	if s.Map()["touch"] != nil {
		return
	}
	c:=appengine.NewContext(r)
	s.Map()["touch"]=true

	uri := r.URL.RequestURI()
	clk := Click{r.RemoteAddr,"",time.Now().UnixNano(),uri,r.UserAgent()}
	ch <- &send{&clk,c}
}
func init(){
	ch = make(chan *send)
	go forever()
}
func forever(){
	var last int64
	last = 0
	for sendc := range ch{
		if(sendc.clk.When - last)> 2000000000 {
			last = sendc.clk.When
			_,err:=datastore.Put(sendc.c,datastore.NewIncompleteKey(sendc.c,"Click",nil),sendc.clk)
			if err != nil {
				sendc.c.Errorf("%v",err)
			}
		}
	}
}
