package main

import (
	"fmt"
	"strings"
)


func main() {
	words := []string{"Now", "is", "the", "winter", "of", "our", "discontent", "made", "glorious", "summer", "by", "this", "sun", "of", "York"}
	result := fullJustify(words, 20)
	for _, v := range result {
		fmt.Println(v)
	}

}

func fullJustify(words []string, maxWidth int) []string {
	paragraph := []string{}
	wordPos := 0
	var line []string
	for wordPos < len(words) {
		wordPos, line, _ = longestPhrase(words, maxWidth, wordPos)
		paragraph = append(paragraph, padd(line, maxWidth))
	}
	// remove the last line, and add it back un padded
	paragraph = paragraph[:len(paragraph)-1]
	paragraph = append(paragraph, unpadd(line, maxWidth))
	return paragraph
}

func longestPhrase(words []string, maxWidth, wordPos int) (int,
	[]string, bool) {

	// only one word in the line
	remaining := words[wordPos:]
	if len(remaining) == 1 {
		return wordPos + 1, remaining, true
	}

	line := []string{remaining[0]}
	prev := line

	count := 1
	for i := 1; i < len(remaining); i++ {
		line = append(line, remaining[i])
		if lineLength(line) > maxWidth {
			break
		}
		prev = line
		count++
	}
	return wordPos + count, prev, false
}

func lineLength(line []string) int {
	l := 0
	for _, w := range line {
		w := strings.Trim(w, " ")
		l += len(w) + 1
	}
	return l - 1
}

func unpadd(line []string, length int) string {
	for k, _ := range line {
		line[k] = strings.Trim(line[k], " ")
	}

	a := strings.Join(line, " ")
	size := len(a)

	for i := 0; i < length-size; i++ {
		a = a + " "
	}
	return a
}

func padd(line []string, length int) string {
	if len(line) == 1 {
		size := len(line[0])
		for i := 0; i < length-size; i++ {
			line[0] = line[0] + " "
		}
		return line[0]
	}
	spaces := len(line) - 1
	if spaces == 0 {
		return strings.Join(line, " ")
	}

	extraSpaces := length - lineLength(line)

	for i := 0; i < extraSpaces; i++ {
		pos := i % (spaces)
		line[pos] = line[pos] + " "
	}

	return strings.Join(line, " ")
}
