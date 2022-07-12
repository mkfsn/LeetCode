package random_pick_index

type Solution struct {
	seed  int
	table map[int][]int
}

func Constructor(nums []int) Solution {
	s := Solution{seed: 0, table: make(map[int][]int)}
	for i, v := range nums {
		s.table[v] = append(s.table[v], i)
	}
	return s
}

func (this *Solution) Pick(target int) int {
	arr := this.table[target]
	return arr[this.getSeed()%len(arr)]
}

func (this *Solution) getSeed() int {
	this.seed++
	if this.seed > 2*10_000 {
		this.seed %= 2 * 10_000
	}
	return this.seed
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
