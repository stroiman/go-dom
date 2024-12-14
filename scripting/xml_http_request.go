package scripting

type JSXmlHttpRequest struct {
	host *ScriptHost
}

func NewJSXmlHttpRequest(host *ScriptHost) JSXmlHttpRequest {
	return JSXmlHttpRequest{host}
}
