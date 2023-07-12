package conf

import "flag"

var (
	Port      string
	APPID     string
	APIKey    string
	APISecret string
	APIURL    string
)

func init() {
	flag.StringVar(&Port, "port", "1234", "监听端口")
	flag.StringVar(&APPID, "appid", "", "讯飞控制台星火大模型页面的AppID")
	flag.StringVar(&APIKey, "apikey", "", "讯飞控制台星火大模型页面的APIKey")
	flag.StringVar(&APISecret, "apisecret", "", "讯飞控制台星火大模型页面的APISecret")
	flag.StringVar(&APIURL, "url", "ws://spark-api.xf-yun.com/v1.1/chat", "讯飞星火大模型接入地址")
	flag.Parse()
}
