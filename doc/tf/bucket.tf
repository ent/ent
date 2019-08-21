resource "aws_s3_bucket" "website" {
  bucket = local.domain_name
  acl    = "private"

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["GET", "HEAD"]
    allowed_origins = ["*"]
    expose_headers  = ["ETag"]
    max_age_seconds = 3600
  }
}


data "aws_iam_policy_document" "website" {
  statement {
    actions = [
      "s3:GetObject",
      "s3:ListBucket",
    ]

    resources = [
      aws_s3_bucket.website.arn,
      format("%s/*", aws_s3_bucket.website.arn),
    ]

    principals {
      identifiers = [aws_cloudfront_origin_access_identity.website.iam_arn]
      type        = "AWS"
    }
  }

  statement {
    actions = [
      "s3:PutObjectAcl",
      "s3:PutObject",
      "s3:ListBucketMultipartUploads",
      "s3:ListBucket",
      "s3:GetObject",
      "s3:GetBucketLocation",
      "s3:DeleteObject",
      "s3:AbortMultipartUpload",
    ]

    resources = [
      aws_s3_bucket.website.arn,
      format("%s/*", aws_s3_bucket.website.arn),
    ]

    principals {
      identifiers = [aws_iam_user.deployer.arn]
      type        = "AWS"
    }
  }
}

resource "aws_s3_bucket_policy" "website" {
  bucket = aws_s3_bucket.website.id
  policy = data.aws_iam_policy_document.website.json
}

resource "aws_iam_user" "deployer" {
  name = "${local.name}.deployer"
  path = "/${local.name}/"
}
