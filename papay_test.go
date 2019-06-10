/**
* Author: zhanggaoyuancn@163.com
* Date: 2019-06-09
* Time: 21:16
* Software: GoLand
 */

package wxpay

import (
	"testing"
)

func TestContratOrder(t *testing.T) {
	t.Log("========== UnifiedOrder ==========")
	var p = ContratOrderParam{}
	p.Body = "test product"
	p.NotifyUrl = "http://www.test.com"
	p.TradeType = K_TRADE_TYPE_MWEB
	p.SpbillCreateIp = "202.105.107.18"
	p.TotalFee = 101
	p.OutTradeNo = "test-111111122sdf"

	result, err := client.ContratOrder(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.PrepayId, result)
}
