package suite

import (
	"net/url"

	"github.com/KonishiLee/wechat-crypter"
	"github.com/KonishiLee/wechat-qy/base"
)

// 应用套件相关操作的 API 地址
const (
	suiteTokenURI    = "https://qyapi.weixin.qq.com/cgi-bin/service/get_suite_token"
	preAuthCodeURI   = "https://qyapi.weixin.qq.com/cgi-bin/service/get_pre_auth_code"
	authURI          = "https://qy.weixin.qq.com/cgi-bin/loginpage"
	permanentCodeURI = "https://qyapi.weixin.qq.com/cgi-bin/service/get_permanent_code"
	authInfoURI      = "https://qyapi.weixin.qq.com/cgi-bin/service/get_auth_info"
	getAgentURI      = "https://qyapi.weixin.qq.com/cgi-bin/service/get_agent"
	setAgentURI      = "https://qyapi.weixin.qq.com/cgi-bin/service/set_agent"
	corpTokenURI     = "https://qyapi.weixin.qq.com/cgi-bin/service/get_corp_token"
)

// Suite 结构体包含了应用套件的相关操作
type Suite struct {
	id             string
	secret         string
	ticket         string
	token          string
	encodingAESKey string
	msgCrypter     crypter.MessageCrypter
	tokener        *base.Tokener
	client         *base.Client
}

func (s *Suite) GetChatAuthURI(appIDs []int, redirectURI, state string) (string, error) {
	preAuthCodeInfo, err := s.getPreAuthCode(appIDs)
	if err != nil {
		return "", err
	}

	qs := url.Values{}
	qs.Add("suite_id", s.id)
	qs.Add("pre_auth_code", preAuthCodeInfo.Code)
	qs.Add("redirect_uri", redirectURI)
	qs.Add("state", state)

	return authURI + "?" + qs.Encode(), nil
}
