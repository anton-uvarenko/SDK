# README

To use this sdk write these command to your terminal

>go get -u github.com/anton-uvarenko/SDK

##  Usage

To start with sdk you should create sdk, that returns an interface Lotr:
```go
func main() {
	sdk := lotr_sdk.NewSdk()
}
```

Simple request for getting books with parametrs:
```go
func main() {
	sdk := lotr_sdk.NewSdk()
	
	books := sdk.GetAllBooks("limit=100", "offset=50")
}
```
