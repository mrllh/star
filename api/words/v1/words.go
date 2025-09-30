package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ProficiencyLevel uint

const (
	ProficiencyLevel1 ProficiencyLevel = iota + 1
	ProficiencyLevel2
	ProficiencyLevel3
	ProficiencyLevel4
	ProficiencyLevel5
)

type CreateReq struct {
	g.Meta             `path:"/words" method:"post" summary:"创建" tags:"单词"`
	Word               string           `json:"word" v:"required|length:1,100" dc:"单词"`
	Definition         string           `json:"definition" v:"required|length:1,300" dc:"释义"`
	ExampleSentence    string           `json:"example_sentence" v:"required|length:1,300" dc:"例句"`
	ChineseTranslation string           `json:"chinese_translation" v:"required|length:1,300" dc:"中文翻译"`
	Pronunciation      string           `json:"pronunciation" v:"required|length:1,100" dc:"发音"`
	ProficiencyLevel   ProficiencyLevel `json:"proficiency_level" v:"required|between:1,5" dc:"熟练度,1-5"`
}

type CreateRes struct {
}

type UpdateReq struct {
	g.Meta             `path:"/words/{id}" method:"put" summary:"更新" tags:"单词"`
	Id                 uint             `json:"id" v:"required"`
	Word               string           `json:"word" v:"required|length:1,100" dc:"单词"`
	Definition         string           `json:"definition" v:"required|length:1,300" dc:"释义"`
	ExampleSentence    string           `json:"example_sentence" v:"required|length:1,300" dc:"例句"`
	ChineseTranslation string           `json:"chinese_translation" v:"required|length:1,300" dc:"中文翻译"`
	Pronunciation      string           `json:"pronunciation" v:"required|length:1,100" dc:"发音"`
	ProficiencyLevel   ProficiencyLevel `json:"proficiency_level" v:"required|between:1,5" dc:"熟练度,1-5"`
}

type UpdateRes struct {
}

type List struct {
	Id               uint             `json:"id"`
	Word             string           `json:"word"`
	Definition       string           `json:"definition"`
	ProficiencyLevel ProficiencyLevel `json:"proficiency_level"`
}

type ListReq struct {
	g.Meta `path:"/words" method:"get" summary:"列表" tags:"单词"`
	Word   string `json:"word" v:"length:1,100" dc:"单词"`
	Page   int    `json:"page" v:"min:1" d:"1" dc:"页码"`
	Size   int    `json:"size" v:"between:1,100" dc:"每页数量"`
}

type ListRes struct {
	List  []List `json:"list"`
	Total uint   `json:"total"`
}

// words/{id}是一种模糊路由匹配方式，它会识别类似words/1，words/2，words/abc这种路由，并将匹配到的值赋予Id字段。
type DetailReq struct {
	g.Meta `path:"/words/{id}" method:"get" summary:"详情" tags:"单词"`
	Id     uint `json:"id" v:"required"`
}

type DetailRes struct {
	Id                 uint             `json:"id"`
	Word               string           `json:"word"`
	Definition         string           `json:"definition"`
	ExampleSentence    string           `json:"exampleSentence"`
	ChineseTranslation string           `json:"chineseTranslation"`
	Pronunciation      string           `json:"pronunciation"`
	ProficiencyLevel   ProficiencyLevel `json:"proficiencyLevel"`
	CreatedAt          *gtime.Time      `json:"createdAt"`
	UpdatedAt          *gtime.Time      `json:"updatedAt"`
}

type DeleteReq struct {
	g.Meta `path:"/words/{id}" method:"delete" summary:"删除" tags:"单词"`
	Id     uint `json:"id" v:"required"`
}

type DeleteRes struct {
}
