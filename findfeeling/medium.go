package findfeeling

import (
	"container/list"
	"findfeeling/model"
)

type TreeNode = model.TreeNode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}

	flag := false
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		tmp := []int{}
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, (node.Left))
			}
			if node.Right != nil {
				queue = append(queue, (node.Right))
			}
			tmp = append(tmp, node.Val)
		}
		if flag {
			for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
				tmp[i], tmp[j] = tmp[j], tmp[i]
			}
		}
		flag = !flag
		if len(tmp) != 0 {
			res = append(res, tmp)
		}
	}
	return res
}

func postOrderTraversalBT(root *TreeNode) []int {

	if root == nil {
		return nil
	}
	res := []int{}
	stack := make([]*TreeNode, 0)
	cur := root
	// stack = append(stack, cur)
	// postOrderTraversalBT root left right
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = (cur.Right)
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = node.Left
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func postOrderTraversalBT2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	cur := root
	stack := list.New()
	stack.PushFront(cur)
	res := []int{}
	for stack.Len() != 0 {
		element := stack.Front()
		stack.Remove(element)
		node := element.Value.(*TreeNode)
		res = append([]int{node.Val}, res...)
		if node.Left != nil {
			stack.PushFront(node.Left)
		}
		if node.Right != nil {
			stack.PushFront(node.Right)
		}
	}
	return res
}

func Quicksort(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	quicksortHelper(nums, 0, len(nums)-1)
	return nums

}
func quicksortHelper(nums []int, l, r int) {
	if l > r {
		return
	}
	// k := rand.Intn(r-l) + l
	// nums[k], nums[r] = nums[r], nums[k]
	p := paratition(nums, l, r)
	quicksortHelper(nums, l, p[0]-1)
	quicksortHelper(nums, p[1]+1, r)
}

// 分割
func paratition(nums []int, l, r int) []int {
	less, more := l-1, r
	for l < more {
		if nums[l] < nums[r] {
			l++
			nums[l], nums[r] = nums[r], nums[l]
			r++
		} else if nums[l] > nums[r] {
			more--
			nums[more], nums[r] = nums[r], nums[more]
		} else {
			r++
		}
	}
	nums[r], nums[more] = nums[more], nums[r]
	return []int{less + 1, more}
}

func Heapsort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		heapsortInsert(nums, i)
	}
	size := len(nums)
	for size > 0 {
		size--
		nums[0], nums[size] = nums[size], nums[0]
		heapsortIfy(nums, 0, size)
	}
	return nums
}

func heapsortIfy(nums []int, index, size int) {
	left, right := 2*index+1, 2*index+2
	for left < size {
		largest := left
		if right < size && nums[right] > nums[left] {
			largest = right
		}
		if nums[largest] < nums[index] {
			break
		}
		nums[largest], nums[index] = nums[index], nums[largest]
		index = largest
		left, right = 2*index+1, 2*index+2
	}
}

func heapsortInsert(nums []int, index int) {
	parent := (index - 1) >> 1
	for parent >= 0 && nums[index] > nums[parent] {
		nums[index], nums[parent] = nums[parent], nums[index]
		index = parent
		parent = (index - 1) >> 1
	}
}

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mergeHelper(nums, 0, len(nums)-1)
	return nums
}

func mergeHelper(nums []int, l, r int) {
	if l > r {
		return
	}
	mid := l + ((r - l) >> 1)
	mergeHelper(nums, l, mid)
	mergeHelper(nums, mid+1, r)
	merge(nums, l, mid, r)

	return
}

func merge(nums []int, l, mid, r int) {
	p1, p2, helper := l, mid+1, make([]int, 0, r-l+1)
	var index int
	for p1 <= mid && p2 <= r {
		if nums[p1] <= nums[p2] {
			helper[index] = nums[p1]
		} else {
			helper[index] = nums[p2]
		}
		index++
	}
	helper = append(helper, nums[p1:mid+1]...)
	helper = append(helper, nums[p2:r+1]...)
	copy(nums[l:r+1], helper[:])
}
