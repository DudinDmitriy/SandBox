package main

import (
	"sort"
	"fmt"
	"os"
)

type myFileInfo []os.FileInfo

func (mFI myFileInfo) Len() int{
	return len(mFI)
}

func (mFI myFileInfo) Less(i,j int) bool{
	return mFI[i].Name()<mFI[j].Name()
}

func (mFI myFileInfo) Swap(i,j int) {
	mFI[i],mFI[j] = mFI[j],mFI[i]
}

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
	sort.Sort(myFileInfo(fi)))
	fmt.Println(fi[0])
	fmt.Println(fi[1])
	fmt.Println(fi[2])
	fmt.Println(fi[3])
}