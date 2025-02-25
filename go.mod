module go.einride.tech/unit

go 1.21

toolchain go1.24.0

require gotest.tools/v3 v3.5.2

require github.com/google/go-cmp v0.7.0

// Old misversioned releases
retract (
	v1.2.0
	v1.0.0
	v0.1.0
)
