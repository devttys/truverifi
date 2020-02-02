# truverifi
Golang Client for Truverifi API


```golang

import (
  "github.com/devttys/truverifi"
  "log"
)

func main() {
  client := truverifi.NewClient("your-api-key")
  account, err := client.GetAccount()
  
  if err != nil {
    log.Fatal(err)
  }
  
  log.Println(account.Balance)
}
```


# Updating list of services

- Go to https://app.truverifi.com/
- Run in console: `document.write(JSON.stringify(Object.values($('.usecases-option').map((x,y) => { return y.attributes.value && y.attributes.value.nodeValue; })).filter(x => { return typeof(x) === 'string'; }).sort().map(x => { return x.toLowerCase(); })))`
- Update `Services` variable in `models.go`
