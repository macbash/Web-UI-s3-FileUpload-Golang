# s3-fileUpload-golang
Simple web UI and Golang Backend to upload file in AWS S3

## Steps
1. Start Golang Backend
   ```
   go run main.go
   ```
2. Open index.html in the browser, select the file wish to upload in s3

Note:-
a) Pl export the AWS Creds/IAM Role instance profile before running the script.
b) Ensure the s3 privileges required to upload objects in buckets.
c) Add the S3 Bucket name and region in the main.go file
   
