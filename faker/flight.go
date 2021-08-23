package faker

import (
	"fmt"
	"github.com/91go/faker/utils"
	"math/rand"
	"strings"
)

func (g *Generator) Voyage() string {
	var random string
	n := rand.Intn(10)
	if n%2 == 0 {
		random = fmt.Sprintf("%04d", utils.Number(1000, 9999))
	} else {
		random = fmt.Sprintf("%03d", utils.Number(100, 999))
	}
	voyage_prefix := utils.GetRandValue([]string{"flight", "airline_code"})
	voyage := voyage_prefix + random
	return voyage
}

func (g *Generator) AirlineName() string {
	airline := utils.GetRandValue([]string{"flight", "airline_name"})
	return airline
}

func (g *Generator) AirlineInfo() map[string]string {
	airlineInfo := utils.GetRandValue([]string{"flight", "airline_info"})
	splitAirSlice := strings.Split(airlineInfo, ",")
	code := strings.Split(splitAirSlice[0], "=")[1]
	name := strings.Split(splitAirSlice[1], "=")[1]
	return map[string]string{"code": code, "name": name}
}
