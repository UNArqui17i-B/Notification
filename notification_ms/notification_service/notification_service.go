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
    `+notification.From +` shared with you the following file:
    <br>
    http://blinkboxunal.com/file/`+notification.File_id+`
    <br>
    with the following message <br>
    <h1>`+notification.Message+`</h1>
    </html>`
    m.SetBody("text/html", template)

    d := gomail.NewPlainDialer(user_auth.EmailServer, user_auth.Port, user_auth.Username, user_auth.Password)

    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }

}