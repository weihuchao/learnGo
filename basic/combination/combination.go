package main

import (
	"fmt"
	"imooc.com/weihc/learngo/tree"
)

func basic() {
	node := tree.CreateExampleNode()
	node.Traverse()
}

/*
扩展接口体的方法:
	1. 组合, 参见myNode
	2. 别名, 参见queue
	3. 内嵌(Embedding), 参见
*/
type myNode struct {
	node *tree.Node
}

func (node *myNode) postOrder() {
	contentNode := node.node
	if node == nil || contentNode == nil {
		return
	}

	left := myNode{contentNode.Left}
	left.postOrder()

	right := myNode{contentNode.Right}
	right.postOrder()

	contentNode.Print()
}

/*
内嵌实际上是一个语法糖, 相当于把tree.Node的内容全部都平铺到了myEmbeddingNode里头
要想调用这个tree.Node的时候, 可以使用名字Node(也就是tree.Node的实际内容的名字)

但是这不是继承, 它只是一个语法糖, myEmbeddingNode和tree.Node是两个不同的类型

另外, myEmbeddingNode可以类似继承可以重载tree.Node的一些函数如Traverse()
*/
type myEmbeddingNode struct {
	*tree.Node
}

func (node *myEmbeddingNode) preOrder() {
	if node == nil || node.Node == nil {
		return
	}

	node.Print()

	left := myNode{node.Left}
	left.postOrder()

	right := myNode{node.Right}
	right.postOrder()
}

func (node *myEmbeddingNode) Traverse() {
	fmt.Println("\nThis is myEmbeddingNode Traverse.\nGet tree.Node.Traverse():")
	node.Node.Traverse()
}

func main() {
	basic()
	// 0 2 3 0 5

	fmt.Println()
	myRoot := myNode{tree.CreateExampleNode()}
	myRoot.postOrder()
	// 2 0 0 5 3

	fmt.Println()
	myNewRoot := myEmbeddingNode{tree.CreateExampleNode()}
	myNewRoot.preOrder()
	myNewRoot.Traverse()
}
