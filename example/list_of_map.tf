// Typical invocation
resource "conditional_list_of_map" "typical" {
  if = "${true == true}"

  then = [{
    result = true
  }]

  else = [{
    result = false
  }]
}
