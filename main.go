package main

import (
	"fmt"
	"reflect"
)

// ConfigInfo 配置测试
type ConfigInfo struct {
	filePath string // 默认记录文件路径的信息
	Name     string
	Age      int
	Price    float32
}

func checkIDNotZero(id interface{}) bool {
	var pass bool
	switch v := id.(type) {
	case int:
		if v > 0 {
			pass = true
		}
	case int8:
		if v > 0 {
			pass = true
		}
	case int16:
		if v > 0 {
			pass = true
		}
	case int32:
		if v > 0 {
			pass = true
		}
	case int64:
		if v > 0 {
			pass = true
		}
	case uint, uint8, uint16, uint32, uint64:
		//fmt.Println(reflect.ValueOf(id).Type().Name())
		//fmt.Println(reflect.ValueOf(id).Uint())
		if reflect.ValueOf(v).IsZero() {
			pass = true
		}
	default:
		pass = false
	}

	return pass
}

func checkTest(id int) bool {
	var gap int
	gap = id - 10
	switch gap {
	case 50, 30, 20:
		fmt.Println("gap 50/30/20:", gap)
	case -1:
		fmt.Println("gap -1:", gap)
	default:
		fmt.Println("gap default:", gap)
	}
	return true
}

// Tencent 结构
type Tencent struct {
	Wish  int `json:"Wish"`
	All   int `json:"All"`
	Coder int `json:"Coder"`
	Be    int `json:"Be"`
	Happy int `json:"Happy"`
	Every int `json:"Every"`
	Day   int `json:"Day"`
}

func main() {
	/*
		var myCfg ConfigInfo
		err := config.LoadTomlConfig("./test.toml", &myCfg)
		if err != nil {
			fmt.Println("load config err: ", err.Error())
			os.Exit(1)
		}
		logger.InitConf("./", "TEST", logger.LOG_DEBUG, logger.LOG_ALTER_HOUR)
		logger.Println(logger.LOG_DEBUG, "my name is %s, i am %d years old", string("ericlang"), 20)
		fmt.Println(config.String(myCfg))

		tencent := Tencent{1, 0, 0, 2, 0, 0, 4}
		tencentTag := reflect.TypeOf(&tencent)
		tagCnt := tencentTag.Elem().NumField()
		for i := 0; i < tagCnt; i++ {
			fmt.Print(tencentTag.Elem().Field(i).Tag.Get("json"), " ")
		}
		fmt.Println("!")
	*/
	checkIDNotZero(uint(4294836225))
	checkIDNotZero(uint8(255))
	checkIDNotZero(uint16(65535))
	checkIDNotZero(uint32(4294836225))
	checkIDNotZero(uint32(0))
}
