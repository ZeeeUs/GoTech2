package main

import (
	"fmt"
	"sort"
	"strings"
)

// sortCharsInWord сортирует символы в слове для построения корректной map
func sortCharsInWord(s string) string {
	s = strings.ToLower(s)
	container  := strings.Split(s, "")
	sort.Strings(container)
	return strings.Join(container, "")
}

// builtSet формирует слайс только из уникальных значений
func builtSet(bucket []string) []string{
	// Реализовали set
	set := make(map[string]bool)
	var setToSlice []string
	for _, v := range bucket{
		v = strings.ToLower(v)
		set[v] = true
	}
	for i := range set{
		setToSlice = append(setToSlice, i)
	}
	sort.Strings(setToSlice)
	return setToSlice
}

// Функция printAnagrams предназначена для красивой печати
func printAnagrams(bucket map[string][]string) {
	for k, v := range bucket{
		fmt.Printf("%v : %v\n", k, v)
	}
}

// finedAnagram является главной функцией, которая ищет во входном массиве анаграммы
func finedAnagram(arr *[]string) {
	bucket := map[string][]string{}
	for _, v := range *arr{
		normal := sortCharsInWord(v)
		bucket[normal] = append(bucket[normal], v)
	}

	for k, v := range bucket {
		bucket[k] = builtSet(v)
	}
	printAnagrams(bucket)
}

func main(){
	words := []string{"пятка", "пятак", "тяпка", "листок", "слиток", "столик", "ЛИСТОК", "ТяПкА"}
	finedAnagram(&words)
}