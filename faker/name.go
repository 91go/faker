package faker

import (
	"github.com/91go/faker/utils"
	"strings"
)

func (g *Generator) Name() string {
	lang := g.Locale_
	first := strings.Join([]string{lang, "first"}, "_")
	last := strings.Join([]string{lang, "last"}, "_")
	switch lang {
	case "en_US":
		return utils.GetRandValue([]string{"person", first}) + " " + utils.GetRandValue([]string{"person", last})
	case "zh_CN":
		return utils.GetRandValue([]string{"person", last}) + utils.GetRandValue([]string{"person", first})
	default:
		return "未知的语言类型[姓名报错]"
	}
}

func (g *Generator) FirstName() string {
	first := strings.Join([]string{g.Locale_, "first"}, "_")
	return utils.GetRandValue([]string{"person", first})
}

func (g *Generator) LastName() string {
	last := strings.Join([]string{g.Locale_, "last"}, "_")
	return utils.GetRandValue([]string{"person", last})
}
