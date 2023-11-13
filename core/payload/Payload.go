package payload

import (
	"ApricusFuzz/core/Encoder"
	"log"
	"strings"
)

var PayLoaderMap = map[string]payloadLoader{
	"list":  new(ListPayloadLoader),
	"stdin": new(StdinPayloadLoader),
	"file":  new(FilePayloadLoader),
	"range": new(RangePayloadLoader),
}
var EncoderMap = map[string]Encoder.Encoder{
	"default": new(Encoder.DefaultEncoder),
	"md5":     new(Encoder.Md5Encoder),
	"base64":  new(Encoder.Base64Encoder),
}

type Payload struct {
	Original string
	Data     []string
	Encoder  Encoder.Encoder
}

func (p *Payload) Load(loader payloadLoader, EncoderType string, value string) {
	for _, data := range loader.Load(value) {
		p.Data = append(p.Data, Encoder.EncodeLoader(EncoderType, data))
	}
}
func NewPayload(data string) Payload {
	var (
		SourceValue string
		LoadType    string
		payload     Payload
		EncoderType string
	)
	split := strings.Split(data, ",")
	splitLen := len(split)
	switch splitLen {
	case 1:
		LoadType = strings.ToLower(split[0])
		EncoderType = "default"
	case 2:
		LoadType = strings.ToLower(split[0])
		SourceValue = split[1]
		EncoderType = "default"
	case 3:
		LoadType = strings.ToLower(split[0])
		SourceValue = split[1]
		EncoderType = split[2]

	}
	if loader, ok := PayLoaderMap[LoadType]; ok {
		payload = Payload{
			Original: data,
		}
		payload.Load(loader, EncoderType, SourceValue)
	} else {
		log.Fatalf("Unknown 	Payloadertype Error : %s\n", data)
	}
	return payload
}
