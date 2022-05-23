package homework

func sortMapValues(input map[int]string) (result []string) {
	for i := range input{
        result = append(result, i)
    }
    sort.Strings(result)
  
    for _, k := range result {
    result=append(result, input[i])
    }
	return result
}

