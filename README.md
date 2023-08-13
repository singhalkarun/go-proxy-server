# Basic Reverse Proxy Server in Go

### Note

Run as root to be able to bind port 80

### Adding a configuration

Add a new file to conf directory

Format

```json
{   
    "hostname":"localhost",
    "routes": [
        {
            "location":"/api",
            "proxy_url": "http://localhost:3001"
        },
        {
            "location":"/",
            "proxy_url": "http://localhost:3000"
        }
    ]
}
```

Note:

* While proxying matching of prefixes will be in the same order as specified in routes
