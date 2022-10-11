module go.einride.tech/unit

go 1.17

require gotest.tools/v3 v3.4.0

require github.com/google/go-cmp v0.5.5

// Old misversioned releases
retract (
	v1.2.0
	v1.0.0
	v0.1.0
)
