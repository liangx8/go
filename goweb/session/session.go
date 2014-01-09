package session

import (
	"time"
	"appengine"



)
var instId int64
type Session interface{
	Id() int64
	Map() map[string]interface{}
}
var chsession = make(chan map[int64]*session_,1)

type session_ struct{
	id int64
	due *time.Timer
	data map[string]interface{}
}
func (s session_)Id() int64{
	return s.id
}
func (s session_)Map() map[string]interface{}{
	return s.data
}
func New(c appengine.Context) Session{
	id :=time.Now().UnixNano()
	s :=&session_{id,time.AfterFunc(15*time.Minute,func(){
// c is not available when deploy to GAE
//		c.Infof("session %d invialidate, instance %d",id,instId)
		ss := <-chsession
		delete(ss,id)
		chsession<- ss
	}),make(map[string]interface{})}
	c.Infof("new session(%d) instance %d",id,instId)
	ss := <-chsession
	ss[id]=s
	chsession<- ss
	return s
}
func Get(id int64) Session{
//	log.Printf("find session %d",id)
	ss := <-chsession
	s,exists :=ss[id]
	if exists{
		s.due.Reset(15*time.Minute)
	}
	chsession<- ss
	if !exists {
//		log.Println("   **** NOT FOUND ****")
		// 返回nil值时，必须这样写，详情参考FAQ Why is my nil error value not equal to nil?
		return nil
	}
//	log.Println(" FOUND ")
	return s
}
func init(){
	//chsession = make(chan map[int64]*session_,1)
	if len(chsession)==0 {
		chsession<- make(map[int64]*session_)
	}
	instId = time.Now().UnixNano()
/*
	log.Printf("***********************************************************")
	log.Printf("*              initializing session,%d                  *",k)
	log.Printf("***********************************************************")
	go func(){
		idx := 0
		for {
			idx ++
			m := <-chsession
			log.Printf("%d]%d:%v",k,idx,m)
			chsession<- m
			time.Sleep(19*time.Second)
		}
	}()
*/
}
func List() []int64{
	m := make([]int64,0,40)
	m=append(m,instId)
	ss := <-chsession
	for k := range ss {
		m=append(m,k)
	}
	chsession <- ss
	return m
}
/*
func Log(c appengine.Context){
	ss := <-chsession
	for k,v := range ss {
		c.Infof("%v=%v",k,v)
	}
	chsession <- ss
}
*/
