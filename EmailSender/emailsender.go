package EmailSender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"net/smtp"
	configs "trade-platform/Configs"
	entities "trade-platform/Entities"
)

var conn, _ = amqp.Dial("amqp://guest:guest@localhost:5672/")
var ch, _ = conn.Channel()

func Start() {
	defer conn.Close()
	defer ch.Close()

	ch.ExchangeDeclare(
		"emails_direct", // name
		"direct",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	ch.QueueDeclare(
		"emails", true, false, false, false, nil)
	ch.QueueBind(
		"emails", "emails", "emails_direct", false, nil)
	ch.QueueDeclare("emails_delayed", true, false, false, false,
		map[string]interface{}{
			"x-message-ttl":             60000,
			"x-dead-letter-exchange":    "emails_direct",
			"x-dead-letter-routing-key": "emails",
		})
	ch.QueueBind("emails_delayed", "emails_delayed", "emails_direct", false, nil)

	msgs, err := ch.Consume("emails", "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var content entities.EmailContent
			json.NewDecoder(bytes.NewReader(d.Body)).Decode(&content)
			if !SendEmail(content.CustomerEmail, content.GameKey) {
				SendEmailToDelayQueue(content)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func SendEmail(customerEmail string, key string) bool {
	auth := smtp.PlainAuth("", configs.SmtpClientEmail, configs.SmtpClientPassword, configs.SmtpClientHost)
	to := []string{customerEmail}
	msg := []byte("To: " + customerEmail + "\r\n" +
		"Subject: Trade platform!\r\n" +
		"\r\n" +
		"Thanks for your purchase! Here's your key: " + key + "\r\n")
	err := smtp.SendMail(configs.SmtpClientAddress, auth, configs.SmtpClientEmail, to, msg)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
	return err == nil
}

func SendEmailToDelayQueue(content entities.EmailContent) {
	body, _ := json.Marshal(content)
	ch.Publish("emails_direct", "emails_delayed", false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
}

func SendEmailMessage(customerEmail string, key string) {
	var emailContent = entities.EmailContent{CustomerEmail: customerEmail, GameKey: key}
	body, err := json.Marshal(emailContent)
	failOnError(err, "Unable to marshall")
	err = ch.Publish("emails_direct", "emails", false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")
}
