package main

import "fmt"

func swap(x,y *int) {
    z := *x
	*x = *y
	*y = z
}
func main() {
    x := 1 
	y := 2
	swap(&x, &y)
    fmt.Println(x,y)
}