package notification_model

type Notification struct {
        File_id   string
        Message   string 
        From 	  string
        Emails 	 []string 
}

type Confirmation struct {
        Conf_url    string
        Email       string


}

type Recover struct{
    Recover_url     string
    User            string
    Email           string

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