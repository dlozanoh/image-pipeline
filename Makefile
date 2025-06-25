APP_NAME=image-processing-pipeline
S3_BUCKET=my-sam-artifacts-bucket
STACK_NAME=image-processing-stack
REGION=eu-west-1
PROFILE=default

.PHONY: all build deploy delete logs

build:
	sam build -b deployments

deploy:
	sam deploy \
		--stack-name $(STACK_NAME) \
		--s3-bucket $(S3_BUCKET) \
		--capabilities CAPABILITY_IAM CAPABILITY_NAMED_IAM \
		--region $(REGION) \
		--profile $(PROFILE) \
		--confirm-changeset \
		--resolve-s3 \
		--parameter-overrides \
			AppName=$(APP_NAME)

delete:
	aws cloudformation delete-stack \
		--stack-name $(STACK_NAME) \
		--region $(REGION) \
		--profile $(PROFILE)

logs:
	sam logs -n $(FUNCTION) --stack-name $(STACK_NAME) --tail --profile $(PROFILE)
	``
update:
	go mod tidy

build:
	sam build -b deployments