package student

func StrRev(outtext string) string {

	Text := []byte(outtext)
	iindex := 0

	for range Text {
		iindex++
	}

	for index, word := range []byte(outtext) {
		Text[iindex-index-1] = word

	}

	return string(Text)
}