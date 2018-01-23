// Typical invocation
resource "conditional_map" "typical" {
  if = "${true == true}"

  then = {
    result = true
  }

  else = {
    result = false
  }
}
