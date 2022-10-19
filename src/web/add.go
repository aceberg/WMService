package web

import (
	"github.com/aceberg/WMService/conf"
	. "github.com/aceberg/WMService/models"
	"html/template"
	"net/http"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon



	tmpl, _ := template.ParseFiles("templates/add.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", guiData)
	tmpl.ExecuteTemplate(w, "add", guiData)
}