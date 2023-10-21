package site

import (
	"fmt"
	"time"
)

type Monitoring struct {
	Monitoramentos int
	Delay          time.Duration
}

func (m *Monitoring) Start() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < m.Monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			IsRunning(site)
		}
		time.Sleep(m.Delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func (m *Monitoring) Config() (int, time.Duration) {
	fmt.Print("Digite quantos monitoramentos deseja: ")
	fmt.Scan(&m.Monitoramentos)

	fmt.Print("Digite o intervalo de tempo entre monitoramentos: ")
	fmt.Scan(&m.Delay)

	return m.Monitoramentos, m.Delay
}
