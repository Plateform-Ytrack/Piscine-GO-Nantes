package student

func BasicJoin(strs []string) string {
	temp := ""
	for i := range strs {
		temp += strs[i]
	}
	return temp
}