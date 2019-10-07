package main

import (
	"fmt"
	"sort"
)

func changestr(a, b []int) bool {
    mi := func(iii int) int{
		if(iii<0){
			return -iii
		}
		
		return iii
	}
	res := true
	len_a := len(a)
	len_b := len(b)
	i_a := 0
	i_b := 0
	if len_a == len_b && len(a)==0 {
		return false
	}

	if len_a == 0 && len_b > 0 || len_b == 0 && len_a > 0 {
		return false
	}

	sort.Slice(a,func (ii,jj int) bool{
		return mi(a[ii])>mi(a[jj])
	})
	sort.Slice(b,func (ii,jj int) bool{
		return b[ii]>b[jj]
	})

	for {
		if a[i_a]*a[i_a] != b[i_b] {
			return false
		}
		
		lA:
		for i_a<(len_a-1){
		  if mi(a[i_a])==mi(a[i_a+1]) {
			  i_a++
			} else {
				break lA
			}
		}
		
		lB:
		for i_b<(len_b-1){
			if b[i_b]==b[i_b+1]{
				i_b++
			  } else {
				  break lB
			  }
		  }
		  if i_a == (len_a-1) && i_b == (len_b-1) {
			return true
		}
		if i_a == (len_a-1) && i_b != (len_b-1) || i_a != (len_a-1) && i_b == (len_b-1){
			return false
		}
		
		i_a++
		i_b++
	}

	return res
}

func main() {
	// a = [121, 144, 19, 161, 19, 144, 19, 11]  
	// b = [132, 14641, 20736, 361, 25921, 361, 20736, 361]
	var a,b []int
	a = []int{121, 144, 19, 161, 19, 144, 19, 11}
	b = []int{131, 14641, 20736, 361, 25921, 361, 20736, 361}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("result: %s \n",changestr(a,b))

	a = []int{121, 144, 19, 161, 19, 144, 19,19,1,1}
	b = []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("result: %s \n",changestr(a,b))

	a = []int{121, 144, 19, 161, 19, 144, 19, 11,11}
	b = []int{ 121, 14641, 20736, 361, 25921,  361, 20736, 361}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("result: %s \n",changestr(a,b))

	a = nil
	b = []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("result: %s \n",changestr(a,b))

	a = nil
	b = []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("result: %s \n",changestr(a,b))

	a = []int{-5,3,6,-7}
	b = []int{25,49,9,36,-36}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("result: %s \n",changestr(a,b))
}
