package _sum_ii

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	n := len(nums1)
	cnt := 0

	hash := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			hash[nums1[i]+nums2[j]]++
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val := nums3[i] + nums4[j]
			if c, ok := hash[-val]; ok {
				cnt += c
			}
		}
	}

	return cnt
}
