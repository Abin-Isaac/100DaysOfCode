package main

func main() {
	age := 29      // age is int
	age = "naveen" // error since we are trying to assign a string to a variable of type int
// Normal way 
	var z bool = true
// short hand way 
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
// Using && and || logic operator
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)
	fmt.Println(z)
}