# Contributing to the Fivetran SDK for Go

Thank you for your interest in contributing to the Fivetran SDK for Go. As this is an initial
release of the Fivetran SDK for Go, we're not ready yet to accept external contributions. 
Please get in touch with us through our [Support Portal](https://support.fivetran.com/) if you 
have any comments, suggestions, support requests, or bug reports.  

This documentation will be updated once we're ready to accept external contributions.

## Guidelines

Even not being ready to accept external contributions yet, the Fivetran SDK for Go is developed
following some code guidelines:

- We only use the standard library. There are no external dependencies.
- We don't do business rules validation. This is a REST API role.
    - But we do some basic validation on URL path construction.
- Each service has three struct types: Service, Request, and Response:
    - Service: exported; all fields are pointers.
    - Request: unexported; all fields are pointers; it's optional - if the HTTP request to the REST API has no payload, there is no need for a Request type.
    - Response: exported, all fields are values. 
- Services can have subtypes, like Config, Auth. 
- Subtypes can have subtypes as well.
- Each subtype has three struct types: The subtype type, Request and Response:
    - Subtype: exported; all fields are pointers.
    - Request: unexported; all fields are pointers.
    - Response: exported; all fields are values.
- Each service must have a valid [example](examples/).
