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
  instance_type = "t2.micro"

  dynamic "ebs_block_device" {
    for_each = var.blocks
    content {
      device_name = ebs_block_device.value["device_name"]
      volume_size = ebs_block_device.value["volume_size"]
      volume_type = ebs_block_device.value["volume_type"]
    }
  }
  tags = {
    Name = "HelloWorld-Terraform"
  }
}