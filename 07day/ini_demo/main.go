package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// MysqlConfig mysql 配置文件
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Database string `ini:"database"`
	Password string `ini:"password"`
	Test     bool   `init:"test"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 1.参数校验
	// 1.1传进来的data参数必须是指针，因为需要在函数对其赋值
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data params should be a pointer.")
		return
	}
	// 1.2传进来的data参数必须是结构体类型指针，因为配置文件中需要各种键值对需哟啊赋值给结构体的字段
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data params should be a struct pointer.")
		return
	}
	// 2.读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	// string(b) // 将字节类型的文件内容转换成字符串
	// 将文件内容按照字符串切割
	lineSlice := strings.Split(string(b), "\n")
	fmt.Printf("%#v\n", lineSlice)
	// 3.一行一行读取数据
	var structName string
	for idx, line := range lineSlice {
		// 3.1如果是注释就跳过
		// 去掉字符串的首尾空格
		line = strings.TrimSpace(line)
		// 如果是空行就直接跳过
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 3.2如果是[开头就表示节 section
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line: %d syntax error", idx)
				return
			}
			// 把这一行首尾的[]去掉，取到中间的内容把首尾的空格去掉拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line: %d syntax error", idx)
				return
			}

			// 根据字符串sectionName 去data里面根据反射找到对应的结构体

			for i := 0; i < t.Elem().NumField(); i++ {
				filed := t.Elem().Field(i)
				if sectionName == filed.Tag.Get("ini") {
					// 说明找到了对应的嵌套结构体， 把字段名记录下来
					structName = filed.Name
					fmt.Printf("找到了 %s 对应的嵌套结构体 %s \n", sectionName, structName)
					break
				}
			}
		} else {
			// 3.3如果不是[开头就是=分割的键值对
			// 1、= 分割这一行，左边是key，右边是 value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") || strings.HasSuffix(line, "=") {
				err = fmt.Errorf("line:%d, syntas error.", idx)
				return
			}
			subIndex := strings.Index(line, "=")
			key := strings.TrimSpace(line[:subIndex])
			value := strings.TrimSpace(line[(subIndex + 1):])
			fmt.Printf("key: %s, value: %s \n", key, value)
			// 2、根据structName 去 data里面把对应的嵌套结构体取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到结构体的值信息
			sType := sValue.Type()                     // 拿到结构体的类型信息

			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data 中的 %s 字段应该是一个结构体", structName)
				return
			}
			// 3、遍历嵌套结构体中每一个字段，判断tag是不是等于key
			var filedName string
			var filedType reflect.StructField
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i)
				filedType = filed
				if filed.Tag.Get("ini") == key {
					// 找到了结构体中对应字段
					filedName = filed.Name
					break
				}

			}
			// 4、如果key == tag，给这个字段赋值
			// 根据 filedName 取出这个字段进行复制
			if len(filedName) == 0 {
				// 在结构体中找不到对应字段
				continue
			}
			fileObj := sValue.FieldByName(filedName)
			fmt.Println(filedName, filedType.Type.Kind())
			switch filedType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line: %d value type  error.", idx+1)
					return
				}
				fileObj.SetInt(valueInt)

			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line: %d value type  error.", idx+1)
					return
				}
				fileObj.SetBool(valueBool)

			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line: %d value type  error.", idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)

			}
		}

	}
	return
}

func main() {
	var cfg Config
	// var x = new(int)
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Println("load ini file failed, err:", err)
		return
	}
	fmt.Printf("%#v\n", cfg)
}
