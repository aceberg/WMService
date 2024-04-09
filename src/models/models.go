package models

type Conf struct {
	DbPath  string
	GuiIP   string
	GuiPort string
	Theme   string
	Show    string
}

type Item struct {
	Id       int    `db:"_id"`
	Date     string `db:"DATE"`
	Mark     string `db:"MARK"`
	Model    string `db:"MODEL"`
	Trouble  string `db:"TROUBLE"`
	Trouble1 string `db:"TROUBLE1"`
	Street   string `db:"STREET"`
	House    string `db:"HOUSE"`
	Flat     string `db:"FLAT"`
	Phone    string `db:"PHONE"`
	Other    string `db:"OTHER"`
	Repair   string `db:"REPAIR"`
	Repair1  string `db:"REPAIR1"`
	War      string `db:"WAR"`
	Sum      string `db:"SUM"`
	Note     string `db:"NOTE"`
}

type GuiData struct {
	Config   Conf
	Heads    []string
	Icon     string
	ItemList []Item
	OneItem  Item
	Sum      int
	Len      int
	Themes   []string
	Mark     []string
	ListName string
}
