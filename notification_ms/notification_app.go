package main

import (
        "github.com/ant0ine/go-json-rest/rest"
        "log"
        "net/http"
        "app/notification_resource"
)

func main() {
        api := rest.NewApi()
        api.Use(rest.DefaultDevStack...)
        router, err := rest.MakeRouter(
                rest.Post("/notification/messages/:file_id", notification_resource.PostSendNotificationResource),
                rest.Post("/notification/confirmation",notification_resource.PostSendConfirmationResource),
        )
        if err != nil {
                log.Fatal(err)
        }
        api.SetApp(router)
        log.Fatal(http.ListenAndServe(":4010", api.MakeHandler()))
}

