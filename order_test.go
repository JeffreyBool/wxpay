package wxpay

import (
	"testing"
	"time"
)

func TestWXPay_UnifiedOrder(t *testing.T) {
	t.Log("========== UnifiedOrder ==========")
	var p = UnifiedOrderParam{}
	p.Body = "test product"
	p.NotifyURL = "http://www.test.com"
	p.TradeType = K_TRADE_TYPE_NATIVE
	p.SpbillCreateIP = "202.105.107.18"
	p.TotalFee = 101
	p.OutTradeNo = "test-111111125"

	result, err := client.UnifiedOrder(p)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("result: %#v \n",result)
}

func TestWXPay_AppPay(t *testing.T) {
	t.Log("========== AppPay ==========")
	var p = UnifiedOrderParam{}
	p.Body = "test product"
	p.NotifyURL = "http://www.test.com"
	p.SpbillCreateIP = "202.105.107.18"
	p.TotalFee = 101
	p.OutTradeNo = "test-111111125"

	result, err := client.AppPay(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("以下为 TradeType 为 APP 时附加的信息")
	t.Log("AppId", result.AppId)
	t.Log("PartnerId", result.PartnerId)
	t.Log("PrepayId", result.PrepayId)
	t.Log("Package", result.Package)
	t.Log("NonceStr", result.NonceStr)
	t.Log("TimeStamp", result.TimeStamp)
	t.Log("Sign", result.Sign)
}

func TestJSAPIPay(t *testing.T)  {
	var p = UnifiedOrderParam{}
	p.Body = "test product"
	p.NotifyURL = "http://www.test.com"
	p.SpbillCreateIP = "202.105.107.18"
	p.TotalFee = 1
	p.OutTradeNo = time.Now().Format("20060102150405")
	p.OpenId= "oLDyb1HoSBxz2ViWhR16C9vzosZI"
	p.DeviceInfo = "WEB"
	result, err := client.JSAPIPay(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: %#v \n",result)
}

func TestWXPay_OrderQuery(t *testing.T) {
	t.Log("========== OrderQuery ==========")
	var p = OrderQueryParam{}
	p.OutTradeNo = "test-11111112"

	result, err := client.OrderQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.TradeState, result.OutTradeNo, result.TransactionId)
}

func TestWXPay_CloseOrder(t *testing.T) {
	t.Log("========== CloseOrder ==========")
	var p = CloseOrderParam{}
	p.OutTradeNo = "test-11111112"

	result, err := client.CloseOrder(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.ReturnCode, result.ReturnMsg)
}

func TestWXPay_DownloadBill(t *testing.T) {
	t.Log("========== DownloadBill ==========")
	var p = DownloadBillParam{}
	p.BillDate = "20190108"
	p.BillType = "ALL"
	p.TarType = "GZIP"

	result, err := client.DownloadBill(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.ReturnCode, result.ReturnMsg)
	t.Log(string(result.Data))
}