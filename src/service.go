package src

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func numberRandom(max, min int) int {
	return rand.Intn(max-min+1) + min
}

func generateChars(size int, vocal bool) string {
	size = If(size > 0, size, 1).(int)
	result := make([]byte, size)
	var alphabet = []string{"aeiouy", "bcdfghjkmnpqrstvwxz"}
	var statusVocal int
	for index := range result {
		statusVocal = If(vocal, 0, 1).(int)
		result[index] = alphabet[statusVocal][rand.Intn(len(alphabet[statusVocal]))]
	}
	return string(result)
}

func makeWord(setNumber bool) string {
	return fmt.Sprintf("%s%s%s%s%v",
		generateChars(1, false),
		generateChars(1, true),
		generateChars(2, false),
		generateChars(1, true),
		If(setNumber,
			strconv.Itoa(numberRandom(9, 1)),
			generateChars(1, false),
		).(string),
	)
}

func GeneratePass(groupWord int) string {
	groupWord = If(groupWord > 0, groupWord, 1).(int)
	result := make([]string, groupWord)
	//PrintSlice(result)
	positionNumber := numberRandom(groupWord-1, 0)
	for index := range result {
		result[index] = makeWord(positionNumber == index)
	}
	//PrintSlice(result)
	return strings.Join(result, "-")
}
