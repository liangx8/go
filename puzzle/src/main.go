package main
import (
	"fmt"
	"flag"
	"os"
	"io"
	"math/rand"
	"time"
	"bufio"
	"strconv"

	"sortable"

)



type C int

type CSet struct {
	Cs []C
}
// return true if exists
func (cs *CSet)Exists(c C) bool{
	for _,v := range cs.Cs {
		if c == v { return true }
	}
	return false
}
func (cs CSet) String() string {
	return fmt.Sprint(cs.Cs)
}
// return false if duplicated add
func (cs *CSet)Add(c C) bool{
	if cs.Exists(c) { return false }
	cs.Cs = append(cs.Cs,c)
	return true
}
func (cs CSet)Size() int {
	return len(cs.Cs)
}
type M struct{
	seq,weight int
	src,dst C
}
func (m M) String() string{
	return fmt.Sprintf("M%d\tC%d\tC%d\t%d",m.seq,m.src,m.dst,m.weight)
}

func sort(ms []*M,cmp func(l,r *M) int){

}

func usage(){
	fmt.Printf("Usage %s <-G n>|<filename>\n",os.Args[0])
		flag.PrintDefaults()
		return
}
func generate(n int,w io.Writer){
	rand.Seed(time.Now().Unix())
	x :=1
	for i:=0;i<n;i++ {
		for j:=0;j<n;j++ {
			if i==j {continue}
			fmt.Fprintf(w,"M%d\tC%d\tC%d\t%d\n",x,i+1,j+1,rand.Int()%10000)
			x ++
		}
	}
}
func load(f *os.File) ([]interface{},*CSet,error){
	scn := bufio.NewScanner(f)
	scn.Split(bufio.ScanWords)
	var ms []interface{}
	var cs CSet
	start := 0
	var pm *M
	for scn.Scan() {
		txt := scn.Text()

		switch start {
		case 0:
			pm = new(M)
			if txt[0] != 'M' {
				return nil,nil,fmt.Errorf("expecting `M'")
			}
			i,err := asInt(txt[1:])
			if err != nil { return nil,nil,err }
			pm.seq=i

		case 1:
			if txt[0] != 'C' {
				return nil,nil,fmt.Errorf("1 expecting `C'")
			}
			i,err := asInt(txt[1:])
			if err != nil { return nil,nil,err }
			pm.src=C(i)
			cs.Add(pm.src)
		case 2:
			if txt[0] != 'C' {
				return nil,nil,fmt.Errorf("2 expecting `C'")
			}
			i,err := asInt(txt[1:])
			if err != nil { return nil,nil,err }
			pm.dst=C(i)
			cs.Add(pm.dst)
		case 3:
			start=0
			i,err := asInt(txt)
			if err != nil { return nil,nil,err }
			pm.weight=i
			ms = append(ms,pm)
			continue
		}
		start = start + 1
	}
	return ms,&cs,nil
}
func asInt(v string) (int,error){
	i64,err := strconv.ParseInt(v,10,32)
	return int(i64),err
}
func main(){
	num := flag.Int("G",0, "种类数量")
	if len(os.Args)==1 {
		usage()
		return
	}
	flag.Parse()
	if *num > 1 {
		if len(flag.Args())>0 {
			f,err :=os.Create(flag.Arg(0))
			if err != nil {
				fmt.Println(err)
				return
			}
			generate(*num,f)
		} else {
			generate(*num,os.Stdout)
		}


		return
	}
	if len(flag.Args())>0{
		f,err:=os.Open(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		ms,cs,err := load(f)
		if err!=nil {
			fmt.Println(err)
			return
		}
		byWeight :=func(l,r interface{}) (int,error){
			return l.(*M).weight - r.(*M).weight, nil
		}
		bySeq :=func(l,r interface{}) (int,error){
			return l.(*M).seq - r.(*M).seq, nil
		}
		sortable.Bubble(ms,byWeight)
		ms = solution(ms,cs)

		sum :=0
		sortable.Bubble(ms,bySeq)
		for _,m := range ms {
			fmt.Println(m)
			sum += m.(*M).weight
		}
		fmt.Println(sum)
		c,ok := testSolution(ms,cs)
		if ok {
			fmt.Printf("%v 完成\n",c)
			return
		}
		fmt.Printf("%v 未完成\n",c)
		return
	}
	usage()
}

func solution(ms []interface{},cs *CSet) []interface{}{
	solu := make([]interface{},0,cs.Size()*2)
	left(&solu,ms,cs)
	right(&solu,ms,cs)
	return solu
}
func left(ps *[]interface{},ms []interface{},cs *CSet){
	var cc CSet
	for _,m := range *ps {
		cc.Add(m.(*M).src)
	}
	if cc.Size()==cs.Size() { return }
	for _,m := range ms {
		if !cc.Exists(m.(*M).src) {
			*ps = append(*ps,m)
			cc.Add(m.(*M).src)
		}
		if cc.Size()== cs.Size() {return }

	}
}
func right(ps *[]interface{},ms []interface{},cs *CSet){
	var cc CSet
	for _,m := range *ps {
		cc.Add(m.(*M).dst)
	}
	if cc.Size()==cs.Size() { return }
	for _,m := range ms {
		if !cc.Exists(m.(*M).dst) {
			*ps = append(*ps,m)
			cc.Add(m.(*M).dst)
		}
		if cc.Size()== cs.Size() {return }

	}
}


func testSolution(ms []interface{},cs *CSet) ([]*CSet,bool) {
	var cc *CSet
	var ccc []*CSet
	
	for _,c := range cs.Cs {
		if cc != nil {
			if cc.Exists(c) { continue }
		}
		cc=new(CSet)
		cc.Add(c)
		for _,m := range ms {
			pm := m.(*M)
			if pm.src != c {continue}
			cc.Add(pm.dst)
			walk(pm,ms,cc)
			if cc.Size() == cs.Size() {
				return append(ccc,cc),true
			}
		}
		ccc = append(ccc,cc)
	}

	return ccc,false
}


func walk(step *M,ms []interface{},cs *CSet){

	for _,m := range ms {
		pm := m.(*M)
		if pm == step { continue }
		if pm.src == step.dst {
			if cs.Exists(pm.dst) {
				continue
			}
			cs.Add(pm.dst)
			walk(pm,ms,cs)
		}
	}
}
