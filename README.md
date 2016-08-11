HTML Error
==========

Simple library to generate HTML error pages during development.


Usage
-----

```go
func Index(w http.ResponseWriter, r *http.Request) {
	// ... code ...
	if err != nil {
		htmlerror.Error(w, r, err)
		return
	}
}
```


Screenshot
----------

![screenshot](http://i.imgur.com/pwEZdZW.png)
