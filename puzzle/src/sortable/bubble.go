package sortable

func Bubble(data []interface{},cmp CompareType)error{
	n := len(data)
	if n < 2 { return nil }

	isOrder := true
	for i:=0;i<n-1;i++ {
		for x:=0;x<n-i-1;x++ {
			rv,err := cmp(data[x],data[x+1])
			if err != nil { return err }
			if rv > 0 {
				isOrder = false
				data[x],data[x+1]=data[x+1],data[x]
			}
		}
		if isOrder { break }
	}
	return nil
}
