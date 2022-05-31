package sort

func Insertionsort(nums []int) {
	i := 1
	for i < len(nums) {
		x := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > x {
			nums[j + 1] = nums[j]
			j--
		}
		nums[j+1] = x
		i++
	}
}
