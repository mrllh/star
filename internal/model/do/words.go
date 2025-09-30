// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Words is the golang structure of table words for DAO operations like Where/Data.
type Words struct {
	g.Meta             `orm:"table:words, do:true"`
	Id                 any         //
	Uid                any         //
	Word               any         //
	Definition         any         //
	ExampleSentence    any         //
	ChineseTranslation any         //
	Pronunciation      any         //
	ProficiencyLevel   any         //
	CreatedAt          *gtime.Time //
	UpdatedAt          *gtime.Time //
}
