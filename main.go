package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Envelope struct {
	Body Body `xml:"body"`
}
type Body struct {
	InquireUnbilledCreditCardTransactionsResponse InquireUnbilledCreditCardTransactionsResponse `xml:"inquireUnbilledCreditCardTransactionsResponse"`
}

type InquireUnbilledCreditCardTransactionsResponse struct {
	Return Return `xml:"return"`
}

type Return struct {
	Status      Status `xml:"status"`
	ResponseXML string `xml:"responseXML"`
}

type Status struct {
	// Envelope                                      string `xml:"envelope"`
	// Body                                          string `xml:"body"`
	// InquireUnbilledCreditCardTransactionsResponse string `xml:"inquireUnbilledCreditCardTransactionsResponse"`
	//Return                                        string `xml:"return"`
	// ResponseXML string `xml:"responseXML"`
	//Status                                        string      `xml:"status"`
	ErrorCode               string      `xml:"errorCode"`
	ExtendedReply           string      `xml:"extendedReply"`
	ExternalReferenceNo     string      `xml:"externalReferenceNo"`
	InternalReferenceNumber string      `xml:"internalReferenceNumber"`
	IsOverriden             bool        `xml:"isOverriden"`
	PostingDate             PostingDate `xml:"postingDate"`
	DateString              string      `xml:"dateString"`
	ReplyCode               string      `xml:"replyCode"`
	ReplyText               string      `xml:"replyText"`
}

type PostingDate struct {
	DateString string `xml:"dateString"`
}

func main() {
	Data, err := ioutil.ReadFile("demo.xml")
	if err != nil {
		fmt.Println(err)
	}

	response := &Return{}

	err = xml.Unmarshal([]byte(Data), &response)
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("error: %v", err)
	}

	fmt.Println("ResponseXML: ", response.ResponseXML)
	fmt.Println("ExtendedReply: ", response.Status.ExtendedReply)
	fmt.Println("ExternalReferenceNo: ", response.Status.ExternalReferenceNo)
	fmt.Println("InternalReferenceNumber: ", response.Status.InternalReferenceNumber)
	fmt.Println("IsOverriden: ", response.Status.IsOverriden)
	fmt.Println("PostingDate: ", response.Status.PostingDate)
	fmt.Println("DateString: ", response.Status.PostingDate.DateString)
	fmt.Println("ReplyCode: ", response.Status.ReplyCode)
	fmt.Println("ReplyText: ", response.Status.ReplyText)
}
