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

type Flyer interface{
    fly() string
}
 
type Walker interface{
    walk() string
}
 
type Bird struct{
    Name string
}
 
func (b *Bird)fly() string{
    return "Flying..."
}
 
func (b *Bird)walk() string{
    return "Walking..."
}
 
func main() {
    var i interface{}
    var h = Human{"Glarbouered"}

    fmt.Println(i) 
    fmt.Println(h.greet())

    isAPerson(&h)

    var b = Bird{"Chirper"}
     
    fmt.Println(b.fly())
    fmt.Println(b.walk())
}
