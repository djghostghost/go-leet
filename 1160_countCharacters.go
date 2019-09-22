package main

func countCharacters(words []string, chars string) int {

	dest := makeCharacterTable(chars)

	result := 0

	for _, word := range words {
		src := makeCharacterTable(word)
		if compareTable(src, dest) {
			result += len(word)
		}
	}
	return result

}

func makeCharacterTable(word string) []int {

	result := make([]int, 26)
	for _, ch := range word {
		result[ch-'a']++
	}
	return result
}

func compareTable(source, dest []int) bool {

	for i := 0; i < 26; i++ {

		if source[i] > dest[i] {
			return false
		}
	}
	return true

}
