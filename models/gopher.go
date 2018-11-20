package models

import (
	"time"
)

type Gopher struct {
	ID        uint64    `json:"id"`
	FullName  string    `json:"fullName"`
	Headline  string    `json:"headline"`
	AvatarURL string    `json:"avatarUrl"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (gopher *Gopher) FindById() error {
	return db.Where("id = ?", gopher.ID).First(&gopher).Error
}

func (gopher *Gopher) Create() error {
	return db.Create(&gopher).Error
}

func (gopher *Gopher) UpdateMeta() error {
	return db.Table("gopher").Where("id = ?", gopher.ID).Updates(map[string]interface{}{
		"full_name":  gopher.FullName,
		"headline":   gopher.Headline,
		"avatar_url": gopher.AvatarURL,
	}).Error
}
