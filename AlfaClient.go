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

type AlfaSendBitModel struct {
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

// Provides MD5 hashing for password
func (ac AlfaClient) HashMd5(source string) string {
	hasher := md5.New()
	hasher.Write([]byte(source))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Create new instance of Alphamodel
func (ac AlfaClient) New(appName string, api_secret string, apiPassword string) AlfaModel {
	key := ac.HashMd5(apiPassword)
	return AlfaModel{appName, api_secret, key}
}

func (ac AlfaClient) NewSendBitModel(appName string, api_secret string, apiPassword string, recipient_email string, recipient_name string, recipientAddress string, senderAddress string) AlfaSendBitModel {
	key := ac.HashMd5(apiPassword)
	return AlfaSendBitModel{appName, api_secret, key, recipient_email,recipient_name, recipientAddress, senderAddress }
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
func (ac AlfaClient) ParseBitSend(am AlfaSendBitModel, coinType int, amount float64) string {

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
func (ac AlfaClient) GetBalance(am AlfaModel) string {
	jsonCredential := ac.ParseCredential(am)
	fmt.Println(jsonCredential)
	request := gorequest.New()
	_, body, _ := request.Post(getBalanceUrl).Set(jsonContentType, jsonMineType).Send(jsonCredential).End()
	return body;
}

// SendBit - BitSend primary use to payout salaries for staff or making direct deposits
// to different cryptocurrency addresses
func (ac AlfaClient) SendBit(am AlfaSendBitModel, coinType int, amount float64) string {
	jsonData := ac.ParseBitSend(am, coinType, amount)
	fmt.Print(jsonData)
	request := gorequest.New()

	_, body, _ := request.Post(sendBitUrl).Set(jsonContentType, jsonMineType).Send(jsonData).End()
	return body
}

// Bit Status - BitSend status primary use to get information of bitsend payout
func (ac AlfaClient) GetSendBitStatus(am AlfaModel, sendBitId int) string {

	modelstr := ac.ParseBitSendStatus(am, sendBitId)
	fmt.Println(modelstr)
	request := gorequest.New()
	_, body, _ := request.Post(sendBitStatusUrl).Set(jsonContentType, jsonMineType).Send(modelstr).End()
	return body
}

// Order status
func (ac AlfaClient) GetOrderStatus(am AlfaModel, txtId int) string {

	modelstr := ac.ParseOrderStatus(am, txtId)
	fmt.Println(modelstr)
	request := gorequest.New()
	_, body, _ := request.Post(orderStatusUrl).Set(jsonContentType, jsonMineType).Send(modelstr).End()
	return body
}

// Refund
func (ac AlfaClient) RefundOrder(am AlfaModel, txtId int, address string, amount float64) string {

	modelstr := ac.ParseRefundOrder(am, txtId, address, amount)
	request := gorequest.New()
	_, body, _ := request.Post(refundUrl).Set(jsonContentType, jsonMineType).Send(modelstr).End()
	return body
}

// Get Rate - Get rates for pair
func (ac AlfaClient) GetRates() string {
	request := gorequest.New()
	_, body, _ := request.Get(ratesUrl).Set(jsonContentType, jsonMineType).End()
	return body
}

// Get Rates - Get rate for all available pairs