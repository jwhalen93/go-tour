package tour

// fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	first := 0
	second := 0
	return func() int {
		if first == 0 && second == 0 {
			first = 1
			return 0
		}
		first, second = second, first+second
		return second
	}
}

/*func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}*/
