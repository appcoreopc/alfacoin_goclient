# alfacoin_goclient




Go client SDK that allow you to consume alfacoin gateway. 

## Install 

go get github.com/appcoreopc/alfacoin_goclient

## Get Rates 

Example :-

ac := AlfaClient{}
ratesData := ac.GetRates()
  
## Send BitCoin

Example -

ac := AlfaClient{}
sendBitcoinModel := ac.NewSendBitModel(appName, apiSecret, apiPassword, "kepung@gmail.com", "jeremy", "xxxx111", "aaaaaaaaa")

ratesData := ac.SendBit(sendBitcoinModel, 0, 2300) 

NewSendBitModel accepts appname, api secret, api_password, recipient_email, recipient_name, recipient_address, sender_address)

- Please note :- all the fields above are mandatory, otherwise alphacoins will throw error 

SendBit(sendBitcoinModel, coin_type, amount) 

coin_type - 0 bitcoin, 1 - litecoin, 2 - ethereum, 3 - dash 

## Get SendBit Status

Example :-

ac := AlfaClient{}
model := ac.New(appName, apiSecret, apiPassword)
ratesData := ac.GetSendBitStatus(model, sendBitId)

GetSendBitStatus(model, sendBitId)
model -  consist of appname, apisecret and api key
sendBitId - is of type int. It is the transaction id that you get after calling SendBit. 

## GetOrderStatus

Example :-

ac := AlfaClient{}
model := ac.New(appName, apiSecret, apiPassword)
orderStatusData := ac.GetOrderStatus(model, 12345)

GetOrderStatus 
model -  consist of appname, apisecret and api key after calling ac.New()
sendBitId - an int representing transaction id.

## RefundOrder 

Example :- 

ac := AlfaClient{}
model := ac.New(appName, apiSecret, apiPassword)
orderStatusData := ac.RefundOrder(model, 12345, recipient_address, 200)


RefundOrder(model, txtId, recipient_address, amount)
model -  consist of appname, apisecret and api key after calling ac.New()
txtId - transaction id you obtain after getting txtId
recipient_address - beneficiary of the refund
amount - to refund.

 

