# cocola


## Error

### Useage:
```golang
    import "github.com/maxjkfc/cocola/errors"


    func main()  {
        err  := errors.New(1 , "Error: Not Find it")
        fmt.Println(err.Error())
    }
```

## Log
> Use the  go.uber.org/zap package 

### Useage:
```golang
    import "github.com/maxjkfc/cocola/log"

    func main(){
        log.NewZapConfig().Env(log.DEV).Level(log.DEBUG).Build()
        log.Log().Info("This is the zap log pkg" , nil)
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


