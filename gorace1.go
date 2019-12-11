// Sample program to show how to create race conditions in 
// our programs. We dont want to do this.

package main

import (
       "fmt"
       "sync"
)

// Counter is a variable incremented by all goroutines.
var counter int
 
func main() {

   // Number of goroutines to use.
   const grs = 2
   // wg is used to manage concurrency
   var wg sync.WaitGroup
   wg.Add(grs)
   
   // Create two goroutines.
   for g := 0; g < grs; g++ {
      go func() {
         for i := 0; i < 2; i++ {
            // Capture the value of counter.
            value := counter
   
            // Increment our local value of counter.
            value++
   
            // Store the value back into counter
            counter = value
         }
         wg.Done()
      }()
   }
   // Wait for the goroutines to finish.
   wg.Wait()
   fmt.Println("Final Counter:", counter)
}
