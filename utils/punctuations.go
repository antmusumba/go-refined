package utils
func Punctuation(s []string) []string {
	punList := []string{",", ".", "!", "?", ":", ";"}
	// cow .goat sheep
	for i, word := range s {
		for _, pun := range punList {
			if string(word[0]) == pun {
				s[i-1] += pun
				s[i] = word[1:]
			}
		}
	}
	// cow goat sheep .
	for  i ,word := range s{
		for _, pun := range punList {

		}
	}
	// cow ... goat
	for i, word := range s {
		for _, punc := range punList {
			if word[0] == punc && len(word)-1 == punc && word != len(s)-1 {
				s[i-1] += word
				s = append(s[:i],s[i+1:]...)
			}
		}
	}
	// apostrophe
	coun := 0
	for i, word := range s {
		if word == "'" && coun == 0 {
			s[i+1] = word + s[i+1]
			s = append(s[:i],s[i+1:]...)

		}
	}
	for i , word := range s {
		if word == "'" && coun != 0 {
			s[i-1] += word
			s = append(s[:i],s[i+1:]...)
		}
	}

	return s
}