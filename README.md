# alfacoin_goclient




Go client SDK that allow you to consume alfacoin gateway. 

## Install 

go get github.com/appcoreopc/alfacoin_goclient

## Get Rates 

ac := AlfaClient{}
ratesData := ac.GetRates()
  
## Send BitCoin

ac := AlfaClient{}
sendBitcoinModel := ac.NewSendBitModel(appName, apiSecret, apiPassword, "kepung@gmail.com", "jeremy", "xxxx111", "aaaaaaaaa")
ratesData := ac.SendBit(sendBitcoinModel, 0, 2300) 


NewSendBitModel accepts appname, api secret, api_password, recipient_email, recipient_name, recipient_address, sender_address)

- All the fields above are mandatory, otherwise alphacoins will throw error 

SendBit(sendBitcoinModel, coin_type, amount) 

coin_type - 0 bitcoin, 1 - litecoin, 2 - ethereum, 3 - dash 


## Get SendBit Status

ac := AlfaClient{}
model := ac.New(appName, apiSecret, apiPassword)
ratesData := ac.GetSendBitStatus(model, 12345)

GetSendBitStatus(model, sendBitId)
sendBitId is the transaction id that you get after calling SendBit. 





