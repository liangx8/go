package array

type Vector struct {
	e []interface{}
}

func (v *Vector)Sort(c Comparator){
	size:=len(v.e)
	if (size < 2) {return}
	for i:=0;i<size-1;i++ {
		order := true
		for j:=0; j<size-i-1;j++{
			if(c(v.e[i],v.e[i+1])>0){
				v.e[j],v.e[j+1] = v.e[j+1],v.e[j]
				order = false
			}
		}
		if order { break }
	}
}
func (v *Vector)ElementOf(index int) interface{}{
	return v.e[index]
}
func NewVector(size,capicity int) Vector{
	return Vector{make([]interface{},size,capicity)}
}
func (v *Vector)PushBack(e interface{}){v.e=append(v.e,e)}
func (v *Vector)PushFront(e interface{}){
        a := make([]interface{},1,cap(v.e))
        a[0]=e
        for _,ee := range v.e {
                a = append(a,ee)
        }
        v.e=a
}
func (v *Vector)PopFront() interface{}{
	if v.Size()==0 {
		return nil
	}
	e := v.e[0]
	v.e = v.e[1:]
	return e
}
func (v *Vector)PopBack() interface{}{
	if v.Size()==0 {
		return nil
	}
	last := len(v.e)-1
	e:=v.e[last]
	v.e=v.e[:last]
	return e
}
func (v *Vector)Size() int{return len(v.e)}
func (v *Vector)Clear(){
	v.e=make([]interface{},0,cap(v.e))
}


