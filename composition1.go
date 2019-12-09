// ERROR CASE
// This is an example of using type hierarchies with a oop pattern
// This is not something we want to do in GO. Go does not have the 
// concept of sub-typing. All types are their own and the concepts 
// of base and derived types do not exist in Go. This pattern does not 
// provide a good design principle in Go program.

package main

import "fmt"

// Animal contains all the base fields for animals.
type Animal struct {
   Name string
   IsMammal bool
}

// Speak provides generic behavior for all animals and
// how they speak.
func (a *Animal) Speak() {
   fmt.Printf("UGH! My name is %s, It is %t Iam a mammal\n", a.name,a.IsMammal,)
}

// Dog contains everything an animal is but specific attributes
// that only a Dog has
type Dog struct {
   Animal
   PackFactor int
}

// Speak knows how to speak like a dog.
func (d *Dog) Speak() {
   fmt.Printf("Woof! My name is %s, it is %t Iam a mammal with a pack factor of %d.\n", d.Name, d.IsMammal, d.PackFactor,)
}

// Cat contains everything an Animal is but specific attributes 
// that only a Cat has.
type Cat struct {
   Animal
   ClimbFactor int
}

// Speak knows how to speak like a cat.
func (c *Cat) Speak() {
   fmt.Printf("Meow! My name is %s, it is %t Iam a mammal with a climb factor of %d.\n", c.Name, c.IsMammal, c.ClimbFactor,)
}

func main() {
   // Create a list of animals that know how to speak.
   animals := []Animal{
           // Create a Dog by initializing its Animal parts
           // and then its specific Dog attributes.
           Dog{
              Animal: Animal{
                      Name: "Fido",
                      IsMammal: true,
              },
              PackFactor: 5,
           },
           // Create a Cat by initializing its animal parts
           // and then its specific Cat attributes.
           Cat{
              Animal: Animal{
                      Name: "Milo",
                      IsMammal: true,
              },
              ClimbFactor: 4,
           },
   }
   
   // Have the animals Speak
   for _, animal := range animals {
          animal.Speak()
   }  
}
