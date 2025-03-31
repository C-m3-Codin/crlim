# Rate Limiting Library for Golang

This Golang library provides client-side rate limiting for external requests. It supports rate limits based on domains, URLs, query parameters, and tokens, making it highly configurable. The library also allows defining rate limits via JSON/YAML configuration files.

## Features
- Configurable rate limits for domains, URLs
- Supports JSON/YAML-based configuration.
- Wrapper around the `http.Client` to enforce rate limits.
- Efficient request filtering and enforcement.
- Extensible design for future improvements.

## Installation
```sh
go get github.com/c-m3-codin/crlim
```

## Usage

### 1. Load Configuration
```go
config, err := ratelimiter.LoadConfig("config.json") // or "config.yaml"
if err != nil {
    log.Fatal("Error loading config:", err)
}
client := ratelimiter.NewRateLimitedClient(config.RateLimits)
```

### 2. Example Configuration Files
#### JSON (`config.json`)
```json
{
  "rate_limits": {
    "api.example.com/v1/users": { "requests_per_second": 10, "burst_size": 5 },
    "api.example.com/v1/orders": { "requests_per_second": 5, "burst_size": 2 }
  }
}
```

#### YAML (`config.yaml`)
```yaml
rate_limits:
  api.example.com/v1/users:
    requests_per_second: 10
    burst_size: 5
  api.example.com/v1/orders:
    requests_per_second: 5
    burst_size: 2
```

## Roadmap
- [ ] Support wildcard-based rate limiting.
- [ ] Implement automatic config reloading.
- [ ] Add real-time request monitoring dashboard.
- [ ] Configurable ratelimits query parameters, and tokens.

## License
MIT
