This repository contains generated golang code for:
* Event Family: https://desp.kroger.com/event-family/c6e47021-4c3c-3f57-a9f8-5245f4159412
* Version: v0.1.3

To use this in your go client, add the following requirements to your go.mod file.

```
module github.com/krogertechnology/my-go-client-application

go 1.13

require (
	github.com/actgardner/gogen-avro/v9 v9.0.0 // indirect
	github.com/rhammonds1-kr/desp-catalog-go-order-orderrepository2 v0.1.3
	...
```