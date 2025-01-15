package main

import (
	"io"
	"net/http"
)

const form = `<html>
    <head>
    <title></title>
    </head>
    <body>
        <form action="/" method="post">
            <label>Логин</label><input type="text" name="login">
            <label>Пароль<input type="password" name="password">
            <input type="submit" value="Login">
        </form>
    </body>
</html>`

func mreain() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(authenticationHandler))
	if err != nil {
		panic(err)
	}
}

func Auth(login, password string) bool {
	return login == "sanzhar@mail.ru" && password == "sanzhar2005"
}

func authenticationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		if Auth(login, password) {
			io.WriteString(w, "Welcome, Sanzhar!")
		} else {
			http.Error(w, "Wrong login or password", http.StatusUnauthorized)
		}
		return
	}
	io.WriteString(w, form)
}
