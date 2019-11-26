package main

import (
   "fmt"
   "runtime"
   "strconv"
   "unsafe"
)

type T struct {
   B uint8 // is a byte
   I int //it is int depends on arch
   P *int //it depends on arch
   S string
   SS []string
}

var p = fmt.Println

func memUsage(m1, m2 *runtime.MemStats) {
   p("Alloc:", m2.Alloc-m1.Alloc,
     "TotalAlloc:", m2.TotalAlloc-m1.TotalAlloc,
     "HeapAlloc:", m2.HeapAlloc-m1.HeapAlloc)
}
func main() {
   const PtrSize = 32 << uintptr(^uintptr(0)>>63)
   p("PtrSize=", PtrSize)
   p("IntSize=", strconv.IntSize)
   
   var m1, m2, m3, m4, m5, m6 runtime.MemStats
   runtime.ReadMemStats(&m1)
   t := T{}
   runtime.ReadMemStats(&m2)
   p("sizeof(uint8)", unsafe.Sizeof(t.B),
     "offset=", unsafe.Offsetof(t.B))
   p("sizeof(int)", unsafe.Sizeof(t.I),
     "offset=", unsafe.Offsetof(t.I))
   p("sizeof(*int)", unsafe.Sizeof(t.P),
     "offset=", unsafe.Offsetof(t.P))
   p("sizeof(string)", unsafe.Sizeof(t.S),
     "offset=", unsafe.Offsetof(t.S))
   p("sizeof([]string)", unsafe.Sizeof(t.SS),
     "offset=", unsafe.Offsetof(t.SS))
   p("sizeof(T)", unsafe.Sizeof(t))
   
   memUsage(&m1,&m2)
   
   runtime.ReadMemStats(&m3)
   t2 := "abc"
   runtime.ReadMemStats(&m4)
   
   memUsage(&m3,&m4)
   
   runtime.ReadMemStats(&m5)
   t3 := map[int]string{1: "x"}
   runtime.ReadMemStats(&m6)
   
   memUsage(&m5,&m6)
   
   fmt.Println(t2, t3) // prevent compiler error
}
