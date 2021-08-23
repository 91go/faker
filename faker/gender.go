package faker

import (
	"github.com/91go/faker/utils"
)

func (g *Generator) Gender() string {
	return utils.GetRandValue([]string{"gender", g.Locale_})
}
