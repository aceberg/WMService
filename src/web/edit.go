package web

import (
	. "github.com/aceberg/WMService/models"
	"html/template"
	"net/http"
	"strconv"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	var item Item
	var guiData GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	item.Id, _ = strconv.Atoi(r.FormValue("id"))
	item.Date = r.FormValue("date")
	item.Mark = r.FormValue("mark")
	item.Model = r.FormValue("model")
	item.Trouble = r.FormValue("trouble")
	item.Trouble1 = r.FormValue("trouble1")
	item.Street = r.FormValue("street")
	item.House = r.FormValue("house")
	item.Flat = r.FormValue("flat")
	item.Phone = r.FormValue("phone")
	item.Other = r.FormValue("other")
	item.Repair = r.FormValue("repair")
	item.Repair1 = r.FormValue("repair1")
	item.War = r.FormValue("war")
	item.Sum = r.FormValue("sum")
	item.Note = r.FormValue("note")

	guiData.OneItem = item

	tmpl, _ := template.ParseFiles("templates/add.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", guiData)
	tmpl.ExecuteTemplate(w, "add", guiData)
}
