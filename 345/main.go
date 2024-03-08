package main

func reverseVowels(s string) string {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	chars := []rune(s)
	i, j := 0, len(chars)-1
	for i < j {
		if !contains(vowels, chars[i]) {
			i++
			continue
		}
		if !contains(vowels, chars[j]) {
			j--
			continue
		}
		chars[i], chars[j] = chars[j], chars[i]
		i++
		j--
	}
	return string(chars)
}

func contains(s []rune, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}

func main() {
	s := "hello"
	reversed := reverseVowels(s)
	println(reversed) // "holle"
}
