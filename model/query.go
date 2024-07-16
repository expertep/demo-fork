package model

type Query struct {
	Limit  int64    `form:"limit" default:10`
	Page   int64    `form:"page" default:20`
	Sort   string   `form:"sort"`
	Order  int      `form:"order"`
	Search string   `form:"search"`
	Name   string   `form:"name"`
	Range  []string `form:"range[]"`
}

type QueryTime struct {
	CreateAt string   `form:"create_at"`
	Range    []string `form:"range[]"`
	Sort     string   `form:"sort"`
	Order    int      `form:"order"`
	Page     int64    `form:"page" default:20`
	Limit    int64    `form:"limit" default:10`
}
type QueryS3 struct {
	Limit                 int64  `form:"limit" default:10`
	Sort                  string `form:"sort"`
	Order                 int    `form:"order"`
	NextContinuationToken string `form:"next_continuation_token"`
}
