package main

import (
	"fmt"
	"mygoedu/myconfigurator"
)

type Confstruct struct {
	TS      string  `xml:"testString" json:"testString" name:"testString"`
	TB      bool    `xml:"testBool" json:"testBool" name:"testBool"`
	TF      float64 `xml:"testFloat" json:"testFloat" name:"testFloat"`
	TestInt int
}

func main() {
	configstruct := new(Confstruct)
	//errConfig := myconfigurator.GetConfiguration(myconfigurator.CUSTOM, configstruct, "configfile.conf")
	errConfig := myconfigurator.GetConfiguration(myconfigurator.JSON, configstruct, "configfile.json")
	//errConfig := myconfigurator.GetConfiguration(myconfigurator.XML, configstruct, "configfile.xml")
	if errConfig != nil {
		fmt.Println(errConfig)
	}
	fmt.Println(*configstruct)

	if configstruct.TB {
		fmt.Println("bool is true")
	}

	fmt.Println(float64(4.8 * configstruct.TF))

	fmt.Println(5 * configstruct.TestInt)

	fmt.Println(configstruct.TS)
}
