package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	check := func(e error) bool {
		if e == io.EOF {
			return true
		} else if e != nil {
			log.Fatal(err)
			return true
		}
		return false
	}

	reader := bufio.NewReader(input)
	var calibrationValues []int

	for {
		line, _, err := reader.ReadLine()
		if check(err) {
			break
		}

		lineReader := bytes.NewReader(line)
		buf := new(strings.Builder)
		for {
			char, _, err := lineReader.ReadRune()
			if check(err) {
				break
			}
			if unicode.IsNumber(char) {
				buf.WriteRune(char)
			}
		}

		digits := buf.String()
		number := string(digits[0]) + string(digits[buf.Len()-1])
		if numberInt, err := strconv.Atoi(number); !check(err) {
			fmt.Println(number)
			calibrationValues = append(calibrationValues, numberInt)
		}
	}

	calibrationSum := 0
	for _, number := range calibrationValues {
		calibrationSum += number
	}
	fmt.Println("-----")
	fmt.Println(calibrationSum)
}
