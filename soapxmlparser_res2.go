package main

import (
	"fmt"
	"os"

	mxj "github.com/clbanning/mxj/v2"
)

func main() {
	fmt.Println("DECODING XML...")
	xmlFile, err := os.Open("Res2.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	mv, err := mxj.NewMapXmlReader(xmlFile)
	if err != nil {
		fmt.Println(err)
	}

	paths := mv.LeafPaths()
	var resXML interface{}
	for _, path := range paths {
		if path == "Envelope.Body.inquireUnbilledCreditCardTransactionsResponse.return.responseXML" {
			resXML, _ = mv.ValueForPath(path)
		} else {
			value, _ := mv.ValueForPath(path)
			fmt.Println(path, " has value ", value)
		}
	}
	//fmt.Println(resXML.(string))
	fmt.Println("Response XML is being parsed")
	mvRes, _ := mxj.NewMapXml([]byte(resXML.(string)))
	paths = mvRes.LeafPaths()
	for _, path := range paths {
		value, _ := mvRes.ValueForPath(path)
		fmt.Println(path, " has value ", value)
	}
}
