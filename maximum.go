package search_maximum

import (
	"errors"
	"reflect"
	"sort"
)

const countTypesInOneSlice = 1

func FindMax(sl []interface{}, less func(i, j int) bool) (interface{}, error) {
	if len(sl) <= 0 {
		return nil, errors.New("slice cannot be empty")
	}

	if !hasOneType(sl) {
		return nil, errors.New("slice values must be of the same type")
	}

	return comparator(sl, less), nil
}

func hasOneType(sl []interface{}) bool {
	types := map[string]int{}

	for _, v := range sl {
		types[reflect.TypeOf(v).String()]++
	}

	return len(types) == countTypesInOneSlice
}

func comparator(sl []interface{}, less func(i, j int) bool) interface{} {
	itemType := reflect.TypeOf(sl[0])
	newSlice := reflect.MakeSlice(reflect.SliceOf(itemType), 0, len(sl))
	for _, v := range sl {
		newSlice = reflect.Append(newSlice, reflect.ValueOf(v))
	}
	sort.SliceStable(newSlice.Interface(), less)
	item := newSlice.Index(newSlice.Len() - 1)

	return item.Interface()
}
