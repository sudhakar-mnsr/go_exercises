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
   waitForFinished()
}

// waitForFinished: Think about being a manager and hiring a new
// employee. In this scenario, you want your new employee to perform a 
// taskimmediately when they are hired, and you need to wait for the result
// of their work. You need to wait because you cant move on until you know 
// they are but you dont need anything from them.
func waitForFinished() {
   ch := make(chan struct{})

   go func() {
      time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
      close(ch)
      fmt.Println("employee : sent signal")
   }()

   _, wd := <-ch
   fmt.Println("manager : recv'd signal :", wd)
   time.Sleep(time.Second)
   fmt.Println("---------------------------------------------------")
}
