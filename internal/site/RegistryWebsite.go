package site

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func RegistryWebsites() error {
	var url string

	arquivo, err := os.OpenFile("sites.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("erro ao abrir o arquivo: %v", err)

	}

	var nRegistryWebsites int
	fmt.Print("Digite quantos websites deseja registrar: ")
	fmt.Scan(&nRegistryWebsites)

	for i := 0; i < nRegistryWebsites; i++ {
		fmt.Print("Escreva a URL do site que deseja registrar: ")
		fmt.Scan(&url)

		url = ValidateWebsiteURL(url)
		url = url + "\n"

		writer := bufio.NewWriter(arquivo)
		_, err = writer.WriteString(url)
		writer.Flush()
		if err != nil {
			return fmt.Errorf("erro ao abrir o arquivo: %v", err)
		}

		fmt.Println("Website registrado com sucesso")
	}

	defer arquivo.Close()

	return nil
}

func ValidateWebsiteURL(url string) string {
	url = strings.TrimSpace(url)
	regex := `(https:\/\/www\.|http:\/\/www\.|https:\/\/|http:\/\/)([a-zA-Z]{1,})(\.[a-zA-Z]{2,3})(\.[a-zA-Z]{2,3})?(\/[a-zA-Z0-9]{2,})?`

	matched, err := regexp.MatchString(regex, url)
	if err != nil {
		fmt.Printf("erro ao registrar o site: %v", err)
		os.Exit(-1)
	}

	if !matched {
		fmt.Println("URL invalida!")
		fmt.Println("Tente uma url com o protocolo http ou https, como por exemplo: https://nomedosite.dominio ou https://www.nomedosite.dominio")
		os.Exit(-1)
	}

	return url
}
