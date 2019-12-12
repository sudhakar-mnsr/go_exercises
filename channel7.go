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
   drop()
}   

// drop: You are a manager and you hire a new employee. Your new employee
// doesnt know immediately what they are expected to do and waits for
// you to tell them what to do. You prepare the work and send it to them.
// The amount of time they wait is unknown because you need a guarantee that 
// the work your sending is received by the employee. You wont wait for the 
// employee to take the work if they are not ready to recieve it. In that
// case you drop the work on the floor and try again with the next piece of
// work.
func drop() {
   const cap = 100
   ch := make(chan string, cap)
   
   go func() {
      for p := range ch {
         fmt.Println("employee : recvd signal :", p)
      }
   }()
   const work = 2000
   for w := 0; w < work; w++ {
      select {
      case ch <- "paper":
         fmt.Println("manager : sent signal :", w)
      default:
         fmt.Println("manager : dropped data:", w)   
      }
   }
   close(ch)
   fmt.Println("manager : sent shutdown signal")
   time.Sleep(time.Second)
   fmt.Println("------------------------------------------------")
}
