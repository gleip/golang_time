package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// Numbers Массив чисел в ASCII
var Numbers = [][]string{
	{"000000", "00  00", "00  00", "00  00", "000000"},
	{"1111  ", "  11  ", "  11  ", "  11  ", "111111"},
	{"222222", "     2", "222222", "2     ", "222222"},
	{"333333", "    33", "333333", "    33", "333333"},
	{"44  44", "44  44", "444444", "    44", "    44"},
	{"555555", "55    ", "555555", "    55", "555555"},
	{"666666", "66    ", "666666", "66  66", "666666"},
	{"777777", "    77", "    77", "    77", "    77"},
	{"888888", "88  88", "888888", "88  88", "888888"},
	{"999999", "99  99", "999999", "    99", "999999"},
}

func main() {
	for {
		var (
			hourNow    = ""
			minutsNow  = ""
			secondsNow = ""
			number     = ""
		)
		timeNow := time.Now()

		if timeNow.Hour() < 10 {
			hourNow = "0" + strconv.Itoa(timeNow.Hour())
		} else {
			hourNow = strconv.Itoa(timeNow.Hour())
		}

		if timeNow.Minute() < 10 {
			minutsNow = "0" + strconv.Itoa(timeNow.Minute())
		} else {
			minutsNow = strconv.Itoa(timeNow.Minute())
		}

		if timeNow.Second() < 10 {
			secondsNow = "0" + strconv.Itoa(timeNow.Second())
		} else {
			secondsNow = strconv.Itoa(timeNow.Second())
		}

		number = hourNow + minutsNow + secondsNow

		for key := range Numbers[0] {
			line := ""
			for column := range number {
				digit := number[column] - '0'
				if 0 <= digit && digit <= 9 {
					line += Numbers[digit][key]
					if column != 0 && column%2 > 0 && key == 1 {
						line += " ++ "
					} else {
						line += "    "
					}
				} else {
					log.Fatal("Неправильное число")
				}
			}
			fmt.Println(line)
		}
		time.Sleep(time.Second)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
