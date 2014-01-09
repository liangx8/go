// virtual path
package store
import (
	"appengine"
	"appengine/datastore"
	"time"
)
type Vpath struct {
	Name string
	Key appengine.BlobKey
	Update time.Time
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
	vp.Update=time.Now()
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
func (vpu *VpathUtil)All(cb func(*Vpath)) {
	var vp Vpath
	q := datastore.NewQuery("Vpath")
	itr := q.Run(vpu.c)
	for {
		_,err := itr.Next(&vp)
		if err == datastore.Done {
			break
		} else if err != nil {
			vpu.c.Errorf("%v",err)
			break
		}
		cb(&vp)
	}
}
