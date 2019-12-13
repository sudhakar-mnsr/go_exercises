// Sample program to show when a Context is canceled, All context
// derived from it are also cancelled

package main

import (
   "context"
   "fmt"
   "sync"
)

// Need a key type
type myKey int

// Need a key value
const key myKey = 0

func main() {
   // Create a context that can be cancelled.
   ctx, cancel := context.WithCancel(context.Background())
   defer cancel()
   
   // Use the Waitgroup for orchestration
   var wg sync.WaitGroup
   wg.Add(10)
   
   // Create ten goroutines that will derive a Context from
   // the one created above.
   for i := 0; i < 10; i++ {
      go func(id int) {
         defer wg.Done()
         // Derive a new context for this goroutine from the Context
         // owned by the main function.
         ctx := context.WithValue(ctx, key, id)
   
         // Wait until the context is cancelled.
         <-ctx.Done()
         fmt.Println("Cancelled:", id)
      }(i)
   }
   
   // Cancel the context and any derived contexts as well
   cancel()
   wg.Wait()
}
