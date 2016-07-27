package main

import (
    "fmt"
    "time"
)
    func sleepwalker(n int) int{
	time.After(time.Second*time.Duration(n))
		return n
        }

	
func main() {
    var i int
	fmt.Println("Luke Sleepwalker likes to walk in his sleep. Enter 0 for stop sleeping or enter duration (seconds) that Luke will walk! #SleepIsWalkHype!")
	fmt.Scanf("%d", &i)
	for i != 0 {
	fmt.Println("Walking", sleepwalker(i), " seconds! Enter 0 for stop sleeping or enter sleep duration (seconds)")
	fmt.Scanf("%d", &i)
	}
}
