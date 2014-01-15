package web

import (
	"reflect"
	"net/http"
	"fmt"
	"strconv"

	"appengine"

	"session"
	"aop"
)

func handleFunc(pattern string,handler interface{}){
	var signiture func(http.ResponseWriter,*http.Request)
	t := reflect.TypeOf(handler)
	numIn := t.NumIn()
	in := make([]reflect.Type,numIn)
	for i:=0;i<numIn;i++ {
		in[i]=t.In(i)
	}

	aop.Inject(&signiture,handler,createRequestInjector(in),nil)
	http.HandleFunc(pattern,signiture)
}
// argtype are only available for ResponseWriter, Request, and Session
func createRequestInjector(argtype []reflect.Type) func([]reflect.Value) []reflect.Value{
	return func(in []reflect.Value) []reflect.Value{
		// in[0] expected http.ResponseWriter
		// in[1] expected *http.Request
		r,ok := in[1].Interface().(*http.Request)
		if !ok {
			panic(fmt.Errorf("the second argument shoud be *http.Request, but %v",in[1].Type()))
		}
		w,ok := in[0].Interface().(http.ResponseWriter)
		if !ok {
			panic(fmt.Errorf("the first argument shoud be http.ResponseWriter, but %v",in[0].Type()))
		}
		numArg := len(argtype)
		args := make([]reflect.Value,numArg)
		for i:=0;i<numArg;i++ {
			switch {
				default:
					panic(fmt.Errorf("unsupport type %v",argtype[i]))
				case argtype[i] == in[0].Type():
					args[i]=in[0]
				case argtype[i] == in[1].Type():
					args[i]=in[1]
				case argtype[i] == reflect.TypeOf((**session.Session)(nil)).Elem():
					var id int64
					c := appengine.NewContext(r)
					//c.Infof(appengine.AppID(c))
					cookie,err := r.Cookie("SESSIONID")
					if err != nil {
							if err != http.ErrNoCookie {
									c.Errorf("%v",err)
							}
							id =0
					} else {
							id,err = strconv.ParseInt(cookie.Value,10,64)
							if err != nil {
									c.Errorf("%v",err)
									id =0
							}
					}
					s :=&session.Session{Id:id}
					if session.Get(c,s) {
							// for every page has own session
							// w.Header().Add("Set-Cookie",fmt.Sprintf("SESSIONID=%d;path=%s",session.Id(),r.URL.RequestURI()))
							// all page have one session
							w.Header().Add("Set-Cookie",fmt.Sprintf("SESSIONID=%d;path=/",s.Id))
					} else {
						s.Update(c)
					}
					args[i]=reflect.ValueOf(s)

			}
		}
		return args
	}
}
