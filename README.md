# Description

For this project, I've used 3 libraries:
 - `github.com/golang/mock`, for generating interface mocks
 - `github.com/gorilla/mux`, for creating APIs
 - `github.com/stretchr/testify`, for testing assertions

To generate the mocks, I've added an annotation in the top of `interfaces.go` file. To trigger it, simply run:
`go generate ./...` from the root folder.

## How to run the tests

To check the test coverage, run the following command from the root folder:
```
go test $(go list ./... | grep -v /vendor/) -coverprofile cover.out && go tool cover -html=cover.out
```

This will create a `cover.out` file and open a tab in the browser, showing detailed coverage status for each package.
This information is also available in the command line.

## How to use it

To use it, type `go run .` in the root of the folder. The canvas has `height` and `width` set to *50* and it will launch the server in the port `:3000` by default. However, it can be customizable.
To do it, type `go build .` and use the following arguments to override the default values:
 - --height (to set the height)
 - --width (to set the width)
 - --port (to set the port)

For example:
```
go build .
./exercise --height 30 --width 20 --port 4000
```

### Drawing in the canvas
For testing the APIs, I've used [Postman](https://www.postman.com/). The example given in the open-api documentation should work straight away. Just copy and paste it to the body of the POST request.
For example, once the server is running, it should accept requests for drawing as the one below:
```
POST http://localhost:3000/canvas
Body: {
  "coordinates": {
    "x": 5,
    "y": 3
  },
  "height": 10,
  "width": 4,
  "outline": "X",
  "fill": "@"
}
```

The request above is for drawing in the canvas and valid ones should return 201 (Created). If any parameter is incorrect, a 400 (Bad Request) will be returned with some description of the possible issue. The requirements for a valid request are given below:
 - `coordinates` must contain `x` and `y` fields included and its values ranging from:
    - \>= 0 and < `maxWidth` for `x`;
    - \>= 0 and < `maxHeight` for `y`.
 - `height` and `width` values must be greater than 1
 - `outline` and `fill` cannot be empty simultaneously
 - `x` + `width` < `maxWidth` and `y` + `height` < `maxHeight`, i.e., the rectangle to be drawn must be inside the canvas bounds.

### Getting the canvas
To fetch the canvas content, it should accept requests as the one given below:
```
GET http://localhost:3000/canvas
```

The returning content should something similar as following:
```
{
    "canvas": [
        "                              ",
        "                              ",
        "                              ",
        "     XXXX                     ",
        "     X@@X                     ",
        "     X@@X                     ",
        "     X@@X                     ",
        "     X@@X                     ",
        "     X@@X                     ",
        "     X@@X                     ",
        "     X@@X                     ",
        "     X@@X                     ",
        "     XXXX                     ",
        "                              ",
        "                              "
    ]
}
```
