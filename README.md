# console spinner
[![License][license-img]][license] [![Actions Status][action-img]][action] [![GoDoc][godoc-img]][godoc] [![Go Report Card][goreport-img]][goreport] [![Coverage Status][codecov-img]][codecov]

spinner provides basic and simple console spinner functionality

TODO: add more info

Example:
```
	...
import (
	"go.melnyk.org/spinner"
)
	...
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	spin := spinner.NewSpinner(ctx)
	
	spin.Process("Long process message...")
	// long process	
	...
```

## Development Status: In active development
All APIs are in active development and not finalized, and breaking changes will be made in the 0.x series.


[license-img]: https://img.shields.io/badge/license-MIT-blue.svg
[license]: https://github.com/mmelnyk/spinner/blob/master/LICENSE
[action-img]: https://github.com/mmelnyk/spinner/workflows/Test/badge.svg
[action]: https://github.com/mmelnyk/spinner/actions
[godoc-img]: https://godoc.org/github.com/mmelnyk/spinner?status.svg
[godoc]: https://godoc.org/github.com/mmelnyk/spinner
[goreport-img]: https://goreportcard.com/badge/github.com/mmelnyk/spinner
[goreport]: https://goreportcard.com/report/github.com/mmelnyk/spinner
[codecov-img]: https://codecov.io/gh/mmelnyk/spinner/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/mmelnyk/spinner
