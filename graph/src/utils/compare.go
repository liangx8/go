package utils
import (
	"fmt"
)
const (
	errStr ="Both elements must be type %T(%v @ %T,%v @ %T)"
)
type Compare func(l,r interface{}) (int,error)


func PrimitiveCompare(l,r interface{}) (int,error) {

	switch l_value := l.(type){
		case string:
			r_value,ok := r.(string)
			if !ok {
				return 0,fmt.Errorf(errStr,l,l,l,r,r)
			}
			return str_comp(l_value,r_value),nil
		case int:
			r_value,ok := r.(int)
			if !ok {
				return 0,fmt.Errorf(errStr,l,l,l,r,r)
			}
			return l_value - r_value,nil
		case int8:
			r_value,ok := r.(int8)
			if !ok {
				return 0,fmt.Errorf(errStr,l,l,l,r,r)
			}
			return int(l_value - r_value),nil
		case int16:
			r_value,ok := r.(int16)
			if !ok {
				return 0,fmt.Errorf(errStr,l,l,l,r,r)
			}
			return int(l_value - r_value),nil
		case int32:
			r_value,ok := r.(int32)
			if !ok {
				return 0,fmt.Errorf(errStr,l,l,l,r,r)
			}
			return int(l_value - r_value),nil
		case int64:
			r_value,ok := r.(int64)
			if !ok {
				return 0,fmt.Errorf(errStr,l,l,l,r,r)
			}
			return int(l_value - r_value),nil
		case float32:
			r_value,ok := r.(float32)
			if !ok {
				return 0,fmt.Errorf(errStr,l,l,l,r,r)
			}
			var res int
			if l_value > r_value {
				res = 1
			} else {
				res = -1
			}
			return res,nil
		case float64:
			r_value,ok := r.(float64)
			if !ok {
				return 0,fmt.Errorf(errStr,l,l,l,r,r)
			}
			var res int
			if l_value > r_value {
				res = 1
			} else {
				res = -1
			}
			return res,nil
	}
	return 0,fmt.Errorf("Unkonwn type (%v @ %T)",l,l)
}
func str_comp(l,r string) int{
	if l==r {
		return 0
	} else if l>r {
		return 1
	} else {return -1}
}


