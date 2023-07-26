package golangtest

func Absolute(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}

func Add(a, b int) int {
	return a + b
}
/*

		a			b
		negative	negative
		negative	positif
		positif 	negatif
		positif 	positif



*/