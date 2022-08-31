# lambda 使用流程
   1.创建函数
   2.创建触发器 apiGateway (resetApi)
   3.设置路由 /{proxy+}
   4.设置vpc
   5.设置环境变量(数据库配置....)

# 部署
 - 下载aws cli
    https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-help.html
 - aws configure   
    AWS Access Key ID [****************6V7A]: PotoPrintPhp.env.S3_KEY
    AWS Secret Access Key [****************qKHa]: PotoPrintPhp.env.S3_SECRET
    Default region name [us-east-1]: us-east-1
    Default output format [json]: json
- 直接运行 build.sh  

- (单独) 编译 压缩项目 命令
    GOOS=linux go build main
    zip function.zip main
- (单独) 更新lambda
    aws lambda update-function-code --function-name go-poto --zip-file "fileb://function.zip"


# 坑
  给role权限之后 信任策略一定要更新

# gin
- https://segmentfault.com/a/1190000013297683

# 测试url
 - https://ibbv4y3f1j.execute-api.us-east-1.amazonaws.com/default/hello-world

# 本地测试

 - 停止项目 command/control + c
 - 运行  go run main.go

# 本地配置文件
   conf/app.ini
