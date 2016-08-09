package main

import (
 "fmt"
 "strconv"
)
type Student struct {
 Name string
}
func (stu *Student) Present() {
 fmt.Println("I am " + stu.Name)
}
func (stu *Student) Leave() {
 fmt.Println(stu.Name + " Leaving")
}