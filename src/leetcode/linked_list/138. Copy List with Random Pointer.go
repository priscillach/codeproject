package linked_list

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

func copyRandomList(head *RandomNode) *RandomNode {
	old2New := make(map[*RandomNode]*RandomNode)
	cur := head
	newHead := &RandomNode{}
	cur2 := newHead
	for cur != nil {
		newNode := &RandomNode{
			Val: cur.Val,
		}
		cur2.Next = newNode
		old2New[cur] = newNode
		cur2 = cur2.Next
		cur = cur.Next

	}
	cur = head
	for cur != nil {
		if cur.Random != nil {
			old2New[cur].Random = old2New[cur.Random]
		}
		cur = cur.Next
	}
	return newHead.Next
}

func copyRandomListV2(head *RandomNode) *RandomNode {
	if head == nil {
		return nil
	}

	cur := head
	for cur != nil {
		newNode := &RandomNode{Val: cur.Val, Next: cur.Next}
		cur.Next = newNode
		cur = newNode.Next
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	oldHead := head
	newHead := head.Next
	curOld := oldHead
	curNew := newHead

	for curOld != nil {
		curOld.Next = curOld.Next.Next
		if curNew.Next != nil {
			curNew.Next = curNew.Next.Next
		} else {
			curNew.Next = nil
		}
		curOld = curOld.Next
		curNew = curNew.Next
	}

	return newHead
}
