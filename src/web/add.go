package web

import (
	// "fmt"
	"github.com/aceberg/WMService/db"
	. "github.com/aceberg/WMService/models"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon
	guiData.Mark = db.SelectList(AppConfig.DbPath, "rs_mark")
	guiData.Themes = db.SelectList(AppConfig.DbPath, "rs_trouble")
	guiData.Heads = db.SelectList(AppConfig.DbPath, "rs_repair")

	currentTime := time.Now()
	guiData.OneItem.Date = currentTime.Format("2006-01-02")

	tmpl, _ := template.ParseFiles("templates/add.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", guiData)
	tmpl.ExecuteTemplate(w, "add", guiData)
}

func saveItem(w http.ResponseWriter, r *http.Request) {
	var item Item

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

	if item.Id == 0 {
		db.AddItem(AppConfig.DbPath, item)
	} else {
		db.UpdateItem(AppConfig.DbPath, item)
	}

	http.Redirect(w, r, "/", 302)
}
