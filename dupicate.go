package main

func duplicate() []int {
	nums := []int{1, 2, 3, 4, 5, 2, 1, 3}
	result := []int{}
	for _, num := range nums {
		exists := false
		for _, val := range result {
			if val == num {
				exists = true
				break
			}
		}
		if !exists {
			result = append(result, num)
		}
	}
	return result
}
