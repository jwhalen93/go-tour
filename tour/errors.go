package tour

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func SqrtError(x float64) (float64, error) {
	if x < 0.0 {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 6; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

/*func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}*/
