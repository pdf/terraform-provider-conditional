// Typical invocation
//
// conditional_list_of_string.typical.result == ["hello", "universe"]
resource "conditional_list_of_string" "typical" {
  if   = "${true == true}"
  then = ["hello", "universe"]
  else = ["hello", "alternative universe"]
}

// It is valid to have an empty `else` field, in which case:
//
// conditional_list_of_string.empty_else.result == []
resource "conditional_list_of_string" "empty_else" {
  if   = "${true == false}"
  then = ["hello", "universe"]
}

// Count is also supported
//
// conditional_list_of_string.with_count.*.result[0] == ["true", "0"]
// conditional_list_of_string.with_count.*.result[1] == ["false", "1"]
// conditional_list_of_string.with_count.*.result[2] == ["true", "2"]
locals {
  conditions = [
    true,
    false,
    true,
  ]

  // List of lists needs to be constructed using `list` function due to:
  // https://github.com/hashicorp/terraform/issues/15971

  trueValues  = "${list(list("true", "0"), list("true", "1"), list("true", "2"))}"
  falseValues = "${list(list("false", "0"), list("false", "1"), list("false", "2"))}"
}

resource "conditional_list_of_string" "with_count" {
  if   = "${local.conditions[count.index]}"
  then = "${local.trueValues[count.index]}"
  else = "${local.falseValues[count.index]}"

  count = 3
}
