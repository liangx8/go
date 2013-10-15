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
)

type ResultCallback interface{
	Head(isEmpty bool)
	Each(k *datastore.Key,e interface{},num int) error
	Tail(count int)
}
type Resource struct{
	Md5str string
	Key appengine.BlobKey
}

type ResourceUtil interface{
	SaveUnique(r io.Reader,mimeType string) (*Resource,error)
	All(cb ResultCallback) error
}
type resu struct{
	c appengine.Context
}

func NewResourceUtil(c appengine.Context) ResourceUtil{
	return &resu{c}
}
func (rs *resu)All(callback ResultCallback) error {
	var res Resource
	count := 0
	q := datastore.NewQuery("Resource")
	itr := q.Run(rs.c)
	key,err:=itr.Next(&res)

	if err == datastore.Done {
		callback.Head(true)
	} else if err !=nil {
		return err
	} else {
		callback.Head(false)
		for {
			callback.Each(key,&res,count)
			count ++
			key,err = itr.Next(&res)
			if err == datastore.Done {break}
		}
	}
	callback.Tail(count)
	return nil
}
func (rs *resu)SaveUnique(r io.Reader,mimeType string) (*Resource,error){
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
		res = &Resource{md5str,bkey}
		_,err = datastore.Put(rs.c,datastore.NewIncompleteKey(rs.c,"Resource",nil),res)
		if err != nil {
			return nil,err
		}
	}
	return res,nil
}
func (rs *resu)byMd5str(md5str string) (*Resource,error){
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
