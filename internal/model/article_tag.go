package model

type ArticleTag struct {
	*Mdoel
	TagId     uint32 `json:"tag_id"`
	ArticleId uint32 `json:"article_id"`
	State     uint8  `json:"state"`
}

func (AT ArticleTag) TableName() string {
	return "blog_article_tag"
}
