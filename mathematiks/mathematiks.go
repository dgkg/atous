package mathematiks

import "fmt"

// Créer un nouveau type appelé myInt de type int32
type myInt int32

// Créer les méthodes suivantes :
// Divide : retourner la division avec un nombre n de type int passé en paramètre
func (m myInt) Divide(n int) (int, error) {
	if n == 0 {
		return 0, fmt.Errorf("Division par 0 impossible")
	}
	return int(m) / n, nil
}

// Add : retourner la valeur ajouté par n de type int passé en paramètre
func (m myInt) Add(n int) (int, error) {
	return int(m) + n, nil
}

// Sub : retourner la valeur soustraite avec n toujours passé en paramètre
func (m myInt) Sub(n int) (int, error) {
	return int(m) - n, nil
}

// Multiply : retourner la valeur multiplié des deux paramètres de type int en myInt
func (m myInt) Multiply(n int) (int, error) {
	return int(m) * n, nil
}
