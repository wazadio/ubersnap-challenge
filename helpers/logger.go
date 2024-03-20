package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strings"
)

const (
	ERROR_ROCOVER_INFROMATION string = "Unable to recover log information"
	INFO                      string = "INFO"
	ERROR                     string = "ERROR"
	reset                     string = "\033[0m"
	red                       string = "\033[31m"
	green                     string = "\033[32m"
)

func Log(level, data string) {
	var information strings.Builder
	pc, fileName, line, ok := runtime.Caller(1)
	if !ok {
		log.Println(ERROR_ROCOVER_INFROMATION)
	}

	if level == ERROR {
		information.WriteString(fmt.Sprintf("%s%s%s: ", red, level, reset))
	} else {
		information.WriteString(fmt.Sprintf("%s%s%s: ", green, level, reset))
	}

	informationDetail, _ := json.Marshal(map[string]any{
		"package.function_name": runtime.FuncForPC(pc).Name(),
		"file_ame":              fileName,
		"line":                  line,
		"data":                  data,
	})

	information.WriteString(string(informationDetail))
	log.Println(information.String())

}
