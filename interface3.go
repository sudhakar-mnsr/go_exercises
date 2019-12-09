// Sample program to show how the concrete value assigned to 
// the interface is what is stored inside the interface

package main

import "fmt"

// printer displays information.
type printer interface {
   print()
}

// user defines a user in the program.
type user struct {
   name string
}

// print displays the user's name.
func (u user) print() {
   fmt.Printf("User Name: %s\n", u.name)
}

func main() {
   // Create values of type user and admin.
   u := user{"Bill"}
   
   // Add the values and pointers to the slice of 
   // printer interface values.
   entities := []printer{
      // Store a copy of the user value in the interface
      u,
      // Store a copy of the address of the user value
      &u,
   }
   
   // Change the name field on the user value.
   u.name = "Bill_CHG"
   
   // Iterate over the slice of entities and call
   // print against the copied interface value.
   for _, e := range entities {
      e.print()
   }
   
   // When we store a value, the interface value has its own
   // copy of the value. Changes to the original value will
   // not be seen.
   
   // When we store a pointer, the interface value has its own
   // copy of the address. Changes to the original value will
   // be seen.
}
