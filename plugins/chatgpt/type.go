package chatgpt

type ReplyMessage struct {
	Replytext    string   `json:"replytext"`
	Replyurl     []string `json:"replyurl"`
	ReplyContent string   `json:"content"`
}
type RequestData struct {
	AppCode  string    `json:"app_code"`
	Messages []Message `json:"messages"`
}

type ResponseData struct {
	Choices []Choices   `json:"choices"`
	Usage   Usage       `json:"usage"`
	Model   interface{} `json:"model"`
	TraceID string      `json:"trace_id"`
	Agent   Agent       `json:"agent"`
}
type Message struct {
	Role    string      `json:"role"`
	Content string      `json:"content"`
	Refusal interface{} `json:"refusal"`
}

type FinishDetails struct {
	Type string `json:"type"`
	Stop string `json:"stop"`
}
type Choices struct {
	Index         int           `json:"index"`
	Message       Message       `json:"message"`
	Logprobs      interface{}   `json:"logprobs"`
	FinishDetails FinishDetails `json:"finish_details"`
	ImgUrls       []string      `json:"img_urls"`
	TextContent   string        `json:"text_content"`
}
type PromptTokensDetails struct {
	CachedTokens int `json:"cached_tokens"`
}
type CompletionTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"`
}
type Usage struct {
	PromptTokens            int                     `json:"prompt_tokens"`
	CompletionTokens        int                     `json:"completion_tokens"`
	TotalTokens             int                     `json:"total_tokens"`
	PromptTokensDetails     PromptTokensDetails     `json:"prompt_tokens_details"`
	CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"`
}
type Chain struct {
	PluginName  string `json:"plugin_name"`
	PluginIcon  string `json:"plugin_icon"`
	PluginInput string `json:"plugin_input"`
}
type Agent struct {
	Status          string  `json:"status"`
	Chain           []Chain `json:"chain"`
	NeedShowPlugin  bool    `json:"need_show_plugin"`
	NeedShowThought bool    `json:"need_show_thought"`
}
