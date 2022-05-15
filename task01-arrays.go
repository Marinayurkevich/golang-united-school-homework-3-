package homework

func average(input [15]float32) (result float32) {
	var sum float32 = 0
	for _, input := range input{
	sum+=input}
	N:=float32 (len(input))	
	return sum/N
}
