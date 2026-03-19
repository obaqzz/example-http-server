package main

import (
    "html/template"
    "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("views/index.html")

    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, nil) // replace nil with data struct if you want to pass dynamic data to the template
}