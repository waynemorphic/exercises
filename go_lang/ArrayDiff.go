package main

// Implement a difference function, which subtracts one list from another and returns the result.

// It should remove all values from list a, which are present in list b keeping their order.

// array_diff({1, 2}, 2, {1}, 1, *z) == {2} (z == 1)
// If a value is present in b, all of its occurrences must be removed from the other:

// array_diff({1, 2, 2, 2, 3}, 5, {2}, 1, *z) == {1, 3} (z == 2)

func ArrayDiff(a, b []int) []int {
	var res []int

	for _, i := range(a) {
		var found = false

		for _, j := range(b) {
			if i == j{
				found = true
				break
			}
		}
		if !found {
			res = append(res, i)
		}
	}
	return res
}