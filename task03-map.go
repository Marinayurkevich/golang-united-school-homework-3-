package homework

func sortMapValues(input map[int]string) (result []string) {
	
	for k, v := range input {
	result =append(result, k)
	}
	sort.Ints(result)
	return result
}
