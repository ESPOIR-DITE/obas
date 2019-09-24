package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
)

// Route Path
func Register(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", RegisterHandler(app))
	return r

}

func RegisterHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "/register/register.page.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())

		}

	}

}
