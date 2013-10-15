package open

import(
	"fmt"
	"log"
	"net/http"
	"time"
	"regexp"
	"strconv"

	"appengine"
	"appengine/datastore"

	"rcapps"
)

type Remote struct{
	Ip string
	Name string
	History []Click
}
type Click struct{
	Agent string
	When int64
	Request string
}

func init(){
	http.HandleFunc("/",handler)
	log.SetPrefix("SRC\t")
	log.SetFlags(log.LstdFlags)
}

func handler(w http.ResponseWriter, r *http.Request){
	uri:=r.URL.RequestURI()
	re,err := regexp.Compile("^\\/click\\.*(.*)$")
	if err != nil{
		log.Print(err)
		return
	}
	remoteAddr:=r.RemoteAddr
	log.Printf(remoteAddr)
	c := appengine.NewContext(r)
	switch {
	case uri == "/":
		counter(uri,get_user_agent(w,r),remoteAddr,c)
		http.Redirect(w,r,"/s/index.html",http.StatusMovedPermanently)
		//fmt.Fprintf(w,uri)
// match "click" page
	case re.MatchString(uri):
		var page int
		result := re.FindStringSubmatch(uri)
		if len(result)>1 {
			page, _ = strconv.Atoi(result[1])
		}

		show_click(w,r,page)
	default:
		counter(uri,get_user_agent(w,r),remoteAddr,c)
//		http.Redirect(w,r,"/s"+uri,http.StatusMovedPermanently)
		rcapps.Gate(w,r,uri)
	}
}
func get_user_agent(w http.ResponseWriter, r *http.Request) string{
	h:=r.Header
	for key,value := range h{
		if key == "User-Agent"{
			return value[0]
		}
	}
	return ""
}
func show_click(w http.ResponseWriter, r *http.Request,page int){
	q := datastore.NewQuery("Remote")
	c := appengine.NewContext(r)

	var rs []Remote
	_,err := q.GetAll(c,&rs)
	// reserv the page parameter, using click.<page number> for page in furture
	fmt.Fprint(w,HTML_HEAD)
	fmt.Fprintf(w,"Page %d<br />",page+1)
	fmt.Fprint(w,"<table width=\"100%\" border=\"1\" cellspacing=\"1\" cellpadding=\"1\">")
	if err == nil {
		fmt.Fprint(w,"<TR bgcolor=\"#30abfd\"><TD>Seq</TD><TD>IP</TD><TD>HOST NAME</TD><TD>Click</TD><TD>USER AGENT</TD><TD>REQUEST</TD><TD>LATEST</TD></TR>")
		for key,value := range rs {
			clk := value.History[len(value.History)-1]
			fmt.Fprintf(w,"<TR bgcolor=\"#c1e6ff\"><TD>%d</TD><TD>%s</TD><TD>%s</TD><TD>%d</TD><TD>%s</TD><TD>%s</TD><TD>%s</TD></TR>",key,value.Ip,value.Name,len(value.History),clk.Agent,clk.Request,int64_time(clk.When))
		}
	}
	fmt.Fprint(w,"</TABLE><a href=\"http://appengine.google.com\">")
	fmt.Fprint(w,GOOGLE_LOGO)
	fmt.Fprint(w,"</a>")
	fmt.Fprint(w,HTML_TAIL)
}
func int64_time(i int64) time.Time{
	sec := i/1000000000
	nano := i - sec * 1000000000
	return time.Unix(sec,nano)
}
func counter(uri,agent,ip string,c appengine.Context){
	cur_time := time.Now().UnixNano()
	q := datastore.NewQuery("Remote").Filter("Ip =",ip)
	var rs []Remote
	keys, err := q.GetAll(c,&rs)
	if err == nil {
		if len(rs)==0 {
			h := make([]Click,1,20)
			r := Remote{ip,"",h}
			h[0]= Click{agent,cur_time,uri}
			_, err:= datastore.Put(c,datastore.NewIncompleteKey(c,"Remote",nil),&r)
			if err != nil {
				//has error
				log.Print(err)
			}
		} else {
			r := rs[0]
			h := r.History
			h=append(h,Click{agent,cur_time,uri})
			r.History=h
			_, err := datastore.Put(c,keys[0],&r)
			if err != nil {
				//has error
				log.Print(err)
			}
		}
	} else {
		log.Print(err)
	}
}
const (
	HTML_HEAD="<HTML><HEAD><TITLE>Statistic</TITLE></HEAD><BODY>"
	HTML_TAIL="</BODY></HTML>"
	GOOGLE_LOGO=`<img src="https://developers.google.com/appengine/images/appengine-silver-120x30.gif" alt="Powered by Google App Engine" />`
)

