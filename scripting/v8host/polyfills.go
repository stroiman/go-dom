package v8host

import (
	_ "embed"
	"errors"
)

//go:embed polyfills/xpath.js
var xpath []byte

func installPolyfills(context *V8ScriptContext) error {
	errs := []error{
		context.Run(`
		FormData.prototype.forEach = function(cb) {
			return Array.from(this).forEach(([k,v]) => { cb(v,k) })
		}
	`),
		context.Run(string(xpath)),
		context.Run(`
			const { XPathExpression, XPathResult } = window;
			const evaluate = XPathExpression.prototype.evaluate;
			XPathExpression.prototype.evaluate = function (context, type, res) {
				return evaluate.call(this, context, type ?? XPathResult.ANY_TYPE, res);
			};
			Element.prototype.scrollIntoView = function() {};
			Node.ELEMENT_NODE = 1;
			Node.ATTRIBUTE_NODE = 2;
			Node.TEXT_NODE = 3;
			Node.CDATA_SECTION_NODE = 4;
			Node.ENTITY_REFERENCE_NODE = 5;
			Node.ENTITY_NODE = 6;
			Node.PROCESSING_INSTRUCTION_NODE = 7;
			Node.COMMENT_NODE = 8;
			Node.DOCUMENT_NODE = 9;
			Node.DOCUMENT_TYPE_NODE = 10;
			Node.DOCUMENT_FRAGMENT_NODE = 11;
			Node.NOTATION_NODE = 12;
			Node.DOCUMENT_POSITION_DISCONNECTED = 0x01;
			Node.DOCUMENT_POSITION_PRECEDING = 0x02;
			Node.DOCUMENT_POSITION_FOLLOWING = 0x04;
			Node.DOCUMENT_POSITION_CONTAINS = 0x08;
			Node.DOCUMENT_POSITION_CONTAINED_BY = 0x10;
			Node.DOCUMENT_POSITION_IMPLEMENTATION_SPECIFIC = 0x20;

	`),
	}
	return errors.Join(errs...)
}
