package faker

import (
	"fmt"
	"github.com/91go/faker/utils"
	"math/rand"
)

func (g *Generator) Mail() string {
	var mail string
	splitString := []string{"_", ".", "-"}
	random := fmt.Sprintf("%04d", utils.Number(10, 3000))
	flag := rand.Intn(10)

	if i := flag % 2; i == 0 {
		mail = utils.GetRandValue([]string{"person", "en_US_first"}) + utils.GetRandValue([]string{"person", "en_US_last"}) + random
	} else {
		mail = utils.GetRandValue([]string{"person", "en_US_first"}) + splitString[rand.Intn(len(splitString))] + utils.GetRandValue([]string{"person", "en_US_last"}) + random
	}

	mail += utils.GetRandValue([]string{"email", "postfix"})

	return mail
}
