package utils

import (
	"errors"
	"regexp"
	"strings"
)

type ParsedPacket struct {
	Method      string
	URL         string
	Protocol    string
	Host        string
	Body        string
	Headers     map[string]string
	QueryParams map[string]string
}

var (
	requestLineRegex = regexp.MustCompile(`^([A-Z]+) (\S+) (HTTP/\d\.\d)$`)
	headerRegex      = regexp.MustCompile(`^([^:]+): (.+)$`)
)

func ParsePacket(packet string) (*ParsedPacket, error) {
	parsedPacket := &ParsedPacket{
		Headers:     make(map[string]string),
		QueryParams: make(map[string]string),
	}
	lines := strings.Split(packet, "\n")
	requestLineMatches := requestLineRegex.FindStringSubmatch(lines[0])
	if len(requestLineMatches) != 4 {
		return nil, errors.New("invalid request line")
	}
	parsedPacket.Method = requestLineMatches[1]
	parsedPacket.URL = requestLineMatches[2]
	parsedPacket.Protocol = requestLineMatches[3]
	var currentPart *string
	for _, line := range lines[1:] {
		if line == "" && currentPart == nil {
			currentPart = &parsedPacket.Body
			continue
		}
		if currentPart == nil {
			if matches := headerRegex.FindStringSubmatch(line); matches != nil {
				parsedPacket.Headers[matches[1]] = matches[2]
			}
			continue
		}
		*currentPart += line + "\n"
	}
	if host, ok := parsedPacket.Headers["Host"]; ok {
		parsedPacket.Host = host
	}
	if parsedPacket.URL != "" {
		urlParts := strings.SplitN(parsedPacket.URL, "?", 2)
		parsedPacket.URL = urlParts[0]
		if len(urlParts) > 1 {
			queryParams := strings.Split(urlParts[1], "&")
			for _, param := range queryParams {
				keyValue := strings.SplitN(param, "=", 2)
				if len(keyValue) == 2 {
					parsedPacket.QueryParams[keyValue[0]] = keyValue[1]
				}
			}
		}
	}
	parsedPacket.Body = strings.TrimSuffix(parsedPacket.Body, "\n")
	for key, value := range parsedPacket.Headers {
		parsedPacket.Headers[key] = strings.TrimSuffix(value, "\n")
	}
	return parsedPacket, nil
}

type ParsedRequestLine struct {
	Method   string
	URL      string
	Protocol string
	Body     string
}
