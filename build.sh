
#!/bin/bash
echo "编译..."
GOOS=linux  go build main.go
echo "压缩..."
zip function.zip main
echo "更新lambda..."
aws lambda update-function-code --function-name go-* --zip-file "fileb://function.zip"
rm -rf main
rm -rf function.zip
echo "end"
