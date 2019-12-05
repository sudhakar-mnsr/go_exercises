// Sample program to show that you cannot take the address 
// of a element of a map.
package main

// player represents someone playing our game
type player struct {
   name string
   score int
}

func main() {
   // Declare a map with initial values using a map literal.
   players := map[string]player{
      "anna": {"Anna", 42},
      "jacob": {"Jacob", 21},
   }
   
   // Trying to take the address of a map element fails.
   anna := &players["anna"]
   anna.score++
   
   //./map5.go:19:12: cannot take the address of players["anna"] 
   
   // Instead take the element, modify it, and put it back.
   player := players["anna"]
   player.score++
   players["anna"] = player
}  
