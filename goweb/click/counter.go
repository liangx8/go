package click
import (
	"net/http"
	"time"
	"appengine"
	"appengine/datastore"
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
func Counter(r *http.Request){
	uri := r.URL.RequestURI()
	if uriFilter(uri){return}
	clk := Click{r.RemoteAddr,"",time.Now().UnixNano(),uri,r.Header.Get("User-Agent")}
	ch <- &send{&clk,appengine.NewContext(r)}
	return
}
func init(){
	ch = make(chan *send)
	go forever()
}
func forever(){
	var last int64
	last = 0
	for {
		sendc := <- ch
		if(sendc.clk.When - last)> 2000000000 {
			last = sendc.clk.When
			_,err:=datastore.Put(sendc.c,datastore.NewIncompleteKey(sendc.c,"Click",nil),sendc.clk)
			if err != nil {
				sendc.c.Errorf("%v",err)
			}
		}
	}
}
func uriFilter(uri string) bool{
	if uri == "/favicon.ico" {return true}
	if uri[:7] == "/images" {return true}
	return false
}
