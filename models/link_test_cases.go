package models

import "crawler/errors"

var testCasesNewLink = []struct {
	URL      string
	Hostname string
	Error    error
}{
	{
		"example.com",
		"",
		errors.INVALID_URL,
	},
	{
		"http://example.com",
		"example.com",
		nil,
	},
}

var testCasesParseURL = []struct {
	URL      string
	Hostname string
	Ok       bool
}{
	{
		"example.com",
		"",
		false,
	},
	{
		"ws://example.com",
		"",
		false,
	},
	{
		"http://example.com",
		"example.com",
		true,
	},
	{
		"https://example.com",
		"example.com",
		true,
	},
	{
		"http://example.com/",
		"example.com",
		true,
	},
	{
		"http://example.com/some/path?a=1&b=2#qwerty",
		"example.com",
		true,
	},
	{
		"http://external.asd1230-123.asd_internal.asd.gm-_ail.com/some/path",
		"external.asd1230-123.asd_internal.asd.gm-_ail.com",
		true,
	},
	{
		"http://abc",
		"example.com",
		false,
	},
	{
		"ws://abc",
		"",
		false,
	},
	{
		"http://!@#$%.com/some/path",
		"",
		false,
	},
	{
		"",
		"",
		false,
	},
	{
		"http:/ya.ru",
		"",
		false,
	},
}
