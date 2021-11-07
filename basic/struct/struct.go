package main

import "fmt"

/*
Go语言只支持封装, 不支持继承和多态
	封装: private/public这样的访问权限
	继承: 继承父类, 子类可以使用父类的东西
	多态: 子类继承父类, 重写了父类的方法, 如果一个父类型的变量(实际是子对象), 调用重写还是会调到子类的方法

Go封装
	public: 首字母大写
	private: 首字母小写
	针对package来说的
*/

type Node struct {
	Value       int
	Left, Right *Node
}

/*
给结构体定义方法
	前面放上(node Node)就可以, 叫接受者, 相当于self/this
	可以直接被node.print()点出来
	(node Node)这个是值传递
	需要引用传递需要定义(node *Node)

	必须要放在一个包内, 但是可以是不同的文件

值接收者 VS 指针接收者
	要改变内容, 用指针
	结构过大, 用指针
	保持一致, 如果有指针了, 最后都是指针

*/

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node Node) SetValue(value int) {
	node.Value = value
}

/*
虽然是要求Node的指针, 但是实际上还是可以用node直接点
甚至node的指针也是可以直接点的
*/

func (node *Node) SetValuePoint(value int) {
	node.Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

/*
结构体可以用方法来创建, 这个时候会返回一个局部变量的地址给别人用, 在Go中这样的做法是允许的.

堆栈
	栈: 局部变量, 函数退出之后就会被回收
	堆: 全局变量等, 创建了之后根据语言不同可能需要手动释放, c++就需要手动回收

	Java的做法是用new的方式, 创建的变量都在堆上, 这样就会触发垃圾回收机制
	Go语言不需要顾及这些问题, 编译器会自动识别变量的作用, 自动判断是在堆上还是在栈上创建变量
*/
func createNode(value int) *Node {
	return &Node{Value: value}
}

func createExampleNode() *Node {
	return &Node{
		Value: 3,
		Left: &Node{
			Value: 0,
			Left:  nil,
			Right: &Node{Value: 2},
		},
		Right: &Node{
			Value: 5,
			Left:  &Node{Value: 0},
			Right: nil,
		},
	}
}

/*
结构体的创建有多种方式:
	类似变量的处理
	new()
*/
func basic() {
	var root Node
	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil}
	root.Left.Right = new(Node)
	root.Right.Left = createNode(2)

	nodes := []Node{
		{Value: 3},
		{},
		{6, nil, nil},
	}
	fmt.Println(root)  // {3 0xc0000a6018 0xc0000a6030}
	fmt.Println(nodes) // [{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> <nil>}]

	root.SetValue(5)
	root.Print() // 3
	root.SetValuePoint(5)
	root.Print() // 5

	pRoot := &root
	pRoot.SetValue(10)
	pRoot.Print() // 5
	pRoot.SetValuePoint(20)
	pRoot.Print() // 20
}

func main() {
	basic()

	root := createExampleNode()
	root.Traverse()
}
