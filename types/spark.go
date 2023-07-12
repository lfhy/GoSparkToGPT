package types

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

type SparkWsParam struct {
	APPID     string
	APIKey    string
	APISecret string
	Host      string
	Path      string
	URL       string
}

func (p *SparkWsParam) GetURL() string {
	now := time.Now().UTC()
	date := now.Format(time.RFC1123)

	signatureOrigin := fmt.Sprintf("host: %s\ndate: %s\nGET %s HTTP/1.1", p.Host, date, p.Path)
	h := hmac.New(sha256.New, []byte(p.APISecret))
	h.Write([]byte(signatureOrigin))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	authorizationOrigin := fmt.Sprintf(`api_key="%s", algorithm="hmac-sha256", headers="host date request-line", signature="%s"`, p.APIKey, signature)
	authorization := base64.StdEncoding.EncodeToString([]byte(authorizationOrigin))

	v := url.Values{}
	v.Set("authorization", authorization)
	v.Set("date", date)
	v.Set("host", p.Host)

	return fmt.Sprintf("%s?%s", p.URL, v.Encode())
}

// 新建讯飞连接参数
func NewSparkWsParam(APPID, APIKey, APISecret, URL string) *SparkWsParam {
	u, _ := url.Parse(URL)
	return &SparkWsParam{
		APPID:     APPID,
		APIKey:    APIKey,
		APISecret: APISecret,
		Host:      u.Host,
		Path:      u.Path,
		URL:       URL,
	}
}
