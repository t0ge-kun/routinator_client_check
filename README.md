# routinator_client_check

This program uses the RoutinatorAPI to process the results.  
Output a number of connections per client and where the client belongs.

## Overview

this program use RoutinatorAPI `/api/v1/status`.  
[Routinator API Endpoints link](https://routinator.docs.nlnetlabs.nl/en/stable/api-endpoints.html#:~:text=the%20following%20paths%3A-,/api/v1/status,-Returns%20exhaustive%20information)  

Output a number of clients and an organization to which the client belongs in JSON format for client addresses with "connections" of 1 or more.

Example Output Data

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

## Edit config

Edit in config.json.  
Specify the address or host name (FQDN) and port number of the target Routinator API.

```json
{
    "api_url": "192.168.10.1",   <-- Target Routinator Address 
    "api_port": "9556"  <-- Target Routinator API port
}
```

Only one destination can be specified.

## Requirement

- golang ver:  > 1.18.0
- Routinator ver: 0.12.1
- Enable your(target ?) Routinator's metrics option
  - [`--rtr-client-metrics`](https://routinator.docs.nlnetlabs.nl/en/stable/monitoring.html#:~:text=Metrics%20for%20each%20RTR%20client%20is%20available%20if%20the%20%2D%2Drtr%2Dclient%2Dmetrics%20option%20is%20provided)

## Usage

`$ go run routinator_connect_client`
