package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	_, err := w.Write([]byte("Hello from Snippetbox"))
	if err != nil {
		app.logger.Error(err.Error())
	}
}
