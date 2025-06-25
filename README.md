# 🖼️ Image Processing Pipeline (Go + AWS Step Functions)

This repository contains a serverless image processing pipeline built with **AWS Lambda**, **Step Functions**, and **Go** using the **AWS Serverless Application Model (SAM)**.

## 📌 Overview

This project demonstrates how to orchestrate multiple AWS services to create an image processing workflow that:

1. **Triggers** when an image is uploaded to S3.
2. **Generates thumbnails** using a Go Lambda function.
3. **Analyzes image labels** using AWS Rekognition.
4. **Stores metadata** in DynamoDB.
5. Optionally integrates **Redis** for caching or task tracking.

---

## 🛠️ Tech Stack

- 🧠 **Go** — for all Lambda functions
- ⚙️ **AWS Lambda** — serverless compute
- 🔄 **AWS Step Functions** — to coordinate the workflow
- 🗂️ **Amazon S3** — to store original images and thumbnails
- 👁️ **Amazon Rekognition** — to detect labels in images
- 📇 **Amazon DynamoDB** — to store metadata
- 🧱 **AWS SAM** — to define and deploy the infrastructure
- 🐳 **Docker** — for local testing with SAM CLI

---

## 🗺️ Features

 - Serverless architecture
 - Go Lambda functions
 - Image resizing using imaging package
 - Label detection using AWS Rekognition
 - Metadata storage in DynamoDB
 - Easily extendable via Step Functions
 - Infrastructure-as-code with AWS SAM

---

## 📸 Sample Use Case

1. Upload photo.jpg to your S3 bucket.
2. A thumbnail is generated and stored in another S3 location.
3. Image labels are detected and logged.
4. Metadata (file name, labels, timestamp) is stored in DynamoDB.

---

## 📥 Contributions
Feel free to fork the repo and submit pull requests. Ideas for improvement are welcome!

---

## 📝 License
This project is open source under the MIT License.

---

## 📧 Contact
Created by: [David Lozano](mailto:david.lozano.hurtado@gmail.com) — feel free to connect!
