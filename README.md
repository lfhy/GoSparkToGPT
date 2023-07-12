# GoSparkToGPT
用Go1.19写的将讯飞星火大模型转为chatGPT的completions接口返回体

# 构建
```
go mod tidy
go build -o sparktogpt main.go
```

# 使用方法
输入命令行参数后启动程序，访问```/v1/chat/completions```接口即可
```
sparktogpt 命令行参数:
  -apikey string
        讯飞控制台星火大模型页面的APIKey
  -apisecret string
        讯飞控制台星火大模型页面的APISecret
  -appid string
        讯飞控制台星火大模型页面的AppID
  -port string
        监听端口 (default "1234")
  -url string
        讯飞星火大模型接入地址 (default "ws://spark-api.xf-yun.com/v1.1/chat")
```