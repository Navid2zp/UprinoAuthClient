# Uprino Authentication Client
Golang client library for Uprino authentication API.

### Install
```
go get github.com/Navid2zp/UprinoAuthClient
```

### Usage

```
import (
  "fmt"
  "github.com/Navid2zp/UprinoAuthClient"
)

func main() {
    token := UprinoAuthClient.TokenCheck{
      Token: "TOKEN-TO-CHECK",
      PublicKey: "YOUR-PUBLIC-API-KEY",
      PrivateKey: "YOUR-PRIVATE-API-KEY",
    }
    tokenData, err := token.Check()
    if err != nil {
      // Connection error or ...
    }
    
    // NOTE: Use token.Checked if you're not sure that token.Check() is called before.
    if token.Valid {
      fmt.Println("Hey,", tokenData.UserData.Username)
    } else {
      fmt.Println("I don't know you :(")
    }
}
```

|      Arg      |          Description           |           Required         |
|---------------|--------------------------------|----------------------------|
|Token          |User token to check             |             YES            |
|PublicKey      |Only for your public calls      |             YES*           |
|PrivateKey     |Only for private calls (backend)|             YES*           |

> *You should provide at least one of the keys.

### Notes:
- Login/Register/Password reset is only available on the authentication service and there is no API for that.
- To login a user for your API key, send the user to `https://auth.uprino.com/login?api_key=YOUR-PUBLIC-API-KEY` and user will be redirected to your callback URL with a token.
