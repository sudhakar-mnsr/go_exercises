// This is an example of using composition and interfaces. This is
// something we want to do in Go. We will group common types by 
// their bahavior and not by their state. This pattern does 
// provide a good design principle in a Go program.
package main

import "fmt"

// Speaker provide a common behavior for all concrete types
// to follow if they want to be a part of this group. This is
// a contract for these concrete types to follow.
type Speaker interface {
   Speak()
}

// Dog contains everything a Dog needs.
type Dog struct {
   Name string
   IsMammal bool
   PackFactor int
}

// Seapk knows how to speak like a dog.
// This makes a Dog now part of group of concrete
// types that know how to speak.
func (d *Dog) Speak() {
   fmt.Printf("Woof! My name is %s, it is %t Iam a mammal with a pack factor of %d.\n", d.Name, d.IsMammal, d.PackFactor,)
}

// Cat contains everything a Cat needs.
type Cat struct {
   Name string
   IsMammal bool
   ClimbFactor int
}

// Speak knows how to speak like a cat
// This makes a cat now part of a group of concrete 
// types that know how to speak.
func (c *Cat) Speak() {
   fmt.Printf("Meow! My name is %s, it is %t Iam a mammal with a climb factor of %d.\n", c.Name, c.IsMammal, c.ClimbFactor,)
}

func main() {

   // Create a list of Animals that knows how to speak.
   speakers := []Speaker{
               // Create a Dog by initializing its animal parts
               // and then its specific Dog attributes.
               &Dog{
                  Name: "Fido",
                  IsMammal: true,
                  PackFactor: 5,
               },
               // Create a cat by initializing its animal parts
               // and then its specific cat attributes.
               &Cat{
                  Name: "Milo",
                  IsMammal: true,
                  ClimbFactor: 4,
               },
   }
   
   // Have the animals speak.
   for _, spkr := range speakers {
           spkr.Speak()
   }
}
  

