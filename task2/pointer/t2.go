package main

func DoubleSlice(s *[]int) {
	for i, v := range *s {
		(*s)[i] = v * 2
	}
}
