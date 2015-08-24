## Upgrading from httpstress-go < 4.1

`httpstress-go` is `httpstress` now. If you're building from source, you should:

* Delete `$GOPATH/src/chillum/httpstress`
* Rename `$GOPATH/src/chillum/httpstress-go` to `$GOPATH/src/chillum/httpstress`

## Upgrading from httpstress < 1.3

Change the import path from `"github.com/chillum/httpstress"` to `"github.com/chillum/httpstress/lib"`
