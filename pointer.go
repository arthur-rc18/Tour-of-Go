package main

import "fmt"

var p *int //When used to reference itself, the value will be 'nil'

func point() {
	i := 42         //i has the value of 42
	var x = &i      //The & operator generates a pointer to its operand
	fmt.Println(*x) //Read i through the pointer x
	fmt.Println(x)  //When printed alone, will show the memory address of i

	*p = 21 //Set i through the pointer p
	//This is known as "dereferecing" or "indirecting"

}
