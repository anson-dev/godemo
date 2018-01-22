package main

/*
#include <stdio.h>
int test() {
    return 2016;
}
*/
import "C"
import "fmt"

func main() {
	    fmt.Println(C.test())
}
