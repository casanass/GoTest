package bo

import "github.com/jinzhu/gorm"

type deviceInfo struct {
	gorm.Model
	num  int    `gorm:"AUTO_INCREMENT"`
	name string `gorm:"unique;not null"`
	prio int8   `gorm:"not null"`
}
