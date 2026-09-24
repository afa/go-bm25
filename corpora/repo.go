package corpora

import (
	"app/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CorpusResult struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	ExternalKey string `json:"external_key"`
	State       string `json:"state"`
}

func List(db *gorm.DB) (list []CorpusResult, err error) {
	list = make([]CorpusResult, 0)
	_ = db.Model(&models.Corpus{}).Order("id asc").Find(&list)
	return
}

func ByKey(db *gorm.DB, key string) (item CorpusResult, err error) {
	_ = db.Model(&models.Corpus{}).Where("external_key = ?", key).First(&item)
	return
}

func Create(db *gorm.DB, params map[string]interface{}) (item CorpusResult, err error) {
	rez := db.Model(models.Corpus{}).Clauses(clause.Returning{}).Create(&params)
	if rez.Error != nil {
		return CorpusResult{}, rez.Error
	}
	rez = db.Model(models.Corpus{}).Where("id = ?", params["id"]).Take(&item)
	return item, rez.Error
}
