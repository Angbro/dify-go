package dify

import (
	"encoding/json"
	"fmt"
)

// APIError Dify API 错误
type APIError struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Status     int    `json:"status"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("dify api error: status=%d, code=%s, message=%s", e.StatusCode, e.Code, e.Message)
}

// ParseAPIError 解析 API 错误响应
func ParseAPIError(statusCode int, body []byte) error {
	var apiErr APIError
	if err := json.Unmarshal(body, &apiErr); err != nil {
		return fmt.Errorf("http error: status=%d, body=%s", statusCode, string(body))
	}
	apiErr.StatusCode = statusCode
	return &apiErr
}

// 常见错误码
const (
	ErrCodeNoAPIKey              = "no_api_key"
	ErrCodeInvalidAPIKey         = "invalid_api_key"
	ErrCodeAppUnavailable        = "app_unavailable"
	ErrCodeProviderNotInitialize = "provider_not_initialize"
	ErrCodeProviderQuotaExceeded = "provider_quota_exceeded"
	ErrCodeModelCurrentlyNotSupport = "model_currently_not_support"
	ErrCodeCompletionRequestError = "completion_request_error"
	ErrCodeNotFound              = "not_found"
	ErrCodeNotChatApp            = "not_chat_app"
	ErrCodeNotCompletionApp      = "not_completion_app"
	ErrCodeConversationCompleted = "conversation_completed"
	ErrCodeFileNotFound          = "file_not_found"
	ErrCodeFileTooLarge          = "file_too_large"
	ErrCodeUnsupportedFileType   = "unsupported_file_type"
	ErrCodeS3ConnectionFailed    = "s3_connection_failed"
	ErrCodeS3PermissionDenied    = "s3_permission_denied"
	ErrCodeS3FileTooLarge        = "s3_file_too_large"
)
