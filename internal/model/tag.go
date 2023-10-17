package model

type Tag struct {
	*Mdoel
	Name string `json:"name"`
	State uint8 `json:"state"`
}

func (t Tag) TableName () string {
	return "blog_tag"
}
