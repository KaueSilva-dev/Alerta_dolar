🔔 Automatizando Alertas de Cotação do Dólar com Go no macOS (Notificações + WhatsApp)

✨ Introdução: Do evento à ideia

Durante a GopherCon LatAm 2025, participei de um workshop liderado por ninguém menos que Bill Kennedy, referência global no ecossistema Go e membro da Ardan Labs.

O evento foi repleto de palestras incríveis, e quem passou pelo stand da Ardan Labs recebeu um benefício irrecusável: 50% de desconto em um curso avançado de Go — um dos mais renomados do mundo.

Com o desconto em mãos e a inspiração do evento ainda fresca, só havia uma questão:📉 Quando seria o melhor dia para aproveitar o dólar mais baixo e efetuar a compra?

Foi então que decidi resolver isso do jeito que nós, desenvolvedores, gostamos:automatizando!

⚙️ O que o script faz?

Criei um script em Go que:

✅ Consulta a cotação do dólar dos últimos 7 dias úteis

✅ Compara o valor atual com o menor do período

✅ Emite uma notificação no macOS (mesmo com Ventura/Sonoma)

✅ Dispara um WhatsApp automático com o resumo da cotação

✅ É agendado para rodar diariamente via Automator + Calendário

💻 O código em Go

📅 Consulta à AwesomeAPI:

resp, err := http.Get("https://economia.awesomeapi.com.br/json/daily/USD-BRL/7")

🔎 Avaliação e recomendação:

if currentBid < minLow+0.01 {
    fmt.Println("Recomendação: comprar agora.")
} else {
    fmt.Println("Recomendação: aguardar.")
}

📟 Notificação no macOS com timeout (sem botões):

func notificarComGivingUp(titulo, mensagem string) {
    script := fmt.Sprintf(`
display dialog "%s" with title "%s" giving up after 5
`, mensagem, titulo)
    cmd := exec.Command("osascript")
    cmd.Stdin = strings.NewReader(script)
    cmd.Run()
}

📲 Envio de WhatsApp com CallMeBot

Ative seu número enviando “I allow callmebot to send me messages” para o número +34 603 21 25 97.

Use a API:

url := fmt.Sprintf("https://api.callmebot.com/whatsapp.php?phone=%s&text=%s&apikey=%s", telefone, msg, apiKey)
http.Get(url)

[Mais informações sobre Call Me Bot](https://www.callmebot.com/blog/free-api-whatsapp-messages/)

🗓️ Agendamento via Automator + Calendário

Crie um .app no Automator com ação "Executar Shell Script" chamando seu binário Go.

Crie um evento recorrente no Calendário do macOS, com ação "Abrir Arquivo" → selecione o .app.

Pronto: agendamento diário sem terminal, sem cron, sem complicação.

📸 Resultado final
![Captura de Tela 2025-05-12 às 01 57 46](https://github.com/user-attachments/assets/6aeba0c5-b190-4fbb-b5cb-0d9b4d38d98b)

![Captura de Tela 2025-05-12 às 01 58 37](https://github.com/user-attachments/assets/cbf5c888-1707-48a6-9419-e54a161dc3bf)

<img width="551" alt="Captura de Tela 2025-05-12 às 01 58 07" src="https://github.com/user-attachments/assets/d16baf4d-708b-437e-af12-fecb3c75dc74" />


🚀 Conclusão

Essa automação é um ótimo exemplo de como podemos usar Go para resolver problemas reais do dia a dia — indo de um desafio prático (qual o melhor dia para comprar em dólar?) a uma solução automatizada que combina APIs, scripts e ferramentas nativas do sistema.

Porque tecnologia é isso: transformar decisões em soluções. 😎
