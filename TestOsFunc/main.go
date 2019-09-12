package main

import (
	"sort"
	"fmt"
	"os"
)

func main(){
	pathdir := "testdata"
	f,err :=os.OpenFile(pathdir,os.O_RDONLY,os.ModeDir)
	if err!=nil{
		fmt.Printf("Ошибка при открытии папки")
	}
	
	fi,err := f.Readdir(0)

	if err!=nil{
		fmt.Printf("Ошибка при чтении директории")
	}
    sortfilename := func(f1,f2 os.FileInfo)(bool){
		if f1.Name()>f2.Name() {
			return true
		}
		return false
	} 
	sort.Sort(struct{
		fi,
		sortfilename
	})
}