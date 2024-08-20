package controller

func jsonError(message string) string {
	return `{"message": "` + message + `"}`
}
