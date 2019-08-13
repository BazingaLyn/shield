package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type conf struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}

func main() {

	var c conf
	currentConf := c.getConf()
	fmt.Println(currentConf)

	bytes, err := json.Marshal(currentConf)

	if err != nil {
		fmt.Println("err:\t", err.Error())
		return
	}

	//最终以json格式，输出
	fmt.Println("data:\t", string(bytes))
}

func (c *conf) getConf() *conf {

	yamlFile, error := ioutil.ReadFile("/Users/liguolin/go/src/shield/yaml/test.yaml")
	if error != nil {
		fmt.Println(error.Error())
	}

	error = yaml.Unmarshal(yamlFile, c)

	if error != nil {
		fmt.Println(error.Error())
	}

	return c

}
