package main
import (
   "fmt"
   "unsafe"
)

type example struct {
   flag bool
   counter int64
   flag2 bool
   counter2 int64
   flag3 bool
   counter3 int64
   pi float32
}

func main() {
   // Declare a variable of type example set to its zero value
   var e1 example
   // Display the value
   fmt.Printf("%+v\n", e1)
   // Declare a variable of type example and init using a struct literal
   e2 := example{
      flag:  true,
      counter: 10,
      flag2:  true,
      counter2: 10,
      flag3:  true,
      counter3: 10,
      pi: 3.141592,
   }
   
   // Display the field values.
   fmt.Println("Flag", e2.flag)
   fmt.Println("Counter", e2.counter)
   fmt.Println("Pi",e2.pi)
   fmt.Println("size of first struct e1", unsafe.Sizeof(e1))
   fmt.Println("size of second struct e2", unsafe.Sizeof(e2))
}
