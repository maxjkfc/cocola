# cocola

## Log
> Use the  go.uber.org/zap package 

### Useage:
```golang
    import "github.com/maxjkfc/cocola/log"

    func main(){
        log.NewZapConfig().Env(log.DEV).Level(log.DEBUG).Build()
        log.Zlogs().Info("This is the zap log pkg")
    }
```

## Token
> Use the github.com/dgrijalva/jwt-go

### Useage:

#### Create Token
```golang
    import "github.com/maxjkfc/cocola/token"
    import "fmt"


    func main(){
        token.SetKey("123")
        token.SetIssuer("xxx")
        j :=token.New()      
        j.Create(map[string]interface{}{
                    "name":"maxjkfc",
                    "phone":192939102,
                })
        fmt.Println(j.Get())
    }
```

#### Parse Token

```golang
    import "github.com/maxjkfc/cocola/token"
    import "fmt"


    func main(){
        token.SetKey("123")
        token.SetIssuer("xxx")
        j :=token.New()

        x := make(map[string]interface{})

        j.Parse(token , x)
        fmt.Println(x)
    }
```


