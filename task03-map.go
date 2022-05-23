package homework

func sortMapValues(input map[int]string) (result []string) {
	for k := range input{
        keys = append(keys, k)
    }
    sort.Strings(keys)
  
    for _, k := range keys {
    result=append(result, input[k])
    }
	return result
}

