generate swagger yaml file

```
make swagger
```

generate client from the swagger.yaml

```
cd sdk
swagger generate client -f ../swagger.yaml -A product-api
```
