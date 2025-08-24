package models

type Department struct {
	Cod    string `json:"dipCod" gorm:"primaryKey;size:10"`
	DipDes string `json:"dipDesc"`
}
