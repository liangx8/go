package main

import (
	"fmt"
	"os"
	"regexp"
	"bufio"
)
const (
	search_path = "/home/arm/git/go/goweb"
)
var filter func(p string) bool
var search func(content string)bool
func main(){
	filter = c_filter
	search = content
	err := proc(search_path)
	if(err != nil){
		fmt.Println(err)
	}
}
func content(line string) bool{
	re,err := regexp.Compile("log")
	if err != nil { return false}
	v:=re.FindString(line)
	return v != ""
}
func proc(p string) error{
//	fmt.Println(p)
	fi,err := os.Stat(p)
	if(err != nil){
		return err
	}
	if (fi.IsDir()){
		return proc_dir(p)
	} else {
		return proc_file(p)
	}
	return nil
}

func proc_dir(p string) error{
	f,err := os.Open(p)
	if(err != nil){
		return err
	}
	defer f.Close()
	fis,err := f.Readdir(-1)
	if(err !=nil){
		return err
	}
	if(p== "/") {
		p=""
	}
	for _,fi := range fis {

		err = proc(p+"/"+fi.Name())
		if err != nil { return err }
	}
	return nil
}
func proc_file(p string) error{
	if filter(p){
		return nil
	}
	show_name := 0
	//fmt.Println(p)
	f,err := os.Open(p)
	if(err !=nil){
		if (os.IsPermission(err)){
			fmt.Println(err)
			return nil
		}
		return err
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	line_count :=0
	for {
		line,err :=rd.ReadString('\n')
		line_count++
		if err != nil {
			break
		}
		if(search(line)){
			if(show_name ==0){
				fmt.Printf("====>%s\n",p)
				show_name ++
			}
			fmt.Printf("%4d>%s",line_count,line)
		}
	}
	return nil
}
func c_filter(p string) bool {
	//re,err := regexp.Compile(`.c$|.cpp$|.h$`)
	re,err := regexp.Compile(`.go$`)
	if (err == nil){
		v := re.FindString(p)
		if (v == "") {
			return true
		} else {
			return false
		}
	} else {
		fmt.Println(err)
	}
	return true
}
