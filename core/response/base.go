package response

import (
	"github.com/go-resty/resty/v2"
)

type Response struct {
	RespChannel   chan *resty.Response
	FuzzResponses *FuzzResponses
}

type FuzzResponses struct {
	Responses []resty.Response
}

type Matcher struct {
	MatcherType         string
	ExprType            string
	Scope               string
	Condition           string
	Group               []string
	GroupEncoding       string
	Negative            bool
	SubMatcherCondition string
	SubMatchers         []*Matcher
}
