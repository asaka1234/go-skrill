package go_skrill

func getHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"charset":      "utf-8",
	}
}
