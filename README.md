# SSE supported reverse proxy

sseproxy is a [SSE (Server-Sent Events)][sse] supported HTTP revese proxy.

`sseproxy.Proxy` keeps HTTP connections too long time, so it should be better
to accept minimum connections.

[sse]:https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events

## Usage

It can be used just like `httputil.NewSingleHostReverseProxy`.

```go
import (
	"net/http"
	"net/url"
	"github.com/tama-go/sseproxy"
)

u, _ := url.Parse("http://example.com")
p := sseproxy.New(u)
http.ListenAndServe(":8080", p)
```
