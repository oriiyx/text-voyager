package main

import (
	"net/url"
	"time"

	"github.com/oriiyx/text-voyager/formatter"
	"github.com/oriiyx/text-voyager/parser"
)

type Request struct {
	SearchQuery          string
	Url                  string
	Data                 string
	Headers              string
	ResponseHeaders      string
	RawResponseBody      []byte
	ContentType          string
	Duration             time.Duration
	Formatter            formatter.ResponseFormatter
	GoogleParser         parser.GoogleResponseParser
	ResultNavigationData []ResultNavigationData
}

func (r Request) URLEncodedSearchQuery() string {
	return url.QueryEscape(r.SearchQuery)
}
