package link

import (
	"testing"
)

func TestLink(t *testing.T) {
	var head *Node
	stu1 := Node{Student{101, "xiaoming1"}, nil}
	stu2 := Node{Student{102, "xiaoming2"}, nil}
	stu3 := Node{Student{103, "xiaoming3"}, nil}
	stu4 := Node{Student{104, "xiaoming4"}, nil}

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
