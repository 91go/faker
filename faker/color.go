package faker

import (
	"github.com/91go/faker/utils"
)

func (g *Generator) Color() string {
	//return g.NewGenerator("color")
	return utils.GetRandValue([]string{"color", g.Locale_})
}
