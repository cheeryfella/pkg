# specification

A generic Go implementation of the [Specification pattern](https://en.wikipedia.org/wiki/Specification_pattern), enabling composable, reusable business rules expressed as first-class objects.

Specifications encapsulate a single predicate (`IsSatisfiedBy`) and can be combined using boolean logic operators — `And`, `Or`, `Not`, `NAnd`, and `NOr` — to build complex rules without scattering conditional logic across your codebase.

## Installation

```sh
go get github.com/cheeryfella/pkg/specification
```

## Usage

### 1. Define a specification

Embed `specification.Spec[T]` in your struct, implement `IsSatisfiedBy`, and call `Associate` in your constructor to wire up the composable methods.

```go
package main

import "github.com/cheeryfella/pkg/specification"

type Order struct {
    Total    float64
    Customer string
}

// IsHighValueOrder checks whether an order exceeds a value threshold.
type IsHighValueOrder struct {
    specification.Spec[Order]
    threshold float64
}

func (s *IsHighValueOrder) IsSatisfiedBy(o Order) bool {
    return o.Total > s.threshold
}

func NewIsHighValueOrder(threshold float64) *IsHighValueOrder {
    s := &IsHighValueOrder{threshold: threshold}
    s.Associate(s)
    return s
}

// IsPreferredCustomer checks whether an order belongs to a preferred customer.
type IsPreferredCustomer struct {
    specification.Spec[Order]
    preferred map[string]bool
}

func (s *IsPreferredCustomer) IsSatisfiedBy(o Order) bool {
    return s.preferred[o.Customer]
}

func NewIsPreferredCustomer(names ...string) *IsPreferredCustomer {
    m := make(map[string]bool, len(names))
    for _, n := range names {
        m[n] = true
    }
    s := &IsPreferredCustomer{preferred: m}
    s.Associate(s)
    return s
}
```

### 2. Compose specifications

Once defined, specifications can be combined freely using the chainable operators.

```go
highValue       := NewIsHighValueOrder(500)
preferredCustomer := NewIsPreferredCustomer("alice", "bob")

// Order must be high-value AND from a preferred customer.
eligibleForDiscount := highValue.And(preferredCustomer)

// Order must be high-value OR from a preferred customer.
eligibleForReview := highValue.Or(preferredCustomer)

// Order must NOT be high-value.
standardOrder := highValue.Not()

// NAND / NOR are also available.
neitherCondition := highValue.NOr(preferredCustomer)
notBothConditions := highValue.NAnd(preferredCustomer)
```

### 3. Evaluate

```go
order := Order{Total: 750, Customer: "alice"}

if eligibleForDiscount.IsSatisfiedBy(order) {
    fmt.Println("apply discount")
}
```

## Operators

| Method | Behaviour |
|--------|-----------|
| `And(b)` | `a && b` |
| `Or(b)` | `a \|\| b` |
| `Not()` | `!a` |
| `NAnd(b)` | `!(a && b)` |
| `NOr(b)` | `!(a \|\| b)` |

## Type parameter

All types are generic over `T any`. The type is inferred from your `IsSatisfiedBy` implementation, so a single package supports specifications over any domain type without casting.