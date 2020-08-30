package src

func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		num = num & (num - 1) //将最后为1的位设置为0
		res++
	}
	return res
}
