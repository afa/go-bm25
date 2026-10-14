package documents

import (
	"app/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DocumentResult struct {
	Id          int64  `json:"id"`
	ExternalKey string `json:"external_key"`
	State       string `json:"state"`
	Corpus      string `json:"corpus"`
	Text        string `json:"text"`
}

func List(db *gorm.DB) (list []DocumentResult, err error) {
	list = make([]DocumentResult, 0)
	_ = db.Model(&models.Document{}).Order("id asc").Find(&list)
	return
}

func ByKey(db *gorm.DB, key string) (item DocumentResult, err error) {
	_ = db.Model(&models.Document{}).Where("external_key = ?", key).First(&item)
	return
}

func Create(db *gorm.DB, params map[string]interface{}) (item DocumentResult, err error) {
	rez := db.Model(models.Document{}).Clauses(clause.Returning{}).Create(&params)
	if rez.Error != nil {
		return DocumentResult{}, rez.Error
	}
	rez = db.Model(models.Document{}).Where("id = ?", params["id"]).Take(&item)
	return item, rez.Error
}
