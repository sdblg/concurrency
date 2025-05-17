package main

import (
	largearrayaccess "app/examples/large-array-access"
	primenumbergenerate "app/examples/primenumber-generate"
	webfetch "app/examples/web-fetch"
)

func main() {
	webfetch.CallConcurrently()
	webfetch.Call()

	primenumbergenerate.Do()
	primenumbergenerate.DoPuneet()

	largearrayaccess.Do()

}