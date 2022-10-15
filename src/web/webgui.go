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
	http.ListenAndServe(address, nil)
}
