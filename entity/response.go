package entity

type Response struct {
	ID           uint64 `gorm:"primary_key;auto_increment"`
	RequestedURL string `gorm:"type:varchar(256);UNIQUE"`
	Body         string `gorm:"type:text"`
	StatusCode   int
}
