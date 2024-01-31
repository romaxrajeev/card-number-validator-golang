package pkg

import (
	"errors"
	"strconv"
)

func LuhnValidation(cardNumber *[]int) (bool, error) {
	sum := 0
	for index, number := range *cardNumber {
		if (index+1)%2 == 0 {
			additionNumber := 0
			if number*2 > 9 {
				additionNumber = number*2 - 9
			} else {
				additionNumber = number * 2
			}

			sum += additionNumber
		} else {

			sum += number
		}

	}
	sum = sum * 9
	if sum%10 != 0 {
		return false, errors.New("Checksum failed for sum: " + strconv.Itoa(sum))
	}
	return true, nil
}
