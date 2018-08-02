# jibaku-go
A bot destroyer for golang

ボットが暴走しないように1秒間に3回以上発話しようとすると自動でプロセスをkillする自爆プログラム

## usage

```
go get "github.com/garicchi/jibaku-go/jibaku"
```

```
import "github.com/garicchi/jibaku-go/jibaku"

// before post from your bot
jibaku.Check()
```