package site

import (
	"fmt"
	"net/http"
)

func IsRunning(url string) {
	resp, err := http.Get(url)
	websiteLogs := WebsiteLogs{
		Url:        url,
		StatusCode: resp.StatusCode,
	}

	if err != nil {
		fmt.Printf("erro ao monitorar o site: %v", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", url, "carregado com sucesso")
		websiteLogs.Registry()

	} else {
		fmt.Println("Erro ao carregar site:", url)
		websiteLogs.Registry()
	}
}
