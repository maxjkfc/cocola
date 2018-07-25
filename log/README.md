# Log

It's a log module base on zap log, easy for use and config it

## Config

```golang
    import "github.com/maxjkfc/cocola/log"

    func main() {
        // NewZapConfig - new the zap config
        // Env - set the environment for log
        // Level - set the log Level
        // Build - build the log instrace
	    if _ , err := NewZapConfig().Env(DEV).Level(DEBUG).Build(); err != nil {
	    	log.Error(err)
	    }
        log.Zlog().Info("This is a log")
    }
```


## Feature


## TODO

- [] Notice 
- [] Interface 
