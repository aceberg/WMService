package web

import (
	"github.com/aceberg/WMService/db"
	. "github.com/aceberg/WMService/models"
	"html/template"
	"net/http"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {

	field1 := r.FormValue("field1")
	field2 := r.FormValue("field2")
	search1 := r.FormValue("search1")
	search2 := r.FormValue("search2")

	if search1 != "" {
		var guiData GuiData
		var foundList []Item

		ItemList = db.SelectAll(AppConfig.DbPath)

		foundList = searchTable(ItemList, field1, search1)

		if search2 != "" {
			foundList = searchTable(foundList, field2, search2)
		}

		guiData.Config = AppConfig
		guiData.Icon = Icon
		guiData.ItemList = foundList
		guiData.Heads = Heads
		guiData.Sum = getSum(foundList)
		guiData.Len = len(foundList)

		tmpl, _ := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
		tmpl.ExecuteTemplate(w, "header", guiData)
		tmpl.ExecuteTemplate(w, "index", guiData)
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
