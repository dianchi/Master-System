package src

func StringArraryNoneEmpty(arr []string) []string {
	var final []string
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			final = append(final, arr[i])
		}
	}
	return final
}
