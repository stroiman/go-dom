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

[Software license](./LICENSE.txt)

While the SPA[^1] dominates the web today, some applications still render
server-side HTML, and HTMX is gaining in popularity. Go has some popularity as a
back-end language for HTMX.

In Go, writing tests for the HTTP handler is easy if all you need to do is
verify the response.

But if you need to test at a higher lever, for example verify how any JavaScript
code effects the page; you would need to use browser automation, like
[Selenium](https://www.selenium.dev/), and this introduces a significant overhead.

You need to control launching your server, tests need to wait for the browser to
launch, and communication with the browser has the overhead of out-of-process
communication.

This overhead is so massive that it discourages a TDD loop.

The purpose of this project is to support a TDD feedback loop for code
delivering HTML, and where merely veryifying the HTTP response isn't enough, but
you want to verify:

- JavaScript code has the desired behaviour
- General browser behaviour is verified, e.g. clicking a `<button type="submit">`,
  submits a form, and a redirect response is followed.

Some advantages of a native headless browser are:

- No need to wait for a browser to launch.
- Everything works in-process, so interacting with the browser from test does
  not incur the overhead of out-of-process communication, and you could for
  example redirect all console output to go code easily.
- You wouln't actually need to start an HTTP server. You could just pass an
  `HTTPHandler` instance to the "browser", and instead of making HTTP requests,
  this calls your handler.
- You can run parallel tests in isolation as each can create their own _instance_
  of the HTTP handler.[^2]

Some disadvantages to e.g. Selenium.

- You cannot verify how things look; e.g. you cannot get a screenshot of a failing test
  - This means you cannot create snap-shot tests detect undesired UI changes.[^3]
- You cannot verify that everything works in _all supported browsers_.

## Project status

Coding has just begun. Planned first steps

1. Create a simple HTML parser, that can parse just the basic HTML tags.
2. Integrate with an [`http.Handler`](https://pkg.go.dev/net/http#Handler)
3. Embed a v8 engine, to allow running javascript code.
4. Expose the DOM objects in a way compatible with v8.

There is much to do, which includes (but this is not a full list):

- Support all DOM elements
- Handle bad HTML gracefully (browsers don't generate an error if an end tag is
  missing or misplaced)
- Implement all standard JavaScript classes that a browser should support; but
  not provided by the V8 engine.
  - JavaScript polyfills would be a good starting point, where they exist.
  - Over time, the most commonly used, in particular `fetch` could be replaced
    with a Go implementation for performance optimization.
- Implement default browser behaviour for user interaction, e.g. pressing 
  <key>enter</key> when an input field has focus should submit the form.

## Out of scope.

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
