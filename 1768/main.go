package main

func main() {
	word := solutionOne("abc", "")
	println(word)
}

func solutionOne(word1 string, word2 string) string {
	word1Length := len(word1)
	word2Length := len(word2)

	if word1Length == 0 && word2Length == 0 {
		return ""
	}

	if word1Length == 0 {
		return word2
	}
	if word2Length == 0 {
		return word1
	}

	var newWord string

	for i := range word2 {
		newWord += string(word1[i])
		newWord += string(word2[i])

		if i == word1Length-1 {
			newWord += string(word2[i+1:])
			break
		}
		if i == word2Length-1 {
			newWord += string(word1[i+1:])
			break
		}
	}
	return newWord
}
