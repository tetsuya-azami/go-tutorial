package main

import "fmt"

type Vertex struct {
	X, Y int
	str  string
}

func main() {
	// printPointer()
	// checkNew()
	// checkStruct()
	structPointer()
}

func printPointer() {
	var n int = 100
	fmt.Println(n)
	fmt.Println(&n) // 0x1400000e0f8(ポインタの表示)(表示される値は毎回変わる)

	var p *int = &n // pはポインタ変数と呼ばれる
	fmt.Println(p)  // 0x1400000e0f8(ポインタの表示)
	fmt.Println(*p) // 100(ポインタが指す値)
	*p = 200        // ポインタが指す値を変更
	fmt.Println(*p) // 200
}

func checkNew() {
	var p *int = new(int)
	fmt.Println(p)
	fmt.Println(*p) // 0(初期値)
	var p2 *int
	fmt.Println(p2)
	fmt.Println(*p2) // panic: runtime error: invalid memory address or nil pointer dereference
}

func checkStruct() {
	v := Vertex{1, 2, "test"}
	fmt.Println(v)
	fmt.Println(v.X)
	fmt.Println(v.Y)
	v.X = 100
	fmt.Println(v)
	v2 := Vertex{Y: 1}
	fmt.Println(v2)

	v4 := Vertex{}
	fmt.Printf("%T %+v\n", v4, v4) // main.Vertex {X:0 Y:0 str:}
	var v5 Vertex
	fmt.Printf("%T %+v\n", v5, v5) // main.Vertex {X:0 Y:0 str:}
	// newと&はポインタになる
	v6 := new(Vertex)
	fmt.Printf("%T %+v\n", v6, v6) // *main.Vertex &{X:0 Y:0 str:}
	v7 := &Vertex{}
	fmt.Printf("%T %+v\n", v7, v7) // *main.Vertex &{X:0 Y:0 str:}
}

func structPointer() {
	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println("structPointer", v)
}

func changeVertex(v Vertex) {
	v.X = 1000
	fmt.Println("changeVertex: ", v)
}
