# go-dom - Headless browser for Go

The Go-to headless browser for TDD workflows.

```go
browser := NewBrowserFromHandler(pkg.RootHttpHandler)
window, err := browser.Open("/example")
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
> This is still in development, and has not yet reached a level of usability.
>
> Expected alpha release for a version supporting forms in the spring 2025.

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

## Join the "community"

[Join my discord server](https://discord.gg/rPBRt8Rf) to chat with me, and stay
up-to-date on progress: 

- https://discord.gg/rPBRt8Rf

## Project background

While the SPA[^1] dominates the web today, some applications still render
server-side HTML, and HTMX is gaining in popularity. Go has some popularity as a
back-end language for HTMX.

While Go has great tooling for verifying request/responses of HTTP applications,
if you need to test at a higher level, for example verify how any JavaScript
code effects the page; you would need to use browser automation which introduce
a significant overhead; not only from out-of-process communication with the
browser, but also the necessity of launching your server.

This overhead discourages a TDD loop.

The purpose of this project is to support a TDD feedback loop for code
delivering HTML, and where merely verifying the HTTP response isn't enough, but
you want to verify:

- JavaScript code has the desired behaviour
- General browser behaviour is verified, e.g. 
  - clicking a `<button type="submit">` submits the form and a redirect response
    is followed.

Some advantages of a native headless browser are:

- No need to wait for a browser to launch.
- Everything works in-process, so interacting with the browser from test does
  not incur the overhead of out-of-process communication, and you could for
  example redirect all console output to go code easily.
- You can request application directly through the 
  [`http.Handler`](https://pkg.go.dev/net/http#Handler); so no need to start an
  HTTP server.
- You can run parallel tests in isolation as each can create their own _instance_
  of the HTTP handler.[^2]

Some disadvantages compared to browser automation.

- You cannot verify how it look; e.g. you cannot get a screenshot of a failing test
  - This means you cannot create snap-shot tests detect undesired UI changes.[^3]
- You cannot verify that everything works in _all supported browsers_.

This isn't intended as a replacement for the cases where an end-2-end test is
the right choice. It is intended as a tool to help when you want a smaller
isolated test, e.g. mocking out part of the behaviour;

## Code structure

This is still in early development, and the structure may still change.

The main library is in a subfolder, `browser`, to separate it from the code
generator that generates pieces of code from IDL specs. The browser and code
generator bases do not have any inter-dependencies, and they have different sets
of unrelated external dependencies.

```sh
browser/
  dom/ # Core DOM implementation
  html/ # Window, HTMLDocument, HTMLElement, 
  # ...
  scripting/ # Client-side script support
    v8host/ # v8 engine, and bindings
    gojahost/ # goja javascript engine,
  browser.go # Main module
code-gen/
  webref/ # Git submodule -> https://github.com/w3c/webref
```

The subfolders under `browser/` reflects the [web
APIs](https://developer.mozilla.org/en-US/docs/Web/API), and the naming
reflects the corresponding idl files. E.g., `browser/dom/` will have types
corresponding to the types specified in `code-gen/webref/ed/idl/dom.idl`.
`browser/html/` corresponds to `html.idl`, etc.[^4]

The `webref/` folder is not necessary for normal use, only when working with the
code-generator.

### Modularisation

Although the code isn't modularised yet, it is an idea that you should be able
to include the modules relevant to your app. E.g., if your app deals with
location services, you can add a module implementing location services.

This helps keep the size of the dependencies down for client projects; keeping
build times down for the TDD loop.

It also provides the option of alternate implementations. E.g., for location
services, the simple implementation can provide a single function to set the
current location / accuracy. The advanced implementation can replay a GPX track.

### Building the code generator.

To build the code generator, you need to build a _curated_ set of files first.
You need [node.js](https://nodejs.org) installed.

```sh
$ cd webref
$ npm install # Or your favourite node package manager
$ npm run curate
```

This build a set of files in the `curated/` subfolder.

## Project status

The browser is currently capable of loading an simple HTMX app; which can fetch
new contents and swap as a reaction to simple events, such as click.

The test file [htmx_test.go](./browser/scripting/htmx_test.go) verifies that
content is updated. The application being tested is [found
here](./browser/internal/test/README.md).

Client-side script is executed using the v8 engine.[^5]

Experimental work is done to also support [goja](https://github.com/dop251/goja)
for client-side script; but this version is not fully compatible yet.

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

The following two areas are the next focus of attention

- Navigation. Some actions, e.g. clicking a link, or submitting a normal
  (non-JS) form should result in a new HTTP reqest, and the response loaded in a
  new script context (global state reset).
- Form handling. Add code supporting typeing form values, and submitting the
  form, building a request body.
- Replace early hand-written JS wrappers with auto-generated code, helping drive
  a more complete implementation.

### Future goals

There is much to do, which includes (but this is not a full list):

- Support all DOM elements, including SVG elements and other namespaces.
- Support web-sockets and server events.
- Implement all standard JavaScript classes that a browser should support; but
  not provided by the V8 engine.
  - JavaScript polyfills would be a good starting point. This is used for xpath
    evaluator.
      - Conversion to native go implementations would be prioritized on usage, e.g.
        [`fetch`](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API) 
        would be high in the list of priorities.
  - A proper event loop with time travel. `setTimeout` and `setImmediate` are
    not implemented by v8. When testing code that has to wait, it is very useful
    to be able to fast forward simulated time.
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

---

[^1]: Single-Page app
[^2]: This approach allows you to mock databases, and other external services;
A few integration tests that use a real database, message bus, or other external
services, is a good idea. Here, isolation of parallel tests may be
non-trivial; depending on the type of application.
[^3]: I generally dislike snapshot tests; as they don't _describe_ expected
behaviour, only that the outcome mustn't change. There are a few cases where
where snapshot tests are the right choice, but they should be avoided for a TDD
process.
[^4]: This code structure may not be completely possible due to circular
dependencies between web APIs. E.g., `HTMLFormElement` and `FormData` have
circular dependencies.
[^5]: The engine is based on the v8go project by originally by @rogchap, later
kept up-to-date by @tommie; who did a remarkale job of automatically keeping the
v8 dependencies up-to-date. But many necessary features of V8 are not exported;
which I am adding in my own fork.
