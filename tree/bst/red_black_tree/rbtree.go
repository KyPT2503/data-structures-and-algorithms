package red_black_tree

import (
	"log"
)

const (
	red   = 0
	black = 1
)

func main() {
	log.Println("Start ...")

	tree := &RBTree{}
	for val := 0; val < 100; val++ {
		tree.Insert(val)
		tree.ShowInf()
	}
	log.Println(tree.Flatten())
}

type node struct {
	val   int
	col   int // 0 (red) and 1(black)
	left  *node
	right *node
}

type RBTree struct {
	root *node
}

func (tree *RBTree) Insert(val int) {
	if tree.root == nil {
		tree.root = &node{val: val, col: black}
		return
	}
	insert(tree.root, val)
}

func insert(n *node, val int) *node {
	if n == nil {
		return &node{val: val} // color is 1 (red)
	}
	if val < n.val {
		n.left = insert(n.left, val)
		if n.left.col == red {
			if n.left.left.col == red {
				if n.right.col == black {
					x := n.left
					n.left = x.left
				}
			} else if n.left.right.col == red {

			}
		}
	} else {
		n.right = insert(n.right, val)
		if n.right.col == red {
			// todo
		}
	}
	return n
}

func (t *RBTree) Delete(val int) {

}

// ============================
// extensions
// ============================

func (t *RBTree) Flatten() []int {
	return flatten(t.root)
}
func flatten(n *node) []int {
	if n == nil {
		return nil
	}
	return append(append(flatten(n.left), n.val), flatten(n.right)...)
}

func (t *RBTree) Get(n int) bool {
	panic("unimplemented")
}

func (t *RBTree) ShowInf() {
	inf := checkHeight(t.root)
	log.Printf("- height: %d, number of nodes: %d\n", inf[1], inf[0])
}
func checkHeight(root *node) []int {
	if root == nil {
		return []int{0, 0}
	}
	var height int
	queue := []*node{root}
	var off int
	for off < len(queue) {
		l := len(queue)
		for off < l {
			n := queue[off]
			if n.left != nil {
				queue = append(queue, n.left)
			}
			if n.right != nil {
				queue = append(queue, n.right)
			}
			off++
		}
		height++
	}
	return []int{len(queue), height}
}
