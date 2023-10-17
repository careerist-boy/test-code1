package model

type Mdoel struct {
	ID uint32 `gorm:"primary_key" json:"id"`
	CreateBy string `json:"create_by"`
	ModifyBy string `json:"modified_by"`
	CreateAt uint32 `json:"create_at"`
	ModifyAt uint32 `json:"modify_at"`
}
