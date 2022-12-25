package utils

func For[T interface{}](array []T, item func(i int, v T)) {
	for i, v := range array {
		item(i, v)
	}
}
