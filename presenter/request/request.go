package request

type PageOptions struct {
	Page   int64  `query:"page" json:"page"`
	Limit  int64  `query:"limit" json:"limit"`
	Sort   string `query:"sort" json:"sort"`
	Search string `query:"search" json:"search"`
}
