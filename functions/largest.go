package main

import "fmt"
func largest(args ...int) int {
max := args[0]
for _, v := range args {
if v > max {
max = v
} else {}
}
return max
}

func main() {
fmt.Println("Largest number:", largest(48,96,86,68,57,82,63,70,37,34,83,27,19,97, 9,17,))
}