ALTER TABLE "aws_apigatewayv2_vpc_links"
    ADD COLUMN "vpc_link_id" TEXT;

UPDATE "aws_apigatewayv2_vpc_links"
SET "vpc_link_id" = "id";