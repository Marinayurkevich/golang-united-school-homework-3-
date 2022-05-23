package homework

func sortMapValues(input map[int]string) (result []string) {
	result := make([]string, 0, len(input))
	for k, v := range input {
	result =append(result, v)
	}
	sort.Strings(result)
	return result
}
