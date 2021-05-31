
//dont know how to add layer version policy
resource "aws_lambda_layer_version" "lambda_layer" {
  filename = "lambda_layer_code/layer.zip"
  layer_name = "test_lambda_layer"

  compatible_runtimes = [
    "nodejs12.x"]
}
