# go-dom - Headless browser for Go

> [!NOTE] 
>
> This is only a POC
>
> This readme file was written before any code was added, and describes the
> intent of the project; why it was conceived, and the benefits that it brings.
>
> Also this starts as a hobby project for learning; there is no guarantee that
> it will ever become a useful tool; unless it will gain traction and support of
> multiple developers.

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
delivering HTML, and where merely veryifying the HTTP response isn't enough, but
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
isolated test, e.g. mocking out part of the behaviour; but 

## Project status

Coding has just begun, and I have the following extremely basic features:

- Simple streaming tokenizer and parser, can only parse the text `<html></html>`
- Ability to receive directly to an `http.Handler`
- Experimental embed of v8 engine.[^4]
  - This is not in the main branch, as it significantly increase build times.
    [embed-v8-engine branch](https://github.com/stroiman/go-dom/tree/embed-v8-engine)
  - This was done early to identify the risk; or impact this could have on how
    the DOM objects should be implemented.

Planned first steps:

1. Create a simple HTML parser, that can parse just the basic HTML tags, and
   generate a valid subset of a [`Document`](https://developer.mozilla.org/en-US/docs/Web/API/Document)
2. Integrate with an [`http.Handler`](https://pkg.go.dev/net/http#Handler)
3. Embed a v8 engine, to allow running javascript code. (partially done)
4. Expose the DOM objects to v8.

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
  - Avoid the locking the account due to _suspiscious activity_.
  - The IDP may use a Captcha or 2FA that can be impossible; or difficult to
    control from tests, and would cause a significant slowdown to the test
    suite.
- For applications like map providers
  - Avoid being billed for API use during testing.

## Help

This project will likely die without help. If you are interested in this, I
would welcome contributions. Particularly if:

- You have experience building tokenisers and parsers, especially HTML.
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
- You have experience working with the v8 engine, particularly exposing internal
  objects to JavaScript (which is then External to JavaScript).
  - In particular, if you've done this from Go.


## Out of scope.

### Visual Rendering

It is not a goal to be able to provide a visual rendering of the DOM.

### Accessibility tree

It is not currently planned that this library should maintain the accessibility
tree; nor provide higher level testing capabilities like what
[Testing Library](https://testing-library.com) provides for JavaScript.

These problems _should_ eventually be solved, but could easily be implemented in
a different library with dependency to the DOM alone.

[^1]: Single-Page app
[^2]: This approach allows you to mock databases, and other external services;
A few integration tests that use a real database, message bus, or other external
services, is a good idea. Here, isolation of parallel tests may be
non-trivial; depending on the type of application.
[^3]: I generally dislike snapshot tests; as they don't _describe_ expected
behaviour, only that the outcome mustn't change. There are a few cases where
where snapshot tests are the right choice, but they should be avoided for a TDD
process.
[^3]: The engine is based on the v8go project by originally by @rogchap, later
kept up-to-date by @tommie. The project is missing some features in dealing with
external objects; which I work on adding in my own fork.
