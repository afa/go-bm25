package documents

import (
	"app/models"
	"fmt"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestList(t *testing.T) {
	db := initDocs("host=localhost dbname=fw_test sslmode=disable")
	tx := setupDocs(db)
	defer clean(tx)
	list, err := List(tx)
	if err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	}
	// if list[0].Name != "c1" {
	// 	t.Fail()
	// }
	if len(list) != 2 {
		t.Fail()
	}
}

func TestByKey(t *testing.T) {
	db := initDocs("host=localhost dbname=fw_test sslmode=disable")
	tx := setupDocs(db)
	defer clean(tx)
	_, err := ByKey(tx, "c2")
	if err != nil {
		t.Fail()
	}
	// if item.Name != "c2" {
	// 	t.Fail()
	// }
}

func TestCreate(t *testing.T) {
	db := initDocs("host=localhost dbname=fw_test sslmode=disable")
	tx := setupDocs(db)
	defer clean(tx)
	item, err := Create(tx,
		map[string]interface{}{"external_key": "c3", "text": "c3 doc", "corpus": "c1"})
	if err != nil {
		t.Fail()
	}
	// if item.Name != "c3" {
	// 	t.Fail()
	// }
	rez := tx.Model(&models.Corpus{}).Where("external_key = ?", "c3").First(&item)
	if rez.Error != nil {
		t.Fail()
	}
}

func initDocs(dbName string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("db connect error")
	}
	db.AutoMigrate(&models.Corpus{})
	err = truncate(db)
	return db
}
func setupDocs(db *gorm.DB) *gorm.DB {
	tx := db.Begin()
	corpus := models.Corpus{ExternalKey: "c1", Name: "c1", State: "ready"}
	db.Create(corpus)
	docs := []*models.Document{
		{Corpus: "c1", Text: "idkfa iddqd", ExternalKey: "c1", State: "init"},
		{Corpus: "c1", Text: "other text", ExternalKey: "c2", State: "published"},
	}
	db.Create(docs)
	return tx
}

func truncate(db *gorm.DB) (err error) {
	tx := db.Exec("truncate table document")
	if tx.Error != nil {
		return tx.Error
	}
	return
}

func clean(tx *gorm.DB) {
	tx.Rollback()
	con, _ := tx.DB()
	con.Close()
	return
}
