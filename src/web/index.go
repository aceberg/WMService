package web

import (
	// "fmt"
	"github.com/aceberg/WMService/db"
	. "github.com/aceberg/WMService/models"
	"html/template"
	"net/http"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	ItemList = db.SelectAll(AppConfig.DbPath)

	show, _ := strconv.Atoi(AppConfig.Show)
	length := len(ItemList)
	if show > length {
		show = length
	}

	guiData.Config = AppConfig
	guiData.Icon = Icon
	guiData.ItemList = ItemList[0:show]
	guiData.Heads = Heads
	guiData.Sum = getSum(ItemList[0:show])
	guiData.Len = length

	tmpl, _ := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", guiData)
	tmpl.ExecuteTemplate(w, "index", guiData)
}
