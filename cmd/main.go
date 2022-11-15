package main

import (
	"flag"
	"log"
)

func main() {

	// token = flags.Get(token)
	////////// Он выташит из запроса token клиента, что бы общаться с ним


	var token string  = mustToken()

	// tgClient = telegram.New(token)
	////////// Это клиент который мы общаться

	// fetcher = fetcher.New(tgClient)
	////////// он нужен для того что бы получать новые событие

	// proccessor 'processor.New(tgClient)
	////////// Для того что бы отправлять новые сообшении
	////////// То есть кидать ссылки или ругать что не понял комманду

	// consumer.Start(fetcher, processor)
	////////// Consumer  - он получает событие и обрабытвает их
	////////// для того что-бы его получать нужен fetcher
	////////// для обработки processor

}

func mustToken() string {
	var token *string = flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	if *token == "" {
		log.Fatal("User token is empty")
	}

	return *token 
}
