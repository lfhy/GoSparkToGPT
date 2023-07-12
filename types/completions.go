package types

import "encoding/json"

// 请求体
type CompletionsRequest struct {
	Accept           string               `json:"Accept"`
	ContentType      string               `json:"Content-Type"`
	FrequencyPenalty *float64             `json:"frequency_penalty,omitempty"` // -2.0 和 2.0 之间的数字。正值会根据新标记在文本中的现有频率对其进行惩罚，从而降低模型逐字重复同一行的可能性。; [查看有关频率和存在惩罚的更多信息。](https://platform.openai.com/docs/api-reference/parameter-details)
	LogitBias        interface{}          `json:"logit_bias"`                  // 修改指定标记出现在完成中的可能性。  接受一个 json 对象，该对象将标记（由标记器中的标记 ID 指定）映射到从 -100 到 100; 的关联偏差值。从数学上讲，偏差会在采样之前添加到模型生成的 logits 中。确切的效果因模型而异，但 -1 和 1 之间的值应该会减少或增加选择的可能性；像 -100 或; 100 这样的值应该导致相关令牌的禁止或独占选择。
	MaxTokens        *int64               `json:"max_tokens,omitempty"`        // 聊天完成时生成的最大令牌数。  输入标记和生成标记的总长度受模型上下文长度的限制。
	Messages         []CompletionsMessage `json:"messages"`                    // 以[聊天格式](https://platform.openai.com/docs/guides/chat/introduction)生成聊天完成的消息。
	Model            string               `json:"model"`                       // 要使用的模型的 ID。有关哪些模型适用于聊天 API; 的详细信息，请参阅[模型端点兼容性表。](https://platform.openai.com/docs/models/model-endpoint-compatibility)
	N                *int64               `json:"n,omitempty"`                 // 为每个输入消息生成多少个聊天完成选项。
	PresencePenalty  *float64             `json:"presence_penalty,omitempty"`  // -2.0 和 2.0 之间的数字。正值会根据到目前为止是否出现在文本中来惩罚新标记，从而增加模型谈论新主题的可能性。; [查看有关频率和存在惩罚的更多信息。](https://platform.openai.com/docs/api-reference/parameter-details)
	Stop             *string              `json:"stop,omitempty"`              // API 将停止生成更多令牌的最多 4 个序列。
	Stream           *bool                `json:"stream,omitempty"`            // 如果设置，将发送部分消息增量，就像在 ChatGPT; 中一样。当令牌可用时，令牌将作为纯数据[服务器发送事件](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format)`data:; [DONE]`发送，流由消息终止。[有关示例代码](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_stream_completions.ipynb)，请参阅; OpenAI Cookbook 。
	Temperature      *int64               `json:"temperature,omitempty"`       // 使用什么采样温度，介于 0 和 2 之间。较高的值（如 0.8）将使输出更加随机，而较低的值（如 0.2）将使输出更加集中和确定。; 我们通常建议改变这个或`top_p`但不是两者。
	TopP             *int64               `json:"top_p,omitempty"`             // 一种替代温度采样的方法，称为核采样，其中模型考虑具有 top_p 概率质量的标记的结果。所以 0.1 意味着只考虑构成前 10% 概率质量的标记。; 我们通常建议改变这个或`temperature`但不是两者。
	User             *string              `json:"user,omitempty"`              // 代表您的最终用户的唯一标识符，可以帮助 OpenAI; 监控和检测滥用行为。[了解更多](https://platform.openai.com/docs/guides/safety-best-practices/end-user-ids)。
}

type CompletionsMessage struct {
	Content string `json:"content,omitempty"`
	Role    string `json:"role,omitempty"`
}

// 返回体
type CompletionsResponse struct {
	Choices []CompletionsChoice `json:"choices"`
	Created int64               `json:"created"`
	ID      string              `json:"id"`
	Object  string              `json:"object"`
	Usage   CompletionsUsage    `json:"usage"`
}

type CompletionsChoice struct {
	FinishReason string             `json:"finish_reason,omitempty"`
	Index        int64              `json:"index,omitempty"`
	Message      CompletionsMessage `json:"message,omitempty"`
}

type CompletionsUsage struct {
	CompletionTokens int `json:"completion_tokens"`
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func (chunk *CompletionsResponse) String() string {
	resp, _ := json.Marshal(chunk)
	return string(resp)
}
