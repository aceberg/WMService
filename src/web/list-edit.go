package web

import (
	"github.com/aceberg/WMService/db"
	. "github.com/aceberg/WMService/models"
	"html/template"
	"net/http"
	"log"
	"os"
	"io"
	"strings"
)

func listEdit(w http.ResponseWriter, r *http.Request) {
	var guiData GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	guiData.ListName = r.FormValue("listname")
	guiData.Mark = db.SelectList(AppConfig.DbPath, guiData.ListName)

	tmpl, _ := template.ParseFiles("templates/list.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "header", guiData)
	tmpl.ExecuteTemplate(w, "list", guiData)
}

func listAdd(w http.ResponseWriter, r *http.Request) {

	list := r.FormValue("listname")
	name := r.FormValue("name")

	db.ListInsert(AppConfig.DbPath, list, name)

	address := "/list_edit/?listname=" + list

	http.Redirect(w, r, address, 302)
}

func listDel(w http.ResponseWriter, r *http.Request) {

	list := r.FormValue("listname")
	name := r.FormValue("name")

	db.ListDelete(AppConfig.DbPath, list, name)

	address := "/list_edit/?listname=" + list

	http.Redirect(w, r, address, 302)
}

func listUpload(w http.ResponseWriter, r *http.Request) {

	list := r.FormValue("listname")

	uploadFile, _, err := r.FormFile("listfile")
	if err != nil {
		log.Println("ERROR: Upload error:", err)
	} else {
		defer uploadFile.Close()

		tmppath := "/tmp/" + list
		newFile, _ := os.Create(tmppath)
		defer newFile.Close()

		io.Copy(newFile, uploadFile)

		file, _ := os.ReadFile(tmppath)
		fileString := string(file)

		perString := strings.Split(fileString, "\n")
		for _, str := range perString {
			db.ListInsert(AppConfig.DbPath, list, str)
		}

		log.Println("INFO: list file uploaded to", list)
	}

	address := "/list_edit/?listname=" + list

	http.Redirect(w, r, address, 302)
}