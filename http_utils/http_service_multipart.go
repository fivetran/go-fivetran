package httputils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
)

func (s HttpService) DoMultipart(
	ctx context.Context,
	method,
	url string,
	fieldName string,
	fileName string,
	fileContent io.Reader,
	queries map[string]string,
	expectedStatus int,
	response any) error {

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		return fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := io.Copy(part, fileContent); err != nil {
		return fmt.Errorf("failed to write file content: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close multipart writer: %w", err)
	}

	s.CommonHeaders["Content-Type"] = writer.FormDataContentType()

	r := Request{
		Method:           method,
		Url:              s.BaseUrl + url,
		Body:             buf.Bytes(),
		Queries:          queries,
		Headers:          s.CommonHeaders,
		Client:           s.Client,
		HandleRateLimits: s.HandleRateLimits,
		MaxRetryAttempts: s.MaxRetryAttempts,
	}

	respBody, respStatus, err := r.Do(ctx)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		return err
	}

	if respStatus != expectedStatus {
		return fmt.Errorf("status code: %v; expected: %v", respStatus, expectedStatus)
	}
	return nil
}
