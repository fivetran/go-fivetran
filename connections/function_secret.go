package connections

type FunctionSecret struct {
	key   *string
	value *string
}

type functionSecretRequest struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type FunctionSecretResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (fs *FunctionSecret) request() *functionSecretRequest {
	return &functionSecretRequest{
		Key:   fs.key,
		Value: fs.value,
	}
}

func (fs *FunctionSecret) Key(value string) *FunctionSecret {
	fs.key = &value
	return fs
}

func (fs *FunctionSecret) Value(value string) *FunctionSecret {
	fs.value = &value
	return fs
}
