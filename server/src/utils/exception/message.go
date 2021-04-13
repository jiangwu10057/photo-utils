package exception

import (
	"net/http"
	"server/conf"
)

// MsgFlags 状态码map
var MsgFlags = map[int]string{
	http.StatusOK:                   "ResponseCode.Success",
	http.StatusInternalServerError:  "ResponseCode.Fail",
	http.StatusCreated:              "修改密码成功",
	http.StatusNonAuthoritativeInfo: "该第三方账号未绑定",

	http.StatusBadRequest:      "Error.BadRequest",
	ERROR_EXIST_NICK:           "已存在该昵称",
	ERROR_EXIST_USER:           "已存在该用户名",
	ERROR_NOT_EXIST_USER:       "该用户不存在",
	ERROR_NOT_COMPARE:          "帐号密码错误",
	ERROR_NOT_COMPARE_PASSWORD: "两次密码输入不一致",
	ERROR_FAIL_ENCRYPTION:      "加密失败",
	ERROR_NOT_EXIST_PRODUCT:    "该商品不存在",
	ERROR_NOT_EXIST_ADDRESS:    "该收货地址不存在",
	ERROR_EXIST_FAVORITE:       "已收藏该商品",

	ERROR_AUTH_CHECK_TOKEN_FAIL:       "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:    "Token已超时",
	ERROR_AUTH_TOKEN:                  "Token生成失败",
	http.StatusUnauthorized:           "Error.Unauthorized",
	ERROR_AUTH_INSUFFICIENT_AUTHORITY: "权限不足",
	ERROR_READ_FILE:                   "读文件失败",
	ERROR_SEND_EMAIL:                  "发送邮件失败",
	ERROR_CALL_API:                    "调用接口失败",
	ERROR_UNMARSHAL_JSON:              "解码JSON失败",

	ERROR_DATABASE: "Error.Database",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		msg = MsgFlags[http.StatusBadRequest]
	}

	return conf.T(msg)
}
