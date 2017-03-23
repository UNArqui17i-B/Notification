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
                rest.Post("/notification/:file_id", notification_resource.PostSendNotificationResource),
        )
        if err != nil {
                log.Fatal(err)
        }
        api.SetApp(router)
        log.Fatal(http.ListenAndServe(":4010", api.MakeHandler()))
}

