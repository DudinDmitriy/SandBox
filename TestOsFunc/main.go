package main

import (
	"fmt"
	"os"
	"sort"
)

type myFileInfo []os.FileInfo

func (mFI myFileInfo) Len() int {
	return len(mFI)
}

func (mFI myFileInfo) Less(i, j int) bool {
	return mFI[i].Name() < mFI[j].Name()
}

func (mFI myFileInfo) Swap(i, j int) {
	mFI[i], mFI[j] = mFI[j], mFI[i]
}

func main() {
	//pathdir := "testdata"
	pathdir := "D:\\Develop\\GoLang\\src\\SandBox\\TestOsFunc"
	//pathdir = "testdatastatic"
	strres := makestring("", pathdir)
	fmt.Println(strres)
}

func makestring(pref string, pathdir string) (res string) {

	res = ""
	var pref1, pref2 string
	//pref1 := fmt.Sprintf(pref + "\t├───")
	f, err := os.OpenFile(pathdir, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Printf("Ошибка при открытии папки")
	}

	fi, err := f.Readdir(0)

	if err != nil {
		fmt.Printf("Ошибка при чтении директории")
	}
	sort.Sort(myFileInfo(fi))
	lenfi := len(fi) - 1

	for el := range fi {
		if el == lenfi {
			pref1 = fmt.Sprintf(pref + "└───")
			pref2 = fmt.Sprintf(pref + " 	")
		} else {
			pref1 = fmt.Sprintf(pref + "├───")
			pref2 = fmt.Sprintf(pref + "│	")
		}
		if fi[el].IsDir() {
			res = res + fmt.Sprintf(pref1+"%s\n", fi[el].Name())
			res = res + makestring(pref2, pathdir+string(os.PathSeparator)+fi[el].Name())
		} else {
			sizefile:=fi[el].Size()
			strsizefile:="(empty)"
			if sizefile != 0{
				strsizefile = fmt.Sprintf("(%db)",sizefile)
			}
			res = res + fmt.Sprintf(pref1+"%s %s\n", fi[el].Name(),strsizefile)
		}
	}
	return
}
