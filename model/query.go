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
