package advanced

import (
	"fmt"
	"math"
)

// 定义常量π
const pi = math.Pi

// Shape 定义了一个形状接口，包含计算面积和周长的方法
// 该接口用于统一不同几何形状的计算规范
type Shape interface {

	// Area 方法于计算形状的面积
	Area()

	// Perimeter 方法用于计算形状的周长
	Perimeter()
}

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
// 矩形结构体
type Rectangle struct {
	// width 宽度
	Width float64
	// height 高度
	Height float64
}

// 矩形实现接口计算面积
func (r *Rectangle) Area() {
	area := r.Width * r.Height
	fmt.Println("矩形的面积是：", area)
}

// 矩形实现接口计算周长
func (r *Rectangle) Perimeter() {
	perimeter := 2 * (r.Width + r.Height)
	fmt.Println("矩形的周长是：", perimeter)
}

// 定义圆形结构体
type Circle struct {
	// radius 半径
	Radius float64
}

func (c *Circle) Area() {
	area := pi * c.Radius * c.Radius
	fmt.Println("圆形的面积是：", area)
}
func (c *Circle) Perimeter() {
	perimeter := 2 * pi * c.Radius
	fmt.Println("圆形的周长是：", perimeter)
}
