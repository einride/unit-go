# Unit Go

Go SDK for modeling physical units and conversion between them.

## Installation

```
go get go.einride.tech/unit
```

## Example

```go
angle := 3 * unit.Radian
fmt.Println(angle.Degrees())
// Output: 171.88733853924697
```
