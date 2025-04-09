variable "subnet_ids" {
  type = list(string)
}

variable "sg_ecs_task_id" {
  type = string
}

variable "ecr_repository_url" {
  type = string
}

variable "target_group_arn" {
  type = string
}

variable "ddb_table_name" {
  type    = string
  default = "Leetcode-04102025"
}

variable "github_token" {
  type      = string
  sensitive = true
}

variable "gemini_apikey" {
  type      = string
  sensitive = true
}
