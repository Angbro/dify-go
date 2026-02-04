package dify

// Usage 表示 token 使用量
type Usage struct {
	PromptTokens        int     `json:"prompt_tokens"`
	PromptUnitPrice     string  `json:"prompt_unit_price"`
	PromptPriceUnit     string  `json:"prompt_price_unit"`
	PromptPrice         string  `json:"prompt_price"`
	CompletionTokens    int     `json:"completion_tokens"`
	CompletionUnitPrice string  `json:"completion_unit_price"`
	CompletionPriceUnit string  `json:"completion_price_unit"`
	CompletionPrice     string  `json:"completion_price"`
	TotalTokens         int     `json:"total_tokens"`
	TotalPrice          string  `json:"total_price"`
	Currency            string  `json:"currency"`
	Latency             float64 `json:"latency"`
}

// RetrieverResource 表示知识库引用资源
type RetrieverResource struct {
	Position     int     `json:"position"`
	DatasetID    string  `json:"dataset_id"`
	DatasetName  string  `json:"dataset_name"`
	DocumentID   string  `json:"document_id"`
	DocumentName string  `json:"document_name"`
	SegmentID    string  `json:"segment_id"`
	Score        float64 `json:"score"`
	Content      string  `json:"content"`
}

// Metadata 响应元数据
type Metadata struct {
	Usage              Usage               `json:"usage"`
	RetrieverResources []RetrieverResource `json:"retriever_resources,omitempty"`
}

// FileInput 文件输入
type FileInput struct {
	Type           string `json:"type"`
	TransferMethod string `json:"transfer_method"`
	URL            string `json:"url,omitempty"`
	UploadFileID   string `json:"upload_file_id,omitempty"`
}

// ========== Chat 相关类型 ==========

// ChatRequest 对话消息请求
type ChatRequest struct {
	Query            string                 `json:"query"`
	Inputs           map[string]interface{} `json:"inputs,omitempty"`
	ResponseMode     string                 `json:"response_mode"`
	User             string                 `json:"user"`
	ConversationID   string                 `json:"conversation_id,omitempty"`
	Files            []FileInput            `json:"files,omitempty"`
	AutoGenerateName bool                   `json:"auto_generate_name,omitempty"`
}

// ChatResponse 对话消息响应 (blocking 模式)
type ChatResponse struct {
	MessageID      string   `json:"message_id"`
	ConversationID string   `json:"conversation_id"`
	Mode           string   `json:"mode"`
	Answer         string   `json:"answer"`
	Metadata       Metadata `json:"metadata"`
	CreatedAt      int64    `json:"created_at"`
}

// Conversation 会话信息
type Conversation struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Inputs    map[string]interface{} `json:"inputs"`
	Status    string                 `json:"status"`
	CreatedAt int64                  `json:"created_at"`
	UpdatedAt int64                  `json:"updated_at"`
}

// ConversationListResponse 会话列表响应
type ConversationListResponse struct {
	Data    []Conversation `json:"data"`
	HasMore bool           `json:"has_more"`
	Limit   int            `json:"limit"`
}

// Message 消息信息
type Message struct {
	ID                 string                 `json:"id"`
	ConversationID     string                 `json:"conversation_id"`
	Inputs             map[string]interface{} `json:"inputs"`
	Query              string                 `json:"query"`
	Answer             string                 `json:"answer"`
	MessageFiles       []MessageFile          `json:"message_files"`
	Feedback           *Feedback              `json:"feedback"`
	RetrieverResources []RetrieverResource    `json:"retriever_resources"`
	CreatedAt          int64                  `json:"created_at"`
}

// MessageFile 消息文件
type MessageFile struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	URL       string `json:"url"`
	BelongsTo string `json:"belongs_to"`
}

// MessageListResponse 消息列表响应
type MessageListResponse struct {
	Data    []Message `json:"data"`
	HasMore bool      `json:"has_more"`
	Limit   int       `json:"limit"`
}

// Feedback 反馈信息
type Feedback struct {
	Rating string `json:"rating"`
}

// FeedbackRequest 反馈请求
type FeedbackRequest struct {
	Rating string `json:"rating"`
	User   string `json:"user"`
}

// FeedbackResponse 反馈响应
type FeedbackResponse struct {
	Result string `json:"result"`
}

// RenameRequest 重命名请求
type RenameRequest struct {
	Name         string `json:"name,omitempty"`
	AutoGenerate bool   `json:"auto_generate,omitempty"`
	User         string `json:"user"`
}

// RenameResponse 重命名响应
type RenameResponse struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Inputs    map[string]interface{} `json:"inputs"`
	Status    string                 `json:"status"`
	CreatedAt int64                  `json:"created_at"`
	UpdatedAt int64                  `json:"updated_at"`
}

// SuggestedResponse 建议问题响应
type SuggestedResponse struct {
	Result string   `json:"result"`
	Data   []string `json:"data"`
}

// ========== Completion 相关类型 ==========

// CompletionRequest 文本生成请求
type CompletionRequest struct {
	Inputs       map[string]interface{} `json:"inputs"`
	ResponseMode string                 `json:"response_mode"`
	User         string                 `json:"user"`
	Files        []FileInput            `json:"files,omitempty"`
}

// CompletionResponse 文本生成响应 (blocking 模式)
type CompletionResponse struct {
	MessageID string   `json:"message_id"`
	Mode      string   `json:"mode"`
	Answer    string   `json:"answer"`
	Metadata  Metadata `json:"metadata"`
	CreatedAt int64    `json:"created_at"`
}

// ========== Workflow 相关类型 ==========

// WorkflowRequest 工作流执行请求
type WorkflowRequest struct {
	Inputs       map[string]interface{} `json:"inputs"`
	ResponseMode string                 `json:"response_mode"`
	User         string                 `json:"user"`
	Files        []FileInput            `json:"files,omitempty"`
}

// WorkflowResponse 工作流执行响应 (blocking 模式)
type WorkflowResponse struct {
	WorkflowRunID string       `json:"workflow_run_id"`
	TaskID        string       `json:"task_id"`
	Data          WorkflowData `json:"data"`
}

// WorkflowData 工作流数据
type WorkflowData struct {
	ID          string                 `json:"id"`
	WorkflowID  string                 `json:"workflow_id"`
	Status      string                 `json:"status"`
	Outputs     map[string]interface{} `json:"outputs"`
	Error       string                 `json:"error,omitempty"`
	ElapsedTime float64                `json:"elapsed_time"`
	TotalTokens int                    `json:"total_tokens"`
	TotalSteps  int                    `json:"total_steps"`
	CreatedAt   int64                  `json:"created_at"`
	FinishedAt  int64                  `json:"finished_at"`
}

// WorkflowRunResponse 工作流执行状态响应
type WorkflowRunResponse struct {
	ID          string                 `json:"id"`
	WorkflowID  string                 `json:"workflow_id"`
	Status      string                 `json:"status"`
	Inputs      map[string]interface{} `json:"inputs"`
	Outputs     map[string]interface{} `json:"outputs"`
	Error       string                 `json:"error,omitempty"`
	TotalSteps  int                    `json:"total_steps"`
	TotalTokens int                    `json:"total_tokens"`
	CreatedAt   int64                  `json:"created_at"`
	FinishedAt  int64                  `json:"finished_at"`
	ElapsedTime float64                `json:"elapsed_time"`
}

// ========== 流式事件类型 ==========

// StreamEvent 流式事件基础结构
type StreamEvent struct {
	Event string `json:"event"`
}

// MessageStreamEvent 消息流事件
type MessageStreamEvent struct {
	Event          string `json:"event"`
	TaskID         string `json:"task_id"`
	MessageID      string `json:"message_id"`
	ConversationID string `json:"conversation_id,omitempty"`
	Answer         string `json:"answer"`
	CreatedAt      int64  `json:"created_at"`
}

// MessageEndStreamEvent 消息结束流事件
type MessageEndStreamEvent struct {
	Event          string   `json:"event"`
	TaskID         string   `json:"task_id"`
	MessageID      string   `json:"message_id"`
	ConversationID string   `json:"conversation_id,omitempty"`
	Metadata       Metadata `json:"metadata"`
}

// WorkflowStartedEvent 工作流开始事件
type WorkflowStartedEvent struct {
	Event         string              `json:"event"`
	TaskID        string              `json:"task_id"`
	WorkflowRunID string              `json:"workflow_run_id"`
	Data          WorkflowStartedData `json:"data"`
}

// WorkflowStartedData 工作流开始数据
type WorkflowStartedData struct {
	ID          string `json:"id"`
	WorkflowID  string `json:"workflow_id"`
	SequenceNum int    `json:"sequence_number"`
	CreatedAt   int64  `json:"created_at"`
}

// NodeStartedEvent 节点开始事件
type NodeStartedEvent struct {
	Event         string          `json:"event"`
	TaskID        string          `json:"task_id"`
	WorkflowRunID string          `json:"workflow_run_id"`
	Data          NodeStartedData `json:"data"`
}

// NodeStartedData 节点开始数据
type NodeStartedData struct {
	ID                string                 `json:"id"`
	NodeID            string                 `json:"node_id"`
	NodeType          string                 `json:"node_type"`
	Title             string                 `json:"title"`
	Index             int                    `json:"index"`
	PredecessorNodeID string                 `json:"predecessor_node_id"`
	Inputs            map[string]interface{} `json:"inputs"`
	CreatedAt         int64                  `json:"created_at"`
}

// NodeFinishedEvent 节点完成事件
type NodeFinishedEvent struct {
	Event         string           `json:"event"`
	TaskID        string           `json:"task_id"`
	WorkflowRunID string           `json:"workflow_run_id"`
	Data          NodeFinishedData `json:"data"`
}

// NodeFinishedData 节点完成数据
type NodeFinishedData struct {
	ID                string                 `json:"id"`
	NodeID            string                 `json:"node_id"`
	NodeType          string                 `json:"node_type"`
	Title             string                 `json:"title"`
	Index             int                    `json:"index"`
	PredecessorNodeID string                 `json:"predecessor_node_id"`
	Inputs            map[string]interface{} `json:"inputs"`
	ProcessData       map[string]interface{} `json:"process_data"`
	Outputs           map[string]interface{} `json:"outputs"`
	Status            string                 `json:"status"`
	Error             string                 `json:"error,omitempty"`
	ElapsedTime       float64                `json:"elapsed_time"`
	ExecutionMetadata map[string]interface{} `json:"execution_metadata"`
	CreatedAt         int64                  `json:"created_at"`
}

// WorkflowFinishedEvent 工作流完成事件
type WorkflowFinishedEvent struct {
	Event         string       `json:"event"`
	TaskID        string       `json:"task_id"`
	WorkflowRunID string       `json:"workflow_run_id"`
	Data          WorkflowData `json:"data"`
}

// ErrorStreamEvent 错误流事件
type ErrorStreamEvent struct {
	Event     string `json:"event"`
	TaskID    string `json:"task_id"`
	MessageID string `json:"message_id,omitempty"`
	Status    int    `json:"status"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

// ========== 通用接口类型 ==========

// FileUploadResponse 文件上传响应
type FileUploadResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Extension string `json:"extension"`
	MimeType  string `json:"mime_type"`
	CreatedBy string `json:"created_by"`
	CreatedAt int64  `json:"created_at"`
}

// AppParametersResponse 应用参数响应
type AppParametersResponse struct {
	OpeningStatement              string              `json:"opening_statement"`
	SuggestedQuestions            []string            `json:"suggested_questions"`
	SuggestedQuestionsAfterAnswer SuggestedQAConfig   `json:"suggested_questions_after_answer"`
	SpeechToText                  SpeechToTextConfig  `json:"speech_to_text"`
	TextToSpeech                  TextToSpeechConfig  `json:"text_to_speech"`
	RetrieverResource             RetrieverConfig     `json:"retriever_resource"`
	AnnotationReply               AnnotationConfig    `json:"annotation_reply"`
	UserInputForm                 []UserInputFormItem `json:"user_input_form"`
	FileUpload                    FileUploadConfig    `json:"file_upload"`
	SystemParameters              SystemParamsConfig  `json:"system_parameters"`
}

// SuggestedQAConfig 建议问题配置
type SuggestedQAConfig struct {
	Enabled bool `json:"enabled"`
}

// SpeechToTextConfig 语音转文字配置
type SpeechToTextConfig struct {
	Enabled bool `json:"enabled"`
}

// TextToSpeechConfig 文字转语音配置
type TextToSpeechConfig struct {
	Enabled  bool   `json:"enabled"`
	Voice    string `json:"voice"`
	Language string `json:"language"`
}

// RetrieverConfig 知识库配置
type RetrieverConfig struct {
	Enabled bool `json:"enabled"`
}

// AnnotationConfig 标注回复配置
type AnnotationConfig struct {
	Enabled bool `json:"enabled"`
}

// UserInputFormItem 用户输入表单项
type UserInputFormItem map[string]FormItemConfig

// FormItemConfig 表单项配置
type FormItemConfig struct {
	Label     string   `json:"label"`
	Variable  string   `json:"variable"`
	Required  bool     `json:"required"`
	Default   string   `json:"default"`
	MaxLength int      `json:"max_length,omitempty"`
	Options   []string `json:"options,omitempty"`
}

// FileUploadConfig 文件上传配置
type FileUploadConfig struct {
	Image ImageUploadConfig `json:"image"`
}

// ImageUploadConfig 图片上传配置
type ImageUploadConfig struct {
	Enabled         bool     `json:"enabled"`
	NumberLimits    int      `json:"number_limits"`
	TransferMethods []string `json:"transfer_methods"`
}

// SystemParamsConfig 系统参数配置
type SystemParamsConfig struct {
	FileSizeLimit      int `json:"file_size_limit"`
	ImageFileSizeLimit int `json:"image_file_size_limit"`
	AudioFileSizeLimit int `json:"audio_file_size_limit"`
	VideoFileSizeLimit int `json:"video_file_size_limit"`
}

// AppMetaResponse 应用元信息响应
type AppMetaResponse struct {
	ToolIcons map[string]interface{} `json:"tool_icons"`
}

// StopRequest 停止响应请求
type StopRequest struct {
	User string `json:"user"`
}

// StopResponse 停止响应
type StopResponse struct {
	Result string `json:"result"`
}

// AudioToTextResponse 语音转文字响应
type AudioToTextResponse struct {
	Text string `json:"text"`
}
