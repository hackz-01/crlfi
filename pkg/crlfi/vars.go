package crlfi

var (
	appendList = [...]string{
		"",
		"crlfi",
		"?crlfi=",
		"#",
	}

	escapeList = [...]string{
		"%00",
		"%0a",
		"%0a%20",
		"%0d",
		"%0d%09",
		"%0d%0a",
		"%0d%0a%09",
		"%0d%0a%20",
		"%0d%20",
		"%20",
		"%20%0a",
		"%20%0d",
		"%20%0d%0a",
		"%23%0a",
		"%23%0a%20",
		"%23%0d",
		"%23%0d%0a",
		"%23%oa",
		"%25%30",
		"%25%30%61",
		"%2e%2e%2f%0d%0a",
		"%2f%2e%2e%0d%0a",
		"%2f..%0d%0a",
		"%3f",
		"%3f%0a",
		"%3f%0d",
		"%3f%0d%0a",
		"%e5%98%8a%e5%98%8d",
		"%e5%98%8a%e5%98%8d%0a",
		"%e5%98%8a%e5%98%8d%0d",
		"%e5%98%8a%e5%98%8d%0d%0a",
		"%e5%98%8a%e5%98%8d%e5%98%8a%e5%98%8d",
		"%u0000",
		"%u000a",
		"%u000d",
		"\r",
		"\r%20",
		"\r\n",
		"\r\n%20",
		"\r\n\t",
		"\r\t",
	}
)
