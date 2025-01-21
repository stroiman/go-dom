# Script engine

There are two script hosts in the code base

- v8 - Default engine
- goja - WORK IN PROGRESS

V8 is the JavaScript engine in Chrome. It is a C++ API, and requires a lot of
CGo code to work, making this library troublesome for some. However, this
library is only intended for _testing_, not production. So your production
system does not inherit a CGo dependency.

[Goja](https://github.com/dop251/goja) is a pure Go JavaScript engine. I didn't
know about this when I started. I wish I had. I'm slowly implementing support
for existing features to Goja. Had I known about it, I would have used it
instead of v8.

### Default engine

The _default_ host is currently v8. When the goja host is ready, and all
features are supported, it will become the new default.

V8 will continue to be supported. As long as V8 powers Chrome, you woulndn't use
JavaScript features that are not supported in the engine.

## About v8go

V8 is based on the v8go project, Originally created by [Github user rogchap](https://github.com/rogchap/v8go). It wasn't kept up to date, and now [Tommie's branch](https://github.com/tommie/v8go) is the best maintained. 

However, many v8 necessary v8 features were not implemented in the branch. [My own fork](https://github.com/stroiman/v8go/tree/go-dom-feature-dev) contains 

Thank's to Tommie, v8go is automatically being updated with the latest v8 engine
from the chromium sources.
