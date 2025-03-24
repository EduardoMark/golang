package kata

func Maps(x []int) []int {
	newArr := []int{}

	for _, num := range x {
		newArr = append(newArr, num*2)
	}

	return newArr
}
