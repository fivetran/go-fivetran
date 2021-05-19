# go-fivetran

## Some notes (to document)

- API Chain
- No external dependencies (only std library)
- Error handling
- Naming: REST API user_id => UserID
- Field ID => Fid, etc...
- we don't validate business rules (unless it is strictly necessary) // in this case the result would be an empty UserDetails{} with a 200 status code... URL paths are validated.
- tests
- context
- json annotations, omitempty. Explain when it should be used.
- type DestinationConfig, Fport interface{} // Type should change to int when https://fivetran.height.app/T-97508 fixed
