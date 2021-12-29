package alg

// 单调栈反向模板
// 仅保存数字（iterator point 为当前位置）
func nextGreaterElements1(nums []int) []int {
	n, res, stack := len(nums), make([]int, 0, len(nums)), make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
	}
	return res
}

// 单调栈反向模板
// post process 保存之前 element 的 index
func nextGreaterElements2(nums []int) []int {
	n, res, stack := len(nums), make([]int, 0, len(nums)), make([]int, 0)
	for i, _ := range stack {
		stack[i] = -1
	}
	for i := 0; i < n; i++ {
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i] {
			res[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
