# Feature list

Go-dom is still in a very early stage. This is a brief list of supported
features (which may be outdated because I forgot to update it)

Overall, the browser can parse the DOM and execute scripts, but has a very
limited DOM

**Do NOT expect a feature to work if not explicitly mentioned here**

## Handling of `<script>`

Go-dom supports blocking script execution; both inline, and downloaded. There is
no support for `async` and `defer` attributes; such scripts are still blocking.

ES modules are NOT supported yet. 

See also: [Script engines](./README_SCRIPT_ENGINE.md)

### execution

Scripts are executed while during DOM construction, so only already constructed
elements are accessible. But the entire response body has already been
processed, so implementing `document.write` wouldn't be possible without a
complete rewrite of the HTML parser.

> [!NOTE]
>
> If you create a `Window` directly, you need to pass a script engine for
> scripts to execute. If you create a `Browser`, a default script host will be
> created.

## Content loading

Cookies _should_ be supported, but is untested (there is no API to access the
cookie jar)

The browser errors if the server does not respond with a 200 (window loading,
not XHR requests). Redirects are not followed.

## DOM in general

- `Element.outerHTML` and `innerHTML` - works, but output is not escaped.

### No namespace support

There is no support for namespaces. While the document has a `CreateElementNS`,
that is so far _only_ implemented to prevent SVG elements from being constructed
as HTML elements. Other `...NS` functions do not exist, neither the ability to
query namespaces.

## HTML Elements

All elements have the correct class in JavaScript code, but not all element
behaviours are implemented. In general, if a type exists for an element in the
`html` package, some behaviour specific to that element is implemented.

### Forms

There is simple form support, 

- Calling `click()` on an `<input type="submit">` or `<button type="submit">`
submits the form.
- Form submits trigger a `formdata` event, but not `requestsubmit` event.
- Attributes on the submitter does not override the form's method/action
- Forms have a `submit` method, but not `requestsubmit` 

(requestsubmit is in dev, so maybe it exisst, but I didn't update this doc)

Clever input mechanisms are not supported, you must set the `value` attribute on
the field.

Reset behaviour is not implemented!

### Other elements

- `<a>` navigates to a new pages when calling `click()`.
- `<template>` has a `content` attribute containing a fragment where the HTML
child nodes are found.

## XHR, no fetch

There is an XHR implementation. It doesn't support all events yet. `send(body)`
only accepts a `FormBody`.

Fetch is not supported but you can add a polyfill to your project if you need
fetch, or you could make a PR:

- https://github.com/stroiman/go-dom/tree/main/browser/scripting/v8host/polyfills
- https://github.com/stroiman/go-dom/blob/main/browser/scripting/v8host/polyfills.go 
