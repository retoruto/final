package main

import (
	"agency"
	"fmt"
	"time"
)
func main()  {
        var a string
        fmt.Println("There are two modes to be compare")
        fmt.Println("choose N to run the Naive")
        fmt.Println("choose O to run the Optimized")
        fmt.Scanln(&a)
        if a=="N" {
        t1 := time.Now()
        agency.Naive()
        elapsed1 := time.Since(t1)
        fmt.Println("Naive Approach Time: ", elapsed1)
        } else if a == "O" {
        t2 := time.Now()
        agency.Optimized()
        elapsed2 := time.Since(t2)
        fmt.Println("Optimized Approach Time: ", elapsed2)
        } else {
        fmt.Println("Error!Chosse one of two modes:")
        fmt.Println("choose N to run the Naive")
        fmt.Println("choose O to run the Optimized")
        }

}
