package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

type DataRow struct{
	Id     int
	Name   string
	Age    int
	About  string
	Gender string
}

func main(){
	
	f,err :=os.Open("./dataset.xml")
	if err!=nil{
		fmt.Println("Can`t open file")
		return
	}
	fileCon,err := ioutil.ReadAll(f)
	xml.Unmarshal(fileCon,dataset)
}