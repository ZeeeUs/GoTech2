package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortCharsInWords(t *testing.T) {
	table := []struct {
		in  string
		out string
	}{
		{
			"дгвба",
			"абвгд",
		},
		{
			"привет",
			"веипрт",
		},
		{
			"Hello",
			"ehllo",
		},
	}

	for _, v := range table {
		result := SortCharsInWord(v.in)
		assert.Equal(t, result, v.out)
	}
}

func TestBuiltSet(t *testing.T) {
	table := []struct {
		in  []string
		out []string
	}{
		{
			[]string{"Hello", "World", "hello.html", "WoRlD"},
			[]string{"hello.html", "world"},
		},
		{
			[]string{"One", "two", "thrEe", "onE", "TWO", "three"},
			[]string{"one", "three", "two"},
		},
		{
			[]string{"Bred", "banana", "tomato", "tomato", "BANANA", "Ban", "BanaN", "TOMATO"},
			[]string{"ban", "banan", "banana", "bred", "tomato"},
		},
	}

	for _, v := range table {
		result := BuiltSet(v.in)
		assert.Equal(t, result, v.out)
	}
}

func TestFinedAnagram(t *testing.T) {
	table := []struct {
		in  *[]string
		out map[string][]string
	}{
		{
			&[]string{"пятка", "пятак", "тяпка", "листок", "слиток", "столик", "ЛИСТОК", "ТяПкА"},
			map[string][]string{
				"пятка":  []string{"пятак", "пятка", "тяпка"},
				"листок": []string{"листок", "слиток", "столик"},
			},
		},
		{
			&[]string{"импортер", "пирометр", "реимпорт", "каретник", "катерник", "кретинка"},
			map[string][]string{
				"импортер" : []string{"импортер", "пирометр", "реимпорт"},
				"каретник" : []string{"каретник", "катерник", "кретинка"},
			},
		},
		{
			&[]string{"шрам", "лоза", "зало", "Зола", "лОза", "марш", "шарм", "ШрАм"},
			map[string][]string{
				"лоза" : []string{"зало", "зола", "лоза"},
				"шрам" : []string{"марш", "шарм", "шрам"},
			},
		},
	}

	for _, v := range table {
		result := FinedAnagram(v.in)
		assert.Equal(t, result, v.out)
	}
}
