#!/bin/bash
# Usage: ./upload_image.sh path/to/image.jpg
aws s3 cp $1 s3://your-upload-bucket-name/
