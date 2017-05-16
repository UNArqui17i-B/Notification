package notification_service

import (
	    "gopkg.in/gomail.v2"
        "app/notification_model"
)

func SendEmail_Service(notification *notification_model.Notification,user_auth *notification_model.EmailUser, subject string){
	m := gomail.NewMessage()
	m.SetAddressHeader("From", user_auth.Username, user_auth.Name)
    m.SetHeader("To", notification.Emails...)
    m.SetHeader("Subject", subject)
    var template = `
    <html>
    <h1>Blinkbox</h1>
    `+notification.From +` compartio contigo el siguiente archivo:
    <br>
    http://blinkboxunal.com/file/`+notification.File_id+`
    <br>
    con el siguiente mensaje <br>
    <h1>`+notification.Message+`</h1>
    </html>`
    m.SetBody("text/html", template)

    d := gomail.NewPlainDialer(user_auth.EmailServer, user_auth.Port, user_auth.Username, user_auth.Password)

    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }

}

func SendEmail_Conf_Service(confirmation *notification_model.Confirmation,user_auth *notification_model.EmailUser, subject string){
    m := gomail.NewMessage()
    m.SetAddressHeader("From", user_auth.Username, user_auth.Name)
    m.SetHeader("To", confirmation.Email)
    m.SetHeader("Subject", subject)
    var template = `
    <html>
    <h1>Blinkbox</h1>
    Para empezar a compartir tus archivos por favor confirma tu cuenta con el siguiente link:
    <br>
    `+confirmation.Conf_url+`
    <br>
    Estamos muy felices de tenerte entre nosotros
    </html>`
    m.SetBody("text/html", template)

    d := gomail.NewPlainDialer(user_auth.EmailServer, user_auth.Port, user_auth.Username, user_auth.Password)

    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }

}

func SendEmail_Recover_Service(recover *notification_model.Recover,user_auth *notification_model.EmailUser, subject string){
    m := gomail.NewMessage()
    m.SetAddressHeader("From", user_auth.Username, user_auth.Name)
    m.SetHeader("To", recover.Email)
    m.SetHeader("Subject", subject)
    var template = `
    <html>
    <h1>Blinkbox</h1>
    Hola `+recover.User+` para recuperar tu cuenta por favor entra al siguiente link:
    <br>
    `+recover.Recover_url+`
    <br>
    Estamos muy felices de tenerte entre nosotros
    </html>`
    m.SetBody("text/html", template)

    d := gomail.NewPlainDialer(user_auth.EmailServer, user_auth.Port, user_auth.Username, user_auth.Password)

    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }

}