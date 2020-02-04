package dummy

import (
	"fmt"
	"strconv"
)

func Add(a int, b int){
	fmt.Print( "A and B sum is : " + strconv.Itoa(a+b))
}