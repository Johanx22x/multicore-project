package server

import (
	"log"
	"net/http"
    "fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type myCompo struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (c *myCompo) Render() app.UI {
	data := map[string]int{
		"Go":         10,
		"JavaScript": 4,
		"Python":     6,
		"C":          8,
	}

	return app.Ul().Body(
		app.Range(data).Map(func(k string) app.UI {
			s := fmt.Sprintf("%s: %v/10", k, data[k])

			return app.Li().Text(s)
		}),
	)
}

func Start() {
	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &myCompo{})

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
