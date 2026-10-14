package models

import (
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

type Corpus struct {
	gorm.Model
	Name        string
	ExternalKey string `gorm:"unique;uniqueIndex:corpora_external_key_uniq_idx;type:string"`
	State       string `gorm:"index:corpora_state_idx;type:string"`
}

type Document struct {
	gorm.Model
	ExternalKey string `gorm:"unique;uniqueIndex:document_external_key_uniq_idx;type:string"`
	Corpus      string `gorm:"index:document_corpus_state_idx;type:string"`
	Text        string
	State       string `gorm:"index:document_corpus_state_idx;type:string"`
}
