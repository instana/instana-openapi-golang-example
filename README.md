# Example code for using OpenAPI clients with Instana

Side note:
We need the tenant and unit from your environment. When using our SaaS offering, you can get these from your environments' URL. If you are accessing your environment through https://qa-acme.instana.io, then acme is the tenant and qa is your unit.

## Running the example

```
INSTANA_TENANT=instana INSTANA_UNIT=test INSTANA_KEY=$your_token$ go run main.go
```

## License

MIT
