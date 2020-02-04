package dum

import (
	"fmt"
	"strconv"
)

//func ain(){
//	fmt.Println("dummy main...")
//	Add(1,2)
//}

func Add(a int, b int){
	fmt.Print( "A and B sum is : " + strconv.Itoa(a+b))
}