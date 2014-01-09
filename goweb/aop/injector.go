package aop

import (
	"reflect"
)


func Inject(fptr,handler interface{},before,after func([]reflect.Value)[]reflect.Value){
        wrap := func(in []reflect.Value) []reflect.Value{
                var b []reflect.Value
                if before == nil {
                        b = in
                } else {
                        b = before(in)
                }
                hv := reflect.ValueOf(handler)
                r := hv.Call(b)
                if after == nil {
                        return r
                } else {
                        return after(r)
                }
        }
        fn := reflect.ValueOf(fptr).Elem()
        v := reflect.MakeFunc(fn.Type(),wrap)
        fn.Set(v)
}

