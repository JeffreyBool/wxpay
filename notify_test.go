/**
* Author: zhanggaoyuancn@163.com
* Date: 2019-07-15
* Time: 15:51
* Software: GoLand
 */

package wxpay

import (
	"testing"
)

func TestGetTradeNotification(t *testing.T)  {
	result, err := client.GetTradeNotification([]byte(`<xml><appid><![CDATA[wxcdc4e97b26ed94a9]]></appid>
<bank_type><![CDATA[CFT]]></bank_type>
<cash_fee><![CDATA[100]]></cash_fee>
<fee_type><![CDATA[CNY]]></fee_type>
<is_subscribe><![CDATA[Y]]></is_subscribe>
<mch_id><![CDATA[1543924041]]></mch_id>
<nonce_str><![CDATA[54i77czqpn7kwu2e7z1elylzqexhlej3]]></nonce_str>
<openid><![CDATA[oLDyb1HoSBxz2ViWhR16C9vzosZI]]></openid>
<out_trade_no><![CDATA[201907150000023777]]></out_trade_no>
<result_code><![CDATA[SUCCESS]]></result_code>
<return_code><![CDATA[SUCCESS]]></return_code>
<sign><![CDATA[0A0797E546886F1FE148589E27BC5B42]]></sign>
<time_end><![CDATA[20190715145759]]></time_end>
<total_fee>100</total_fee>
<trade_type><![CDATA[NATIVE]]></trade_type>
<transaction_id><![CDATA[4200000349201907152838587776]]></transaction_id>
</xml>`))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}