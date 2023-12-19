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
)

type Game struct {
	ID       int
	fields   [][]Count
	possible bool
}

type Count struct {
	color string
	count int
}

func main() {
	input, err := os.Open("testinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	reader := bufio.NewReader(input)

	check := func(e error) bool {
		if e == io.EOF {
			return true
		} else if e != nil {
			log.Fatal(err)
			return true
		}
		return false
	}

	for {
		line, _, err := reader.ReadLine()
		if check(err) {
			break
		}
		str := bytes.NewBuffer(line).String()
		strSplit := strings.Split(str, ":")

		game := new(Game)

		game.ID, _ = strconv.Atoi(strings.TrimSpace(
			strings.Fields(strSplit[0])[1],
		))
		body := strings.TrimSpace(strSplit[1])

		for _, v := range strings.Split(body, ";") {
			str := strings.TrimSpace(v)
			var set []Count

			for _, v := range strings.Split(str, ",") {
				str := strings.TrimSpace(v)
				keys := strings.Fields(str)

				count := new(Count)
				count.count, _ = strconv.Atoi(keys[0])
				count.color = keys[1]

				set = append(set, *count)
			}

			game.fields = append(game.fields, set)
		}
		game.possible = true

		// --------------
		fmt.Printf("ID: %d\n", game.ID)
		for i, v := range game.fields {
			fmt.Printf(" Game %d\n", i+1)
			for setNr, set := range v {
				fmt.Printf("  %d: %d of >%s<\n", setNr, set.count, set.color)
			}
		}
		fmt.Println()
	}
}
