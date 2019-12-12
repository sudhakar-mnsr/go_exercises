// This sample program demonstrates the basic channel mechanics
// for goroutine signaling.
package main

import (
       "context"
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
   cancellation()
}   

// cancellation: You are a manager and you hire a new employee. Your
// new employee knows immediately what they are expected to do and
// starts their work. You sit waiting for the result of the employees
// work. The amount of time you wait on the employee is unknown because
// you need a guarantee that the result sent by the employee is received
// by you. Except you are not willing to wait forever for the employee
// to finish their work. They have a specified amount of time and if they
// are not done, you dont wait and walk away.
func cancellation() {
   duration := 150 * time.Millisecond
   ctx, cancel := context.WithTimeout(context.Background(), duration)
   defer cancel()
   
   ch := make(chan string, 1)
   
   go func() {
      time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
      ch <- "paper"
   }()
   
   select {
   case d := <-ch:
      fmt.Println("work complete", d)
   case <-ctx.Done():
      fmt.Println("work cancelled")
   }
   time.Sleep(time.Second)
   fmt.Println("------------------------------------------------")
}   

