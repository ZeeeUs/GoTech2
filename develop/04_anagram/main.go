package main

import (
	"fmt"
	"sort"
	"strings"
)

// SortCharsInWord сортирует символы в слове для построения корректной map
func SortCharsInWord(s string) string {
	s = strings.ToLower(s)
	container := strings.Split(s, "")
	sort.Strings(container)
	return strings.Join(container, "")
}

// BuiltSet формирует слайс только из уникальных значений
func BuiltSet(bucket []string) []string {
	// Реализовали set
	set := make(map[string]bool)
	var setToSlice []string
	for _, v := range bucket {
		v = strings.ToLower(v)
		set[v] = true
	}
	for i := range set {
		setToSlice = append(setToSlice, i)
	}
	sort.Strings(setToSlice)
	return setToSlice
}

// Функция printAnagrams предназначена для красивой печати
func printAnagrams(bucket map[string][]string) {
	for k, v := range bucket {
		fmt.Printf("%v : %v\n", k, v)
	}
}

// FinedAnagram является главной функцией, которая ищет во входном массиве анаграммы
func FinedAnagram(arr *[]string) map[string][]string {
	bucket := make(map[string][]string, 0)
	keys := make(map[string][]int, 0)
	res := make(map[string][]string, 0)
	buck := make([]string, 0)

	for i, v := range *arr {
		normal := SortCharsInWord(v)
		keys[normal] = append(keys[normal], i)
		bucket[normal] = append(bucket[normal], v)
	}

	for k, _ := range bucket {
		buck = append(buck, k)
	}

	for _, v := range buck {
		min := keys[v][0]
		res[(*arr)[min]] = BuiltSet(bucket[v])
	}

	return res
}

func main() {
	words := []string{"пятка", "пятак", "тяпка", "листок", "слиток", "столик", "ЛИСТОК", "ТяПкА"}
	anagrams := FinedAnagram(&words)
	printAnagrams(anagrams)
}
