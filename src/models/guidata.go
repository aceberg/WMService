package models

type GuiData struct {
	Config   Conf
	Heads    []string
	Icon     string
	ItemList []Item
	OneItem  Item
	Sum      int
	Len		 int
	Themes   []string
}
