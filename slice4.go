// Sample program to show how the for range has both value and pointer semantics

package main

import "fmt"

func main() {
   // Using the value semantic form of the for range
   
   friends := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
   for _, v := range friends {
      friends = friends[:2]
      fmt.Printf("v[%s]\n", v)
   }
   // Using the pointer semantic for of the for range
   friends = []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
   for i := range friends {
      friends = friends[:2]
      fmt.Printf("v[%s]\n", friends[i])
   }
}
