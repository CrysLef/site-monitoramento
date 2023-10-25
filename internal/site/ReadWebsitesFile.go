package site

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadWebsiteFile() []string {
	var sites []string

	arquivo, err := os.OpenFile("sites.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("erro ao abrir o arquivo: %v", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')

		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	arquivo.Close()

	return sites
}
