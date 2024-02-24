package routes

import (
	"net/http"
	"slices"
)

// This whole function is probably a war crime on nature

func SFWRoutes(param string) *http.Response {
	routesParams := [8]string{"hug", "kiss", "slap", "punch", "wink", "pat", "kill", "cuddle"}
	if slices.Contains(routesParams, param) == true {
		// This honestly probably shouldn't be a panic
		panic("not a valid param")
	}

	resp, err := http.Get("" + param)

	// Probably shouldn't panic here and add some additional logging
	if (err != nil) {
		panic("There was an error")
	}

	return resp
}