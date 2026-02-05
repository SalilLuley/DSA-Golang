package leetcode_75

func MaxSubarray() {
	arr := []int{-1, -2, -3, -4}
	println(maxSubArray(arr))
}

func maxSubArray(A []int) int {
	maxi := A[0]
	x := A[0]

	for i := 1; i < len(A); i++ {
		x = max(x+A[i], A[i])
		maxi = max(x, maxi)
	}
	return maxi

}
