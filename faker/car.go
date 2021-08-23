package faker

import (
	"github.com/91go/faker/utils"
)

func (g *Generator) Car() string {
	return utils.GetRandValue([]string{"carbrand", g.Locale_})
}
