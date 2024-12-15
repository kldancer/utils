package dataStructure

import (
	"fmt"
	"testing"
)

func TestRegexpString(t *testing.T) {
	input := `Connecting to host 172.24.20.230, port 5201
[  4] local 172.24.20.230 port 45062 connected to 172.24.20.230 port 5201
[ ID] Interval           Transfer     Bandwidth       Retr  Cwnd
[  4]   0.00-1.00   sec  3.35 GBytes  28804 Mbits/sec    0   1.94 MBytes       
[  4]   1.00-2.00   sec  3.47 GBytes  29839 Mbits/sec    0   2.00 MBytes       
[  4]   2.00-3.00   sec  3.83 GBytes  32925 Mbits/sec    0   2.00 MBytes       
[  4]   3.00-4.00   sec  3.81 GBytes  32751 Mbits/sec    0   2.00 MBytes       
[  4]   4.00-5.00   sec  3.82 GBytes  32790 Mbits/sec    0   2.00 MBytes       
- - - - - - - - - - - - - - - - - - - - - - - - -
[ ID] Interval           Transfer     Bandwidth       Retr
[  4]   0.00-5.00   sec  18.3 GBytes  31422 Mbits/sec    0             sender
[  4]   0.00-5.00   sec  18.3 GBytes  31420 Mbits/sec                  receiver`

	RegexpString(input)

}

func TestStr(t *testing.T) {
	fmt.Println(len(make([]string, 10)))
}

func TestRegexpString2(t *testing.T) {
	RegexpString2("")
}

func TestRegexpString3(t *testing.T) {
	input := `Connecting to host 172.24.20.230, port 5201
[  4] local 172.24.20.230 port 45062 connected to 172.24.20.230 port 5201
[ ID] Interval           Transfer     Bandwidth       Retr  Cwnd
[  4]   0.00-1.00   sec  3.35 GBytes  28804 Mbits/sec    0   1.94 MBytes       
[  4]   1.00-2.00   sec  3.47 GBytes  29839 Mbits/sec    0   2.00 MBytes       
[  4]   2.00-3.00   sec  3.83 GBytes  32925 Mbits/sec    0   2.00 MBytes       
[  4]   3.00-4.00   sec  3.81 GBytes  32751 Mbits/sec    0   2.00 MBytes       
[  4]   4.00-5.00   sec  3.82 GBytes  32790 Mbits/sec    0   2.00 MBytes       
- - - - - - - - - - - - - - - - - - - - - - - - -
[ ID] Interval           Transfer     Bandwidth       Retr
[  4]   0.00-5.00   sec  18.3 GBytes  31422 Mbits/sec    0             sender
[  4]   0.00-5.00   sec  18.3 GBytes  31.420 Mbits/sec                  receiver`
	RegexpString3(input)
}

func TestRegexpString4(t *testing.T) {
	input := `[  5] local 172.28.8.121 port 56970 connected to 172.28.8.122 port 30200
[ ID] Interval           Transfer     Bitrate         Total Datagrams
[  5]   0.00-1.00   sec  81.7 MBytes   686 Mbits/sec  83690  
[  5]   1.00-2.00   sec   103 MBytes   857 Mbits/sec  104980  
[  5]   2.00-3.00   sec  97.1 MBytes   816 Mbits/sec  99390  
[  5]   3.00-4.00   sec   104 MBytes   870 Mbits/sec  106080  
[  5]   4.00-5.00   sec   101 MBytes   845 Mbits/sec  103190  
- - - - - - - - - - - - - - - - - - - - - - - - -
[ ID] Interval           Transfer     Bitrate         Jitter    Lost/Total Datagrams
[  5]   0.00-5.00   sec   486 MBytes   815 Mbits/sec  0.000 ms  0/497330 (0%)  sender
[  5]   0.00-5.04   sec   453 MBytes   753 Mbits/sec  0.006 ms  33688/497329 (6.8%)  receiver

iperf Done.`
	RegexpString4(input)
}

func TestTime(t *testing.T) {
	timeStr()
}

func TestParseTimestamp(t *testing.T) {
	time := 1638422847.28721
	parseTimestamp(time)
}

func TestRegexpUse(t *testing.T) {
	regexpUse()
}

func TestSplitNUse(t *testing.T) {
	splitNUse()
}

func TestTimeUse(t *testing.T) {
	timeUse()
}

func TestRemoveSubstrings(t *testing.T) {
	// 示例字符串
	inputStr := []string{"vir16", "vir08", "vir04", "vir02", "vir01",
		"vir02_1c", "vir04_3c", "vir04_3c_ndvpp", "vir04_4c_dvpp"}

	// 调用函数并打印结果
	for _, s := range inputStr {
		result := removeSubstrings(s)
		fmt.Println(result)
	}
}

func TestGetDeviceTypeByChipName(t *testing.T) {
	a := GetDeviceTypeByChipName("910ProB")
	fmt.Printf("a:%s", a)
}

func TestStr2(t *testing.T) {
	str2()
}

func TestStr3(t *testing.T) {
	a := 2500*6 + 1500
	fmt.Printf("当前:%d ", a)

	moneys := []int{
		1200,
		1300,
		1400,
		1500,
		1600,
		1700,
	}

	for _, money := range moneys {
		fmt.Printf("-------\n")
		fmt.Printf("房租：%d\n", money)

		str3(a, money)
	}
}

func TestAtoi(t *testing.T) {
	atoi("32.8369140625")
}

func TestTimet(t *testing.T) {
	timet()
}

func TestConstructVNPUyTpe(t *testing.T) {
	// 示例字符串
	inputStr := []string{"vir16", "vir08", "vir04", "vir02", "vir01",
		"vir02_1c", "vir04_3c", "vir04_3c_ndvpp", "vir04_4c_dvpp"}

	// 调用函数并打印结果
	for _, s := range inputStr {
		result := ConstructVNPUyTpe(s)
		fmt.Println(result)
	}

	fmt.Println("---")
	inputStr2 := []string{"vir12_3c_32g", "vir06_1c_16g", "vir03_1c_8g"}

	for _, s := range inputStr2 {
		result := ConstructVNPUyTpe(s)
		fmt.Println(result)
	}
}

func TestIntersection(t *testing.T) {
	tests := [][]string{
		[]string{
			"eno1", "ens4f0", "ens4f1", "ens5f1",
		},
		[]string{
			"eno1", "ens4f0", "ens5f1",
		},
		[]string{
			"eno1", "ens4f0", "ens5f1",
		},
	}
	res := Intersection(tests...)
	fmt.Println(res)
}
