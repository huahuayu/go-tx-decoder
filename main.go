package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/huahuayu/go-tx-decoder/abi"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

var (
	abiPath string
	txInput string
)

func init() {
	flag.StringVar(&abiPath, "abi", "", "abi path")
	flag.StringVar(&txInput, "input", "", "transaction input")
	flag.Parse()
}

func main(){
	if abiPath == "" || txInput == ""{
		logrus.Error("abi & input should be given")
		return
	}
	// read the whole abi file into bytes
	fileBytes, err := ioutil.ReadFile(abiPath)
	if err != nil {
		logrus.Error(err)
		return
	}
	a, err := abi.NewAbi(string(fileBytes))
	if err != nil {
		logrus.Error(err)
		return
	}
	method, decodeInput, err := a.Decode(txInput)
	if err != nil {
		logrus.Error(err)
		return
	}
	res := struct{
		Function string `json:"function"`
		Id string `json:"id"`
		Data interface{} `json:"data"`
	}{
		Function: method.String(),
		Id: hexutil.Encode(method.ID),
		Data:       decodeInput,
	}
	bs, err := json.Marshal(res)
	if err != nil {
		logrus.Error(err)
	}
	fmt.Println(jsonPrettyPrint(bs))
}

func jsonPrettyPrint(bs []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, bs, "", "    ")
	if err != nil {
		return string(bs)
	}
	return out.String()
}
