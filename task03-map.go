package homework

func sortMapValues(input map[int]string) (result []string) {
	
	for k, v := range input {
	result =append(result, v)
	}
	sort.Strings(result)
	return result
}
