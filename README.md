# routinator_client_check

This program uses the RoutinatorAPI to process the results.

## Overview

this program use RoutinatorAPI `/api/v1/status`.  
[Routinator API Endpoints link](https://routinator.docs.nlnetlabs.nl/en/stable/api-endpoints.html#:~:text=the%20following%20paths%3A-,/api/v1/status,-Returns%20exhaustive%20information)  

Obtain rtr.client information from the API output results.  
Identify the IP address owner of the client whose "connection" value is 1 or more, and output it in "description".


```json
{
    "xxx.xxx.xxx.xxx": {
       "connections": 1,
       "description": "ASxxxx TEST corpration" 
    },
    "yyy.yyy.yyy.yyy": {
        "connections": 3,
        "description" : "ASxxxx TEST Inc."
    },
    {
        :
       snip
    },
    "total_connections": 100
}
```

## Requirement

- golang  ver > 1.18.0 
- Enable your(target ?) Routinator API

## Usage

`$ go run routinator_connect_client`


## Edit in config.json

```json
{
    "apiURL": "192.168.10.1",   <-- Target Routinator Address 
    "apiPort": "9556"  <-- Target Routinator API port
}
```
