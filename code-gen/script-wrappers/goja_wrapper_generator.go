package wrappers

import "io/fs"

func NewGojaWrapperModuleGenerator(idlSources fs.FS) ScriptWrapperModulesGenerator {
	specs := NewWrapperGeneratorsSpec()
	// xhrModule := specs.Module("xhr")
	// xhr := xhrModule.Type("XMLHttpRequest")
	// xhr.SkipPrototypeRegistration = true
	// xhr.InnerTypeName = "XmlHttpRequest"
	// xhr.Receiver = "xhr"
	//
	// xhr.MarkMembersAsNotImplemented(
	// 	"readyState",
	// 	"responseType",
	// 	"responseXML",
	// )
	// xhr.Method("open").SetCustomImplementation()
	// xhr.Method("upload").SetCustomImplementation()
	// xhr.Method("getResponseHeader").HasNoError = true
	// xhr.Method("setRequestHeader").HasNoError = true
	//
	// urlSpecs := specs.Module("url")
	// url := urlSpecs.Type("URL")
	// url.InnerTypeName = "Url"
	// url.Receiver = "u"
	// url.MarkMembersAsNotImplemented(
	// 	"SetHref",
	// 	"SetProtocol",
	// 	"username",
	// 	"password",
	// 	"SetHost",
	// 	"SetPort",
	// 	"SetHostname",
	// 	"SetPathname",
	// 	"searchParams",
	// 	"SetHash",
	// 	"SetSearch",
	// )

	domSpecs := specs.Module("dom")
	domSpecs.SetMultipleFiles(true)

	// domTokenList := domSpecs.Type("DOMTokenList")
	// domTokenList.InnerTypeName = "DomTokenList"
	// domTokenList.Receiver = "u"
	// domTokenList.RunCustomCode = true
	// domTokenList.Method("item").SetNoError()
	// domTokenList.Method("contains").SetNoError()
	// domTokenList.Method("remove").SetNoError()
	// domTokenList.Method("toggle").SetCustomImplementation()
	// domTokenList.Method("replace").SetNoError()
	// domTokenList.Method("supports").SetNotImplemented()

	domNode := domSpecs.Type("Node")
	domNode.Method("nodeType").SetCustomImplementation()
	domNode.Method("contains").SetNoError()
	domNode.Method("getRootNode").SetNoError().Argument("options").HasDefault()
	domNode.Method("previousSibling").SetNoError()
	domNode.Method("nextSibling").SetNoError()

	domNode.Method("hasChildNodes").Ignore()
	domNode.Method("normalize").Ignore()
	domNode.Method("cloneNode").Ignore()
	domNode.Method("isEqualNode").Ignore()
	domNode.Method("isSameNode").Ignore()
	domNode.Method("compareDocumentPosition").Ignore()
	domNode.Method("lookupPrefix").Ignore()
	domNode.Method("lookupNamespaceURI").Ignore()
	domNode.Method("isDefaultNamespace").Ignore()
	domNode.Method("replaceChild").Ignore()
	domNode.Method("baseURI").Ignore()
	domNode.Method("parentNode").Ignore()
	domNode.Method("parentElement").Ignore()
	domNode.Method("lastChild").Ignore()
	domNode.Method("nodeValue").Ignore()
	domNode.Method("textContent").Ignore()

	// htmlSpecs := specs.Module("html")
	// htmlSpecs.SetMultipleFiles(true)
	//
	// htmlTemplateElement := htmlSpecs.Type("HTMLTemplateElement")
	// htmlTemplateElement.InnerTypeName = "HtmlTemplateElement"
	// htmlTemplateElement.Method("shadowRootMode").SetNotImplemented()
	// htmlTemplateElement.Method("shadowRootDelegatesFocus").SetNotImplemented()
	// htmlTemplateElement.Method("shadowRootClonable").SetNotImplemented()
	// htmlTemplateElement.Method("shadowRootSerializable").SetNotImplemented()
	//
	// window := htmlSpecs.Type("Window")
	// window.InnerTypeName = "Window"
	// window.CreateWrapper()
	//
	// window.Method("window").SetCustomImplementation()
	// window.Method("location").Ignore()
	// window.Method("parent").Ignore() // On `Node`
	//
	// window.Method("prompt").SetNotImplemented()
	// window.Method("close").SetNotImplemented()
	// window.Method("stop").SetNotImplemented()
	// window.Method("focus").SetNotImplemented()
	// window.Method("blur").SetNotImplemented()
	// window.Method("open").SetNotImplemented()
	// window.Method("alert").SetNotImplemented()
	// window.Method("confirm").SetNotImplemented()
	// window.Method("postMessage").SetNotImplemented()
	// window.Method("print").SetNotImplemented()
	// window.Method("self").SetNotImplemented()
	// window.Method("name").SetNotImplemented()
	// window.Method("personalbar").SetNotImplemented()
	// window.Method("locationbar").SetNotImplemented()
	// window.Method("menubar").SetNotImplemented()
	// window.Method("scrollbars").SetNotImplemented()
	// window.Method("statusbar").SetNotImplemented()
	// window.Method("status").SetNotImplemented()
	// window.Method("toolbar").SetNotImplemented()
	// window.Method("history").SetNotImplemented()
	// window.Method("navigation").SetNotImplemented()
	// window.Method("customElements").SetNotImplemented()
	// window.Method("closed").SetNotImplemented()
	// window.Method("frames").SetNotImplemented()
	// window.Method("navigator").SetNotImplemented()
	// window.Method("frames").SetNotImplemented()
	// window.Method("top").SetNotImplemented()
	// window.Method("opener").SetNotImplemented()
	// window.Method("frameElement").SetNotImplemented()
	// window.Method("clientInformation").SetNotImplemented()
	// window.Method("originAgentCluster").SetNotImplemented()
	// window.Method("length").SetNotImplemented()

	return ScriptWrapperModulesGenerator{
		IdlSources:       idlSources,
		Specs:            specs,
		PackagePath:      goja,
		TargetGenerators: GojaTargetGenerators{},
	}
}
