package main

import (
	inputreader "adventOfCode2019/datautils"
	"fmt"
	"strconv"
	"strings"
)

type requiredPositions struct {
	firstPosition  int
	secondPosition int
}

type passwordInfo struct {
	letter   string
	required requiredPositions
	password []string
	valid    bool
}

func main() {
	inputList := inputreader.ReadStringFile("../data/day02b.txt", "\n")
	validCount := countValidPasswords(validatePasswords(collatePasswords(inputList)))
	fmt.Println("validCount", validCount)
}

func countValidPasswords(passwordList []passwordInfo) int {
	count := 0
	for _, passwordInfo := range passwordList {
		if passwordInfo.valid {
			count++
		}
	}
	return count
}

func validatePasswords(passwordList []passwordInfo) []passwordInfo {
	validatedPasswords := []passwordInfo{}
	for _, pInfo := range passwordList {
		validatePassword := validatePassword(pInfo)
		validatedPasswords = append(validatedPasswords, validatePassword)
	}
	return validatedPasswords
}

func validatePassword(pInfo passwordInfo) passwordInfo {
	usage := 0
	if len(pInfo.password) < pInfo.required.secondPosition {
		return pInfo
	}
	passPosition1 := pInfo.password[pInfo.required.firstPosition-1]
	passPosition2 := pInfo.password[pInfo.required.secondPosition-1]
	if passPosition1 == pInfo.letter {
		usage++
	}
	if passPosition2 == pInfo.letter {
		usage++
	}
	if usage == 1 {
		pInfo.valid = true
	}
	return pInfo
}

func collatePasswords(inputList []string) []passwordInfo {
	passwordInfoList := []passwordInfo{}
	for _, input := range inputList {
		passwordInfo := parsePassword(input)
		passwordInfoList = append(passwordInfoList, passwordInfo)
	}
	return passwordInfoList
}

func parsePassword(password string) passwordInfo {
	splitPassword := strings.Split(password, " ")
	unsantizedLetterPositions, unsantizedLetter, pword := splitPassword[0], splitPassword[1], splitPassword[2]
	sanitzedLetterSlice := strings.Split(unsantizedLetterPositions, "-")
	sanitizedLetter := strings.ReplaceAll(unsantizedLetter, ":", "")
	firstPos, _ := strconv.Atoi(sanitzedLetterSlice[0])
	secondPos, _ := strconv.Atoi(sanitzedLetterSlice[1])
	pInfo := passwordInfo{
		letter: sanitizedLetter,
		required: requiredPositions{
			firstPosition:  firstPos,
			secondPosition: secondPos,
		},
		password: strings.Split(pword, ""),
	}
	return pInfo
}
