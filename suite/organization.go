package suite

import (
	"net/url"

	"github.com/KonishiLee/wechat-crypter"
	"github.com/KonishiLee/wechat-qy/base"
)

// 应用套件相关操作的 API 地址

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

func (s *Suite) GetOrganizationAuthURI(appIDs []int, redirectURI, state string) (string, error) {
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
