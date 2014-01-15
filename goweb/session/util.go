package session
import (

	"fmt"
)

func marshal(src interface{})([]byte,error){
	buf := make([]byte,0,100)
	//ss,ok := src.(map[int64]*Session)
	sp,ok := src.(*sessionpool)
	if !ok {
		return nil,fmt.Errorf("can handle type %T",src)
	}
	buf = append(buf,sp.CountSerialize()...)
	for _,s := range sp.S{
		buf=append(buf,s.Serialize()...)
	}
	return buf,nil
}

func unmarshal(src []byte,dest interface{})error{
	//pss,ok := dest.(*map[int64]*Session)
	psp,ok := dest.(*sessionpool)
	if !ok {
		return fmt.Errorf("can handle type %T",dest)
	}
	idx:=0
	size:=len(src)
	n:=psp.CountUnserialize(src)
	idx = idx + n
	for {
		buf := src[idx:]
		s:=&Session{}
		n=s.Unserialize(buf)
		idx=idx+n
		psp.S[s.Id]=s
		if idx == size {break}
	}
	return nil
}
