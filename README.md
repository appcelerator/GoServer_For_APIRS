# GoServer for Axway API Runtime Services
This is a sample golang based rest api server that can be deployed to Axway API Runtime Services.

## Logging

In order to get application access log with *appc cloud accesslog <appname>*, the log directories need to be setup at app start time in start_app script. Access log should be written to /ctlog/requests.log in json format and one request per line.

```
{
  "time": "2017-07-14T19:34:22.587409155Z",  // Time of request
  "response_time": 2,                        // Duration of request
  "req": {
    "method": "GET",                         // HTTP method
    "url": "/test",                          // HTTP endpoint
    "headers": {
      "accept": "*/*",
      "host": "goserver.cloud.appctest.com", // Hostname
      "user-agent": "curl/7.29.0",
      "x-forwarded-for": "54.214.34.215" .   // Client IP
    }
  },
  "res": {
    "statusCode": 200                        // Response code
  }
}
``` 


