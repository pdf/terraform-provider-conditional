## _DEPRECATED: As of Terraform v0.12.0-alpha1 you should be able to use conditionals with list types._

# Terraform Provider Conditional
Terraform currently lacks the ability to act on a ternary condition with anything but string values. This provider aims to be a stop-gap, providing conditionals for more complex types, until Terraform either gains broader conditional support, or support for custom functions (with HCL2).

## Requirements
- Terraform 0.11.x
- Go 1.9+ and [dep](https://github.com/golang/dep) (for building)

## Build the Provider
```bash
mkdir -p "${GOPATH:-${HOME}/go}/src/github.com/pdf"
cd "${GOPATH:-${HOME}/go}/src/github.com/pdf"
git clone https://github.com/pdf/terraform-provider-conditional.git && cd terraform-provider-conditional
dep ensure
go install
```

## Using the Provider
If you're building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing it in your plugins directory, run `terraform init` to initialize it.

See the [example](https://github.com/pdf/terraform-provider-conditional/tree/master/example) directory for example declarations, below is a typical use-case:

```hcl
/*
 * Declare the conditional.
 *
 * Typically the `if` field (and possibly `then`/`else` fields) will be
 * populated from variables within your module. When the value of `if` is
 * `true`, the value from `then` will populate the `result` field, otherwise
 * `result` will contain the value of the (optional) `else` field.
 *
 */
resource "conditional_list_of_string" "typical" {
  if   = "${true == true}"
  then = ["hello", "universe"]
  else = ["hello", "alternative universe"]
}

resource "some_resource" "instance" {
	list_field = ["${conditional_list_of_string.typical.result}"]
}
```

The currently available conditional resources are:

| Resource name | Description | Example `then`/`else`/`result` values |
|---------------|-------------|----------------|
| **`conditional_list_of_string`** | A list of strings | `["list", "of", "string"]` |
| **`conditional_list_of_list_of_string`** | A list of lists of strings | `[["a", "list"], ["another", "list"]]` |
| **`conditional_list_of_map`** | A list of maps | `[{a = "map"}, {another = "map"}]` |
| **`conditional_map`** | A map | `{a = "map"}` |

Due to the type system, condtionals supporting new data types for the `then`/`else`/`result` fields must be added explicitly.  This is easy to do though, so if you require additional types for results, please feel free to open an issue.
