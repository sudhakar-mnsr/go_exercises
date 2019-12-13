// Sample program to show how to use the WithDeadline function.

package main

import (
   "context"
   "fmt"
   "time"
)

type data struct {
   UserID string
}

func main() {
   // set a deadline
   deadline := time.Now().Add(150 * time.Millisecond)
   
   // Create a context that is both manually cancellable and will signal
   // a cancel at the specified date/time
   ctx, cancel := context.WithDeadline(context.Background(), deadline)
   defer cancel()
   
   // Create a channel to recieved a signal that work is done.
   ch := make(chan data, 1)
   
   // Ask the goroutine to do some work for us
   go func() {
      // simulate work.
      time.Sleep(200 * time.Millisecond)
      
      // Report the work is done.
      ch <- data{"123"}
   }()
   
   // Wait for the work to finish. If it takes too long moveon.
   select {
   case d := <-ch:
      fmt.Println("work complete", d)
   case <-ctx.Done():
      fmt.Println("work cancelled")
   }
}

