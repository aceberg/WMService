package web

import (
	"github.com/aceberg/WMService/conf"
	. "github.com/aceberg/WMService/models"
	"html/template"
	"net/http"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "flatly", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	tmpl, _ := template.ParseFiles("templates/config.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", guiData)
	tmpl.ExecuteTemplate(w, "config", guiData)
}

func saveConfig(w http.ResponseWriter, r *http.Request) {

	AppConfig.Theme = r.FormValue("theme")
	AppConfig.Show = r.FormValue("show")
	conf.WriteConfig(AppConfig)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
