// Sample program to show how what we are doing is not embedding
// a type but just using a type as a field

package main

import "fmt"

// user defines a user in the program.
type user struct {
   name string
   email string
}

// notify implements a method notifies users
// of different events.
func (u *user) notify() {
   fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

// admin represents an admin user with privileves
type admin struct {
   person user // NOT Embedding
   level string
}


func main() {
   // Create an admin user.
   ad := admin{
           person: user{
                   name: "John smith",
                   email: "John@yahoo.com",
           },
           level: "super",
   }      
   
   // We can access fields methods.
   ad.person.notify()
}
