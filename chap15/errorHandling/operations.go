package main

type CategoryError struct {
	requestedCategory string
}

func (e *CategoryError) Error() string {
	return "Category " + e.requestedCategory + " does not exist"
}

func (slice ProductSlice) TotalPrice(category string) (total float64,
	err *CategoryError) {
	productCount := 0
	for _, p := range slice {
		// Products에서 category가 같은 항목을 찾아서 total에 더한다.
		if p.Category == category {
			total += p.Price
			productCount++
		}
	}

	// productCount가 0이라면 해당 category가 같은 항목이 한개도 없기때문에 err를 할당한다.
	if productCount == 0 {
		err = &CategoryError{requestedCategory: category}
	}
	return
}
