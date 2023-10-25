package site

import (
	"fmt"
	"os"
	"time"
)

type WebsiteLogs struct {
	Url        string
	StatusCode int
}

func (log *WebsiteLogs) Registry() {
	arquivo, err := os.OpenFile("sites.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("erro ao registrar log: %v", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + log.Url + " - " + "status code: " + fmt.Sprint(log.StatusCode) + "\n")
}

func (*WebsiteLogs) Read() {
	arquivo, err := os.ReadFile("sites.log")
	if err != nil {
		fmt.Printf("erro ao abrir log: %v", err)
	}
	fmt.Println(string(arquivo))
}
