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
	p.TradeType = K_TRADE_TYPE_JSAPI
	p.SpbillCreateIP = "202.105.107.18"
	p.TotalFee = 1
	p.OutTradeNo = time.Now().Format("20060102150405")
	p.OpenId= "oLDyb1HoSBxz2ViWhR16C9vzosZI"
	p.DeviceInfo = "WEB"

	result, err := client.UnifiedOrder(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result: #%v \n",result)
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
