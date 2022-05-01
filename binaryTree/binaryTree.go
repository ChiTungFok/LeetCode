package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func main() {
	rc := bufio.NewReader(os.Stdin)
	parent, err := rc.ReadString('\n')
	if err != nil {
		return
	}
	p := BuildTree(strings.Split(parent, " "))
	child, err := rc.ReadString('\n')
	if err != nil {
		return
	}
	c := BuildTree(strings.Split(child, " "))
	fmt.Println(isSubStructure(p, c))
}

func BuildTree(vals []string) *TreeNode {
	var nodes []*TreeNode
	for _, val := range vals {
		v, _ := strconv.ParseInt(val, 10, 64)
		nodes = append(nodes, &TreeNode{
			Val: int(v),
		})
	}
	for i := 0; i*2+2 < len(nodes); i++ {
		nodes[i].Left = nodes[i*2+1]
		nodes[i].Right = nodes[i*2+2]
	}
	fmt.Println(nodes)
	if len(nodes) > 0 {
		return nodes[0]
	}
	return nil
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if A.Val == B.Val {
		if res := preOrderWay(A, B); res {
			return res
		}
	}
	if A.Left != nil {
		if res := isSubStructure(A.Left, B); res {
			return res
		}
	}
	if A.Right != nil {
		if res := isSubStructure(A.Right, B); res {
			return res
		}
	}
	return false
}

func preOrderWay(A *TreeNode, B *TreeNode) bool {
	var lr bool = true
	var rr bool = true
	if A.Left != nil && B.Left != nil {
		lr = preOrderWay(A.Left, B.Left)
	}
	if A.Left == nil && B.Left != nil {
		lr = false
	}
	if A.Right != nil && B.Right != nil {
		rr = preOrderWay(A.Right, B.Right)
	}
	if A.Right == nil && B.Right != nil {
		rr = false
	}
	return A.Val == B.Val && lr && rr
}
