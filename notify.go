package wxpay

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/nanjishidu/gomini/gocrypto"
)

//xml 解析
func (this *Client) Unmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}

// GetTradeNotification https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_7&index=3
//微信支付异步验证
func (this *Client) GetTradeNotification(data []byte) (result *TradeNotification, err error) {
	return result, this.verifyResponse(data, &result)
}

// GetRefundNotification 退款结果通知
// docs: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=10
func (this *Client) GetRefundNotification(data []byte) (result *RefundNotification, err error) {
	var param = make(XMLMap)
	if err = xml.Unmarshal(data, &param); err != nil {
		return nil, err
	}

	// 处理错误信息
	var code = param.Get("return_code")
	if code == K_RETURN_CODE_FAIL {
		var msg = param.Get("return_msg")
		if msg == "" {
			msg = param.Get("retmsg")
		}
		return nil, errors.New(msg)
	}

	code = param.Get("result_code")
	if code == K_RETURN_CODE_FAIL {
		var msg = param.Get("err_code_des")
		return nil, errors.New(msg)
	}

	bytes, err := this.DecryptRefundNotifyReqInfo(param.Get("req_info"))
	if err != nil {
		return nil, err
	}

	if err = this.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}

	result.AppId = param.Get("appid")
	result.MCHId = param.Get("mch_id")
	result.NonceStr = param.Get("nonce_str")
	result.ReqInfo = param.Get("req_info")
	result.ReturnCode = code
	result.ReturnMsg = param.Get("return_msg")
	return result, nil
}

//签约解约异步通知
func (this *Client) GetContractNotification(data []byte) (result *ContractNotification, err error) {
	return result, this.verifyResponse(data, &result)
}

//签约扣款异步通知
func (this *Client) GetPayApplyNotification(data []byte) (result *PayApplyNotification, err error) {
	return result, this.verifyResponse(data, &result)
}

// 解密微信退款异步通知的加密数据
// 方法获取的加密数据 req_info
// 文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=10
func (this *Client) DecryptRefundNotifyReqInfo(reqInfo string) (result []byte, err error) {
	key, err := this.getKey()
	if err != nil {
		return nil, err
	}
	b, err := base64.StdEncoding.DecodeString(reqInfo)
	if err != nil {
		return nil, err
	}
	if err = gocrypto.SetAesKey(strings.ToLower(gocrypto.Md5(key))); err != nil {
		return nil, err
	}
	return gocrypto.AesECBDecrypt(b)
}

//验签
func (this *Client) verifyResponse(data []byte, result interface{}) (err error) {
	key, err := this.getKey()
	if err != nil {
		return err
	}
	if _, err := VerifyResponseData(data, key); err != nil {
		return err
	}
	return this.Unmarshal(data, result)
}

func (this *Client) AckNotification(w http.ResponseWriter) {
	AckNotification(w)
}

func AckNotification(w http.ResponseWriter) {
	var v = url.Values{}
	v.Set("return_code", "SUCCESS")
	v.Set("return_msg", "OK")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(URLValueToXML(v)))
}
