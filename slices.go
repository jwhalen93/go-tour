package main

func Pic(dx, dy int) [][]uint8 {
	dySlice := make([][]uint8, dy)
	for i := range dySlice {
		dySlice[i] = make([]uint8, dx)
		for j := range dySlice[i] {
			dySlice[i][j] = uint8(i ^ j)
		}
	}
	return dySlice
}

/*func main() {

	pic.Show(Pic)
}*/
