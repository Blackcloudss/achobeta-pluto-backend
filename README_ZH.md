# AchoBeta 冥王星系统 - 2024年G组 - 后端 - 复试项目
AchoBeta 冥王星系统 - 2024年G组 - 后端 - 复试项目

```
   _______   ___       ____  ____  ___________  ______    
  |   __ "\ |"  |     ("  _||_ " |("     _   ")/    " \   
  (. |__) :)||  |     |   (  ) : | )__/  \\__/// ____  \  
  |:  ____/ |:  |     (:  |  | . )    \\_ /  /  /    ) :) 
  (|  /      \  |___   \\ \__/ //     |.  | (: (____/ //  
 /|__/ \    ( \_|:  \  /\\ __ //\     \:  |  \        /   
(_______)    \_______)(__________)     \__|   \"_____/    
```



## 项目功能
团队管理系统

## 项目技术栈
1. Gin
2. Gorm
3. Mysql
4. Redis
5. zap
6. viper
7. ...
## 项目目录结构
```
├── cmd       程序入口
├── configs   配置实体类
├── global    全局变量和常量
├── initalize 初始化文件
├── internal  核心业务代码
├── log       日志文件配置
├── pkg       第三方通用包
├── test      测试文件
├── utils     工具类
├── config.yaml.template    配置文件模板
├── go.mod
└── go.sum
```
## 开发前

### 1. **设置git hook**
```shell
git config core.hooksPath .githooks 
chmod -R -x .githooks 
```

### 2. **阅读以下开发规范**

分支命名规范
我们必须确认：

1. 分支命名应包含负责人的名字。

2. 分支命名必须清楚地表达分支正在处理的问题。

因此分支命名必须标准化。
```bash
<type>-<name>-<description>
```
例如：
- 如果是开发新功能的分支，命名规范如下
```bash
feature-<name>-<feature description>
例如：feature-jett-dev_log_system
```

- 如果是修复bug：
```bash
bugfix-<name>-<bug name>
例如：bugfix-jett-login_error
```
以及其他类型：
`hotfix`、`release`...


### 提交信息格式
提交信息应尽可能清晰，每次提交只做一件事。

```bash
<type>(<scope>): <subject>

例如：feat: add new api
或：feat(common): add new api
```

### 类型

```text
# 主要类型
feat:      添加新功能
fix:       修复bug

# 特殊类型
docs:      仅更改文档相关内容
style:     不影响代码含义的更改，例如删除空格、更改缩进、添加或删除分号
build:     更改构建工具或外部依赖项，例如webpack、npm
refactor:  重构代码时使用
revert:    执行git revert时打印的消息

# 暂不使用的类型
test:      添加或修改现有测试
perf:      改善性能的更改
ci:        与CI（持续集成服务）相关的更改
chore:     不修改src或test的其他修改，例如更改构建过程或辅助工具
```

### 主题

末尾不加句号或标点符号

例如：
```bash
feat: add new feature
fix: fix a bug
```

### 内容
请删除无用的导入。您还可以通过设置GoLand使用快捷键ctrl + alt + o自动删除无用的导入。

## **你必须知道**
1. **不要**提交任何敏感信息，例如`api_key`、`address`或`password`。
2. 您可以使用配置文件`config.yaml`来存储某些敏感信息，但不要试图提交它。每次修改`config.yaml`的结构后，您必须同步更新`config.yaml.template`。
3. 任何时候不要用 `git push --force` 除非你知道你在干什么。

## 待办事项
- [ ] 项目初始化
    - [ ] 项目初始化
    - [ ] 项目目录结构
    - [ ] 项目配置文件
    - [ ] 项目全局变量和常量
    - [ ] 项目日志配置
    - [ ] 项目初始化文件
- [ ] gin模块搭建
    - [ ] gin框架搭建
    - [ ] gin路由搭建
    - [ ] gin中间件搭建
    - [ ] gin参数绑定
    - [ ] gin返回数据封装
- [ ] 登录和注册模块
    - [ ] 登录和登出功能
    - [ ] 30天内自动登录
    - [ ] 强制下线
    - [ ] 显示常用设备
- [ ] 团队主页模块
    - [ ] 个人中心
    - [ ] 积分概览
    - [ ] 消息模块
    - [ ] 飞书多维表格
- [ ] 团队信息模块
    - [ ] 点赞个人信息
    - [ ] 用户列表
    - [ ] 新增用户
    - [ ] 管理员团队架构管理
    - [ ] 查看和编辑用户信息详情