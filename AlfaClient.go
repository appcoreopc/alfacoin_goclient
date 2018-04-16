package main

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"github.com/parnurzeal/gorequest"
	"strconv"
)

const getBalanceUrl = "https://www.alfacoins.com/api/stats"
const sendBitUrl = "https://www.alfacoins.com/api/bitsend"
const sendBitStatusUrl = "https://www.alfacoins.com/api/bitsend_status"
const createOrderUrl = "https://www.alfacoins.com/api/create"
const orderStatusUrl = "https://www.alfacoins.com/api/status"
const refundUrl = "https://www.alfacoins.com/api/refund"
const jsonContentType = "Content-Type"
const jsonMineType = "application/json"

const ratesUrl = "https://www.alfacoins.com/api/rates"
const ratePairUrl = "https://www.alfacoins.com/api/rate/"

const(
	bitcoin = iota
	litecoin = iota
	Ethereum = iota
	dash = iota
)

type AlfaModel struct {
	Name       string
	Secret_key string
	Password   string
}

type AlfaModelSendBit struct {
	Name       string
	Secret_key string
	Password   string
	Recipient_email string
	Recipient_name string
	RecipientAddress string
	SenderAddress string
}

type AlfaClient struct {

	appName string
	api_secret_key string
	apiPassword string
}

func (ac AlfaClient) HashMd5(source string) string {
	hasher := md5.New()
	hasher.Write([]byte(source))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (ac AlfaClient) New(appName string, api_secret string, apiPassword string) AlfaModel {
	key := ac.HashMd5(apiPassword)
	return AlfaModel{appName, api_secret, key}
}

// Using the json marshal command causessome issue for alfacoint validation
func (ac AlfaClient) ParseCredential(am AlfaModel) string {
	dataJson  := "{\"name\" : \"" + am.Name + "\", \"secret_key\" : \"" + am.Secret_key + "\", " + "\"password\" : \"" + am.Password + "\"}"
	return dataJson
}

func (ac AlfaClient) ParseBitSendStatus(am AlfaModel, sendbitId int) string {
	dataJson  := "{\"name\" : \"" + am.Name + "\", \"secret_key\" : \"" + am.Secret_key + "\", " + "\"password\" : \"" + am.Password + "\", \"bitsend_id\" : " +  strconv.Itoa(sendbitId) +  "}"
	return dataJson
}

func (ac AlfaClient) ParseOrderStatus(am AlfaModel, txtId int) string {
	dataJson  := "{\"name\" : \"" + am.Name + "\", \"secret_key\" : \"" + am.Secret_key + "\", " + "\"password\" : \"" + am.Password + "\", \"txn_id\" : " +  strconv.Itoa(txtId) +  "}"
	return dataJson
}

func (ac AlfaClient) ParseRefundOrder(am AlfaModel, txtId int, recipientAddress string, amount float64) string {
	dataJson  := "{\"name\" : \"" + am.Name + "\", \"secret_key\" : \"" + am.Secret_key + "\", " + "\"password\" : \"" + am.Password + "\", \"txn_id\" : " +  strconv.Itoa(txtId) + ", \"address\" : \"" +  recipientAddress +  "\", " + "\"amount\" : " + strconv.FormatFloat(amount, 'f', 6, 32) + "}"
	return dataJson
}

// Using the json marshal command causessome issue for alfacoint validation
func (ac AlfaClient) ParseBitSend(am AlfaModelSendBit, coinType int, amount float64) string {

	coinStringName := "bitcoin"
	switch coinType {
	case 0:
		coinStringName = "bitcoin"
		break;
	case 1:
		coinStringName = "litecoin"
		break;
	case 2:
		coinStringName = "ethereum"
		break;
	case 3:
		coinStringName = "dash"
		break;
	}

	optionJson := "\"options\" : {\"address\": \"" + am.RecipientAddress + "\" , \"destination_tag\": \"1294967290\"},"
	recipientEmail := "\"recipient_email\" : \"" + am.Recipient_email + "\""
	referenceInfo := "\"reference\" : \"" + am.Recipient_name +  "\""
	recipientName := "\"recipient_name\" : \"" + am.Recipient_name + "\"" + ", " + recipientEmail
	dataJson  := "{\"name\" : \"" + am.Name + "\", \"secret_key\" : \"" + am.Secret_key + "\", " + "\"password\" : \"" + am.Password + "\", " + referenceInfo + "," + recipientName + "," + optionJson + "\"type\" : \"" +  coinStringName +  "\", \"amount\": \"" + ToString(amount) + "\""  +  "  }"
	return dataJson
}

// Get balance
func (ac AlfaClient) GetBalance(am AlfaModel) {
	jsonCredential := ac.ParseCredential(am)
	fmt.Println(jsonCredential)

	request := gorequest.New()
	res, _, _ := request.Post(getBalanceUrl).Set(jsonContentType, jsonMineType).Send(jsonCredential).End()
	fmt.Println(res.Status)
	fmt.Println(res.Body)
}

// SendBit - BitSend primary use to payout salaries for staff or making direct deposits to different cryptocurrency addresses
func (ac AlfaClient) SendBit(am AlfaModelSendBit, coinType int, amount float64) {
	jsonData := ac.ParseBitSend(am, coinType, amount)
	fmt.Print(jsonData)
	request := gorequest.New()

	res, _, _ := request.Post(sendBitUrl).Set(jsonContentType, jsonMineType).Send(jsonData).End()

	fmt.Println(res.Status)
	fmt.Println(res.Body)
}

// Bit Status - BitSend status primary use to get information of bitsend payout
func (ac AlfaClient) SendBitStatus(am AlfaModel, sendBitId int) {

	modelstr := ac.ParseBitSendStatus(am, sendBitId)

	fmt.Println(modelstr)
	request := gorequest.New()
	res, _, _ := request.Post(sendBitStatusUrl).Set(jsonContentType, jsonMineType).Send(modelstr).End()

	fmt.Println(res.Status)
	fmt.Println(res.Body)
}
// Create Order - Pay someone


// Order status
func (ac AlfaClient) GetOrderStatus(am AlfaModel, txtId int) {

	modelstr := ac.ParseOrderStatus(am, txtId)

	fmt.Println(modelstr)
	request := gorequest.New()
	res, _, _ := request.Post(orderStatusUrl).Set(jsonContentType, jsonMineType).Send(modelstr).End()

	fmt.Println(res.Status)
	fmt.Println(res.Body)
}

// Refund
func (ac AlfaClient) RefundOrder(am AlfaModel, txtId int, address string, amount float64) {

	modelstr := ac.ParseRefundOrder(am, txtId, address, amount)

	fmt.Println(modelstr)
	request := gorequest.New()
	res, _, _ := request.Post(refundUrl).Set(jsonContentType, jsonMineType).Send(modelstr).End()

	fmt.Println(res.Status)
	fmt.Println(res.Body)
}

// Get Rate - Get rates for pair
func (ac AlfaClient) GetRates() {
	request := gorequest.New()
	res, _, _ := request.Get(ratesUrl).Set(jsonContentType, jsonMineType).End()
	fmt.Println(res.Status)
	fmt.Println(res.Body)
}

// Get Rates - Get rate for all available pairs