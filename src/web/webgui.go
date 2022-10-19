package web

import (
	// "fmt"
	"github.com/aceberg/WMService/db"
	. "github.com/aceberg/WMService/models"
	"log"
	"net/http"
)

var AppConfig Conf
var ItemList []Item

func Webgui(appConfig Conf) {

	AppConfig = appConfig
	ItemList = db.SelectAll(AppConfig.DbPath)

	address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add/", addHandler)
	http.HandleFunc("/backup/", backupHandler)
	http.HandleFunc("/config/", configHandler)
	http.HandleFunc("/save_config/", saveConfig)
	http.HandleFunc("/search/", searchHandler)
	http.HandleFunc("/upload/", uploadHandler)
	http.ListenAndServe(address, nil)
}
