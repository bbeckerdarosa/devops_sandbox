variable "aws_region" {
  description = "AWS region to launch servers"
  default = "us-east-1"
}

variable "aws_amis" {
  default = {
    us-east-1 = "ami-{number_ami_id}"
    us-east-1 = "ami-{number_ami_id}"
  }
}

variable "aws_access_key" {
  description = "AWS ACCESS KEY"
}

variable "aws_secret_key" {
  description = "AWS SECRET KEY"
}

variable "key_name" {
  description = "Name of SSH keypair to use in AWS"
}

variable "i_ip_range" {
  type = "list"
  description = "IP range or ips for entry or ingress."
}

variable "ip_range_tcp" {
  type = "list"
  description = "IP range or ips for entry or ingress."
}

variable "e_ip_range" {
  type = "list"
  description = "IP or ips range for egress or egress rule."
}