package main

import (  
    "fmt"
    "math"
    "unsafe"
)



func main() {  

  tutorial42();

}

func tutorial1() {
    fmt.Println("Hello World")

    var width, height int
    fmt.Println("width is", width, "height is", height)
    width = 100
    height = 50
    fmt.Println("new width is", width, "new height is ", height)
}

func tutorial2(){

    //  name, age := "naveen", 29 //short hand declaration
    var (
        name   = "naveen"
        age    = 29
        height int
    )
    fmt.Println("my name is", name, ", age is", age, "and height is", height)
}

func tutorial3(){

    x, y := 145.8, 543.8 
    z := math.Min(x, y)
    fmt.Println("minimum value is ", z)  

    a := true
    b := false
    fmt.Println("a:", a, "b:", b)
    c := a && b
    fmt.Println("c:", c)
    d := a || b
    fmt.Println("d:", d)
}

func tutorial41(){

    var a int = 89
    b := 95
    fmt.Println("value of a is", a, "and b is", b)
    fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) //type and size of a
    fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //type and size of b
}

func tutorial42(){
    a, b := 5.67, 8.97
    fmt.Printf("type of a %T b %T\n", a, b)
    sum := a + b
    diff := a - b
    fmt.Println("sum", sum, "diff", diff)

    no1, no2 := 56, 89
    fmt.Println("sum", no1+no2, "diff", no1-no2)

    c1 := complex(5, 7)
    c2 := 8 + 27i
    cadd := c1 + c2
    fmt.Println("sum:", cadd)
    cmul := c1 * c2
    fmt.Println("product:", cmul)

    first := "Naveen"
    last := "Ramanathan"
    name := first +" "+ last
    fmt.Println("My name is",name)

}

func tutorial43(){
    i := 55      //int
    j := 67.8    //float64
    sum := i + int(j) //j is converted to int
    fmt.Println(sum)
}