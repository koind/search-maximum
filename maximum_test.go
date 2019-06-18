package search_maximum

import "testing"

func TestFindMax(t *testing.T) {
	nums := []interface{}{4, 2, 9}
	expected, _ := FindMax(nums, func(i, j int) bool {
		return nums[i].(int) < nums[j].(int)
	})
	if expected.(int) != 9 {
		t.Errorf("values must match: %v - %v", expected, 9)
	}

	strs := []interface{}{"qwe", "gh", "sdf", "z"}
	expected, _ = FindMax(strs, func(i, j int) bool {
		return strs[i].(string) < strs[j].(string)
	})
	if expected.(string) != "z" {
		t.Errorf("values must match: %v - %v", expected, "z")
	}
}
