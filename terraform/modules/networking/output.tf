output "lb_alias" {
  value = {
    dns_name = aws_lb.lb.dns_name
    zone_id  = aws_lb.lb.zone_id
  }
}

output "private_subnet_ids" {
  value = [for i in aws_subnet.private_subnet : i.id]
}

output "sg_ecs_task_id" {
  value = aws_security_group.sg_ecs_task.id
}

output "target_group_arn" {
  value = aws_lb_target_group.tg.arn
}
