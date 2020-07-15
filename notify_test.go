/**
* Author: zhanggaoyuancn@163.com
* Date: 2019-07-15
* Time: 15:51
* Software: GoLand
 */

package wxpay

import (
	"encoding/json"
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

func TestClient_GetRefundNotification(t *testing.T) {
	result, err := client.GetRefundNotification([]byte(`PHhtbD48cmV0dXJuX2NvZGU+U1VDQ0VTUzwvcmV0dXJuX2NvZGU+PGFwcGlkPjwhW0NEQVRBW3d4Y2RjNGU5N2IyNmVkOTRhOV1dPjwvYXBwaWQ+PG1jaF9pZD48IVtDREFUQVsxNTQzOTI0MDQxXV0+PC9tY2hfaWQ+PG5vbmNlX3N0cj48IVtDREFUQVs3NzFhMDM5YjRkNjQ3MTMwMjU4MGY4NDkzYjAwNGZhNF1dPjwvbm9uY2Vfc3RyPjxyZXFfaW5mbz48IVtDREFUQVthbEN3dzlnV1lSQWZhOXY3d1VES3ZIOE5hTVhMK0JCV2VNZW9vbHZhWHdxZDc4OUtPS0dkdkNuTDQzV2d2Z3BTRHhIRURmSytZUlA5NkxmK084YTFBV2IxaEt4SmFMc0R1bEU2cEpidzhpdVA2NERGanBTYWFyWDU2aWNjWElaZTBML2I4UVlTMTNISXAzZzhpVmRPK0VZSUlLSzYvWVhzcTI1RWdtTEJDZFFJSGRXanFLZis3dFJxWTB1bHI1MlZnRUtUSXUrWmZEa09PUUs2dFZ2VEVTaWYxNXFYZUtzUXE3YVFjMENYNEpoU3IwdDloZ2w4enFrWmdady9qdCtMRUs1RXh0Q1hYUTB4dkNmWEFhQ3JoRlFIWmNhOSttVWFmYXJ2MjJVTDljd1NNRXVnTjJoeWZFTFdJeGdISFliMHcrVVdTeTdCWGNGcU8vcDJKTEtCWjRpcTRPKzcrYThlVWRBTDgwNFd3WElJQUVZWjdUVzBsR0QzNldUNlMrTGNwMjZVVkh2d3k0NEJCMEVabW1Tb2pPb1I2SmcwNDZtR25BQkp5NVFTb2xuK2xZcWFFN0xRSmR4UHJCWmpDMGo1NEJxcGJNU25BRFE1bHNtREJmL3paR3ZWTTlSeExPL0hiUllic3BwUlMvT1lWTGo4R212cjdDTlZVNDdzQjBqZG9aRm8wazBmK1dsSG1CQ3ljdElvRTl6ZTBTWEZOanF4bmEwNkRFOGQvTk54R1FmQk5PcXdVdVFLbnlHRzU3cnQrRVh1aThpbzViU1VXY09oaEVhZWVoZmJzM3JtWS80Z3hIVnp0QWdacDZ4NENjMjY3WFlsc3lKOUJ3UWlqVVFXb3NzTjNKVkJtVGVqdVI5N3hncW5jVFNaemxPcVNNOSt2Rm10ZG1BYThCWklnUkFBUkdBcFE4ZFBaWFUyU2ZZMVN3RW5hcnJvaTQ5ZmRGeTIra0dUckdBYXoxLzBvTlJESFpueDIrMVlLV0kyeHNZYXFBK0VOK3phaW9idlRJM2diUmhXRFFBWnZiS1NnRy9XRG1KNzhZanJ3OUFERnBKK250Rmc0WVVBOE5ablZPd3M5d2UzS1NiakxoQklIMGFmSmNwN1p6NS9USGRQWmwwa2FzTGVJcWphcWE0Y1VobkdjZENDYnZLcGJVYmFMazY3TExrVVhocVNUVzN6SEVaV3BYc1Z6d1dnNWNBeHBvbllMRmJJdEtiOTNxMlVSbDF4Y3dtWnAyS0lDM05JRlNoQlRZTjN2S2d3dTVZZE92QmJIdmpyZk40dWRqZGhMRUFsMlpQZnYrVXJEZ0wrekFWRG9mUUJHaGcvdmhyblZjRk1tNm8xaHlZZUdueG5RNkJzQ3ozZHRtU1dQQjNzRWxSanRpZ3VwTWRaZjBKRU5VYXhDOG95QkNOWlozYz1dXT48L3JlcV9pbmZvPjwveG1sPg==`))
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n",bytes)
}