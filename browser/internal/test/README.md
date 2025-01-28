# HTMX Test app

This folder contains a simple HTMX test app, that is being used in the test
suite for gost-dom.

```sh
htmx-app/ # The actual HTML server implementation as an http.Handler
htmx-app-main/ # A runnable executable that exposes the app on a TCP port 
```

The app exposes a simple counter, that is incremented when clicked on.
