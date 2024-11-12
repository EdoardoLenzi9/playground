package main

import "strings"
import "fmt"

func main(){
	a := "a,b,c,d"
	var b = ""
	for i, chunk := range strings.Split(a, ","){
		if i > 0{
			b += ","
		}
	    b += chunk
	}
	c := strings.Join(strings.Split(a, ","), ",")
	if a == b && a == c {
		fmt.Println("assertion passed")
	} else {
		panic("assertion failed")
	}
}