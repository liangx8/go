package utils
import (
	"reflect"
	"fmt"
)
// copy primitive object 

func Cp(src,dst interface{}) error{
	dv := reflect.ValueOf(dst)
	if dv.Kind() != reflect.Ptr {
		return fmt.Errorf("dst %T is not a pointer",dst)
	}
	if !dv.Elem().CanSet() {
		return fmt.Errorf("dst %T can't be set, is it a nil value?",dst)
	}
	tSrc:=reflect.TypeOf(src)
	tDst:=dv.Elem().Type()
	if tSrc != tDst {
		return fmt.Errorf("src type %v does not match dest type %v",tSrc,tDst)
	}
	// panic if dst is not a pointer what point to type of src, or dst is a nil
	dv.Elem().Set(reflect.ValueOf(src))
	return nil
}
