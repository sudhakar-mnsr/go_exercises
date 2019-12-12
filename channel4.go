// This sample program demonstrates the basic channel mechanics
// for goroutine signaling.
package main

import (
//       "context"
       "fmt"
       "math/rand"
       "runtime"
//       "sync"
       "time"
)

func init() {
   rand.Seed(time.Now().UnixNano())
}

func main() {
   pooling()
}   

// pooling: You are a manager and you hire a team of employees.
// None of the new employees know what they are expected to do 
// and wait for you to provide work. When work is provided to the 
// group, any given employee can take it and you dont care who
// it is. The amount of time you wait for any given employee to
// take your work is unknown because you need a guarantee that the 
// work you sending is recieved by and employee
func pooling() {
   ch := make(chan string)
   g := runtime.NumCPU()
   for e := 0; e < g; e++ {
      go func(emp int) {
         for p := range ch {
            fmt.Printf("employee %d : recv'd signal : %s\n", emp, p)
         }
      }(e)
   }
   const work = 100
   for w := 0; w < work; w++ {
      ch <- "paper"
      fmt.Println("manager : sent signal :", w)
   }
   close(ch)
   fmt.Println("manager : sent shutdown signal")
   time.Sleep(time.Second)
   fmt.Println("--------------------------------------------------")
}

