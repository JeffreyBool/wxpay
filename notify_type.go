package wxpay

// https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_7&index=3
type TradeNotification struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppId      string `xml:"appid"`
	MCHId      string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	SignType   string `xml:"sign_type"`
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`

	OpenId             string `xml:"openid"`
	IsSubscribe        string `xml:"is_subscribe"`
	TradeType          string `xml:"trade_type"`
	BankType           string `xml:"bank_type"`
	TotalFee           int    `xml:"total_fee"`
	SettlementTotalFee int    `xml:"settlement_total_fee"`
	FeeType            string `xml:"fee_type"`
	CashFee            int    `xml:"cash_fee"`
	CashFeeType        string `xml:"cash_fee_type"`
	CouponFee          int    `xml:"coupon_fee"`
	CouponCount        int    `xml:"coupon_count"`
	TransactionId      string `xml:"transaction_id"`
	OutTradeNo         string `xml:"out_trade_no"`
	Attach             string `xml:"attach"`
	TimeEnd            string `xml:"time_end"`
}

// RefundNotification 退款结果
// docs: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=10
type RefundNotification struct {
	ReturnCode          string `xml:"return_code"`           // 返回状态码
	ReturnMsg           string `xml:"return_msg"`            // 当return_code为FAIL时返回信息为错误原因 ，例如 签名失败 参数格式校验错误
	AppId               string `xml:"appid"`                 // 微信分配的公众账号ID（企业号corpid即为此appId）
	MCHId               string `xml:"mch_id"`                // 微信支付分配的商户号
	NonceStr            string `xml:"nonce_str"`             // 随机字符串，不长于32位。推荐随机数生成算法
	ReqInfo             string `xml:"req_info"`              // 加密信息请用商户秘钥进行解密，详见解密方式
	TransactionId       string `xml:"transaction_id"`        // 微信订单号
	OutTradeNo          string `xml:"out_trade_no"`          // 商户系统内部的订单号
	RefundId            string `xml:"refund_id"`             // 微信退款单号
	OutRefundNo         string `xml:"out_refund_no"`         // 商户退款单号
	TotalFee            int    `xml:"total_fee"`             // 订单总金额，单位为分，只能为整数，详见支付金额
	SettlementTotalFee  int    `xml:"settlement_total_fee"`  // 当该订单有使用非充值券时，返回此字段。应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	RefundFee           int    `xml:"refund_fee"`            // 退款总金额,单位为分
	SettlementRefundFee int    `xml:"settlement_refund_fee"` // 退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
	RefundStatus        string `xml:"refund_status"`         // 退款状态 SUCCESS-退款成功 CHANGE-退款异常 REFUNDCLOSE—退款关闭
	SuccessTime         string `xml:"success_time"`          // 退款成功时间 资金退款至用户帐号的时间，格式2017-12-15 09:46:01
	RefundRecvAccout    string `xml:"refund_recv_accout"`    // 退款入账账户
	RefundAccount       string `xml:"refund_account"`        // 退款资金来源 REFUND_SOURCE_RECHARGE_FUNDS 可用余额退款/基本账户 REFUND_SOURCE_UNSETTLED_FUNDS 未结算资金退款
	RefundRequestSource string `xml:"refund_request_source"` // 退款发起来源
}

//签约解约结果通知
//docs: https://pay.weixin.qq.com/wiki/doc/api/pap.php?chapter=18_17&index=6
type ContractNotification struct {
	ReturnCode string `xml:"return_code"` //SUCCESS/FAIL 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnMsg  string `xml:"return_msg"`  //返回信息，如非空，为错误原因 签名失败 参数格式校验错误
	ResultCode string `xml:"result_code"` //业务结果 SUCCESS/FAIL

	MchId                   string `xml:"mch_id"`                    //微信支付分配的商户号
	ContractCode            string `xml:"contract_code"`             //签约协议号
	PlanId                  string `xml:"plan_id"`                   //协议模板id，设置路径见开发步骤
	OpenId                  string `xml:"open_id"`                   //同一个AppId下，用户的唯一标识
	Sign                    string `xml:"sign"`                      //签名
	ChangeType              string `xml:"change_type"`               //变更类型 ADD--签约/DELETE--解约 商户可通过该字段判断是签约回调还是解约回调
	OperateTime             string `xml:"operate_time"`              //操作时间
	ContractId              string `xml:"contract_id"`               //签约成功后，微信返回的委托代扣协议id
	ContractExpiredTime     string `xml:"contract_expired_time"`     //协议到期时间，当change_type为ADD时有返回（目前协议时间为长期有效，可以忽略该字段）
	ContractTerminationMode int    `xml:"contract_termination_mode"` //协议解约方式 当change_type为DELETE时有返回 0-未解约/1-有效期过自动解约/2-用户主动解约 /3-商户API解约 /4-商户平台解约 /5-用户帐号注销
	RequestSerial           int64  `xml:"request_serial"`            //商户请求签约时的序列号，要求唯一性。序列号主要用于排序，不作为查询条件，纯数字,范围不能超过Int64的范围（9223372036854775807）。
}

//签约申请扣款结果通知
//docs: https://pay.weixin.qq.com/wiki/doc/api/pap.php?chapter=18_7&index=11
//该链接是通过【申请扣款API】中提交的参数notify_url设置，如果链接无法访问，商户将无法接收到微信通知。

type (
	PayApplyNotification struct {
		ReturnCode string `xml:"return_code"` //SUCCESS/FAIL 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
		ReturnMsg  string `xml:"return_msg"`  //返回信息，如非空，为错误原因 签名失败 参数格式校验错误

		//以下字段在return_code为SUCCESS的时候有返回
		AppId         string `xml:"appid"`
		MchId         string `xml:"mch_id"`
		DeviceInfo    string `xml:"device_info"`
		NonceStr      string `xml:"nonce_str"`
		Sign          string `xml:"sign"`
		ResultCode    string `xml:"result_code"`
		ErrCode       string `xml:"err_code"`
		ErrCodeDes    string `xml:"err_code_des"`
		OpenId        string `xml:"open_id"`
		IsSubscribe   string `xml:"is_subscribe"`
		BankType      string `xml:"bank_type"`
		TotalFee      int    `xml:"total_fee"`
		FeeType       string `xml:"fee_type"`
		CashFee       int    `xml:"cash_fee"`
		CashFeeType   string `xml:"cash_fee_type"`
		TradeState    string `xml:"trade_state"`
		CouponFee     int    `xml:"coupon_fee"`
		CouponCount   int    `xml:"coupon_count"`
		CouponIdn     string `xml:"coupon_id_$n"`
		CouponFeen    int    `xml:"coupon_fee_$n"`
		TransactionId string `xml:"transaction_id"`
		OutTradeNo    string `xml:"out_trade_no"`
		Attach        string `xml:"attach"`
		TimeEnd       string `xml:"time_end"`
		ContractId    string `xml:"contract_id"`
	}
)

//type (
//	//扣款成功通知参数
//	PayApplySuccessNotification struct {
//		ReturnCode string `xml:"return_code"` //SUCCESS/FAIL 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
//		ReturnMsg  string `xml:"return_msg"`  //返回信息，如非空，为错误原因 签名失败 参数格式校验错误
//
//		//以下字段在return_code为SUCCESS的时候有返回
//		AppId         string `xml:"appid"`
//		MchId         string `xml:"mch_id"`
//		DeviceInfo    string `xml:"device_info"`
//		NonceStr      string `xml:"nonce_str"`
//		Sign          string `xml:"sign"`
//		ResultCode    string `xml:"result_code"`
//		ErrCode       string `xml:"err_code"`
//		ErrCodeDes    string `xml:"err_code_des"`
//		OpenId        string `xml:"open_id"`
//		IsSubscribe   string `xml:"is_subscribe"`
//		BankType      string `xml:"bank_type"`
//		TotalFee      int    `xml:"total_fee"`
//		FeeType       string `xml:"fee_type"`
//		CashFee       int    `xml:"cash_fee"`
//		CashFeeType   string `xml:"cash_fee_type"`
//		TradeState    string `xml:"trade_state"`
//		CouponFee     int    `xml:"coupon_fee"`
//		CouponCount   int    `xml:"coupon_count"`
//		CouponIdn     string `xml:"coupon_id_$n"`
//		CouponFeen    int    `xml:"coupon_fee_$n"`
//		TransactionId string `xml:"transaction_id"`
//		OutTradeNo    string `xml:"out_trade_no"`
//		Attach        string `xml:"attach"`
//		TimeEnd       string `xml:"time_end"`
//		ContractId    string `xml:"contract_id"`
//	}
//
//	//扣款失败通知参数
//	PayApplyFailNotification struct {
//		ReturnCode string `xml:"return_code"` //SUCCESS/FAIL 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
//		ReturnMsg  string `xml:"return_msg"`  //返回信息，如非空，为错误原因 签名失败 参数格式校验错误
//
//		//以下字段在return_code为SUCCESS的时候有返回
//		AppId      string `xml:"appid"`
//		MchId      string `xml:"mch_id"`
//		NonceStr   string `xml:"nonce_str"`
//		Sign       string `xml:"sign"`
//		ResultCode string `xml:"result_code"`
//		ErrCode    string `xml:"err_code"`
//		ErrCodeDes string `xml:"err_code_des"`
//		TradeState string `xml:"trade_state"`
//		OutTradeNo string `xml:"out_trade_no"`
//		ContractId string `xml:"contract_id"`
//	}
//)
