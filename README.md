# retry

This is a retry package with backoff.

#useage

**direct use**:

Run until the maximum number of times to stop.When MaxAttempts is 0, it does not stop.
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

Reset the current count to 0, and you can count again.
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

If you don't want to continue, you can cancel.
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