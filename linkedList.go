package linkedList

type Object interface {}

type Node struct {
	Data Object // 定义数据域
	Next *Node  // 定义地址域(指向下一个表的地址)
}

type List struct {
	headNode *Node  // 头节点
}

// 1. 判断单链表是否为空
func (this *List) IsEmpty() bool {
	if this.headNode == nil { // 判断单链表是否为空,只需要判断头节点是否为空即可
		return true
	} else {
		return false
	}
}

// 2.单链表的长度
func (this *List) Length() int{
	// 获取链表头节点
	cur := this.headNode
	// 定义一个计数器, 初始值为0
	count := 0
	for cur != nil{
		// 如果头节点不为空, 则count++
		count ++
		// 对地址进行逐个位移
		cur = cur.Next
	}
	return count
}
// 3.从头部添加元素
func (this *List) Add(data *Object) *Node{
	node := &Node{Data:data}
	head := this.headNode
	node.Next = head
	this.headNode = node
	return  node
}
// 4.从尾部添加元素
func (this *List) Append(data *Object){
	// 创建一个新元素,通过传入data参数进行数据域的赋值
	node := &Node{Data:data}
	if this.IsEmpty(){ // 如果该链表为空, 那么直接将新节点插到头部
		this.headNode = node
	} else { //此链表不为空
		cur := this.headNode // 定义变量用于存储头节点
		for cur.Next != nil{ // 判断是否为尾节点, 如果为nil, 则为尾节点
			cur = cur.Next // 链表进行位移, 直到cur获取到尾节点
		}
		cur.Next = node // 此时cur为尾节点,将其地址指向创建的新节点
	}
}
// 5.在指定位置添加元素
func (this *List) Insert(index int, data *Object) {
	// 在指定位置添加元素, 此处的index指下标, this.headNode的index = 0
	if index < 0 { // 如果index<0, 则进行头部插入
		this.Add(data)
	} else if index > this.Length() {
		this.Append(data)
	} else {
		pre := this.headNode
		count := 0
		for count < (index - 1) { // 用于控制位移的链表数目
			pre = pre.Next
			count++
		}
		// 当循环退出后,pre指向index-1的位置
		node := &Node{Data: data} // 创建新的链表元素
		node.Next = pre.Next      //将新的链表元素地址指向其上一个链表地址域存储的地址,也就是新元素现在的跟随元素
		pre.Next = node           // 上一个链表地址域存储的地址指向新元素地址
	}
}
// 6.删除链表指定值的元素
func (this *List) Remove(data *Object){
	pre := this.headNode
	if pre.Data == data{
		this.headNode = pre.Next
	} else {
		for pre.Next != nil{
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

// 7.删除指定位置的元素
func (this *List) RemoveAtIndex(index int)  {
	pre := this.headNode
	if index <= 0{
		this.headNode = pre.Next
	} else if index > this.Length() {
		return
	} else {
		count := 0
		for count != (index-1) && pre.Next != nil {
			count ++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}
// 8.查看是否包含某个元素
func (this *List) Contain(data *Object) bool {
	cur := this.headNode
	for cur != nil{
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}
// 9.遍历所有元素
func (this *List) ShowList(){
	if !this.IsEmpty() {
		cur := this.headNode
		for {
			if cur.Next != nil{
				cur = cur.Next
			} else {
				break
			}
		}
	}
}