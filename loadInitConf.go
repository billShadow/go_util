package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)


func LoadIniConf(filename string)(data map[string]map[string]interface{}, err error){
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("open file err %s", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var conf_tag string
	data = make(map[string]map[string]interface{})
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if err != nil {
			fmt.Errorf("readstring err %s", err)
		}
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			conf_tag = line[1:len(line)-1]
		} else {
			//没有匹配判断是否存在xxx=xxx格式
			split := strings.Split(line, "=")
			item, ok := data[conf_tag]
			if !ok {
				item = make(map[string]interface{})
			}
			item[split[0]] = split[1]
			data[conf_tag] = item
		}
	}
	return
}

func main(){
	data, err := LoadIniConf("go_util/conf.ini")
	fmt.Println(data["redis"], err)
	result, err := json.Marshal(data)
	fmt.Println(string(result))

}