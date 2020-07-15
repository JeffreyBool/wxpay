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
	result, err := client.GetRefundNotification([]byte(`<xml><return_code>SUCCESS</return_code><appid><![CDATA[wxcdc4e97b26ed94a9]]></appid><mch_id><![CDATA[1543924041]]></mch_id><nonce_str><![CDATA[b123df307f89a510cb14bced1cacf823]]></nonce_str><req_info><![CDATA[alCww9gWYRAfa9v7wUDKvH8NaMXL+BBWeMeoolvaXwoyDG+kXQKWXJK5+tSXiFkfDxHEDfK+YRP96Lf+O8a1AWb1hKxJaLsDulE6pJbw8iuP64DFjpSaarX56iccXIZeC4jM1F0B7PGGRKEW1uXuFtfEVvJSKVK1DSD245Euyu0IHdWjqKf+7tRqY0ulr52VgEKTIu+ZfDkOOQK6tVvTESif15qXeKsQq7aQc0CX4JhSr0t9hgl8zqkZgZw/jt+LEK5ExtCXXQ0xvCfXAaCrhFQHZca9+mUafarv22UL9cwSMEugN2hyfELWIxgHHYb0w+UWSy7BXcFqO/p2JLKBZ4iq4O+7+a8eUdAL804WwXLXhAIYyHFzqrPrs/43/ZKqEeE1hTKs580i5cFvrEfS/uoR6Jg046mGnABJy5QSoln+lYqaE7LQJdxPrBZjC0j54BqpbMSnADQ5lsmDBf/zZGvVM9RxLO/HbRYbsppRS/OYVLj8Gmvr7CNVU47sB0jdoZFo0k0f+WlHmBCyctIoE9ze0SXFNjqxna06DE8d/NNxGQfBNOqwUuQKnyGG57rt+EXui8io5bSUWcOhhEaeehfbs3rmY/4gxHVztAgZp6x4Cc267XYlsyJ9BwQijUQWossN3JVBmTejuR97xgqncTSZzlOqSM9+vFmtdmAa8BZIgRAARGApQ8dPZXU2SfY1SwEnarroi49fdFy2+kGTrGAaz1/0oNRDHZnx2+1YKWI2xsYaqA+EN+zaiobvTI3gbRhWDQAZvbKSgG/WDmJ78Yjrw9ADFpJ+ntFg4YUA8NZnVOws9we3KSbjLhBIH0afJcp7Zz5/THdPZl0kasLeIoe5a/8PKiHT5dnXGabY9tgkhfen1l9a6TUjlYDofPOhpXsVzwWg5cAxponYLFbItKb93q2URl1xcwmZp2KIC3NIFShBTYN3vKgwu5YdOvBbHvjrfN4udjdhLEAl2ZPfv/HEjZHVaenYoijYSBD9Iu0Acl/PezFWAu8R0UQulNny3m4lHaNbsO8oWzkHcOQI1cdZf0JENUaxC8oyBCNZZ3c=]]></req_info></xml>`))
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n",bytes)
}