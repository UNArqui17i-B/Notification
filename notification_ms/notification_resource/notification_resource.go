package notification_resource

import (
        "github.com/ant0ine/go-json-rest/rest"
        "net/http"
        "app/notification_model"
        "app/notification_service"
        "io/ioutil"
        "encoding/json"
        "os"
        "strconv"
)

func PostSendNotificationResource(w rest.ResponseWriter, req *rest.Request) {
		var t *notification_model.Notification
		body, err := ioutil.ReadAll(req.Body)
    	err = json.Unmarshal(body, &t)
    	if err != nil {
        	panic(err)
    	}
    	t.SetFileId(req.PathParam("file_id"))
    	port, err := strconv.Atoi(os.Getenv("PORT"))
        if err != nil {
            panic(err)
        }
        user_auth := notification_model.EmailUser{os.Getenv("NAME"),os.Getenv("USERNAME"), os.Getenv("PASSWORD"),os.Getenv("EMAIL_SERVER"), port}
        notification_service.SendEmail_Service(t,&user_auth,t.From + " compartio un documento contigo")
		w.WriteJson(&t)	
		w.WriteHeader(http.StatusAccepted)
}

func PostSendConfirmationResource(w rest.ResponseWriter, req *rest.Request) {
        var t *notification_model.Confirmation
        body, err := ioutil.ReadAll(req.Body)
        err = json.Unmarshal(body, &t)
        if err != nil {
            panic(err)
        }
        port, err := strconv.Atoi(os.Getenv("PORT"))
        if err != nil {
            panic(err)
        }
        user_auth := notification_model.EmailUser{os.Getenv("NAME"),os.Getenv("USERNAME"), os.Getenv("PASSWORD"),os.Getenv("EMAIL_SERVER"), port}
        notification_service.SendEmail_Conf_Service(t,&user_auth,"Bienvenido a BlinkBox")
        w.WriteJson(&t) 
        w.WriteHeader(http.StatusAccepted)
}