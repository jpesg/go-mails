package mails

import (
	mailer "go-mails/pkg/mails/infra"
	infra "go-mails/pkg/shared/infra"
	"strings"
	"sync"
)

func SendMails() {
	var wg sync.WaitGroup

	senders := infra.GetEnvVariable("SENDERS")
	passwords := infra.GetEnvVariable("PASSWORDS")
	_senders := strings.Split(senders, ",")
	_passwords := strings.Split(passwords, ",")
	wg.Add(len(_senders))

	for i, s := range _senders {
		from := s
		pass := _passwords[i]
		go func(content string, from string, to []string, pass string) {
			mailer.MailSend("content", from, to, pass, &wg)
		}("content", from, _senders, pass)
	}
	wg.Wait()

}
