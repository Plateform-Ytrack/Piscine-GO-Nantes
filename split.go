package student

func Split(str, charset string) []string {
	length := 0
	for i := range str {
		length = i + 1

	}
	len_split := 0
	for i := range charset {
		len_split = i + 1
	}
	for i := 0; i < len_split; i++ {
		str += " "
	}
	prev := false
	len := 0
	for i := 0; i < length; i++ {
		if (str[i:i+len_split] == charset) && !prev {
			prev = true
			len++
		} else {
			prev = false
		}
	}
	len++

	arr := make([]string, len)

	ss := ""
	count := 0
	for i := 0; i < length; i++ {
		if str[i:i+len_split] == charset {
			l := 0
			for i := range ss {
				l = i + 1
			}
			if l == 0 {
				continue
			}
			arr[count] = ss
			count++
			ss = ""
			i = i + len_split - 1
			continue
		}
		ss += string(str[i])
	}

	l := 0
	for i := range ss {
		l = i + 1
	}
	if l != 0 {
		arr[count] = ss
	}
	return arr
}