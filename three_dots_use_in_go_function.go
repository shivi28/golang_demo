//Playground Link : https://play.golang.org/p/zmY-cVGL5MM

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Error : %+v\n",CheckLengthOfArguments("1"))
	fmt.Printf("Success : %+v\n",CheckLengthOfArguments("1","2"))
	fmt.Printf("Error : %+v\n",CheckLengthOfArguments("1","2","3"))
}


func CheckLengthOfArguments(request ...string) error{
	if len(request) != 2 {
		return fmt.Errorf("Number of variables should be exactly 2 not %+v", len(request))
	}
	return nil
}
