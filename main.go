package main

import (
	"log"
	"net/mail"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
	"github.com/pocketbase/pocketbase/tools/mailer"
)

func main() {
	app := pocketbase.New()

	app.OnMailerBeforeRecordVerificationSend().Add(func(e *core.MailerRecordEvent) error {
		// send custom email
		err := e.MailClient.Send(&mailer.Message{
				From:    mail.Address{Address: "support@example.com"},
				To:      mail.Address{Address: e.Record.Email()},
				Subject: "YOUR_SUBJECT...",
				HTML:    "YOUR_HTML_BODY...",
		})

		if err != nil {
				return err
		}

		return hook.StopPropagation
})

	if err := app.Start(); err != nil {
			log.Fatal(err)
	}
}