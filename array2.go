package main

import "fmt"

func main() {
   // Declare an array of 5 integers that is initialized 
   // to its zero value.
   var five [5]int
   
   // Declare an array of 4 integers that is initialized 
   // with some valures.
   four := [4]int{10,20,30,40}
   five1 := [5]int{10,20,30,40,50}
   
   // ./array2.go:16:9: cannot use four (type [4]int) as type [5]int in assignment
   five = four
   // five = five1
   
   fmt.Println(five)
   fmt.Println(four)
}
