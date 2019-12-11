// Sample program to show how to create goroutines and
// how the scheduler behaves.
package main

import (
       "fmt"
       "runtime"
       "sync"
)

func init() {
   // Allocate one logical processor for the scheduler to use.
   runtime.GOMAXPROCS(1)
}

func main() {
// wg is used to manage concurrency.
var wg sync.WaitGroup
wg.Add(2)

fmt.Println("Start Goroutines")

// Create a goroutine from the lowercase function.
go func() {
   lowercase()
   wg.Done()
}()

// Wait for the goroutines to finish.
fmt.Println("Waiting to finish")
wg.Wait()

fmt.Println("\nTerminating Program")
}

// lowercase displays the set of lowercase letters three times.
func lowercase() {
   // Display the alphabet three times
   for count := 0; count < 3; count++ {
      for r := 'a'; r <= 'z'; r++ {
         fmt.Printf("%c ", r)
      }
   }
}

// Uppercase displays the set of uppercase letters three times.
func uppercase() {
   // Display the alphabet three times
   for count := 0; count < 3; count++ {
      for r := 'A'; r <= 'Z'; r++ {
         fmt.Printf("%c ", r)
      }
   }
}   
