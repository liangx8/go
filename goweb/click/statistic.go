package click

import (
	"net/http"
	"appengine/datastore"
	"appengine"
	"time"

)
const pageSize int = 20
func ListCounter(w http.ResponseWriter, c appengine.Context,page int) map[string]interface{}{
	q := datastore.NewQuery("Click").Order("When")
	total,err:=q.Count(c)
	if err != nil {
		c.Errorf("%v",err)
		return nil
	}
	it := q.Offset(page * pageSize).Limit(pageSize).Run(c)
	ch := make(chan *Click)
	go func(){
		for {
			cl := &Click{}
			_,err := it.Next(cl)
			if err == datastore.Done { break }
			if err != nil {
				c.Errorf("%v",err)
				break
			}
			ch <- cl
		}
		close(ch)
	}()
	model := map[string]interface{} { "view":"click.tmpl","data":ch,"total":total,}
	model["start"]=page * pageSize
	model["int64time"]=int64Time
	model["odd"]=func(n int) bool{return n%2 == 0}
	model["add"]=func(n,m int) int{ return n+m }
	model["lastpage"]=total/pageSize

	return model
}
func int64Time(i int64) time.Time {
	sec := i/1000000000
	nano := i - sec * 1000000000
	return time.Unix(sec,nano)
}

