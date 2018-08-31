# jibaku-go
A bot destroyer for golang

ボットが暴走しないように1秒間にn回以上発話しようとすると自動でプロセスをkillする自爆プログラム

## usage

```
go get "github.com/garicchi/jibaku-go/jibaku"
```

1秒間に10回発話したら自爆

```
import "github.com/garicchi/jibaku-go/jibaku"

// before post from your bot
jibaku.Check(time.Second,10)
```
