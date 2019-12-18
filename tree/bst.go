package main

import(
    "fmt"
)

type Node struct{
    key int
    value int
    left *Node
    right *Node
    parent *Node
}

type Tree struct {
    head *Node
    n int
}

func main() {
/*    t := NewTree()
    t.Add(30, 30)
    t.Add(14,14)
    t.Add(16,16)
    t.Add(25,25)
    t.Add(34,34)
    t.Add(6,6)
    t.Add(19,19)
    t.Add(78,78)
    t.Add(45,45)
    t.Add(58,58)
    t.Add(85,85)
    fmt.Println("---------")
    t.transPreOrder()
    fmt.Println("---------")
    t.transInOrder()
    fmt.Println("---------")
    t.transPostOrder()
*/

    pre := []int{30,14,6,16,25,19,34,78,45,58,85}
    in := []int{6,14,16,19,25,30,34,45,58,78,85}
    post := []int{6,19,25,16,14,58,45,85,78,34,30}
    
//    t := TreeByPreAndInOrder(pre, in)    
    t := TreeByPreAndInOrder(pre, in)
    t.transPreOrder()
    t.transInOrder()
    t.transInOrder()

    t = TreeByPostAndInOrder(post, in)
    t.transPreOrder()
    t.transInOrder()
    t.transPostOrder()
}

func NewTree()*Tree {
    return &Tree{
        n: 0,
    }
}

func (t *Tree) Add(key, value int) {
    add := &Node{
        key: key,
        value: value,
    }
    if t.head == nil {
        t.head = add
        return
    }
    t.head.AddNode(add) 
}

func (n *Node) AddNode(add *Node) {
    if add.key<n.key {
        if n.left == nil {
            n.left = add
            add.parent = n
        } else {
            n.left.AddNode(add)
        }
    } else if add.key>n.key {
        if n.right == nil {
            n.right = add
            add.parent = n
        } else {
            n.right.AddNode(add)
        }
    } else {
        n.value = add.value
    }
    return
}


func (t *Tree) Delete(key int) {
    if t.head == nil {
        return    
    } 
    node := t.head
    var isLeft bool
    for node!= nil && node.key != key {
        if key < node.key {
            node = node.left
            isLeft = true
        } else {
            node = node.right
            isLeft = false
        }
    }    
    
    if node == nil {
        return
    }

    if node.right == nil {
        if node.parent == nil {
            t.head = node.left
        } else {
            if isLeft {
                node.parent.left = node.left
            } else {
                node.parent.right = node.left
            } 
        }   
    } else {
        if node.left != nil {
            tmpNode := node.right
            for tmpNode.left != nil {
                tmpNode = tmpNode.left
            }
            tmpNode.left = node.left
        }
    
        if node.parent == nil {
            t.head = node.right
        } else {
            if isLeft { 
                node.parent.left = node.right
            } else {
                node.parent.right = node.right
            }
        } 
    }
}

func (t *Tree) transPreOrder() {
    if t.head == nil {
        return
    }
    t.head.transPreOrder()
}

func (t *Tree) transInOrder() {
    if t.head == nil {
        return
    }
    t.head.transInOrder()
}

func (t * Tree) transPostOrder() {
    if t.head == nil {
        return
    }
    t.head.transPostOrder()
}

func (n *Node) transPreOrder() {
    fmt.Println(n.key)
    if n.left != nil {
        n.left.transPreOrder()
    } 
    if n.right != nil{
        n.right.transPreOrder()
    }
}

func (n *Node) transInOrder() {
    if n.left != nil {
        n.left.transInOrder()
    }
    fmt.Println(n.key)
    if n.right != nil {
        n.right.transInOrder()    
    }    
}

func (n *Node) transPostOrder() {
    if n.left != nil {
        n.left.transPostOrder()
    }
    if n.right != nil {
        n.right.transPostOrder()
    }
    fmt.Println(n.key)
}

func (n *Node) PreOrder(pre []int, in []int, preIndex *int) {
    key := pre[*preIndex]
    index := search(in, key)
    if index == -1 {
       panic("in error")
    }
    left := in[0:index]
    if len(left) > 0 {
        *preIndex++
        n.left = &Node{
            key: pre[*preIndex],
            value:pre[*preIndex],
            parent:n,
        }
        n.left.PreOrder(pre, left, preIndex)
    }
    right := in[index+1:]
    if len(right) > 0 {
        *preIndex++
        n.right = &Node{
            key: pre[*preIndex],
            value:pre[*preIndex],
            parent:n,
        }    
        n.right.PreOrder(pre, right, preIndex)
    }
    return
} 

func (n *Node)PostOrder(post, in []int, postIndex *int) {
    key := post[*postIndex]
    index := search(in, key)
    if index == -1 {
       panic("in error")
    }
    right := in[index+1:]
    if len(right) > 0 {
        *postIndex--
        n.right = &Node{
            key: post[*postIndex],
            value:post[*postIndex],
            parent:n,
        }    
        n.right.PostOrder(post, right, postIndex)
    }

    left := in[0:index]
    if len(left) > 0 {
        *postIndex--
        n.left = &Node{
            key: post[*postIndex],
            value:post[*postIndex],
            parent:n,
        }
        n.left.PostOrder(post, left, postIndex)
    }
    return

}

func TreeByPreAndInOrder(pre, in []int) *Tree{
    preIndex := 0
    
    tree := NewTree()
    root := &Node{
        key: pre[preIndex],
        value: pre[preIndex],
    }    
    tree.head = root
    tree.head.PreOrder(pre, in, &preIndex)
    return tree
}

func TreeByPostAndInOrder(post, in []int) *Tree {
    postIndex := len(post) - 1
    tree := NewTree()
    root := &Node{
        key: post[postIndex],
        value: post[postIndex],
    }
    tree.head = root
    tree.head.PostOrder(post, in, &postIndex)
    return tree
}

func search(nodes []int, key int) int{
    for index, val := range nodes {
        if key == val {
            return index
        }    
    }    
    return -1
}

