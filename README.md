# go-dom - Headless browser for Go

> [!NOTE] 
>
> This is still in development, and has not reached a level of usability.
>
> Given the progress I've had in the short time I spent working on this, I
> believe this _will_ reach a level of usability. But it depends on how much
> time I will have for this
>
> You can help by sponsoring the project. If I could reach a level where I could
> work full-time on this, the Go community could have the most awesome testing
> tool at their disposal within a few months.

[Software license](./LICENSE.txt) (note: I currently distribute the source under
the MIT license, but as this will most likely depend on v8go, I need to verify
that the license is compatible).

While the SPA[^1] dominates the web today, some applications still render
server-side HTML, and HTMX is gaining in popularity. Go has some popularity as a
back-end language for HTMX.

In Go, writing tests for the HTTP handler is easy if all you need to do is
verify the response.

But if you need to test at a higher lever, for example verify how any JavaScript
code effects the page; you would need to use browser automation, like
[Selenium](https://www.selenium.dev/), and this introduces a significant 
overhead; not only from out-of-process communication with the browser, but also
the necessity of launching your server.

This overhead discourages a TDD loop.

The purpose of this project is to support a TDD feedback loop for code
delivering HTML, and where merely verifying the HTTP response isn't enough, but
you want to verify:

- JavaScript code has the desired behaviour
- General browser behaviour is verified, e.g. 
  - clicking a `<button type="submit">` submits the form
  - A and a redirect response is followed.

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

Some disadvantages compared to e.g. Selenium.

- You cannot verify how it look; e.g. you cannot get a screenshot of a failing test
  - This means you cannot create snap-shot tests detect undesired UI changes.[^3]
- You cannot verify that everything works in _all supported browsers_.

This isn't intended as a replacement for the cases where an end-2-end test is
the right choice. It is intended as a tool to help when you want a smaller
isolated test, e.g. mocking out part of the behaviour;

## Project status

The current state of the project is, you can start a "browser" connected to a Go
`http.Handler` (the intention is to use the root handler, that you normally
expose on a TCP port; not a route, but feel free to do whatever you want).

The browser can open an HTML file, execute included scripts (remote scripts are
downloaded). The DOM is exposed as native Go objects, allowing the test code to
inspect or modify them, and JavaScript can also be executed from Go code.

Only a very minimal subset of the DOM specification is implemented.


- HTML parsing is done in 2 steps
  - Step 1 parsing of HTML using [x/net/html](https://pkg.go.dev/golang.org/x/net/html)
      - Using `x/net/html` gives HTML rendering (i.e., support for `outerHTML`) out
        of the box.
      - Libraries exist implementing XPath on top of this.
  - 2nd pass into my own structure
    - `x/net/html` does not have the interface that a Browser wants, so I wrap
    this to provide the browser DOM both to JavaScript and Go.
    - The library doesn't support the insertion steps, e.g., when a `<script>` is
      connected to the DOM, it should be executed (simplified).
- Embedding of v8 engine.[^4]
  - Expose the native Go objects to JavaScript

### Memory Leaks

The current implementation is leaking memory for the scope of a "Browser". I.e.,
all DOM nodes created and deleted for the lifetime of the browser will stay in
memory until the browser is ready for garbage collection.

The problem here is that this is a marriage between two garbage collected
systems, and what is conceptually _one object_ is split into two, a Go object
and a JavaScript wrapper. As long of them is reachable; so must the other be.

I could join them into one; but that would result in an undesired coupling; the
DOM implementation being coupled to the JavaScript execution engine.

Another solution to this problem involves the use of weak references. This
exists as an `internal` but [was
accepted](https://github.com/golang/go/issues/67552) as a feature.

Because of that, and because the browser is only intended to be kept alive for
the scope of a single short lived test, I have postponed dealing with memory
management.

### Demonstration

An early example showing HTML being loaded with a script. Notice:

- The HTML parser creates a `<head>`, even if missing in the source.
- Whitespace is not inserted in the DOM outside the body (line break and
  indentation is only processed in the body.
- The script is executed when connected to the DOM; which is why it doesn't see
  the `<div>` element after, as well as whitespace.

```go
It("Runs the script when connected to DOM", func() {
    window := ctx.Window()
    window.SetScriptRunner(ctx)
    window.LoadHTML(`
<html>
  <body>
    <script>window.sut = document.documentElement.outerHTML</script>
    <div>I should not be in the output</div>
  </body>
</html>
`,
    )
    Expect(
        ctx.RunTestScript("window.sut"),
    ).To(Equal(`<html><head></head><body>
    <script>window.sut = document.documentElement.outerHTML</script></body></html>`))
})
```

### Next up - HTMX

The next milestone being worked towards is to have the simplest HTMX application
running. One that has a button and a counter displayed; which increments
whenever the user clicks the button.

This drives the need for (not necessarily a complete list, just the parts I know
that HTMX _does_ depend on)

- `location` API - at least in a read-only version.
- `XMLHttpRequest` - work is underway.
- `XPathEvaluator` - I have that in a branch using a polyfill. BUT the necessary
  DOM methods lack implementation.
- `setTimeout`, `setInterval`, and friends.
  - This did surprise me a bit that it's not part of v8. But they're not part of
    ECMAScript specification (surprise to me). Fortunately Go has excellent
    concurrency and synchronization mechanisms.
  - I'm not sure if they are actually necessary for this case, but they
    represent an uncertainty I need to cover quickly.

Following the simple case, the plan is to have more complex cases, e.g. form
handling with redirection.

... and eventually websocket and server event support.

### Future goals

There is much to do, which includes (but this is not a full list):

- Support all DOM elements, including SVG elements and other namespaces.
- Handle bad HTML gracefully (browsers don't generate an error if an end tag is
  missing or misplaced)
- Implement all standard JavaScript classes that a browser should support; but
  not provided by the V8 engine.
  - JavaScript polyfills would be a good starting point, where some exist.
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

## Help

This project will likely die without help. If you are interested in this, I
would welcome contributions. Particularly if:

- ~~You have experience building tokenisers and parsers, especially HTML.~~
  - After first building my own parser, I moved to `x/net/html`, which seems
    like the right choice; at least for now.
- You have intimate knowledge of Go's garbage collection mechanics.
  - If you don't have the time or desire to help _code_ on this project, ~~I would
    appreciate peer reviews on those parts of the code.~~
    - I have postponed solving that problem until Go gets weak references.
    - However, if you do see another solution to the leaking problem, let me
      know.
- You have _intimate knowledge_ of how the DOM works in the browser, and can 
  help detect poor design decisions early. For example:
  - should the internal implementation use `document.CreateElement()`
    when parsing HTML into DOM? 
    - Would it be a big mistake to do so? 
    - Would it be a big mistake to _not_ do so? 
    - Is is it a _doesn't matter_, whatever makes the code clean, issue?
  - Which "objects" should I expose from Go to v8? and where should the
    functions live? The objects themselves, or should I create prototype
    in Go code? (I think I _should_ make prototype objects)
- You have knowledge of the whatwg IDL, and what kind of code could be
  auto-generated from the IDL
- You have experience working with the v8 engine, particularly exposing internal
  objects to JavaScript (which is then External to JavaScript).
  - In particular, if you've done this from Go.

## Out of scope.

## Full Spec Compliance

> A goal is not always meant to be reached, it often serves simply as something to aim at.
> 
> - Bruce Lee

While it is a goal to reach whatwg spec compliance, the primary goal is to have
a useful tool for testing modern web applications. 

Some specs don't really have an effect modern web applications. For example, you
generally wouldn't write an application that depends on quirks mode.

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


[^1]: Single-Page app
[^2]: This approach allows you to mock databases, and other external services;
A few integration tests that use a real database, message bus, or other external
services, is a good idea. Here, isolation of parallel tests may be
non-trivial; depending on the type of application.
[^3]: I generally dislike snapshot tests; as they don't _describe_ expected
behaviour, only that the outcome mustn't change. There are a few cases where
where snapshot tests are the right choice, but they should be avoided for a TDD
process.
[^4]: The engine is based on the v8go project by originally by @rogchap, later
kept up-to-date by @tommie; who did a remarkale job of automatically keeping the
v8 dependencies up-to-date. But many necessary features of V8 are not exported;
which I am adding in my own fork.
