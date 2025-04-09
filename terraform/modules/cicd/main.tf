terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

resource "aws_iam_role" "codepipeline" {
  assume_role_policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Principal" : {
            "Service" : "codepipeline.amazonaws.com"
          },
          "Action" : "sts:AssumeRole"
        }
      ]
    }
  )
}

resource "aws_iam_policy" "codepipeline_policy" {
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "codeconnections:GetConnectionToken",
          "codeconnections:GetConnection",
          "codeconnections:UseConnection",
          "codestar-connections:GetConnectionToken",
          "codestar-connections:GetConnection",
          "codestar-connections:UseConnection"
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "exec_role_attach" {
  role       = aws_iam_role.codepipeline.name
  policy_arn = aws_iam_policy.codepipeline_policy.arn
}

resource "aws_s3_bucket" "artifacts_bucket" {
  bucket        = "codebuild-artifacts-04092025"
  force_destroy = true
}

resource "aws_codestarconnections_connection" "github_connection" {
  name          = "leetcode-ai-connection"
  provider_type = "GitHub"
}

resource "aws_codepipeline" "pipeline" {
  name     = "deploy-pipeline"
  role_arn = aws_iam_role.codepipeline.arn

  artifact_store {
    location = aws_s3_bucket.artifacts_bucket.id
    type     = "S3"
  }

  stage {
    name = "Source"

    action {
      name             = "Source"
      category         = "Source"
      owner            = "AWS"
      provider         = "CodeStarSourceConnection"
      version          = "1"
      output_artifacts = ["source_output"]

      configuration = {
        ConnectionArn    = aws_codestarconnections_connection.github_connection.arn
        FullRepositoryId = "weebNeedWeed/leetcode-ai-explainer"
        BranchName       = "master"
      }
    }
  }

  stage {
    name = "Build"

    action {
      name             = "Build"
      category         = "Build"
      owner            = "AWS"
      provider         = "CodeBuild"
      input_artifacts  = ["source_output"]
      output_artifacts = ["build_output"]
      version          = "1"

      configuration = {
        ProjectName = "test"
      }
    }
  }
}
