package xingchen

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ChatResultStream struct {
	response *http.Response
	isFinish bool
	reader   *bufio.Reader
}

var (
	eventPrefix = []byte("event:")
	dataPrefix  = []byte("data:")
)

func NewChatResultStream(response *http.Response) *ChatResultStream {
	chatResultStream := ChatResultStream{
		response: response,
		isFinish: true,
		reader:   bufio.NewReader(response.Body),
	}
	return &chatResultStream
}

func (o *ChatResultStream) Recv() (ResultDTOChatResultDTO, error) {
	return o.processLine()
}

func (o *ChatResultStream) processLine() (ResultDTOChatResultDTO, error) {
	var hasErr = false
	for {
		rawLine, err := o.reader.ReadBytes('\n')
		if err != nil {
			return *new(ResultDTOChatResultDTO), err
		}
		noSpaceLine := bytes.TrimSpace(rawLine)
		if bytes.HasPrefix(noSpaceLine, eventPrefix) {
			noPrefixLine := bytes.TrimPrefix(noSpaceLine, eventPrefix)
			if string(noPrefixLine) == "error" {
				hasErr = true
			}
		}
		if bytes.HasPrefix(noSpaceLine, dataPrefix) {
			noPrefixLine := bytes.TrimPrefix(noSpaceLine, dataPrefix)
			response := ResultDTOChatResultDTO{}
			if hasErr {
				err = json.Unmarshal(noPrefixLine, &response)
				hasErr = true
				o.isFinish = true
				return response, io.EOF
			}

			chatResult := ChatResultDTO{}
			err = json.Unmarshal(noPrefixLine, &chatResult)
			response.Data = &chatResult
			response.Code = PtrInt32(200)
			response.RequestId = chatResult.RequestId
			response.Success = PtrBool(true)
			if err != nil {
				return response, err
			}
			return response, nil
		}
	}
}

func (o *ChatResultStream) Close() error {
	err := o.response.Body.Close()
	return err
}
