package menu

import "fmt"

func ReadInput() int {
	var digito int
	fmt.Scan(&digito)
	fmt.Println("A sua opcao foi", digito)

	return digito
}
