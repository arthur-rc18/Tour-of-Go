package main

import "fmt"

type Vertex struct {
	sum  int
	sum2 int
}

var p *int //When used to reference itself, the value will be 'nil'

func main() {
	i := 42         //i has the value of 42
	var x = &i      //The & operator generates a pointer to its operand
	fmt.Println(*x) //Read i through the pointer x
	fmt.Println(x)  //When printed alone, will show the memory address of i

	//*p = 21 //Set i through the pointer p
	//This is known as "dereferecing" or "indirecting"

	v := Vertex{} //Struct fields can be accessed through a struct pointer
	p := &v       //p is setting to p, which is a variable of the struct Vertex
	p.sum = 1e9   //Here p is setting v's fields
	p.sum2 = 1748

	fmt.Println(v) //v will print, respectively, 1e9 in sum field and 1748 in sum2 field

}
