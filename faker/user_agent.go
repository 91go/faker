package faker

import (
	"github.com/91go/faker/utils"
	"strconv"
)

func (g *Generator) ChromeUA() string {
	randNum1 := strconv.Itoa(utils.RandIntRange(531, 536)) + strconv.Itoa(utils.RandIntRange(0, 2))
	randNum2 := strconv.Itoa(utils.RandIntRange(36, 40))
	randNum3 := strconv.Itoa(utils.RandIntRange(800, 899))
	return "Mozilla/5.0 " + "(" + randomPlatform() + ") AppleWebKit/" + randNum1 + " (KHTML, like Gecko) Chrome/" + randNum2 + ".0." + randNum3 + ".0 Mobile Safari/" + randNum1
}

func (g *Generator) FirefoxUA() string {
	ver := "Gecko/" + Date().Format("2006-02-01") + " Firefox/" + strconv.Itoa(utils.RandIntRange(35, 37)) + ".0"
	platforms := []string{
		"(" + windowsPlatformToken() + "; " + "en-US" + "; rv:1.9." + strconv.Itoa(utils.RandIntRange(0, 3)) + ".20) " + ver,
		"(" + linuxPlatformToken() + "; rv:" + strconv.Itoa(utils.RandIntRange(5, 8)) + ".0) " + ver,
		"(" + macPlatformToken() + " rv:" + strconv.Itoa(utils.RandIntRange(2, 7)) + ".0) " + ver,
	}

	return "Mozilla/5.0 " + utils.RandString(platforms)
}

func (g *Generator) SafariUA() string {
	randNum := strconv.Itoa(utils.RandIntRange(531, 536)) + "." + strconv.Itoa(utils.RandIntRange(1, 51)) + "." + strconv.Itoa(utils.RandIntRange(1, 8))
	ver := strconv.Itoa(utils.RandIntRange(4, 6)) + "." + strconv.Itoa(utils.RandIntRange(0, 2))

	mobileDevices := []string{
		"iPhone; CPU iPhone OS",
		"iPad; CPU OS",
	}

	platforms := []string{
		"(Windows; U; " + windowsPlatformToken() + ") AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + ver + " Safari/" + randNum,
		"(" + macPlatformToken() + " rv:" + strconv.Itoa(utils.RandIntRange(4, 7)) + ".0; en-US) AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + ver + " Safari/" + randNum,
		"(" + utils.RandString(mobileDevices) + " " + strconv.Itoa(utils.RandIntRange(7, 9)) + "_" + strconv.Itoa(utils.RandIntRange(0, 3)) + "_" + strconv.Itoa(utils.RandIntRange(1, 3)) + " like Mac OS X; " + "en-US" + ") AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + strconv.Itoa(utils.RandIntRange(3, 5)) + ".0.5 Mobile/8B" + strconv.Itoa(utils.RandIntRange(111, 120)) + " Safari/6" + randNum,
	}

	return "Mozilla/5.0 " + utils.RandString(platforms)
}

func (g *Generator) OperaUA() string {
	platform := "(" + randomPlatform() + "; en-US) Presto/2." + strconv.Itoa(utils.RandIntRange(8, 13)) + "." + strconv.Itoa(utils.RandIntRange(160, 355)) + " Version/" + strconv.Itoa(utils.RandIntRange(10, 13)) + ".00"

	return "Opera/" + strconv.Itoa(utils.RandIntRange(8, 10)) + "." + strconv.Itoa(utils.RandIntRange(10, 99)) + " " + platform
}

// linuxPlatformToken will generate a random linux platform
func linuxPlatformToken() string {
	return "X11; Linux " + utils.GetRandValue([]string{"computer", "linux_processor"})
}

// macPlatformToken will generate a random mac platform
func macPlatformToken() string {
	return "Macintosh; " + utils.GetRandValue([]string{"computer", "mac_processor"}) + " Mac OS X 10_" + strconv.Itoa(utils.RandIntRange(5, 9)) + "_" + strconv.Itoa(utils.RandIntRange(0, 10))
}

// windowsPlatformToken will generate a random windows platform
func windowsPlatformToken() string {
	return utils.GetRandValue([]string{"computer", "windows_platform"})
}

// randomPlatform will generate a random platform
func randomPlatform() string {
	platforms := []string{
		linuxPlatformToken(),
		macPlatformToken(),
		windowsPlatformToken(),
	}

	return utils.RandString(platforms)
}
