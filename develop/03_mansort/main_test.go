package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseFile(t *testing.T) {
	type args struct {
		filename string
	}
	table := []struct {
		testName string
		args     args
		want     []string
		wantErr  bool
	}{
		{
			testName: "File exist",
			args:     args{filename: "testingFile.txt"},
			want:     []string{"hello.html", "world", "tomato 3", "bread 2", "cucumber 1", "mars 12", "hello.html"},
		},
		{
			testName: "File doesn't exist",
			args:     args{filename: ""},
			want:     []string{},
			wantErr:  true,
		},
	}

	for _, val := range table {
		result, err := ParseFile(val.args.filename)
		if (err != nil) != val.wantErr {
			t.Errorf("Couldn't read the file!")
			return
		}
		assert.Equal(t, result, val.want)
	}
}

func TestToUniqueize(t *testing.T) {
	table := []struct {
		data []string
		want []string
	}{
		{
			data: []string{"hello.html", "world", "hello.html", "world", "hello.html", "world"},
			want: []string{"hello.html", "world"},
		},
		{
			data: []string{"Russia", "Canada", "Brazil", "USA", "China", "Russia", "USA"},
			want: []string{"Brazil", "Canada", "China", "Russia", "USA"},
		},
		{
			data: []string{"spotify 169", "spotify 85", "appleMusic 199", "spotify 85", "appleMusic 199", "YouTube 299"},
			want: []string{"YouTube 299", "appleMusic 199", "spotify 169", "spotify 85"},
		},
	}

	for _, val := range table {
		result := ToUniqueize(val.data)
		assert.Equal(t, result, val.want)
	}
}

func TestSort(t *testing.T) {
	defaultData := []string{
		"hello.html",
		"world",
		"tomato 3",
		"bread 2",
		"cucumber 1",
		"mars 12",
		"hello.html",
	}
	type args struct {
		data  []string
		flags inputFlags
	}
	table := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Simple sort",
			args: args{data: defaultData, flags: inputFlags{
				column:      1,
				sortByNum:   false,
				reverseSort: false,
				unique:      false,
			}},
			want: []string{"bread 2", "cucumber 1", "hello.html", "hello.html", "mars 12", "tomato 3", "world"},
		},
		{
			name: "Reverse sort",
			args: args{data:defaultData, flags: inputFlags{
				column: 1,
				sortByNum: false,
				reverseSort: true,
				unique: false,
			}},
			want: []string{"world", "tomato 3", "mars 12", "hello.html", "hello.html", "cucumber 1", "bread 2"},
		},
		{
			name: "Sort unique",
			args: args{data:defaultData, flags: inputFlags{
				column: 1,
				sortByNum: false,
				reverseSort: false,
				unique: true,
			}},
			want: []string{"bread 2", "cucumber 1", "hello.html", "mars 12", "tomato 3", "world"},
		},
		{
			name: "Sort by num",
			args: args{data:defaultData, flags: inputFlags{
				column: 2,
				sortByNum: true,
				reverseSort: false,
				unique: false,
			}},
			want: []string{"hello.html", "hello.html", "world", "cucumber 1", "bread 2", "tomato 3", "mars 12"},
		},
		{
			name: "Sort by num & unique",
			args: args{data:defaultData, flags: inputFlags{
				column: 2,
				sortByNum: true,
				reverseSort: false,
				unique: true,
			}},
			want: []string{
				"hello.html",
				"world",
				"cucumber 1",
				"bread 2",
				"tomato 3",
				"mars 12",
			},
		},
		{
			name: "Sort by num & unique",
			args: args{data:defaultData, flags: inputFlags{
				column: 2,
				sortByNum: true,
				reverseSort: true,
				unique: true,
			}},
			want: []string{
				"mars 12",
				"tomato 3",
				"bread 2",
				"cucumber 1",
				"world",
				"hello.html",
			},
		},
	}

	for _, val := range table {
		result := Sort(val.args.data, val.args.flags)
		assert.Equal(t, result, val.want)
	}
}
