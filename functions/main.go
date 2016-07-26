package main

import "fmt"

func sum(xs ...int) int{    
    total := 0
    for _, v := range xs {
        total += v
    }
    return total
}
	
func evenCheck(in int) bool{
if (in%2==0) {
return true
}else {
return false
}
}

func main() {
xs := []int{1,2,3}
var in int = sum(xs ...)
fmt.Println(evenCheck(in))
}
