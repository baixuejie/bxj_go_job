package main

import (
	"bxj_go_job/advanced"
	_ "bxj_go_job/advanced"
	"bxj_go_job/basis"
	"fmt"
)

func job1Test() {
	values := []int{4, 1, 2, 1, 2}
	result := basis.OnlyOneInt(values)
	fmt.Println(result)
}

func job2Test() {
	values := "([)]"
	result := basis.IsValid(values)
	fmt.Println(result)
}

func main() {
	//job1Test()
	//job2Test()
	//result := basis.Job3([]int{1, 9, 3, 9})
	//fmt.Println("result is :", result)

	//basis.Job4()

	//job7
	/*rectangle := advanced.Rectangle{Width: 5, Height: 3}
	circle := advanced.Circle{Radius: 4}
	rectangle.Area()
	rectangle.Perimeter()
	circle.Area()
	circle.Perimeter()*/
	//job8
	/*employee := advanced.Employee{Person: advanced.Person{Name: "张三", Age: 18}, EmployeeID: 1001}
	advanced.PrintInfo(&employee)*/

	//advanced.Job9()
	//job10
	/*param := 1
	advanced.Job10(&param)
	fmt.Println("param is :", param)*/
	//job11
	/*a := 1
	b := 2
	c := 3
	params := []*int{&a, &b, &c}
	job.Job11(params)
	for i := 0; i < len(params); i++ {
		fmt.Println("params[i] is :", *params[i])
	}*/
	//advanced.Job12()
	//advanced.Job13()
	//advanced.Job14()
	advanced.Job16()
}
