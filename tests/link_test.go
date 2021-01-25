package tests

import (
	"testing"

	"github.com/hqd8080/go-algorithms/algorithms"
)

func TestLink(t *testing.T) {
	var head *algorithms.Node
	stu1 := algorithms.Node{algorithms.Student{101, "aa"}, nil}
	stu2 := algorithms.Node{algorithms.Student{102, "bb"}, nil}
	stu3 := algorithms.Node{algorithms.Student{103, "cc"}, nil}
	stu4 := algorithms.Node{algorithms.Student{104, "dd"}, nil}

	// 创建新链表
	head = head.Create()
	head = stu1.Insert(head)
	head = stu2.Insert(head)
	head = stu3.Insert(head)
	head = stu4.Insert(head)

	// 输出链表
	head.PrintLink()

	// 删除节点
	head = stu3.Delete(head)
	head.PrintLink()
}
