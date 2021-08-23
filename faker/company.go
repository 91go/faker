package faker

import (
	"github.com/91go/faker/utils"
)

func (g *Generator) Company() (company string) {
	switch randInt := utils.RandIntRange(1, 3); randInt {
	case 1:
		company = g.LastName() + ", " + g.LastName() + " and " + g.LastName()
	case 2:
		company = g.LastName() + "-" + g.LastName()
	case 3:
		company = g.LastName() + " " + CompanySuffix()
	}

	return
}

// CompanySuffix will generate a random company suffix string
func CompanySuffix() string {
	return utils.GetRandValue([]string{"company", "suffix"})
}

// BuzzWord will generate a random company buzz word string
func BuzzWord() string {
	return utils.GetRandValue([]string{"company", "buzzwords"})
}

// BS will generate a random company bs string
func BS() string {
	return utils.GetRandValue([]string{"company", "bs"})
}
