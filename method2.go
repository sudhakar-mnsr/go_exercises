// Sample program to show how to declare function variables.
package main

import "fmt"

// data is a struct to bing methods to .
type data struct {
   name string
   age int
}

// displayName provides a pretty pring view of name.
func (d data) displayName() {
   fmt.Println("My Name Is", d.name)
}

// setAge sets the age and displays the value.
func (d *data) setAge(age int) {
   d.age = age
   fmt.Println(d.name, "Is Age", d.age)
}

func main() {
   // Declare a variable of type data.
   d := data{
      name: "Bill",
   }
   
   // How we actually call methods in Go.
   d.displayName()
   d.setAge(45)
   
   fmt.Println("\nWhat the compiler is Doing:")
   
   // This is what GO is doing underneath.
   
   data.displayName(d)
   (*data).setAge(&d, 45)
   
   // ========================================================================
   
   fmt.Println("\nCall value reciever methods with Variable:")
   
   // Declare a function variable for the method bound to the d variable.
   // The function variable will get its own copy of d because the method 
   // is using a value receiver.
   f1 := d.displayName
   
   // Call the method via the variable.
   f1()
   
   // Change the value of d.
   d.name = "Joan"
   
   // Call the method via the variable. We dont see the change.
   f1()
   
   // ========================================================================
   fmt.Println("\nCall Pointer receiver method with variable:")
   
   // Declare a function variable for the method bound to the d variable
   // The function variable will get the address of d because the method
   // is using a pointer receiver
   f2 := d.setAge
   
   // Call the method via the variable.
   f2(45)
   
   // Change the value of d.
   d.name = "Sammy"
   
   // Call the method via the variable. We see the change.
   f2(45)
}
