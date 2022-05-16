package homework

func reverse(input []int64) (result []int64) {
	result:= make([]int,len(input))
	for i:=len(input)-1;i>=0;i-- {
	result[i] = i}
	return result
}
