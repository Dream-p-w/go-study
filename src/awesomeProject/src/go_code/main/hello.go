package main

import "fmt"



func add(a int,b int) int {
	var sum int
	sum=a+b
	return sum
}

func main()  {
	var i int
	i = add(10, 20)
	fmt.Print("你好")
	fmt.Print(i)
}

