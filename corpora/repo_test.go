package corpora

import (
	"app/models"
	"fmt"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestList(t *testing.T) {
	db := initCorpora("host=localhost dbname=fw_test sslmode=disable")
	tx := setupCorpora(db)
	defer clean(tx)
	list, err := List(tx)
	if err != nil {
		fmt.Printf("%v", err)
		t.Fail()
	}
	if list[0].Name != "c1" {
		t.Fail()
	}
	if len(list) != 2 {
		t.Fail()
	}
}

func TestByKey(t *testing.T) {
	db := initCorpora("host=localhost dbname=fw_test sslmode=disable")
	tx := setupCorpora(db)
	defer clean(tx)
	item, err := ByKey(tx, "c2")
	if err != nil {
		t.Fail()
	}
	if item.Name != "c2" {
		t.Fail()
	}
}

func TestCreate(t *testing.T) {
	db := initCorpora("host=localhost dbname=fw_test sslmode=disable")
	tx := setupCorpora(db)
	defer clean(tx)
	item, err := Create(tx, map[string]interface{}{"external_key": "c3", "name": "c3"})
	if err != nil {
		t.Fail()
	}
	if item.Name != "c3" {
		t.Fail()
	}
	rez := tx.Model(&models.Corpus{}).Where("external_key = ?", "c3").First(&item)
	if rez.Error != nil {
		t.Fail()
	}
}

func initCorpora(dbName string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("db connect error")
	}
	db.AutoMigrate(&models.Corpus{})
	err = truncate(db)
	return db
}
func setupCorpora(db *gorm.DB) *gorm.DB {
	tx := db.Begin()
	corps := []*models.Corpus{
		{Name: "c1", ExternalKey: "c1", State: "init"},
		{Name: "c2", ExternalKey: "c2", State: "ready"},
	}
	db.Create(corps)
	return tx
}

func truncate(db *gorm.DB) (err error) {
	tx := db.Exec("truncate table corpus")
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func clean(tx *gorm.DB) {
	tx.Rollback()
	con, _ := tx.DB()
	con.Close()
	return
}
