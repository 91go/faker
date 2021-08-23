package faker

import (
	"github.com/91go/faker/utils"
	"strings"
)

// 机场的IATA三位编码
func (g *Generator) AirportIATA() string {
	code := utils.GetRandValue([]string{"airport", "airport_iata_code"})
	return code
}

// 机场的ICAO四位编码
func (g *Generator) AirportICAO() string {
	code := utils.GetRandValue([]string{"airport", "airport_icao_code"})
	return code
}

func (g *Generator) AirportCity() string {
	city := utils.GetRandValue([]string{"airport", "airport_city"})
	return city
}

func (g *Generator) AirportCityPy() string {
	cityPy := utils.GetRandValue([]string{"airport", "airport_city_pinyin"})
	return cityPy
}

func (g *Generator) AirportName() string {
	name := utils.GetRandValue([]string{"airport", "airport_name"})
	return name
}

func (g *Generator) AirportInfo() map[string]string {
	info := utils.GetRandValue([]string{"airport", "airport_info"})
	airPortMap := make(map[string]string)
	splitInfo := strings.Split(info, ",")
	for _, each := range splitInfo {
		splitEach := strings.Split(each, "=")
		key := splitEach[0]
		value := splitEach[1]
		airPortMap[key] = value
	}
	return airPortMap
}
