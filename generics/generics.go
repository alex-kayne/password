package generics

type numberType interface{ int | float32 | float64 }

func Sum[number numberType](a, b number) number {
	return a + b
}
