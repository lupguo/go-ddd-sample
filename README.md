## Go DDD微服务设计分层最佳实践
微服务基于DDD领域驱动设计

### 运行
需要先创建DB，具体参数参考`.env`配置

#### 编译运行HTTP服务
```
// 拉代码
git clone github.com/lupguo/go-ddd-sample
// 编译
cd go-ddd-sample
go build ./cmd/imgupload_server
// 运行
./impgupload_server
```

### DDD分层目录结构

```
$ tree -d -NL 2
.
├── application     // [必须]DDD - 应用层
├── cmd             // [必须]参考project-layout，存放CMD
│   ├── imgupload           // 命令行上传图片
│   └── imgupload_server    // 命令行启动Httpd服务
├── deployments     // 参考project-layout，服务部署相关
├── docs            // 参考project-layout，文档相关
├── domain          // [必须]DDD - 领域层
│   ├── entity      //  - 领域实体
│   ├── repository  //  - 领域仓储接口
│   ├── service     //  - 领域服务，多个实体的能力聚合
│   └── valobj      //  - 领域值对象
├── infrastructure  // [必须]DDD - 基础层
│   └── persistence //  - 数据库持久层
├── interfaces      // [必须]DDD - 接口层
│   └── api         //  - RESTful API接口对外暴露
├── pkg             // [可选]参考project-layout，项目包，还有internel等目录结构，依据服务实际情况考虑
└── tests           // [可选]参考project-layout，测试相关
    └── mock
```
### 其他目录结构
- cmd: 命令文件生成
- deployments: 服务部署相关
- tests: 服务测试相关
- docs: 服务文档相关
- .env: 服务环境配置(可以基于config配置)

### TODO
1. 其余几个接口补充
2. 前端展示
3. 剩余模块服务补充
4. Dockerfile和Makefile补充
5. 更合适的文件或目录命名

### 参考
1. Duck-DDD: https://medium.com/learnfazz/domain-driven-design-in-go-253155543bb1
2. https://dev.to/stevensunflash/using-domain-driven-design-ddd-in-golang-3ee5
3. https://github.com/takashabe/go-ddd-sample
4. https://github.com/victorsteven/food-app-server