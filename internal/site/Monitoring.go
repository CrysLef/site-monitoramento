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
	sites := ReadWebsiteFile()

	if sites[0] != "" {
		for i := 0; i < m.Monitoramentos; i++ {

			for i, site := range sites {
				fmt.Println("Monitorando site", i+1, ":", site)
				IsRunning(site)
			}
			time.Sleep(m.Delay * time.Second)
			fmt.Println("")
		}
		fmt.Println("")
	} else {
		fmt.Println("Sua lista de sites esta vazia, por favor registre pelo menos um site antes de realizar o monitoramento.")
		fmt.Println("")
	}
}

func (m *Monitoring) Config() (int, time.Duration) {

	for m.Monitoramentos == 0 || m.Delay == 0 {
		fmt.Print("Digite quantos monitoramentos deseja: ")
		fmt.Scan(&m.Monitoramentos)

		if m.Monitoramentos == 0 {
			fmt.Println("Ao menos 1 monitoramento deve ser feito")
			continue
		}

		if m.Monitoramentos == 1 {
			break
		}
		fmt.Print("Digite o intervalo de tempo entre monitoramentos: ")
		fmt.Scan(&m.Delay)
	}

	return m.Monitoramentos, m.Delay
}
