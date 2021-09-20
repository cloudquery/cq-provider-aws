ALTER TABLE "aws_apigatewayv2_vpc_links"
    DROP COLUMN "vpc_link_id"

--aws_cloudfront_distributions
ALTER TABLE "aws_cloudfront_distributions"
ADD COLUMN "tags" json
ALTER TABLE "aws_cloudfront_distributions" RENAME COLUMN "aliases_items" TO "aliases";