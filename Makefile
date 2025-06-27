APP_NAME=image-processing-pipeline
S3_BUCKET=sam-build-artifacts-bucket
STACK_NAME=image-processing-stack
REGION=eu-west-3
PROFILE=default
EVENT_DIR=testdata/events

.PHONY: all build deploy delete logs

build:
	sam build -b deployments

deploy:
	sam deploy \
		--template-file deployments/template.yaml \
		--stack-name $(STACK_NAME) \
		--s3-prefix image-processing-artifacts \
		--s3-bucket $(S3_BUCKET) \
		--capabilities CAPABILITY_IAM CAPABILITY_NAMED_IAM CAPABILITY_AUTO_EXPAND \
		--region $(REGION) \
		--profile $(PROFILE) \
		--confirm-changeset \
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

test-detect-labels:
	sam local invoke DetectLabelsFunction --event $(EVENT_DIR)/s3-upload.json

test-generate-thumbnails:
	sam local invoke GenerateThumbnailsFunction --event $(EVENT_DIR)/s3-upload.json

test-store-metadata:
	sam local invoke StoreMetadataFunction --event $(EVENT_DIR)/store-metadata.json
