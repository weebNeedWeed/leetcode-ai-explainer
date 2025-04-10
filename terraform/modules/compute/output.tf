output "ecs_cluster_name" {
  value = aws_ecs_cluster.main.name
}

output "ecs_api_service_name" {
  value = aws_ecs_service.api_go_ecs_service.name
}

output "ecs_api_task_definition" {
  value = aws_ecs_task_definition.api_td.family
}

output "ecs_react_service_name" {
  value = aws_ecs_service.react_ecs_service.name
}

output "ecs_react_task_definition" {
  value = aws_ecs_task_definition.react_td.family
}
