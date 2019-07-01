package singlylinkedlist

import (
	"reflect"
	"testing"
)

func TestListNew(t *testing.T) {
	testCases := []struct {
		name   string
		values []interface{}
		want   []interface{}
	}{
		{"Empty", []interface{}{}, []interface{}{}},
		{"NonEmpty", []interface{}{1, 2, 3}, []interface{}{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list := New(tc.values...)
			got := list.Values()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestListAppend(t *testing.T) {
	testCases := []struct {
		name   string
		values []interface{}
		want   []interface{}
	}{
		{"SingleValue", []interface{}{1}, []interface{}{1}},
		{"MultipleValues", []interface{}{1, 2, 3}, []interface{}{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list := New()
			list.Append(tc.values...)
			got := list.Values()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestListPrepend(t *testing.T) {
	testCases := []struct {
		name   string
		values []interface{}
		want   []interface{}
	}{
		{"SingleValue", []interface{}{1}, []interface{}{1}},
		{"MultipleValues", []interface{}{1, 2, 3}, []interface{}{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list := New()
			list.Append(tc.values...)
			got := list.Values()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestListRemove(t *testing.T) {
	list := New(1, 2, 3, 4)

	testCases := []struct {
		name          string
		indexToRemove int
		want          []interface{}
	}{
		{"MiddleIndex", 1, []interface{}{1, 3, 4}},
		{"NonExistentIndex", 3, []interface{}{1, 3, 4}},
		{"ZeroIndex", 0, []interface{}{3, 4}},
		{"LastIndex", 1, []interface{}{3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list.Remove(tc.indexToRemove)
			got := list.Values()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestListGet(t *testing.T) {
	list := New(1, 2, 3, 4)

	testCases := []struct {
		name       string
		indexToGet int
		want       interface{}
		exists     bool
	}{
		{"MiddleIndex", 1, 2, true},
		{"NonExistentIndex", 4, nil, false},
		{"ZeroIndex", 0, 1, true},
		{"LastIndex", 3, 4, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, exists := list.Get(tc.indexToGet)
			if got != tc.want || exists != tc.exists {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestListValues(t *testing.T) {
	testCases := []struct {
		name   string
		values []interface{}
		want   []interface{}
	}{
		{"Empty", []interface{}{}, []interface{}{}},
		{"NonEmpty", []interface{}{1, 2, 3}, []interface{}{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list := New(tc.values...)
			got := list.Values()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestListSize(t *testing.T) {
	testCases := []struct {
		name   string
		values []interface{}
		want   int
	}{
		{"Empty", []interface{}{}, 0},
		{"NonEmpty", []interface{}{1, 2, 3}, 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list := New(tc.values...)
			got := list.Size()
			if got != tc.want {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestListIsEmpty(t *testing.T) {
	testCases := []struct {
		name   string
		values []interface{}
		want   bool
	}{
		{"Empty", []interface{}{}, true},
		{"NonEmpty", []interface{}{1, 2, 3}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list := New(tc.values...)
			got := list.IsEmpty()
			if got != tc.want {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestListClear(t *testing.T) {
	testCases := []struct {
		name   string
		values []interface{}
	}{
		{"Empty", []interface{}{}},
		{"NonEmpty", []interface{}{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list := New(tc.values...)
			list.Clear()
			if !list.IsEmpty() {
				t.Errorf("list is non-empty %v", list.Values())
			}
		})
	}
}
