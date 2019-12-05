// Sample program to show how maps are reference types
package main
 
import "fmt"

func main() {
   //Initialize a map with values.
   scores := map[string]int{
      "anna": 21,
      "jacob": 12,
   }
   
   // Pass the map to a function to perform some mutation
   double(scores, "anna")
   tripple(scores, "jacob")
   
   // See the change is visible in our map.
   fmt.Println("Score:", scores["anna"])
   fmt.Println("Score:", scores["jacob"])
}

// double finds the score for a specific player and 
// multiplies it by 2
func double(scores map[string]int, player string) {
   scores[player] = scores[player] * 2
}

// triple finds the score for a specific player and
// multiplies it by 3
func tripple(scores map[string]int, player string) {
   scores[player] = scores[player] * 3
}
