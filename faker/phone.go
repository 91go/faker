package faker

import (
	"fmt"
	"github.com/91go/faker/utils"
	"strconv"
	"strings"
)

// 移动电话号码，中文会有前缀，暂时只支持中国的
func (g *Generator) MobilePhone() string {

	lang := g.Locale_
	var phoneNumber string
	if lang == "zh_CN" {
		countryCode := utils.GetRandValue([]string{"phone", lang + "_country_code"})
		areaPrefix := fmt.Sprintf(utils.GetRandValue([]string{"phone", "prefix_format_country"}), countryCode)
		phonePrefix := utils.GetRandValue([]string{"phone", lang + "_mobile"})
		phone8Num := fmt.Sprintf("%08d", utils.Number(1, 99999999))
		phoneNumber = areaPrefix + phonePrefix + phone8Num
	}
	return phoneNumber
}

// 中国长途区号
func (g *Generator) CityCode() string {
	cityCode := utils.GetRandValue([]string{"phone", "zh_CN_citycode"})
	return cityCode
}

// 固定电话号码，暂时只支持中国的（不包括特殊号码）
func (g *Generator) TelPhone() string {

	lang := g.Locale_
	var telPhone string
	if lang == "zh_CN" {
		cityCode := fmt.Sprintf(utils.GetRandValue([]string{"phone", "prefix_format_city"}), g.CityCode()) //随机拼接区号
		phone8Num := fmt.Sprintf("%08d", utils.Number(1, 99999999))
		telPhone = cityCode + phone8Num
	}
	return telPhone
}

// 特殊号码
func (g *Generator) SpecialTellPhone() string {
	speNumber := utils.GetRandValue([]string{"phone", "special"})
	return speNumber
}

// IMSI是15位的十进制数, 结构为: MCC + MNC + MSIN
// 国内为460开头、<=15、纯数字（国内一般为15位，国际规范为不超过15位）
// 暂时只做国内的
func (g *Generator) Imsi() string {
	mcc := "460" // 中国为460
	mncSlice := []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "20"}
	mnc := utils.RandString(mncSlice)
	msin := fmt.Sprintf("%010d", utils.Number(1, 9999999999))
	return mcc + mnc + msin
}

// IMEI（GSM网络）即EQUIREMENT_ID，15-17位、纯数字
func (g *Generator) Imei() string {
	// IMEI = TAC [+FAC] + SNR + CD [+SVN]
	// TAC: 8位数字（早起是6位），中国的前两位是86
	// FAC: 2位数字，仅在早起TAC为6位的手机中存在
	// SNR: 由第9位数字开始的6位数字组成
	// CD:  验证码，由前14位数字通过Luhn算法得出
	// SVN: 软件版本号，仅在部分机型中存在
	tacPrefixSlice := []string{"86", "35", "01"}
	tacPrefix := utils.RandString(tacPrefixSlice)
	tacSuffix := fmt.Sprintf("%06d", utils.Number(1, 999999))
	tac := tacPrefix + tacSuffix
	snr := fmt.Sprintf("%06d", utils.Number(1, 999999))
	preNumStr := tac + snr
	preNumSlice := []int{}
	for _, v := range preNumStr {
		temp, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		preNumSlice = append(preNumSlice, temp)
	}
	checkNum := g.Luhn(preNumSlice)
	checkNumStr := strconv.Itoa(checkNum)
	return preNumStr + checkNumStr
}

// MEID（CDMA网络）：固定14位，16进制，数字+[A-F]或[a-f]，例子：A000008C057C13
// letterType=true返回大写字母, false返回小写字母
func (g *Generator) Meid(letterType bool) string {
	// MEID = RR + XXXXXX + ZZZZZZ + C/CD
	// RR：    范围A0-FF，由官方分配，对应的十进制范围为[160, 255]
	// XXXXXX: 范围000000-FFFFFF，由官方分配，对应的十进制范围为[0, 16777215]
	// ZZZZZZ: 范围000000-FFFFFF，厂商分配给每台终端的流水号，对应的十进制范围为[0, 16777215]
	// C/CD:   0-F，校验码，不参与空中传输（所以暂时忽略）
	rr := fmt.Sprintf("%02x", utils.Number(160, 255))
	xx := fmt.Sprintf("%06x", utils.Number(0, 16777215))
	zz := fmt.Sprintf("%06x", utils.Number(0, 16777215))
	meid := rr + xx + zz
	if !letterType {
		meid = strings.ToUpper(meid)
	}
	return meid
}

// 随机返回大小写，多种分隔符的meid地址
func (g *Generator) RandMeid() string {
	lowUp := utils.RandBool([]bool{true, false})
	return g.Meid(lowUp)
}

// Luhn算法
func (g *Generator) Luhn(preNumArr []int) int {
	total := 0
	temp := 0
	preNumArr = append(preNumArr, 0) // 补充校验数字占位
	length := len(preNumArr)

	for i := length - 1; i > -1; i-- {
		if i%2 != 0 { // 原始数组的第奇数个乘以二，若新数字大于9则减去9
			temp = preNumArr[i] * 2
			if temp > 9 {
				temp = temp - 9
			}
		} else {
			temp = preNumArr[i]
		}
		total += temp
	}
	total = total * 9
	totalStr := strconv.Itoa(total)
	checkNum, _ := strconv.Atoi(totalStr[len(totalStr)-1:])
	return checkNum
}
