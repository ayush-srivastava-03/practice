## Background

View the API definition in API.proto

## TASK
* Create the ViewNetworkService and implement the viewNetworkMembers endpoint.
Feel free to create this service struct with any members you feel are necessary, and with constructors if you need it.
This endpoint presents a list of network members (retrieved from NetworkService/getNetworkMembers), enriched with data 
(from ContactService,InterestsService,UserService).

* Instantiate and register the service in the main function, so that it can be called.

* Write a test for this endpoint, that can be run with `go test`
