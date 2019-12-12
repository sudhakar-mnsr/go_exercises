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
   fanOutSem()
}   

// fanOutSem: You are a manager and you hire one new employee for the exact
// amount of work you have to get done. Each new employee knows immediately
// what they are expected to do and starts their work. However, you dont
// want all the employees working at once. You want to limit how many of 
// them are working at any given time. You sit waiting for all the results 
// of the employees work. The amount of time you wait on the employees is
// unknown because you need a guarantee that all the results sent by employees
// are received by you. No given employee needs an immediate guarantee
// that you received their result.
func fanOutSem() {
   emps := 2000
   ch := make(chan string, emps)
   g := runtime.NumCPU()
   sem := make(chan bool, g)

   for e := 0; e < emps; e++ {
      go func(emp int) {
         sem <- true
         {
            time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
            ch <- "paper"
            fmt.Println("employee: sent signal :", emp)
         }
         <-sem
      }(e)
   }
   for emps > 0 {
      p := <-ch
      emps--
      fmt.Println(p)
      fmt.Println("manager : recvd signal :", emps)
   }
   time.Sleep(time.Second)
   fmt.Println("-----------------------------------------------------")
}
