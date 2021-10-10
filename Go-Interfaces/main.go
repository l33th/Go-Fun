package main
 
import (
    "fmt"
)

type Person interface{
    greet() string
}
 
type Human struct{
    Name string
}
 
func (h *Human)greet() string {
    return "Hi, I am " + h.Name 
}
 
func isAPerson(h Person) {
    fmt.Println(h.greet())
}
 
func main() {
    var i interface{}
    var h = Human{"Glarbouered"}
    fmt.Println(i) 
    fmt.Println(h.greet())

    isAPerson(&h)
}
