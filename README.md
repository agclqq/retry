# retry

This is a retry package with backoff.

#useage
**direct use**
```go
    r := Retry{
        InitialBackoff:    1,
        MaxBackoff:        15,
        BackoffMultiplier: 1.5,
        MaxAttempts:       15,
    }
	
    r.Run(func(step uint) {
        ...
    }
```
**reset**
```go
    r := Retry{
        InitialBackoff:    1,
        MaxBackoff:        15,
        BackoffMultiplier: 1.5,
        MaxAttempts:       5,
    }
	
    r.Run(func(step uint) {
        if yourCondition {
            r.Reset()
        }
        ...
    })
```
**cancel**
```go
    r := Retry{
        InitialBackoff:    1,
        MaxBackoff:        15,
        BackoffMultiplier: 1.5,
        MaxAttempts:       5,
    }
	
    r.Run(func(step uint) {
        if yourCondition {
            r.Cancel()
        }
        ...
    })
```