package faker

import (
	"github.com/91go/faker/utils"
)

func (g *Generator) Job() string {
	return utils.GetRandValue([]string{"job", g.Locale_})
}
