package homework

func sortMapValues(input map[int]string) (result []string) {
	for k := range input{
        result = append(result, k)
    }
    sort.Strings(result)
  
    for _, k := range result {
    result=append(result, input[k])
    }
	return result
}

