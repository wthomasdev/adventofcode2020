package main

import (
	"adventOfCode2020/inputreader"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputList := inputreader.ReadStringFile("../data/day04.txt", "\n\n")
	partTwo(inputList)
}

type validPassport map[string]int

func partOne(inputList []string) {
	validCount := 0
	for _, passport := range inputList {
		validPassport := newPassport()
		validPassport.validateFieldsBasic(passport)
		if validPassport.validatePassport() {
			validCount++
		}
	}
	fmt.Println("validCount", validCount)
}

func partTwo(inputList []string) {
	validCount := 0
	fmt.Println("Number of passports: ", len(inputList))
	for _, passport := range inputList {
		validPassport := newPassport()
		validPassport.validateFieldsAdvanced(passport)
		if validPassport.validatePassport() {
			validCount++
		}
	}
	fmt.Println("validCount", validCount)
}

func (vP validPassport) validateFieldsBasic(passport string) {
	for k := range vP {
		if strings.Contains(passport, k) {
			vP[k] = 1
		}
	}
}

func (vP validPassport) validateFieldsAdvanced(passport string) {
	passPortParts := []string{}
	passportNewLineSplit := strings.Split(passport, "\n")
	for _, v := range passportNewLineSplit {
		chunk := strings.Split(v, " ")
		passPortParts = append(passPortParts, chunk...)
	}
	for _, pPart := range passPortParts {
		value := strings.Split(pPart, ":")[1]
		for k := range vP {
			if strings.Contains(pPart, k) {
				switch k {
				case "byr":
					valNum, _ := strconv.Atoi(value)
					if len(value) == 4 && (valNum >= 1920 && valNum <= 2002) {
						vP[k] = 1
					}
				case "iyr":
					valNum, _ := strconv.Atoi(value)
					if len(value) == 4 && (valNum >= 2010 && valNum <= 2020) {
						vP[k] = 1
					}
				case "eyr":
					valNum, _ := strconv.Atoi(value)
					if len(value) == 4 && (valNum >= 2020 && valNum <= 2030) {
						vP[k] = 1
					}
				case "hgt":
					if strings.Contains(value, "cm") {
						num, _ := strconv.Atoi(strings.Split(value, "cm")[0])
						if num >= 150 && num <= 193 {
							vP[k] = 1
						}
					} else if strings.Contains(value, "in") {
						num, _ := strconv.Atoi(strings.Split(value, "in")[0])
						if num >= 59 && num <= 76 {
							vP[k] = 1
						}
					}
				case "hcl":
					if isHexcolor(value) {
						vP[k] = 1
					}

				case "ecl":
					validColors := "amb blu brn gry grn hzl oth"
					if strings.Contains(validColors, value) {
						vP[k] = 1
					}
				case "pid":
					if len(value) == 9 {
						vP[k] = 1
					}
				}
			}
		}
	}
}

func isHexcolor(str string) bool {
	hexColor := "^#([0-9a-fA-F]{6})$"
	rxHexcolor := regexp.MustCompile(hexColor)
	return rxHexcolor.MatchString(str)
}

func (vP validPassport) validatePassport() bool {
	valid := false
	validCount := 0
	for k := range vP {
		validCount += vP[k]
	}
	requiredCount := len(vP)
	if validCount == requiredCount {
		valid = true
	}
	return valid
}

func newPassport() validPassport {
	vP := validPassport{
		"byr": 0,
		"iyr": 0,
		"eyr": 0,
		"hgt": 0,
		"hcl": 0,
		"ecl": 0,
		"pid": 0,
	}
	return vP
}
