package faker

import (
	"github.com/91go/faker/utils"
)

func (g *Generator) AreaCode() string {
	return utils.GetRandValue([]string{"areacode", g.Locale_})
}
