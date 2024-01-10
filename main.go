package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8080", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		randStr1 := r.FormValue("number1")
		randStr2 := r.FormValue("number2")

		number1, err1 := strconv.Atoi(randStr1)
		number2, err2 := strconv.Atoi(randStr2)

		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid age input", http.StatusBadRequest)
			return
		}

		randGo := rand.Intn(number2-number1) + number1

		data := struct {
			Number1, Number2, Result int
		}{
			Number1: number1,
			Number2: number2,
			Result:  randGo,
		}

		tpl.Execute(w, data)
		return
	}

	tpl.Execute(w, nil)
}
