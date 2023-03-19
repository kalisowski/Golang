package main

func main() {
	var arr1 [20]float64
	var arr2 [20]float64
	var res [20]float64

	for i := range arr1 {
		arr1[i] = 2.0
		arr2[i] = 3.0
		res[i] = arr1[i] * arr2[i]
		println(res[i])
	}
}
