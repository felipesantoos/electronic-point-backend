package mail

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopkg.in/mail.v2"
)

const emailBodyFooter = "<p>Atenciosamente&nbsp;<strong>BACKEND TEMPLATE - TEST MAIL</strong>.</p>"

func SendNewAccountEmail(email, password string) error {
	title := fmt.Sprintf("Acesso ao %s", getFromName())
	body := fmt.Sprintf(`
		<h2>Seja bem vindo ao %s!</h2>
		</br>
		<p>Aqui está a sua senha de acesso à plataforma: <strong>%s</strong></p>
		</br>
		<p>Não a compartilhe com terceiros. Ao entrar na plataforma, aconselhamos alterá-la imediatamente.</p>
		</br>
		</br>
		%s
	`, getFromName(), password, emailBodyFooter)
	return send(title, body, []string{email})
}

func SendPasswordResetEmail(accountName, email, passwordResetLink string) error {
	title := fmt.Sprintf("Solicitação de Atualização de Senha")
	body := fmt.Sprintf(`
		<h2>Olá, %s!</h2>
		</br>
		<p>Recebemos uma solicitação de atualização de senha para o seu usuário.</p>
		<p>Para atualizar a sua senha, acesse o link abaixo:</p>
		</br>
		<a href="%s" target="_blank">%s</a>
		</br>
		<p>Obs.: esse link tem duração de 1 hora.</p>
		</br>
		<p>Caso você não tenha solicitado a atualização de sua senha, por favor, ignore este e-mail.</p>
		</br>
		</br>
		%s
	`, accountName, passwordResetLink, passwordResetLink, emailBodyFooter)
	return send(title, body, []string{email})
}

func send(title, body string, emails []string) error {
	message := mail.NewMessage()
	message.SetHeader("From", getFromName())
	message.SetHeader("To", strings.Join(emails, ","))
	message.SetHeader("Subject", title)
	message.SetHeader("Sender", getFromAddress())
	message.SetBody("text/html", body)
	dialer, err := newDialer()
	if err != nil {
		return err
	}
	return dialer.DialAndSend(message)
}

func newDialer() (*mail.Dialer, error) {
	serviceHost := os.Getenv("MAIL_SMTP_HOST")
	servicePort := os.Getenv("MAIL_SMTP_PORT")
	fromEmail := getFromAddress()
	fromPassword := os.Getenv("MAIL_FROM_PASSWORD")
	if serviceHost == "" {
		return nil, errors.New("you need to define the mail service host!")
	} else if servicePort == "" {
		return nil, errors.New("you need to define the mail service port!")
	} else if fromEmail == "" {
		return nil, errors.New("you need to define the responsible mail email address!")
	} else if fromPassword == "" {
		return nil, errors.New("you need to define the responsible email password!")
	} else if _, err := strconv.Atoi(servicePort); err != nil {
		return nil, errors.New("the mail service port must be an integer!")
	} else if getFromName() == "" {
		return nil, errors.New("you need to define the email sender name!")
	}
	intServicePort, _ := strconv.Atoi(servicePort)
	return mail.NewDialer(serviceHost, intServicePort, fromEmail, fromPassword), nil
}

func getFromAddress() string {
	return os.Getenv("MAIL_FROM_ADDRESS")
}

func getFromName() string {
	return os.Getenv("MAIL_FROM_NAME")
}
