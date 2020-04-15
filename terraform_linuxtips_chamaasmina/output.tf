output "ip_address" {
  // value = "${aws_instance.web[*].public_ip}"
  // o sinal => mostra que existe uma relaçao entre o instance id e instance private ip
  value = {
    for instance in aws_instance.web:
    instance.id => instance.private_ip
  }
}