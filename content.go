package main

import (
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func getContent() []Game {
	mega := getData().Body
	defer mega.Close()

	var content []string

	z := html.NewTokenizer(mega)
	var table [][]string
	var row []string

	for z.Token().Data != "html" {
		tt := z.Next()
		if tt == html.StartTagToken {
			t := z.Token()

			if t.Data == "tr" {
				if len(row) > 0 {
					table = append(table, row)
					row = []string{}
				}
			}

			if t.Data == "td" {
				inner := z.Next()

				if inner == html.TextToken {
					text := string(z.Text())
					t := strings.TrimSpace(text)
					row = append(row, t)
				}
			}

		}
	}
	if len(row) > 0 {
		table = append(table, row)
	}
	for row, data := range table {
		for i := range data {
			if i >= 3 && i <= 8 {
				content = append(content, table[row][i])
			}
		}
	}

	ints := make([]int, len(content))

	for i, s := range content {
		ints[i], _ = strconv.Atoi(s)
	}

	var n int
	var Games []Game

	for i := 1; i <= 60; i++ {
		for _, c := range ints {
			if c == i {
				n++
			}
		}
		Game := Game{
			Ten:  i,
			Freq: n,
			Perc: float32(n) / float32(len(ints)) * 100,
		}
		Games = append(Games, Game)
		n = 0
	}

	sort.SliceStable(Games, func(i, j int) bool {
		return Games[i].Perc > Games[j].Perc
	})

	return Games
}
