package session

import (
	"testing"
	"time"
	"fmt"

	"appengine/aetest"
	"appengine/memcache"
)

func Test_marshal(t *testing.T){
	c,err:=aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	cd := memcache.Codec{Marshal:marshal,Unmarshal:unmarshal}
	ssp := sessionpool{100,make(map[int64]*Session)}
	dsp := sessionpool{0,make(map[int64]*Session)}
	ssp.S[1]=&Session{Id:1,last:time.Now().UnixNano(),used:false,}
	item := &memcache.Item{Key:"count",Object:&ssp,}
	if err = cd.Set(c,item); err != nil {
		t.Fatal(err)
	}
	_, err=cd.Get(c,"count",&dsp)
	if err != nil {
		t.Fatal(err)
	}
	if err = spEqual(&ssp,&dsp);err != nil {
		t.Fatal(err)
	}
	s := ssp.S[1]
	time.Sleep(time.Second)
	s.Update(c)
	_, err=cd.Get(c,"count",&dsp)
	if err != nil {
		t.Fatal(err)
	}
	if sd:=dsp.S[1]; !sessionEqual(s,sd) {
		if sd.last == s.last {
			t.Fatal("Session.Update() fail")
		}else {
			t.Log(s)
			t.Log(sd)
		}
	}
}

func Test_SessionSerialize(t *testing.T){
	var s1,s2 Session
	s1.Id=100
	s1.last=200
	s1.used=true
	buf := s1.Serialize()
	n:=s2.Unserialize(buf)
	if len(buf) != n {
		t.Fatal("Not expected")
	}
	if !sessionEqual(&s1,&s2){
		t.Fatal("not expected")
	}
}

func sessionEqual(s1,s2 *Session) bool {
	return s1.Id == s2.Id && s1.last == s2.last && s1.used == s2.used
}
func spEqual(s1,s2 *sessionpool) error {
	if s1.Count != s2.Count {
		return fmt.Errorf("Count is not equal")
	}
	for id,ss := range s1.S {
		sd,ok:=s2.S[id]
		if !ok {
			return fmt.Errorf("id %d error",id)
		}
		if sessionEqual(ss,sd){
			continue
		}
		fmt.Errorf("id %d error",id)
	}
	return nil
}
