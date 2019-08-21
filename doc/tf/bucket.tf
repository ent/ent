resource "aws_s3_bucket" "website" {
  bucket = local.domain_name
  acl    = "public-read"

  website {
    index_document = "index.html"
  }

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
    ]

    resources = [
      format("%s/*", aws_s3_bucket.website.arn),
    ]

    principals {
      identifiers = ["*"]
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

resource "aws_route53_record" "website" {
  name    = local.domain_name
  zone_id = aws_route53_zone.zone.id
  type    = "A"

  alias {
    evaluate_target_health = false
    name                   = aws_s3_bucket.website.website_domain
    zone_id                = aws_s3_bucket.website.hosted_zone_id
  }
}

resource "aws_iam_user" "deployer" {
  name = "${local.name}.deployer"
  path = "/${local.name}/"
}
