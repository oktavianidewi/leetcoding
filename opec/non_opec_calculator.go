package opec

type Calcul struct {
	result int
}

func Calculate(a, b int) *Calcul {
	addResult := a + b
	return &Calcul{result: addResult}
}
