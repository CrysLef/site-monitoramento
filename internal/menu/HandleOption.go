package menu

import (
	"fmt"
	"os"

	"github.com/CrysLef/site-monitoramento/internal/site"
)

func HandleOption() {

	comando := ReadInput()
	monitoring := site.Monitoring{}

	switch comando {
	case 1:
		monitoring.Config()
		monitoring.Start()

	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Opcao nao valida!")
		os.Exit(-1)
	}
}
