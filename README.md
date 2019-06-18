# search-maximum

Implementing a maximum search on Go.

## Installation

Run the following command from you terminal:


 ```bash
 go get github.com/koind/search-maximum
 ```

## Usage

Package usage example.

```go
package main

import (
	"github.com/koind/search-maximum"
)

func main() {
	nums := []interface{}{4, 2, 9}
	less := func(i, j int) bool {
		return nums[i].(int) < nums[j].(int)
	}
	
	max, err := search_maximum.FindMax(nums, less)
	if err != nil {
		println(err)
	}
	println(max) // 9
}
```

## Available Methods

The following methods are available:

##### koind/search-maximum

```go
FindMax(sl []interface{}, less func(i, j int) bool) (interface{}, error) 
```

## Tests

Run the following command from you terminal:

```
go test -v .
```