package main

import (
        "github.com/ant0ine/go-json-rest/rest"
        "log"
        "net/http"
	    //"strings"
	    "gopkg.in/gomail.v2"
        "encoding/json"
        "io/ioutil"
)

type Notification struct {
        File_id   string
        Message   string 
        From 	  string
        Emails 	 []string 
}

type EmailUser struct {
	Name 		string
    Username    string
    Password    string
    EmailServer string
    Port        int
}

func (f *Notification) SetFileId(File_id string) {
    f.File_id = File_id
}

func (f *Notification) SetEmails(Emails []string) {
    f.Emails = Emails
}

func SendEmail_Service(notification Notification,user_auth EmailUser, subject string){
	m := gomail.NewMessage()
	m.SetAddressHeader("From", user_auth.Username, user_auth.Name)
    m.SetHeader("To", notification.Emails...)
    m.SetHeader("Subject", subject)
    var template = `
    <html>
    <h1>Blinkbox</h1>
    `+notification.From +` shared with you the following message:
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

func PostSendNotificationResource(w rest.ResponseWriter, req *rest.Request) {
		var t Notification
		body, err := ioutil.ReadAll(req.Body)
    	err = json.Unmarshal(body, &t)
    	if err != nil {
        	panic(err)
    	}
    	t.SetFileId(req.PathParam("file_id"))
    	user_auth := EmailUser{"Blinkbox Project","blinkboxunal@gmail.com", "bl1nkb0x","smtp.gmail.com", 587}
    	SendEmail_Service(t,user_auth,"You have been selected to test blinkbox new feature")
		w.WriteJson(&t)	
		
}

func main() {
        api := rest.NewApi()
        api.Use(rest.DefaultDevStack...)
        router, err := rest.MakeRouter(
                rest.Post("/notification/:file_id", PostSendNotificationResource),
        )
        if err != nil {
                log.Fatal(err)
        }
        api.SetApp(router)
        log.Fatal(http.ListenAndServe(":4010", api.MakeHandler()))
}

