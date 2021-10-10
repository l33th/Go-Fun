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

type B struct{
    s string
}

func checkType(i interface{}) {
    switch i.(type) {
    case int:
        fmt.Println("Int")
    case string:
        fmt.Println("String")
    default:
        fmt.Println("Other")
    }
}

func isEqual(i interface{}, j interface{}) {
    if(i == j) {
        fmt.Println("Equal")
    } else {
        fmt.Println("Inequal")
    }
}

func f(i interface{}) {
    fmt.Printf("%T\n", i)
}
 
func main() {
    var i interface{}
    fmt.Println(i)

    var h = Human{"Glarbouered"}
    fmt.Println(h.greet())

    isAPerson(&h)

    var b = Bird{"Chirper"}
    fmt.Println(b.fly())
    fmt.Println(b.walk())

    var bi interface{} = B{"a sample string"}
    fmt.Println(bi.(B))

    var si interface{} = "A string"
     
    checkType(si)

    var ii interface{}
    var ji interface{}
     
    isEqual(ii, ji)   // Equal
     
    var ai interface{} = "A string"
    var bbi interface{} = "A string"
     
    isEqual(ai, bbi)   // Equal

    var ac interface{} = "a string"
    var cc int = 42
     
    f(ac)   // string
    f(cc)   // int
}
