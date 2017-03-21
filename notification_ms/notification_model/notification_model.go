package notification_model

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