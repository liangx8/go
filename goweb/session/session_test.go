package session

import (
	"appengine/aetest"
	"testing"
)

func Test_session(t *testing.T){
	c,err:=aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	c.Infof("success")
}
