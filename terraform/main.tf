resource "aws_cloudwatch_log_group" "hoge" {
  count = var.create_log_group ? 1 : 0
  name  = "hoge"
}

// https://qiita.com/ringo/items/875f08ec550f0826f0dc
// https://www.terraform.io/docs/language/functions/element.html
resource "aws_iam_user" "developer" {
  count = length(var.developer)
  name  = element(var.developer, count.index)
  path  = "/developer"
}

load_balancer_inbound_nat_rules_ids = ["${length(var.lb_natrules) > 0 ? element(concat(var.lb_natrules, list("")), count.index) : ""}"]
