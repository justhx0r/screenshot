# Screenshot
Simple cross-platform pure Go screen shot library. (tested on linux&windows&osx)

<br/>

## Installation
```go
go get github.com/justhx0r/screenshot
```

<br/>

## Basic Usage
Import the package
```go
import (
    "github.com/justhx0r/screenshot"
)
```

```go
func main() {
    img, err := screenshot.CaptureScreen()
    myImg := image.Image(img)
}
```


<br/>

## Dependencies
* **Windows** - None
* **Linux/FreeBSD** - https://github.com/BurntSushi/xgb
* **OSX** - cgo (CoreGraphics,CoreFoundation, that should not be a problem)

<br/>

## Examples
Look at `example/` folder.
