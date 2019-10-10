package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type User struct {
	Id     int
	Name   string
	Age    int
	About  string
	Gender string
}

type DataRow struct {
	Id        int    `xml:"id"`
	FirstName string `xml:"first_name"`
	LastName  string `xml:"last_name"`
	Age       int    `xml:"age"`
	About     string `xml:"about"`
	Gender    string `xml:"gender"`
}
type DataSet struct {
	Rows []DataRow `xml:"row"`
}

func main() {

	f, err := os.Open("./dataset.xml")
	if err != nil {
		fmt.Println("Can`t open file")
		return
	}

	sName := "BoydWolf"
	sAbout := `Nulla cillum enim voluptate consequat laborum esse excepteur occaecat commodo nostrud excepteur ut cupidatat. Occaecat minim incididunt ut proident ad sint nostrud ad laborum sint pariatur. Ut nulla commodo dolore officia. Consequat anim eiusmod amet commodo eiusmod deserunt culpa. Ea sit dolore nostrud cillum proident nisi mollit est Lorem pariatur. Lorem aute officia deserunt dolor nisi aliqua consequat nulla nostrud ipsum irure id deserunt dolore. Minim reprehenderit nulla exercitation labore ipsum.`

	dataset := &DataSet{}
	fileCon, err := ioutil.ReadAll(f)
	//fmt.Println(string(fileCon))

	xml.Unmarshal(fileCon, dataset)
	users := make([]User, 0)

	query := sName + sAbout + string(rune(10))
	query = ""

	for _, d := range dataset.Rows {

		if query == "" || query == (d.FirstName+d.LastName+d.About) {
			usr := User{}
			usr.Name = d.FirstName + d.LastName
			usr.Id = d.Id
			usr.Age = d.Age
			usr.About = d.About
			usr.Gender = d.Gender

			users = append(users, usr)
		}
	}

	type sortfunc func(i, j int) bool
	var sf sortfunc

	sf = func(i, j int) bool {
		if users[i].Id < users[j].Id {
			return true
		}

		if users[i].Id == users[j].Id && users[i].Age < users[j].Age {
			return true
		}

		if users[i].Id == users[j].Id && users[i].Age == users[j].Age && users[i].Name < users[j].Name {
			return true
		}
		return false
	}

	sort.Slice(users, sf)
	fmt.Printf("Data: \n %v+", users)

}

// strSearch := sName + sAbout + string(rune(10))
// ii := sort.Search(len(users), func(i int) bool { return (strSearch) <= (users[i].Name + users[i].About) })
// str1 := users[ii].Name + users[ii].About
// //str2:=sName+sAbout
// bl1 := (ii < len(users))
// bl2 := (str1 == strSearch)
// if bl1 && bl2 {
// 	fmt.Printf("Найден элемент: Id - %d \n", users[ii].Id)
// }
