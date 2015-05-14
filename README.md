Ravelin Code Test
=================

## Summary
We need an http server that will accept any post request (json) from muliple clients' websites. Each request forms part of a struct (for that particular visitor) that will be printed to the terminal when the struct is fully complete. 

For the js part of the test please feel free to use any libraries that may help you **but do not use any non standard library packages for the Go service**.

## Frontend (JS)
Include javascript into the index.html (supplied) that captures and posts data every time one of the below events happens; this means you will be posting multiple times per visitor. Assume only 1 resize occurs.

  - if the screen resizes, the before and after dimensions
  - copy & paste (for each field)
  - time taken from the 1st character typed to clicking the submit button

### Example JSON Requests
```
{
  "eventType": "copyAndPaste",
  "websiteUrl": "https://ravelin.com",
  "sessionId": "123123-123123-123123123",
  "pasted": true,
  "formId": "inputCardNumber"
}

{
  "eventType": "timeTaken",
  "websiteUrl": "https://ravelin.com",
  "sessionId": "123123-123123-123123123",
  "time": 72, // seconds
}

...

```

## Backend (Go)
1. Build a go binary with an http server
2. Accept post requests (json format)
3. Map the json requests to relevant sections of the data struct
4. Print the struct at each stage of it's construction. 
5. Also print the struct when it is complete (i.e. form submit button has been clicked)
6. Use go routines and channel where appropriate

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




