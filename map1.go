package main

import "fmt"

// user represents someone using the program.
type user struct {
   name string
   surname string
}

func main() {
   // Declare and make a map that store values
   // of type user with a key of type string.
   users := make(map[string]user)

   // Add key/value pairs to the map
   users["Roy"] = user{"Rob", "Roy"}
   users["Ford"] = user{"Henry", "Ford"}
   users["Mouse"] = user{"Mickey", "Mouse"}
   users["Jackson"] = user{"Michael", "Jackson"}

   // Read the value at a specific key.
   mouse := users["Mouse"]
   
   fmt.Printf("%+v\n", mouse)

   // Replace the value at the Mouse key.
   users["Mouse"] = user{"Jerry", "Mouse"}

   // Read the Mouse key again.
   fmt.Printf("%+v\n", users["Mouse"])

   // Delete the value at a specific key.
   delete(users, "Roy")
   
   // Check the length of the map. There are only 3 elements.
   fmt.Println(len(users))

   //it is safe to delete an absent key.
   delete(users, "Roy")
   
   // Iterate over map and notice the results are different(order)
   for key := range users {
     fmt.Println(key)
   } 

   fmt.Println("Goodbye.")
} 
