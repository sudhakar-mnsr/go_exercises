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
   waitForTask()
}

// waitForTask: You are a manager and you hire a new employee.
// Your new employee doesnt know immediately what they are expected
// to do and waits for you to tell them what to do. You prepare the work 
// and send it to them. The work your sending is received by the employee
func waitForTask() {
   ch := make(chan string)
   go func() {
      p := <-ch
      fmt.Println("employee : recv'd signal :", p)
   }()
   time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
   ch <- "paper"
   fmt.Println("manager : sent signal")

   time.Sleep(time.Second)
   fmt.Println("-------------------------------------------------")
}

