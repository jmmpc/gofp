package list

import (
	"reflect"
	"slices"
	"strconv"
	"testing"
)

var (
	transformFunc = func(s string) int { return 1 }
	predicateFunc = func(n int) bool { return n%2 == 0 }
)

var mapTestTable = map[int]struct {
	input []string
	want  []int
}{
	1: {
		input: []string{"hello", "world"},
		want:  []int{1, 1},
	},
	2: {
		input: []string{},
		want:  []int(nil),
	},
	3: {
		input: nil,
		want:  []int(nil),
	},
}

func TestMap(t *testing.T) {
	for index, testCase := range mapTestTable {
		got := Map(testCase.input, transformFunc)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, testCase.want)
		}
	}
}

func TestMapSeq(t *testing.T) {
	for index, tc := range mapTestTable {
		got := slices.Collect(MapSeq(slices.Values(tc.input), transformFunc))
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

var filterTestTable = map[int]struct {
	input []int
	want  []int
}{
	1: {
		input: []int{1, 2, 3, 4},
		want:  []int{2, 4},
	},
	2: {
		input: []int{},
		want:  []int(nil),
	},
	3: {
		input: nil,
		want:  []int(nil),
	},
}

func TestFilter(t *testing.T) {
	for index, testCase := range filterTestTable {
		got := Filter(testCase.input, predicateFunc)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, testCase.want)
		}
	}
}

func TestFilterSeq(t *testing.T) {
	for index, testCase := range filterTestTable {
		items := slices.Values(testCase.input)
		filtered := FilterSeq(items, predicateFunc)
		got := slices.Collect(filtered)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, testCase.want)
		}
	}
}

func TestFilterMap(t *testing.T) {
	filterMapTestTable := map[int]struct {
		input []string
		want  []int
	}{
		1: {
			input: []string{"123", "a"},
			want:  []int{123},
		},
		2: {
			input: []string{"a"},
			want:  []int(nil),
		},
		3: {
			input: []string{},
			want:  []int(nil),
		},
		4: {
			input: nil,
			want:  []int(nil),
		},
	}

	for index, testCase := range filterMapTestTable {
		got := FilterMap(testCase.input, strconv.Atoi)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, testCase.want)
		}
	}
}

func TestEvery(t *testing.T) {
	everyTestTable := map[int]struct {
		input []int
		want  bool
	}{
		1: {
			input: []int{1, 2},
			want:  false,
		},
		2: {
			input: []int{2, 4},
			want:  true,
		},
		3: {
			input: nil,
			want:  true,
		},
		4: {
			input: []int{},
			want:  true,
		},
	}

	for index, testCase := range everyTestTable {
		got := Every(testCase.input, predicateFunc)
		if got != testCase.want {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, testCase.want)
		}
	}
}

func TestReduce(t *testing.T) {
	reduceTestTable := map[int]struct {
		input    []string
		initial  int
		callback func(int, string) int
		want     int
	}{
		1: {
			input:    []string{"hello", "my", "dear", "friend"},
			initial:  0,
			callback: func(sum int, s string) int { return sum + len(s) },
			want:     17,
		},
		2: {
			input:    []string{"hello", "my", "dear", "friend"},
			initial:  10,
			callback: func(sum int, s string) int { return sum + 1 },
			want:     14,
		},
		3: {
			input:    nil,
			initial:  0,
			callback: func(sum int, s string) int { return sum + len(s) },
			want:     0,
		},
	}

	for index, testCase := range reduceTestTable {
		got := Reduce(testCase.input, testCase.initial, testCase.callback)
		if got != testCase.want {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, testCase.want)
		}
	}
}

func TestCount(t *testing.T) {
	countTestTable := map[int]struct {
		input []int
		want  int
	}{
		1: {
			input: []int{1, 2, 3, 4, 5, 6},
			want:  3,
		},
		2: {
			input: []int{},
			want:  0,
		},
		3: {
			input: nil,
			want:  0,
		},
	}

	for index, testCase := range countTestTable {
		got := Count(testCase.input, predicateFunc)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, testCase.want)
		}
	}
}
