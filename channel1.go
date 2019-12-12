// This sample program demonstrates the basic channel mechanics
// for goroutine signaling.
package main

import (
//       "context"
       "fmt"
       "math/rand"
//       "runtime"
//       "sync"
       "time"
)

func init() {
   rand.Seed(time.Now().UnixNano())
}

func main() {
   waitForResult()
}   

// waitForResult: You are a manager and you hire a new employee
// employee knows immediately what they are expected to do and starts
// their work. You sit waiting for the result of the employee's work. The 
// amount of time you wait on the employee is unknown because you need a 
// guarantee that the result sent by the employee is received by you.
func waitForResult() {
   ch := make(chan string)
   go func() {
      time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
      ch <- "paper"
      fmt.Println("employee : sent signal")
   }()

   p := <-ch
   fmt.Println("manager : recv'd signal :", p)

   time.Sleep(time.Second)
   fmt.Println("--------------------------------------------------")
}
