// This is for trace the blobs using. a md5 key is defined to keep the blob is unique
// in system
package store
import (
	"appengine"
	"appengine/blobstore"
	"appengine/datastore"
	"io"
	"crypto/md5"
	"fmt"
	"time"

)


type Resource struct{
	Md5str string
	Key appengine.BlobKey
	Update time.Time
}


type ResourceUtil struct{
	c appengine.Context
}

func NewResourceUtil(c appengine.Context) *ResourceUtil{
	return &ResourceUtil{c}
}
func (rs *ResourceUtil)All(cb func(*Resource)) {
	var res Resource
	q := datastore.NewQuery("Resource")
	itr := q.Run(rs.c)
	for {
		_,err:=itr.Next(&res)
		if err == datastore.Done {
			break
		} else if err !=nil {
			rs.c.Errorf("%v",err)
			break
		}
		cb(&res)
	}

}
func (rs *ResourceUtil)SaveUnique(r io.Reader,mimeType string) (*Resource,error){
	blobw,err := blobstore.Create(rs.c,mimeType)
	if err !=  nil {
		return nil,err
	}
	hMd5 := md5.New()
	io.Copy(io.MultiWriter(blobw,hMd5),r)
	md5str := fmt.Sprintf("%x",hMd5.Sum(nil))
	res,err := rs.byMd5str(md5str)
	if res == nil {
		// commit blob
		blobw.Close()
		bkey, err := blobw.Key()
		if err != nil {
			return nil,err
		}
		res = &Resource{md5str,bkey,time.Now()}
		_,err = datastore.Put(rs.c,datastore.NewIncompleteKey(rs.c,"Resource",nil),res)
		if err != nil {
			return nil,err
		}
	}
	return res,nil
}
func (rs *ResourceUtil)byMd5str(md5str string) (*Resource,error){
	q := datastore.NewQuery("Resource").Filter("Md5str =",md5str)
	var res []Resource
	_,err := q.GetAll(rs.c,&res)
	if err != nil {
		return nil,err
	}
	if len(res)>0 {
		return &res[0],nil
	}
	return nil,nil
}
