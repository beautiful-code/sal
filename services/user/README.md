### User Service

#### Steps to Run this service
1. Clone this repository to your $GOPATH/src/github.com
2. Run `cd beautiful-code/sal/services/user`
3. Run `glide install`
4. Run `go run main.go`

#### API Endpoints:
1. Register - "/register"

Example Request:
```
curl -X POST -H "Content-Type: application/json" -H "Cache-Control: no-cache" -d '{
	"data": {
		"firstname": "John",
		"lastname": "Doe",
		"email": "john.doe@example.com",
		"password": "secret"
	}
}' "http://localhost:8080/register"
```

Example Response:
```
HTTP Status: 201 Created
Body:
{
  "msg": "Created user record."
}
```

2. Login - "Login"

Example Request:
```
curl -X POST -H "Content-Type: application/json" -H "Cache-Control: no-cache" -d '{
	"data": {
		"email": "john.doe@example.com",
		"password": "secret"
	}
}' "http://localhost:8080/login"
```

Example Response:
```
HTTP Status: 200 OK
Body:
{
  "data": {
    "user": {
      "ID": 9,
      "CreatedAt": "2017-01-10T19:00:57+05:30",
      "UpdatedAt": "2017-01-10T19:00:57+05:30",
      "DeletedAt": null,
      "FirstName": "John",
      "LastName": "Doe",
      "Email": "john.doe@example.com",
      "Password": ""
    },
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXJyZW50X3VzZXJfZW1haWwiOiJqb2huLmRvZUBleGFtcGxlLmNvbSIsImV4cCI6MTQ4NDA1ODY2MSwiaXNzIjoiYWRtaW4ifQ.fHTASFARMM4j5xRdvDQrocwDc-vJqJGeeNnOmsEWDCT7ZQKsf6ai0owRR3Cb7tvp876An7D2IlOwW7tp16p4TwQtn8GqzpvzMFIKCrgmcYoW8FKACKC_WmFWsbn8IPczc6PrDZ7cEyDPv4djf00c54GVov-GUtNbbcJubEgfYBc"
  }
}
```
