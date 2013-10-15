package dao
import (
	"appengine"
	"appengine/datastore"
)
type utilImpl struct {
	c appengine.Context
}
type Util interface{
	Execute(q *datastore.Query,cb callback,offset,limit int)error
}
func NewUtil(c appengine.Context) Util{
	return &utilImpl{c}
}
type callback interface{
	Head(total int)
	Each(k *datastore.Key,num int)
	Tail(count int)
	Holder() interface{}
}

// there is no limit if limit == 0
func (u *utilImpl)Execute(q *datastore.Query,cb callback,offset,limit int)error{
	holder := cb.Holder()
	count,err:=q.Count(u.c)
	if err != nil {
		return err
	}
	cb.Head(count)
	if count >0 {
		if limit >0 {
			q=q.Offset(offset).Limit(limit)
		}
		itr := q.Run(u.c)
		count =0
		for {
			key,err := itr.Next(holder)
			if err == datastore.Done {
				break
			}
			if err != nil {
				return err
			}
			cb.Each(key,offset+count)
			count ++
		}
	}
	cb.Tail(count)
	return nil
}
