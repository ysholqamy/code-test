Ravelin Code Test
=================

## Summary
We need an http server that will accept any post request (json) from muliple clients' websites. Each request forms part of a struct (for that particular visitor) that will be printed to the terminal when the struct is fully complete. 

For the js part of the test please feel free to use any libraries that may help you **but do not use any non standard library packages for the Go service**.

## Frontend (JS)
Include javascript into the index.html (supplied) that captures and posts every time the following events happens:

  - if the screen resizes, the before and after dimensions
  - copy & paste (for each field)
  - time taken for 1st character to clicking the submit button

### Example JSON Request
```
{
  "eventType": "copyAndPaste",
  "websiteUrl": "https://ravelin.com",
  "sessionId": "123123-123123-123123123",
  "formId": true,
}
```

## Backend (Go)
1. Build a go binary with an http server
2. Accept post requests (json format)
3. Map the json requests to relevant sections of the data struct
4. Only print the struct when it is complete (i.e. form submit button has been clicked)

### Go Struct
```
type Data struct {
	websiteUrl         string
	sessionId          string
	resizeFrom         Dimension
	resizeTo           Dimension
	copyAndPaste       map[string]bool // map[fieldId]true
	formCompletionTime int // Seconds
}

type Dimension struct {
	Width  string
	Height string
}
```




