// data - traz as informaçoes da cloud e guarda dentro do estado
// data - traz a ami da conta do Gomex
data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["IaaSWeek-${var.hash_commit}"]
  }

  owners = ["777015859311"]
}

resource "aws_instance" "web" {
  ami           = data.aws_ami.ubuntu.id
  for_each      = toset(var.instance_type)
  instance_type = each.value
  tags = {
    Name = "HelloWorld-Terraform"
  }
}