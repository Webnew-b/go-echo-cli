package jsonUtil

import "encoding/json"

type JsonError struct {
}

func (j *JsonError) Error() string {
	return "json convert error"
}

func JsonMarshal(data interface{}) ([]byte, *JsonError) {
	conversion, err := json.Marshal(data)
	if err != nil {
		return nil, &JsonError{}
	}
	return conversion, nil
}

func JsonUnmarshal(data []byte, structType any) *JsonError {
	err := json.Unmarshal(data, structType)
	if err != nil {
		return &JsonError{}
	}
	return nil
}
