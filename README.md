# NIF PT Validator

A portuguese NIF (Número de Identificação Fiscal) validator.

## Usage

### Get the package:

```shell
go get github.com/hugorosario/nifptvalidator
```

### Import the validator

```go
    import	"github.com/hugorosario/nifptvalidator"
```

### Use it

```go
    nif := "123456789"
    fmt.Println(nif, "=>", nifptvalidator.IsValidNif(nif))
```

MIT License
