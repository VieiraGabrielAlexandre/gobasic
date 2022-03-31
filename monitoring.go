package main

import (
	"fmt"
	"net/http"
	"time"
)

const monitoramento = 5
const delay = 3

func main() {
	iniciarMonitoramento()
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	site := "https://random-status-code.herokuapp.com/"

	for i := 0; i <= monitoramento; i++ {
		resp, _ := http.Get(site)
		fmt.Println("Status Code:", resp.Status)

		if resp.StatusCode != 200 {
			fmt.Println("Falha")
		} else {
			fmt.Println("Sucesso")
		}

		time.Sleep(delay * time.Second)
	}
}
