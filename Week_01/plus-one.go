package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {

			digits[i] = 0;
		} else {
			digits[i]++
			return digits
		}

	}
	digits = append(digits, 0)
	digits[0] = 1
	return digits

}
