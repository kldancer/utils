package dataStructure

import (
	"fmt"
	"math"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"k8s.io/klog/v2"
)

func RegexpString(input string) {
	// 提取数字部分和带宽单位
	re := regexp.MustCompile(`(\d+(\.\d+)?)(\s*[MG])bits/sec`)
	matches := re.FindStringSubmatch(input)
	if len(matches) > 3 {
		numStr := matches[1]
		unit := matches[3]
		fmt.Println("Parsed Number:", numStr, unit)
	} else {
		fmt.Println("Failed to extract the number.")
	}
}

func Str() {
	str := make([]string, 10)
	fmt.Println(len(str))
}

func RegexpString2(input string) {
	str := "eth1"
	var err error
	var targetIndex int
	if match := regexp.MustCompile(`^eth(\d+)$`).FindStringSubmatch(str); len(match) == 2 {
		targetIndex, err = strconv.Atoi(match[1])
		if err != nil {
			fmt.Println("无法将字符串转换为数字")
			return
		}
	} else {
		fmt.Println("未找到目标数字")
	}
	fmt.Println("目标数字int是:", targetIndex)
}

func RegexpString3(input string) {
	lines := strings.Split(input, "\n")
	re := regexp.MustCompile(`\d+(\.\d+)?\s*Mbits/sec`)
	match := re.FindStringSubmatch(lines[len(lines)-1])
	numStr := ""
	if len(match) > 0 {
		numStr = strings.TrimRight(match[0], " Mbits/sec")
	} else {
		fmt.Println("未找到匹配项")
	}
	num, _ := strconv.ParseFloat(numStr, 64)
	log.Infof("num = %v", num)
	fmt.Printf("num = %v", num)
}

func RegexpString4(input string) {
	lines := strings.Split(input, "\n")
	re := strings.Split(lines[len(lines)-3], " ")
	st := re[len(re)-4]
	lostNum := strings.Split(st, "/")[0]
	sumNUum := strings.Split(st, "/")[1]
	if len(lostNum) > 0 && len(sumNUum) > 0 {
		lost, _ := strconv.ParseFloat(lostNum, 64)
		sum, _ := strconv.ParseFloat(sumNUum, 64)
		lossRate := (lost / sum) * 100
		fmt.Printf("丢包率：%.2f%%", lossRate)
	} else {
		fmt.Println("未找到丢失或总数据包数")
	}
}

func timeStr() {
	RFC3339Nano := "2006-01-02T15:04:05.999999999Z07:00"
	str := "2024-03-06T22:42:35Z"
	if after, err := time.Parse(RFC3339Nano, str); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("after: %v\n", after)
	}

}

func parseTimestamp(timestamp interface{}) string {
	if stringTimestamp, ok := timestamp.(string); ok {

		if stringTimestamp == "" {
			klog.Warningln("Timestamp is empty")
			return ""
		} else if _, err := time.Parse(time.RFC3339Nano, stringTimestamp); err != nil {
			klog.Warningf("Failed to parse timestamp since %s is not RFC3339Nano format", stringTimestamp)
			return ""
		}
		return stringTimestamp

	} else {

		floatTimestamp, ok := timestamp.(float64)
		if !ok {
			klog.Warningf("Failed to parse timestamp since the type of %v is neither string nor float64", timestamp)
			return ""
		}

		stringTimestamp = strconv.FormatFloat(floatTimestamp, 'f', -1, 64)
		t := strings.Split(stringTimestamp, ".")

		sec, err := strconv.ParseInt(t[0], 10, 64)
		if err != nil {
			klog.Warningf("Failed to parse timestamp; %v", err)
			return ""
		}

		var nanoSec int64 = 0
		if len(t) == 2 {
			nanoSec, err = strconv.ParseInt(t[1], 10, 64)
			if err != nil {
				klog.Warningf("Failed to parse timestamp; %v", err)
				return ""
			}
		}

		fmt.Printf("result = %v\n", time.Unix(sec, nanoSec).UTC().Format(time.RFC3339Nano))
		return time.Unix(sec, nanoSec).UTC().Format(time.RFC3339Nano)
	}
}

func regexpUse() {
	filter := `([\w|-]+)\s*=\s*([+-]?\d*(\.\d+)?([Ee][+-]?\d+)?)`
	logTesta := `2024-03-14T02:22:43Z INFO     Epoch[4] Accuracy=0.85000 `
	logTestb := `2024-03-14T02:22:41Z INFO     Epoch[4] loss:0.402045  [57664/60000]`
	reg, _ := regexp.Compile(filter)
	matchStringsa := reg.FindAllStringSubmatch(logTesta, -1)
	fmt.Printf("matchStringsa = %+v\n", matchStringsa)

	matchStringsb := reg.FindAllStringSubmatch(logTestb, -1)
	fmt.Printf("matchStringsb = %+v\n", matchStringsb)
}

func splitNUse() {
	timestamp := time.Time{}.UTC().Format(time.RFC3339)
	fmt.Printf("timestamp = %+v\n", timestamp)
	logline := "2024-03-14T02:22:43Z INFO     Epoch[4] Accuracy=0.85000 1"
	ls := strings.SplitN(logline, " ", 2)

	timestamp = ls[0]
	fmt.Printf("ls= %+v\n", ls)
	fmt.Printf("timestamp = %+v\n", timestamp)
}

func timeUse() {

	now := time.Now()
	str := now.Format(time.DateTime)

	a, _ := time.Parse(time.DateTime, str)
	fmt.Printf("a = %+v\n", a)

	fmt.Printf("now = %+v\n", now.Format(time.DateTime))

	timestamp := time.Time{}.UTC().Format(time.RFC3339)
	fmt.Printf("timestamp = %+v\n", timestamp)

	b, _ := path.Split("/var/test/")
	fmt.Printf("b = %+v\n", b)

}

func GetDeviceTypeByChipName(chipName string) string {

	const (
		// Ascend310 ascend 310 chip
		Ascend310 = "Ascend310"
		// Ascend310B ascend 310B chip
		Ascend310B = "Ascend310B"
		// Ascend310P ascend 310P chip
		Ascend310P = "Ascend310P"
		// Ascend910 ascend 910 chip
		Ascend910 = "Ascend910"
		// Ascend910B ascend 1980B(910B) chip
		Ascend910B = "Ascend910B"

		// Pattern1980B regular expression for 1980B
		Pattern1980B = `^910B\d{1}`
		// Pattern1980 regular expression for 1980
		Pattern1980 = `^910B?`
	)
	if strings.Contains(chipName, "310P") {
		return Ascend310P
	}
	if strings.Contains(chipName, "310B") {
		return Ascend310B
	}
	if strings.Contains(chipName, "310") {
		return Ascend310
	}
	reg910B := regexp.MustCompile(Pattern1980B)
	if reg910B.MatchString(chipName) {
		return Ascend910B
	}
	reg910A := regexp.MustCompile(Pattern1980)
	if reg910A.MatchString(chipName) {
		return Ascend910
	}
	return ""
}

func removeSubstrings(input string) string {
	num, _ := strconv.Atoi(strings.TrimPrefix(strings.Split(input, "_")[0], "vir"))
	output := strconv.Itoa(num) + "c"

	if strings.Contains(input, "_3c_ndvpp") {
		return output + ".3cpu.ndvpp"
	}
	if strings.Contains(input, "_4c_dvpp") {
		return output + ".4cpu.dvpp"
	}

	if strings.Contains(input, "_3c") {
		return output + ".3cpu"
	}

	if strings.Contains(input, "_1c") {
		return output + ".1cpu"
	}
	return output
}

func str2() {
	parts := strings.Split("00000000:1A:00.0", ":")
	if len(parts) >= 2 {
		subBusID := parts[1:]
		result := strings.Join(subBusID, ":")
		fmt.Printf("subBusID: %s\n", subBusID)
		fmt.Printf("result: %s\n", result)
	}
}

func str3(a, money int) {

	a1 := money*6 + 1000
	a2 := (money+100)*6 + 1000
	a3 := money*7 + 1000
	fmt.Printf("一年合同住一年:%d,节约%d元\n", a1, a-a1)
	fmt.Printf("半年合同住半年:%d,节约%d元\n", a2, a-a2)
	fmt.Printf("一年合同住半年:%d,节约%d元\n", a3, a-a3)
}

func formatString(labels map[string]string, labelName string) string {
	labelValue, ok := labels[labelName]
	if !ok {
		labelValue = ""
	}

	return strings.ReplaceAll(labelValue, "\"", "")
}

func formatPciBus(pciBusId string) string {
	parts := strings.Split(pciBusId, ":")
	result := pciBusId
	if len(parts) >= 2 {
		subBusID := parts[1:]
		result = strings.Join(subBusID, ":")
	}

	return strings.ToLower(result)
}

func atoi(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}

	a := math.Round(i*100) / 100
	fmt.Printf("a = %+v\n", a)

	return i
}

func timet() {
	//
	//currentTime, _ := time.Parse(time.RFC3339Nano, "2024-08-27 10:45:49.220641405 +0000 UTC m=+28884.222872387")
	currentTime := time.Now()
	createTime := "2024-08-26 16:30:40 +0000 UTC"
	parsedTime, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", createTime)
	if err != nil {
		klog.Errorf("parse createTime err: %v", err)
	}

	duration := currentTime.Sub(parsedTime)
	fmt.Sprintf("%d days", duration/time.Hour/24)
}

func ConstructVNPUyTpe(templateName string) string {
	num, _ := strconv.Atoi(strings.TrimPrefix(strings.Split(templateName, "_")[0], "vir"))
	output := strconv.Itoa(num) + "c"

	if strings.Contains(templateName, "_3c_ndvpp") {
		return output + ".3cpu.ndvpp"
	}
	if strings.Contains(templateName, "_4c_dvpp") {
		return output + ".4cpu.dvpp"
	}

	if strings.Contains(templateName, "_3c") {
		output += ".3cpu"

		if strings.Contains(templateName, "g") {
			numIndex := strings.Index(templateName, "g")
			lastIndex := strings.LastIndex(templateName, "_") + 1
			output += fmt.Sprintf("." + templateName[lastIndex:numIndex] + "g")
		}
		return output
	}

	if strings.Contains(templateName, "_1c") {
		output += ".1cpu"

		if strings.Contains(templateName, "g") {
			numIndex := strings.Index(templateName, "g")
			lastIndex := strings.LastIndex(templateName, "_") + 1
			output += fmt.Sprintf("." + templateName[lastIndex:numIndex] + "g")
		}
		return output
	}
	return output
}
