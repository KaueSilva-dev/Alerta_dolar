package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Cotacao struct {
	Bid      string `json:"bid"`
	Low      string `json:"low"`
	TimeUnix int64  `json:"timestamp,string"`
}

func main() {
	const (
		diasCotacao = 7
		phone       = "+554796493799"
		apiKey      = "7193559"
	)

	resp, err := http.Get(fmt.Sprintf("https://economia.awesomeapi.com.br/json/daily/USD-BRL/%d", diasCotacao))
	if err != nil {
		log.Fatal("error ao buscar cotação", err)
	}
	defer resp.Body.Close()

	var historico []Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&historico); err != nil {
		log.Fatalf("error ao decodificar json %v", err)
	}

	if len(historico) == 0 {
		log.Fatalf("historico vazio")
	}

	sort.Slice(historico, func(i, j int) bool {
		return historico[i].TimeUnix < historico[j].TimeUnix
	})

	var minLow float64 = 999.99
	var currentBid float64
	var resumo strings.Builder

	resumo.WriteString("Histórico de dólar (último 7 dias úteis): ")

	for i, dia := range historico {
		bid, _ := strconv.ParseFloat(dia.Bid, 64)
		low, _ := strconv.ParseFloat(dia.Low, 64)
		data := time.Unix(dia.TimeUnix, 0).Format("2006-01-02 15:04:05")

		resumo.WriteString(fmt.Sprintf("Dia %d (%s): LOW R$%.2f | BID R$%.2f\n", i+1, data, low, bid))

		if low < minLow {
			minLow = low
		}

		if i == len(historico)-1 {
			currentBid = bid
		}
	}

	resumo.WriteString("__________________________\n")
	resumo.WriteString(fmt.Sprintf("Menor valor encontrado: R$%.2f\n", minLow))
	resumo.WriteString(fmt.Sprintf("Valor atual: R$%.2f\n", currentBid))

	if currentBid < minLow+0.01 {
		notifierMacOS("Alerta do Dólar", "Recomendação: comprar agora - valor está entre os mais baixos recentes!")
		resumo.WriteString("Recomendacao: comprar agora.")
	} else {
		notifierMacOS("Alerta do Dólar", "Recomendação: aguardar - valor acima do menor dos últimas dias")
		resumo.WriteString("Recomendacao: comprar agora.")
	}

	fmt.Println(resumo.String())
	enviarWhatsApp(phone, apiKey, resumo.String())
}

func notifierMacOS(titulo, mensagem string) {
	script := fmt.Sprintf(`
display dialog "%s" with title "%s" giving up after 5
`, mensagem, titulo)

	// Cria o comando osascript sem argumentos (vai ler do stdin)
	cmd := exec.Command("osascript")

	// Passa o script como entrada padrão (stdin)
	cmd.Stdin = strings.NewReader(script)

	// Executa o comando
	if err := cmd.Run(); err != nil {
		fmt.Println("Erro ao executar AppleScript:", err)
	}
}

func enviarWhatsApp(telefone, apiKey, mensagem string) {
	msg := url.QueryEscape(mensagem)

	urlFinal := fmt.Sprintf("https://api.callmebot.com/whatsapp.php?phone=%s&text=%s&apikey=%s", telefone, msg, apiKey)
	resp, err := http.Get(urlFinal)
	if err != nil {
		fmt.Println("Erro ao enviar WhatsApp:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Mensagem enviada via WhatsApp.")
}
