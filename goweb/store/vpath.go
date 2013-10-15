// virtual path
package store
import (
	"appengine"
	"appengine/datastore"
)
type Vpath struct {
	Name string
	Key appengine.BlobKey
}

type VpathUtil struct {
	c appengine.Context
}
func NewVpathUtil(c appengine.Context) *VpathUtil{
	return &VpathUtil{c}
}
// according fn to save or update
func (vpu *VpathUtil)SaveOrUpdate(bk appengine.BlobKey,fn string) (*Vpath,error){
	var vp Vpath
	key,err := vpu.FindOne(fn,&vp)
	if err != nil {
		return nil,err
	}
	vp.Name=fn
	vp.Key=bk
	if key == nil {
		_,err=datastore.Put(vpu.c,datastore.NewIncompleteKey(vpu.c,"Vpath",nil),&vp)
	} else {
		_,err=datastore.Put(vpu.c,key,&vp)
	}
	return &vp,nil
}
func (vpu *VpathUtil)FindOne(fn string,vp *Vpath) (*datastore.Key,error){
	q := datastore.NewQuery("Vpath").Filter("Name =",fn)
	itr := q.Run(vpu.c)
	key,err := itr.Next(vp)
	if err == datastore.Done {
		// no found
		return nil,nil
	} else if err != nil {
		return nil,err
	}
	return key,nil
}
func (vpu *VpathUtil)All(cb ResultCallback) error {
	var vp Vpath
	count := 0
	q := datastore.NewQuery("Vpath")
	itr := q.Run(vpu.c)
	key,err := itr.Next(&vp)
	if err == datastore.Done {
		cb.Head(true)
	} else if err != nil {
		return err
	} else {
		cb.Head(false)
		for {
			cb.Each(key,&vp,count)
			count ++
			key,err = itr.Next(&vp)
			if err == datastore.Done{break}
			if err != nil {
				return err
			}
		}
	}
	cb.Tail(count)
	return nil
}
