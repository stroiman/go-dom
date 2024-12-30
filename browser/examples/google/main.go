package main

import . "github.com/stroiman/go-dom/browser"

func assertNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func Main() {
	browser := NewBrowser()
	window, err := browser.Open("https://google.com")
	assertNoError(err)
	textarea, err := window.Document().QuerySelector("textarea")
	assertNoError(err)
	textarea.SetAttribute("value", "Go")
	button, err := window.Document().QuerySelector("input[type='button']")
	assertNoError(err)
	button.Click()
}
