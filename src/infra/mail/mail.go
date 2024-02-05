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
	return Send(title, body, []string{email})
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
	return Send(title, body, []string{email})
}

func Send(title, body string, emails []string) error {
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

const EmailTitle = "Sua Voz Nossa Lei - Email de Confirmação e Agradecimento!"
const EmailBody = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Sua Voz Nossa Lei - Email de Confirmação e Agradecimento!</title>
    <style>
        * {
            font-family: Arial, serif;
        }
    </style>
</head>
<body>
<div>
    Prezado %s,<br><br>
    Em nome de toda a equipe envolvida no desenvolvimento do Projeto de Lei Eletrônico, queremos expressar nossa sincera gratidão pela sua participação ativa e pelo valioso apoio que nos proporcionou ao submeter o formulário em favor da iniciativa.<br><br>
    Sua assinatura eletrônica representa um passo significativo em direção à modernização e aprimoramento do processo democrático. É com grande apreço que recebemos o seu engajamento, demonstrando o comprometimento com a construção de uma sociedade mais participativa e eficiente.<br><br>
    Cada assinatura é uma peça fundamental na construção desse projeto, e estamos entusiasmados por contar com o seu apoio. Seu gesto solidifica o propósito do projeto em proporcionar uma plataforma segura, prática e autêntica para a expressão da vontade popular.<br><br>
    Manteremos você informado sobre o andamento do projeto e sobre as próximas etapas. Sua confiança é vital para o sucesso desta iniciativa, e estamos dedicados a honrá-la da melhor maneira possível.<br><br>
    Agradecemos mais uma vez por sua colaboração e pela confiança depositada em nosso trabalho. Juntos, estamos construindo um futuro mais participativo e democrático.
</div>
</body>
</html>
`
