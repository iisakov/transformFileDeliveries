package sdk

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
	"transformFileDeliveries/config"
)

var FileInPath, FileOutPath, PatternData, Sep string = "fileIn.txt", "fileOut.txt", "02.01.2006 15:04:05", ";"
var CountOrders = 1000

func ReadFile(file *os.File, sep string) (result map[string][]string) {
	result = make(map[string][]string)
	in := bufio.NewReader(file)
	for {
		line, err := in.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}

		sliceLine := strings.Split(line, sep)
		if _, ok := result[sliceLine[1]]; !ok {
			result[sliceLine[1]] = []string{sliceLine[0]}
		} else {
			result[sliceLine[1]] = append(result[sliceLine[1]], sliceLine[0])
		}
	}
	return
}

func WriteFile(file *os.File, patternData string, countOrders int, requestDateDeliveryShift map[string][]string) {
	out := bufio.NewWriter(file)
	defer out.Flush()

	for requestDate, orderNums := range requestDateDeliveryShift {
		t, _ := time.Parse(patternData, requestDate)
		fmt.Println(t, requestDate)
		file.WriteString(fmt.Sprintf("'%s'", t.Format("2006-01-02")))

		for i, orderNum := range orderNums {
			if i%countOrders == 0 {
				file.WriteString("\n(")
			}
			if i%countOrders == countOrders-1 || i == len(orderNums)-1 {
				file.WriteString("'" + orderNum + "') ")
			} else {
				file.WriteString("'" + orderNum + "', ")
			}
		}
	}
}

func setParam(comands []string) {
	for _, c := range comands {
		cS := strings.Split(c, ":")

		if len(cS) > 1 {
			switch cS[0] {
			case "fileInPath":
				fallthrough
			case "fip":
				FileInPath = cS[1]

			case "fileOutPath":
				fallthrough
			case "fop":
				FileOutPath = cS[1]

			case "patternData":
				fallthrough
			case "pd":
				PatternData = cS[1]

			case "sep":
				fallthrough
			case "s":
				PatternData = cS[1]

			case "countOrders":
				fallthrough
			case "co":
				cSS, err := strconv.Atoi(cS[1])
				if err != nil {
					PrintError("Кажется вы ввели не число.")
					Exit()
				}
				CountOrders = int(cSS)

			default:
				PrintMessage("Команды: " + c + " не существует.")
				Exit()
			}
		}
	}
}

func Handle(comands []string) {
	for _, c := range comands {
		switch c {

		case "--version":
			fallthrough
		case "-v":
			PrintMessage("transform_file [by_artisan] v:" + config.Version)
			Exit()

		case "--param":
			fallthrough
		case "-p":
			setParam(comands)
		default:
		}
	}
}

func PrintMessage(messages ...string) {
	response := ""
	for _, message := range messages {
		response += message + " "
	}
	fmt.Println(response)
}

func Exit(messages ...string) {
	response := ""
	for _, message := range messages {
		response += message + " "
	}
	response += "Спасибо, что воспользовались transform_file от [by_artisan]"

	fmt.Println(response)
	os.Exit(0)
}

func PrintError(messages ...string) {
	response := "\n"
	for _, message := range messages {
		response += message + "\n"
	}
	response += "Утилита закрыта."

	fmt.Println(response)
	time.Sleep(100 + time.Second)
}
