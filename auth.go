package main

import "net/http"

var users = map[string]string{
    "admin": "admin123",
    "user": "password"
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    user := r.FormValue("username")
    pass := r.FormValue("password")
    if users[user] == pass {
        token := generateJWT(user)
        http.SetCookie(w, &http.Cookie{Name: "session", Value: token})
        w.Write([]byte("Login successful"))
    } else {
        w.WriteHeader(403)
    }
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/admin" {
        w.Write([]byte("ADMIN PANEL"))
    }
}
