# AWS S3 File Upload Service

This is a simple file upload service built using **Go** that uploads files to an AWS S3 bucket. It adheres to **SOLID principles**, ensures unique file naming using UUIDs, and is containerized using Docker.

---

## Features

- Upload files to AWS S3 with a unique filename.
- Generate public URLs for uploaded files.
- Handles different file types and maintains the original file extension.
- Adheres to SOLID principles for clean, maintainable code.
- Environment variable support for secure configuration.
- Containerized using Docker.

---

## Prerequisites

- **AWS S3 Bucket** with proper permissions.
- **Docker** (if you plan to run the service in a container)
---

## Environment Variables
You need to configure the following environment variables in a `.env` file or directly in your system:

```env
AWS_REGION=your-aws-region
S3_BUCKET=your-s3-bucket-name
AWS_ACCESS_KEY_ID=your-aws-access-key-id
AWS_SECRET_ACCESS_KEY=your-aws-secret-access-key
```