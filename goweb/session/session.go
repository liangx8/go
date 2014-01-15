// 只是记录当前的SESSION
// id, time,
package session

import (
	"time"
	"appengine"
	"appengine/memcache"
	"encoding/binary"
	"fmt"
)
type Session struct{
	Id, last int64
	used bool
}
type sessionpool struct {
	Count int64
	S map[int64]*Session
}
func (sp sessionpool)CountSerialize()[]byte {
	buf := make([]byte,binary.MaxVarintLen64)
	n := binary.PutVarint(buf,sp.Count) // panic if buf too small
	return buf[:n]
}
func (sp *sessionpool)CountUnserialize(buf []byte)int {
	v,n:=binary.Varint(buf)
	if n <=0 {
		panic("Count unserialized error")
	}
	sp.Count=v
	return n
}
func ListSession(c appengine.Context) *sessionpool{
	return cache(c)
}
func (s Session)Serialize()[]byte{

	buf := make([]byte,0,binary.MaxVarintLen64*3)
	buf1 := make([]byte,binary.MaxVarintLen64)
	idx:=0;
	n :=binary.PutVarint(buf1,s.Id) // panic if buf too small
	idx=n
	buf = append(buf, buf1[:n]...)
	n =binary.PutVarint(buf1,s.last)
	idx=idx+n
	buf = append(buf, buf1[:n]...)

	if s.used {
		buf = append(buf,byte(1))
	} else {
		buf = append(buf,byte(0))
	}
	return buf
}
func (s *Session)Unserialize(buf []byte)int{
	v,n:=binary.Varint(buf)
	idx := 0
	if n <=0 {
		panic("session unserialized error")
	}
	idx = n
	s.Id=v
	buf = buf[n:]
	v,n = binary.Varint(buf)
	if n <=0 {
		panic("session unserialized error")
	}
	idx = idx + n
	s.last=v
	if buf[n]==0 {
		s.used=false
	} else {
		s.used=true
	}
	return idx+1
}
// 获得或者新建
// return true 新SESSION, false 旧SESSION
func Get(c appengine.Context,s *Session)bool{
	sp := cache(c)
	tmp,ok:=sp.S[s.Id]
	if ok {
		s.last,s.used = tmp.last,tmp.used
		return false
	} else {
		id:=time.Now().UnixNano()
		s.Id,s.last,s.used = id,id,false
		sp.S[id]=s
		toCache(c,sp)
		return true
	}
	panic("impossible")
}
// 如果id已经过期，删除
func (s *Session)Update(c appengine.Context) {
	now:=time.Now().UnixNano()
	if now-s.last > int64(5*time.Minute) {
		removeCache(c,s.Id)
	}else {
		s.last=now
		sp:=cache(c)
		sp.S[s.Id]=s
		toCache(c,sp)
	}
}
func (s *Session)IsUsed(c appengine.Context) bool{
	if s.used {return true}
	s.used=true
	sp:=cache(c)
	sp.Count ++
	c.Infof("Total click %d:",sp.Count)
	sp.S[s.Id]=s
	toCache(c,sp)
	return false
}
func (s Session)String() string{
	return fmt.Sprintf("%d,%d,%v",s.Id,s.last,s.used)
}
// 把SESSION池保存到MEMCACHE
func toCache(c appengine.Context,sp *sessionpool){

// delete obsolate session
	dels := make([]int64,0,len(sp.S))
	now :=time.Now().UnixNano()
	for id,s :=range sp.S{
		if now-s.last > int64(5*time.Minute){
			dels = append(dels,id)
		}
	}
	for i:= range dels {
		c.Infof("session %d was removed",dels[i])
		delete(sp.S,dels[i])
	}


	cd := memcache.Codec{Marshal:marshal,Unmarshal:unmarshal}
	item := memcache.Item{
	//	Key:appengine.AppID(c),
		Key:"session_pool",
		Object:sp,
	}
	if err:=cd.Set(c,&item); err != nil {
		panic(err)
	}
}
func cache(c appengine.Context) *sessionpool{
	sp := &sessionpool{Count:0,S:make(map[int64]*Session)}
	cd := memcache.Codec{Marshal:marshal,Unmarshal:unmarshal}
	//_,err:=cd.Get(c,appengine.AppID(c),&ss)
	_,err:=cd.Get(c,"session_pool",sp)
	if err != nil {
		c.Infof("server say:%v",err)
		//toCache(c,sp)
	}
	return sp
}
func removeCache(c appengine.Context,id int64){
	sp:=cache(c)
	delete(sp.S,id)
	toCache(c,sp)
}

