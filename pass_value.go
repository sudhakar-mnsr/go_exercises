package main

func main() {
// Declare variable of type int with a value 10.
count := 10

// Display the value of and address of count
println("count:\tValue Of [", count, "]\t Addr Of [", &count, "]")

// Pass the value of the count
increment(count)

println("count:\tValue Of [", count, "]\t Addr Of [", &count, "]")
}

func increment(inc int) {
inc++
println("inc:\tValue Of [", inc, "]\tAddr Of [", &inc, "]")
}
