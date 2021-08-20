package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type inputFlags struct {
	column      int
	sortByNum   bool
	reverseSort bool
	unique      bool
	filename    []string
}

func parsArguments() inputFlags {
	columnNumber := flag.Int("k", 1, "Number of column for sort")
	num := flag.Bool("n", true, "Sort by number")
	reverse := flag.Bool("r", false, "Reverse sort")
	unique := flag.Bool("u", false, "Don't show repeat string")
	flag.Parse()

	flags := inputFlags{
		column:      *columnNumber,
		sortByNum:   *num,
		reverseSort: *reverse,
		unique:      *unique,
		filename:    flag.Args(),
	}

	return flags
}

func parseFile(inputData string) ([]string, error) {
	var data []string
	file, err := os.Open(inputData)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	return data, nil
}

func readFromFile(inputData []string) ([]string, error) {
	var data []string
	for _, v := range inputData {
		docData, err := parseFile(v)
		if err != nil {
			return []string{}, err
		}
		data = append(data, docData...)
	}

	return data, nil
}

func toUniqueize(data []string) []string {
	set := make(map[string]bool)
	var setToSlice []string
	for _, v := range data {
		set[v] = true
	}

	for i := range set {
		setToSlice = append(setToSlice, i)
	}

	return setToSlice
}

// Sort - главная функция, которая читает ключи и производит сортировку
func Sort(data []string, flags inputFlags) []string {
	if flags.unique {
		data = toUniqueize(data)
	}

	compareAsNumbers := func(i, j string) bool {
		lnum, lerr := strconv.Atoi(i)
		rnum, rerr := strconv.Atoi(j)

		if lerr != nil && rerr != nil {
			return i < j
		}
		if lerr != nil || rerr != nil {
			return lerr == nil
		}
		return lnum < rnum
	}
	compareAsStrings := func(lhs, rhs string) bool {
		return lhs < rhs
	}

	var valueComparator func(string, string) bool
	if flags.sortByNum {
		valueComparator = compareAsNumbers
	} else {
		valueComparator = compareAsStrings
	}

	compareLogic := func(i, j int) bool {
		lhs := strings.Split(data[i], " ")
		rhs := strings.Split(data[j], " ")
		if len(lhs) == 0 {
			return true
		}
		if len(rhs) == 0 {
			return false
		}

		if len(lhs) < flags.column && len(rhs) >= flags.column {
			return true
		}
		if len(lhs) >= flags.column && len(rhs) < flags.column {
			return false
		}

		if len(lhs) < flags.column && len(rhs) < flags.column {
			return valueComparator(lhs[0], rhs[0])
		}
		if len(lhs) >= flags.column && len(rhs) >= flags.column {
			return valueComparator(lhs[flags.column-1], rhs[flags.column-1])
		}
		panic("DEBUG: code should not run here")
	}

	if flags.reverseSort {
		sort.Slice(data, func(i, j int) bool {
			return !compareLogic(i, j)
		})
	} else {
		sort.Slice(data, compareLogic)
	}

	return data
}

func main() {
	par := parsArguments()
	data, err := readFromFile(par.filename)
	if err != nil {
		log.Fatalln(err)
	}

	data = Sort(data, par)

	for _, line := range data {
		fmt.Println(line)
	}
}
