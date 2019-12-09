// Sample program to show how to understand method sets.
package main

import "fmt"

// notifier is an interface that defines notification
// type behavior.
type notifier interface {
   notify()
}

// user defines a user in the program.
type user struct {
   name string
   email string
}

// notify implements the notifier interface with a pointer receiver.
func (u *user) notify() {
   fmt.Printf("Sending User Email to %s,%s>\n", u.name, u.email)
}

func main() {
   // Create a value of type user and send a notification.
   u := user{"Bill", "bill@email.com"}
   
   // Values fo type user do not implement the interface because pointer
   // receivers dont belong to the method set of a value.
   
   // sendNotification(&u)
   sendNotification(u)
   // user does not implement notifier (notify method has pointer receiver)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
   n.notify()
}
