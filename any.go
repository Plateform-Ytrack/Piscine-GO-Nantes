package student

func Any(f func(string) bool, arr []string) bool {

	for _,i:=range arr{
		if f(i){
			return true
		}
	}
	return false
}