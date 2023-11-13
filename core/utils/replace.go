package utils

var ReplaceTagTypeMap = map[string]ReplaceTag{
	"=":    new(AfterAndDataReplaceTag),
	"xml":  new(XMLDataReplaceTag),
	"json": new(JsonReplaceTag),
}

func ReplaceHeaderFuzzTag(packet *ParsedPacket) *ParsedPacket {
	if ReplaceType, ok := ReplaceTagTypeMap["="]; ok {
		packet.Headers["Cookie"] = ReplaceType.ReplaceTag(packet.Headers["Cookie"])

		packet.Headers["cstf-token"] = ReplaceType.ReplaceTag(packet.Headers["cstf-token"])

	}
	return packet
}
func ReplaceBodyFuzzTag(packet *ParsedPacket) *ParsedPacket {
	packet.Body = ReplaceTagTypeMap["="].ReplaceTag(packet.Body)
	packet.Body = ReplaceTagTypeMap["xml"].ReplaceTag(packet.Body)
	packet.Body = ReplaceTagTypeMap["json"].ReplaceTag(packet.Body)
	return packet
}
