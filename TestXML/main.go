package main

import (
	"encoding/xml"
	"io/ioutil"
	"fmt"
	"os"
)

type DataRow struct{
	Id     int `xml:"id"`
	Name   string `xml:"first_name"`
	Age    int `xml:"age"`
	About  string `xml:"about"`
	Gender string `xml:"gender"`
}
type DataSet struct{
	Rows []DataRow `xml:"row"`
}

func main(){
	
	f,err :=os.Open("./dataset.xml")
	if err!=nil{
		fmt.Println("Can`t open file")
		return
	}
	
	dataset := &DataSet{}
	fileCon,err := ioutil.ReadAll(f)
	//fmt.Println(string(fileCon))
	
	xml.Unmarshal(fileCon,dataset)
    
	for _,d := range dataset.Rows {
		fmt.Printf("Strucct: %v /n",d)

	}
}