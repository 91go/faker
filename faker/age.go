package faker

import (
	"github.com/91go/faker/utils"
	"strconv"
)

func (g *Generator) Age() string {
	return strconv.Itoa(utils.Number(0, 105))
}
