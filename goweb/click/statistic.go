package click

import (
	"net/http"
	"fmt"
	"appengine/datastore"
	"appengine"
	"time"

	"dao"
)
const pageSize int = 20
func ListCounter(w http.ResponseWriter, r *http.Request,page int){
	c:=appengine.NewContext(r)
	util:=dao.NewUtil(c)
	q := datastore.NewQuery("Click").Order("When")
	clk = &Click{}
	err := util.Execute(q,&cCallback{w},page * 20,20)
	if  err != nil {
		c.Errorf("%v",err)
	}
}
func int64Time(i int64) time.Time {
	sec := i/1000000000
	nano := i - sec * 1000000000
	return time.Unix(sec,nano)
}

var clk *Click
type cCallback struct {
	w http.ResponseWriter
}
func (cb *cCallback)Head(total int){
	if total>0 {
	lp := total/pageSize
	fmt.Fprintf(cb.w,`Total %d record(s)/ %d pages, <a href="click.%d">last page</a>`,total,lp+1,lp);
	fmt.Fprintf(cb.w,`<table bgcolor="black"><tr bgcolor="grey"><td>x</td><td>IP</td><td>when</td><td>request</td><td>Agent</td></tr>`)
	} else {
		fmt.Fprintf(cb.w,"0 record");
	}
}
func (cb *cCallback)Each(k *datastore.Key, num int) {
	fmt.Fprintf(cb.w,`<tr bgcolor="white"><td>%d</td><td>%s</td><td>%v</td><td>%s</td><td>%s</td></tr>`,num,clk.Ip,int64Time(clk.When),clk.Request,clk.Agent)
}
func (cb *cCallback)Tail(count int){
	fmt.Fprint(cb.w,"</table>")
}
func (cb *cCallback)Holder() interface{} {
	return clk
}
