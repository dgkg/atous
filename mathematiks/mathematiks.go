package mathematiks

import "errors"

// Créer un nouveau type appelé myInt de type int32
type MyInt int32

var (
	ErrTryDivideByZero = errors.New("mathematiks: division par 0 impossible")
)

// Créer les méthodes suivantes :
// Divide : retourner la division avec un nombre n de type int passé en paramètre
func (m MyInt) Divide(n int) (int, error) {
	if n == 0 {
		return 0, ErrTryDivideByZero
	}
	return int(m) / n, nil
}

// Add : retourner la valeur ajouté par n de type int passé en paramètre
func (m MyInt) Add(n int) (MyInt, error) {
	return MyInt(int(m) + n), nil
}

// Sub : retourner la valeur soustraite avec n toujours passé en paramètre
func (m MyInt) Sub(n int) (MyInt, error) {
	return MyInt(int(m) - n), nil
}

// Multiply : retourner la valeur multiplié des deux paramètres de type int en myInt
func (m MyInt) Multiply(n int) (int, error) {
	return int(m) * n, nil
}
