package main

import "fmt"

func main(){
    heads, tails := 0, 0

    switch coinflip() {
        case "heads": heads++
        case "tails": tails++
        default: fmt.Println("landed on edge!")
    } 
}

func coinflip() string{
    return "heads"
}

func TogglelessSwitch(x int) int {
    switch {
        case x > 0:
            return +1
        case x < 0:
            return -1
        default:
            return 0
    }
} 

// standard packages
// https://golang.org/pkg