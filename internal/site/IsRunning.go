package site

import (
	"fmt"
	"net/http"
)

func IsRunning(url string) {
	resp, _ := http.Get(url)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", url, "carregado com sucesso")
		fmt.Println("Status:", resp.StatusCode)
	} else {
		fmt.Println("Erro ao carregar site:", url)
		fmt.Println("Status:", resp.StatusCode)
	}
}
