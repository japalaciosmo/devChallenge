package parser

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"encoding/json"
)
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}


const (
	HOST="https://stooq.com/q/l/?s=%s&f=sd2t2ohlcv&h&e=csv"
)


func ParseMessage(data []byte) (newMessage []byte, resend bool, err error){

	var message Message
	var stockMessage string
	if err := json.Unmarshal(data, &message); err != nil {
		return []byte("fag"), false,err
	}

	r := regexp.MustCompile(`^/stock=`)
	r.MatchString(message.Body)

	if (r.MatchString(message.Body)){
		index := strings.Index(message.Body, "=")
		value:=message.Body[index+1:]
		stockMessage,err=getCsv(value)
		if err != nil {
			return []byte(stockMessage), false,err
		}
		return []byte(stockMessage), true,nil
	}

	return []byte(stockMessage), false,nil
}


func getCsv(code string)(string,error) {
	host:=fmt.Sprintf(HOST,code)

	// Get the data
	resp, err := http.Get(host)
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		return "", err
	}

	if len(lines) >1{
		values:=lines[1]
		if len(values)>6 {
			finalmessage:=fmt.Sprintf("Bot says: %s quote is %s per share",values[0],values[6])
			return finalmessage,nil
		}

	}

	return "",nil
}