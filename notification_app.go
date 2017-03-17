package main

import (
        "github.com/ant0ine/go-json-rest/rest"
        "log"
        "net/http"
        "net/smtp"
	    //"strings"
	    "bytes"
	    "strconv"
        "text/template"
        "encoding/json"
        "io/ioutil"
)

type Notification struct {
        File_id   string
        Message   string 
        Emails 	 []string 
}

type EmailUser struct {
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

func SendMailService(){
		emailUser := &EmailUser{"blinkboxunal@gmail.com", "bl1nkb0x", "smtp.gmail.com", 587}

		auth := smtp.PlainAuth("",
		    emailUser.Username,
		    emailUser.Password,
		    emailUser.EmailServer,
		)

		type SmtpTemplateData struct {
		    From    string
		    To      string
		    Subject string
		    Body    string
		}

		const emailTemplate = `Ola k ase`

		var err error
		var doc bytes.Buffer

		context := &SmtpTemplateData{
		    "SmtpEmailSender",
		    "recipient@domain.com",
		    "You have been seleted to try the new Blinkboxunal-microservice!",
		    "Hello, this is a test e-mail body.",
		}
		templ := template.New("emailTemplate")
		templ, err = templ.Parse(emailTemplate)
		if err != nil {
		    log.Print("error trying to parse mail template")
		}
		err = templ.Execute(&doc, context)
		if err != nil {
		    log.Print("error trying to execute mail template")
		}

		err = smtp.SendMail(emailUser.EmailServer+":"+strconv.Itoa(emailUser.Port), // in our case, "smtp.google.com:587"
		    auth,
		    emailUser.Username,
		    []string{"nrgiraldoc@unal.edu.co","afmesag@unal.edu.co"},
		    doc.Bytes())
		if err != nil {
		    log.Print("ERROR: attempting to send a mail ", err)
		}
        
}

func PostSendNotification(w rest.ResponseWriter, req *rest.Request) {
		body, err := ioutil.ReadAll(req.Body)
		var t Notification
    	err = json.Unmarshal(body, &t)
    	if err != nil {
        	panic(err)
    	}
		t.SetFileId(req.PathParam("file_id"))
		SendMailService()
		w.WriteJson(&t)	
		
}

func main() {
        api := rest.NewApi()
        api.Use(rest.DefaultDevStack...)
        router, err := rest.MakeRouter(
                rest.Post("/notification/:file_id", PostSendNotification),
        )
        if err != nil {
                log.Fatal(err)
        }
        api.SetApp(router)
        log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

