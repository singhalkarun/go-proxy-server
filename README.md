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
            "location":"/api/",
            "proxy_url": "http://localhost:3001/"
        },
        {
            "location":"/",
            "proxy_url": "http://localhost:3000/"
        }
    ]
}
```

Note:

* While proxying matching of prefixes will be in the same order as specified in routes
* While matching the requests as per routes, for matched block, any string that comes after the corresponding **location** will be concatenated with the **proxy_url** and requested

  Examples for reference:

  * localhost/[hello] -> http://localhost:3000/[hello]
  * localhost/api/[hello] -> http://localhost:3001/[hello]
  * localhost/[api123] -> http://localhost:3000/[api123]
  * localhost/[api] -> http://localhost:3000/[api]

  


