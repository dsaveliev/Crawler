package server

var testCasesValidateBody = []struct {
	Body   string
	Result string
	Ok     bool
}{
	{
		"",
		"",
		false,
	},
	{
		"\n",
		"",
		false},
	{
		"\t\r\n\t \n  \t",
		"",
		false},
	{
		"http://ya.ru",
		"http://ya.ru",
		true},
	{
		" \n http://ya.ru\t \r ",
		"http://ya.ru",
		true},
	{
		"http://ya.ru\nhttp://google.com\n",
		"http://ya.ru\nhttp://google.com",
		true},
}
