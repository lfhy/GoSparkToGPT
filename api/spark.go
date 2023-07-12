package api

import (
	"encoding/json"
	"fmt"
	"sparktogpt/conf"
	"sparktogpt/types"

	"github.com/gorilla/websocket"
)

var SparkURL string
var wsParam *types.SparkWsParam

func GetSparkUrl() string {
	wsParam = types.NewSparkWsParam(conf.APPID, conf.APIKey, conf.APISecret, conf.APIURL)
	SparkURL = wsParam.GetURL()
	return SparkURL
}

func ChatToSpark(question []types.CompletionsMessage) string {
	conn, _, err := websocket.DefaultDialer.Dial(GetSparkUrl(), nil)
	if err != nil {
		fmt.Println("### error:", err)
		return err.Error()
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Error occurred while closing connection:", err)
		}
	}(conn)

	done := make(chan struct{})
	var req string
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("### closed ###")
				return
			}
			req += onMessage(message)
		}
	}()

	data := genParams(wsParam.APPID, question)
	err = conn.WriteJSON(data)
	if err != nil {
		fmt.Println("### error:", err)
		return err.Error()
	}

	<-done
	return req
}

// 生成讯飞传输参数
func genParams(appID string, question []types.CompletionsMessage) map[string]interface{} {
	data := map[string]interface{}{
		"header": map[string]interface{}{
			"app_id": appID,
			"uid":    "1234",
		},
		"parameter": map[string]interface{}{
			"chat": map[string]interface{}{
				"domain":           "general",
				"random_threshold": 0.5,
				"max_tokens":       2048,
				"auditing":         "default",
			},
		},
		"payload": map[string]interface{}{
			"message": map[string]interface{}{
				"text": question,
			},
		},
	}

	return data
}

func onMessage(message []byte) string {
	var data map[string]interface{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		fmt.Println("### error:", err)
		return err.Error()
	}

	header := data["header"].(map[string]interface{})
	code := int(header["code"].(float64))
	if code != 0 {
		fmt.Printf("请求错误: %d, %s\n", code, data)
		return err.Error()
	}

	choices := data["payload"].(map[string]interface{})["choices"].(map[string]interface{})
	status := int(choices["status"].(float64))
	content := choices["text"].([]interface{})[0].(map[string]interface{})["content"].(string)
	fmt.Print(content)

	if status == 2 {
		fmt.Println()
	}
	return content
}
