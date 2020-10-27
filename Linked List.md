#### 1. 单链表  
结构体:  
```go
type Node struct {
	value int
	next *Node
}
```  

列表循环:  
```go
func ShowListNode(node *Node){
	for node != nil{
		fmt.Println("node:", node)
		node = node.next
	}
}
```  

头插法:  
```go
	tail1 := &Node{value:0}
	head1 := tail1
	for i := 1; i < 10; i++{
		node := &Node{value:i}
		node.next = head1
		head1 = node
	}
	ShowListNode(head1)
```  

尾插法:  
```go
	head2 := &Node{value:0}
	tail2 := head2
	for i := 1; i < 10; i++  {
		node := &Node{value:i}
		tail2.next = node
		tail2 = node
	}
	ShowListNode(head2)
```  

