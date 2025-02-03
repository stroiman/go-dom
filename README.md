# go-dom - Headless browser for Go

The Go-to headless browser for TDD workflows.

```go
browser := NewBrowserFromHandler(pkg.RootHttpHandler)
window, err := browser.Open("http://example.com/example") // host is ignored
Expect(err).ToNot(HaveOccurred())
doc := window.Document()
button := doc.QuerySelector("button")
targetArea := doc.GetElementById("target-area")
button.Click()
Expect(targetArea).To(HaveTextContent("Click count: 1"))
button.Click()
Expect(targetArea).To(HaveTextContent("Click count: 2"))
```

Go-dom downloads, and executes client-side script, making it an ideal choice to
help build applications using a Go/HTMX stack. 

Being written in Go you can connect directly to an `http.Handler` bypassing the
overhead of a TCP connection; as well as the burden of managing ports and
connections.

This greatly simplifies the ability to replace dependencies during testing, as
you can treat your HTTP server as a normal Go component.

> [!NOTE] 
>
> This still early pre-release. Minimal functionality exists for some basic
> flows, but only just.
>
> [Feature list](./README_FEATURE_LIST.md)

> [!IMPORTANT]
>
> This package currently requires module replacement, check out the
> [Installation](#Installation) section.

> [!WARNING]
>
> The API is not yet stable. use at your own risk.

## Looking for sponsors

If this tool could reach a minimum level of usability, this would be extremely
valuable in testing Go web applications, particularly combined with HTMX, a tech
combination which is becoming increasingly popular.

Progress so far is the result of too much spare time; but that will not last. If
If enough people would sponsor this project, it could mean the difference
between continued development, or death.

## Looking for contributors

This is a massive undertaking, and I would love people to join in.

Particularly if you have experience in the area of building actual browsers (I'm
not talking about skinning Chromium here, like, implementing an actual rendering
engine)

## Join the "community"

- [Join my discord server](https://discord.gg/rPBRt8Rf) to chat with me, and stay
up-to-date on progress
- Participate in the [github discussions](https://github.com/orgs/gost-dom/discussions), and [say
hi!](https://github.com/orgs/gost-dom/discussions)

## Installation

After go getting `github.com/gost-dom/browser`, you need replace some modules:

```sh
go mod edit -replace="github.com/ericchiang/css=github.com/gost-dom/css@latest"
go mod edit -replace="github.com/tommie/v8go=github.com/stroiman/v8go@go-dom-support"
go mod tidy
```

The CSS is just a simple fix that allows the CSS selector to accept tag name
patterns with upper case tag names, which HTMX produces. I've filed a PR, so
hopefully this will get merged to the source repo soon.

For the v8go project, I've added a lot of V8 features that were missing in v8go.
I'm working with tommie, who runs the [best maintained
branch](https://github.com/tommie/v8go), but it may be a while before they are
all merged.

> [!NOTE]
>
> New features _will probably_ be added to my branch, requiring the replacement
> to be updated. If you get build errors that look v8-ish, try running the
> replacement again. Tip: Create a shell script for this.

## Project background

Go and HTMX is gaining in popularity as a stack.

While Go has great tooling for verifying request/responses of HTTP applications,
but for HTMX, or just client-side scripting with server side rendering, you need
browser automation to test the behaviour.

This introduces a significant overhead; not only from out-of-process
communication with the browser, but also the necessity of launching your server.

This overhead discourages a TDD loop.

The purpose of this project is to enable a fast TDD feedback loop these types of
project, where verification depend on

- Behaviour of client-side scripts.
- Browser behaviour when interacting with browser elements, e.g., clicking the
  submit button submits a form, and redirects are followed.

### Unique features

Being written in Go, this library supports consuming an
[`http.Handler`](https://pkg.go.dev/net/http#Handler) directly. This removes the
necessity managing TCP ports, and start a server on a real port. Your HTTP
server is consumed by test code, like any other Go component would, also
allowing you to replace dependencies for the test if applicable.

This also makes it easy to run parallel tests in isolation as each can create
their own _instance_ of the HTTP handler.

### Drawbacks to Browser automation

- You cannot verify how it look; e.g. you cannot get a screenshot of a failing
test, nor use such screenshots for snapshot tests.
- The verification doesn't prove that it works as intended in _all browsers_ you
want to support.

This isn't intended as a replacement for the cases where an end-2-end test is
the right choice. It is intended as a tool to help when you want a smaller
isolated test, e.g. mocking out part of the behaviour;

## Code structure

This is still in early development, and the structure may still change.

```sh
dom/ # Core DOM implementation
html/ # Window, HTMLDocument, HTMLElement, 
scripting/ # Client-side script support
v8host/ # v8 engine, and bindings
gojahost/ # goja javascript engine,
browser.go # Main module
```

The folders, `dom`, and `html` correspond to the [web
APIs](https://developer.mozilla.org/en-US/docs/Web/API). It was the intention to
have a folder for each supported web API, but that may turn out to be
impossible, as there are circular dependencies between some of the specs.

### Modularisation

Although the code isn't modularised yet, it is an idea that you should be able
to include the modules relevant to your app. E.g., if your app deals with
location services, you can add a module implementing location services.

This helps keep the size of the dependencies down for client projects; keeping
build times down for the TDD loop.

It also provides the option of alternate implementations. E.g., for location
services, the simple implementation can provide a single function to set the
current location / accuracy. The advanced implementation can replay a GPX track.

## Project status

Currently, the most basic HTMX app is working, simple click handler with
swapping, boosted links, and form (at least with text fields).

### Memory Leaks

The current implementation is leaking memory for the scope of a browser
`Window`. I.e., all DOM nodes created and deleted for the lifetime of the
window will stay in memory until the window is actively disposed.

**This is not a problem for the intended use case**

#### Why memory leaks

This codebase is a marriage between two garbage collected runtimes, and what is
conceptually _one object_ is split into two, a Go object and a JavaScript
wrapper. As long of them is reachable; so must the other be.

I could join them into one; but that would result in an undesired coupling; the
DOM implementation being coupled to the JavaScript execution engine. Eventually,
a native Go JavaScript runtime will be supported.

A solution to this problem involves the use of weak references. This exists as
an `internal` but [was accepted](https://github.com/golang/go/issues/67552) as a
feature.

For that reason; and because it's not a problem for the intended use case, I
have postponed dealing with that issue.

### Next up

The following are main focus areas ATM

- Handle redirect responses
- Implement a proper event loop, with proper `setTimeout`, `setInterval`, and
their clear-counterparts.
- Implement fast-forwarding of time. 
- Replace early hand-written JS wrappers with auto-generated code, helping drive
  a more complete implementation.

A parallel project is adding support for [Goja](https://github.com/dop251/goja)
, to eventually replace V8 with Goja as the default script engine, resulting in
a pure Go implementation. V8 support will not go away, so there's a fallback, if
important JS features are lacking from Goja.

### Future goals

There is much to do, which includes (but this is not a full list):

- Support web-sockets and server events.
- Implement all standard JavaScript classes that a browser should support; but
  not part of the ECMAScript standard itself.
  - JavaScript polyfills would be a good starting point; which is how xpath is
    implemented at the moment.
    - Conversion to native go implementations would be prioritized on usage, e.g.
      [`fetch`](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API) 
      would be high in the list of priorities.
- Implement default browser behaviour for user interaction, e.g. pressing 
  <key>enter</key> when an input field has focus should submit the form.

### Long Term Goals

#### CSS Parsing

Parsing CSS woule be nice, allowing test code to verify the resulting styles of
an element; but having a working DOM with a JavaScript engine is higher
priority.

#### Mock external sites

The system may depend on external sites in the browser, most notably identity
providers (IDP), where your app redirects to the IDP, which redirects on
successful login; but could be other services such as map providers, etc.

For testing purposes, replacing this with a dummy replacement would have some
benefits:

- The verification of your system doesn't depend on the availability of an
  external service; when working offline
- Avoid tests breaking because of changes to the external system.
- For an identity provider
  - Avoid pollution of dummy accounts to run your test suite.
  - Avoid locking out test accounts due to _"suspiscious activity"_.
  - The IDP may use a Captcha or 2FA that can be impossible; or difficult to
    control from tests, and would cause a significant slowdown to the test
    suite.
- For applications like map providers
  - Avoid being billed for API use during testing.

## Out of scope.

### Full Spec Compliance

> A goal is not always meant to be reached, it often serves simply as something
> to aim at.
> 
> - Bruce Lee

While it is a goal to reach whatwg spec compliance, the primary goal is to have
a useful tool for testing modern web applications. 

Some specs don't really have any usage in modern web applications. For example,
you generally wouldn't write an application that depends on quirks mode.

Another example is `document.write`. I've yet to work on any application that
depends on this. However, implementing support for this feature require a
complete rewrite of the HTML parser. You would need a really good case (or
sponsorship level) to have that prioritised.

### Accessibility tree

It is not currently planned that this library should maintain the accessibility
tree; nor provide higher level testing capabilities like what
[Testing Library](https://testing-library.com) provides for JavaScript.

These problems _should_ eventually be solved, but could easily be implemented in
a different library with dependency to the DOM alone.

### Visual Rendering

It is not a goal to be able to provide a visual rendering of the DOM. 

But just like the accessibility tree, this could be implemented in a new library
depending only on the interface from here.
