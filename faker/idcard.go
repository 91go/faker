package faker

import (
	"math/rand"
	"strconv"
)

func (g *Generator) IdCard() string {
	areaCode := g.getAreaCodeNumber()
	birthday := g.Birthday()
	rNum := getRandomNumber()
	sex := getGenderNumber("男")
	num17 := areaCode + birthday + rNum + sex
	validate := getValidateNumber(num17)
	idcard := num17 + validate
	return idcard
}

func (g *Generator) getAreaCodeNumber() string {
	// SELECT area_code, area_name FROM cnarea_2016 where area_code != '0' and level = 2 and name != '市辖区'
	areacode := g.AreaCode()
	return areacode
}

// 根据前17位数字获取校验码
func getValidateNumber(number string) string {
	weight := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}        // 十七位数字本体码权重
	validate := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"} // mod11,对应校验码字符值
	sum := 0
	for i := 0; i < len(number); i++ {
		temp, _ := strconv.Atoi(string(number[i]))
		temp = temp * weight[i]
		sum = sum + temp
	}
	mode := sum % 11
	return validate[mode]
}

func getGenderNumber(g string) string {
	var sex string
	var male = []string{"1", "3", "5", "7", "9"}
	var female = []string{"0", "2", "4", "6", "8", "10"}
	switch g {
	case "男":
		sex = male[rand.Intn(len(male))]
	case "女":
		sex = female[rand.Intn(len(female))]
	default:
		panic("性别错误, 未知性别")
	}
	return sex
}

func getRandomNumber() string {
	var number string
	first := []string{"0", "1"}
	second := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	number = first[rand.Intn(len(first))] + second[rand.Intn(len(second))]
	return number
}
