package pkg

import "strconv"

func StringToListOfIntConverter(number string) ([]int, error) {
	cardNumber := []int{}
	for _, num := range number {
		intNum, err := strconv.Atoi(string(num))
		if err != nil {
			return []int{}, err
		} else {
			cardNumber = append(cardNumber, intNum)
		}
	}
	return cardNumber, nil
}
