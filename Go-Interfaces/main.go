package main
 
import (
    "fmt"
    "math"
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

/* define an interface */
type Shape interface {
    area() float64
}

/* define a circle */
type Circle struct {
    x,y,radius float64
}

/* define a rectangle */
type Rectangle struct {
    width, height float64
}

/* define a method for circle (implementation of Shape.area())*/
func(circle Circle) area() float64 {
    return math.Pi * circle.radius * circle.radius
}

/* define a method for rectangle (implementation of Shape.area())*/
func(rect Rectangle) area() float64 {
    return rect.width * rect.height
}

/* define a method for shape */
func getArea(shape Shape) float64 {
    return shape.area()
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

circle := Circle{x:0,y:0,radius:5}
   rectangle := Rectangle {width:10, height:5}
   
   fmt.Printf("Circle area: %f\n",getArea(circle))
   fmt.Printf("Rectangle area: %f\n",getArea(rectangle))
}
