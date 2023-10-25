package menu

import "fmt"

func ReadInput() int {
	var digito int
	fmt.Printf("Selecione uma opcao: ")
	fmt.Scan(&digito)

	return digito
}
