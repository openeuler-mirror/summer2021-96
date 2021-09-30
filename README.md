<div align="center">
<br/>

  <svg xmlns="http://www.w3.org/2000/svg" style="display: none;">
      <symbol id="cpu-fill" viewBox="0 0 16 16">
                <path d="M6.5 6a.5.5 0 0 0-.5.5v3a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5v-3a.5.5 0 0 0-.5-.5h-3z"/>
                <path d="M5.5.5a.5.5 0 0 0-1 0V2A2.5 2.5 0 0 0 2 4.5H.5a.5.5 0 0 0 0 1H2v1H.5a.5.5 0 0 0 0 1H2v1H.5a.5.5 0 0 0 0 1H2v1H.5a.5.5 0 0 0 0 1H2A2.5 2.5 0 0 0 4.5 14v1.5a.5.5 0 0 0 1 0V14h1v1.5a.5.5 0 0 0 1 0V14h1v1.5a.5.5 0 0 0 1 0V14h1v1.5a.5.5 0 0 0 1 0V14a2.5 2.5 0 0 0 2.5-2.5h1.5a.5.5 0 0 0 0-1H14v-1h1.5a.5.5 0 0 0 0-1H14v-1h1.5a.5.5 0 0 0 0-1H14v-1h1.5a.5.5 0 0 0 0-1H14A2.5 2.5 0 0 0 11.5 2V.5a.5.5 0 0 0-1 0V2h-1V.5a.5.5 0 0 0-1 0V2h-1V.5a.5.5 0 0 0-1 0V2h-1V.5zm1 4.5h3A1.5 1.5 0 0 1 11 6.5v3A1.5 1.5 0 0 1 9.5 11h-3A1.5 1.5 0 0 1 5 9.5v-3A1.5 1.5 0 0 1 6.5 5z"/>
      </symbol>
  </svg>
  <h1 align="center">
    <svg class="bi me-2" width="30" height="24"><use xlink:href="#cpu-fill"/></svg>
    Observer - 帮你更优雅的配置~
  </h1>
  <h4 align="center">
     适用于 openEuler 的可扩展的应用可视化配置软件
  </h4> 
  <h4 align="center">
     🎉 Summer 2021 - openEuler-No.96 🎉
  </h4>



<div align="center">
<p>ID:210010425</p>
</div>
</div>


## 一、项目说明
在 Linux 中，一切都是基于文件，我们常用到的一些软件的配置也是基于文件来配置的。因此我们可以实现一个 Web 端，然后扫描并读取所需软件的配置文件，抽取相关的配置项目，并显示在 Web 前端页面上。同时除了可视化配置之外，也提供一个“高级配置”，用于直接配置对应的配置文件。

![项目结构图](docs/images/项目结构图.png)

## 二、功能点及完成情况描述

| 功能点             | 完成情况 | 备注                                                                                   |
| ------------------ | -------- | -------------------------------------------------------------------------------------- |
| 前后端基本结构     | 100%     | 后端使用 Golang 的 gin 框架来实现，前端使用 BootStrap@v5 和 Jquery 实现                |
| 配置文件的解析     | 80%      | 目前实现了 (MySQL,Redis,Nginx,Crontab) 配置文件的解析工作（待实现 Crontab,Nginx,httpd 等配置文件解析） |
| 前端实现可视化配置 | 80%      | 实现了(MySQL,Redis,Nginx,Crontab) 的可视化配置及服务重启                                              |
| 高级配置           | 80%      | 实现了 (MySQL,Redis,Nginx,Crontab) 的高级配置及服务重启                                                |
| 插件架构           | /     | 待实现

## 三、项目结构
```shell
observer
├─.idea
├─app
│  ├─conf               # 项目配置
│  ├─handler              # 处理器
│  ├─model                
│  ├─route                # 路由
│  └─utils                # 工具包
│      └─go2parse         # 配置文件解析工具
├─docs                    # 项目目录
│  └─images
├─plugins                 # 插件
├─statics                 # 静态资源
│  ├─config_file
│  │  ├─crontab
│  │  ├─httpd
│  │  ├─iptables
│  │  ├─mysql
│  │  ├─nginx
│  │  └─redis
│  ├─css
│  ├─example_file
│  ├─js
│  └─libs 
└─templates               # 前端模板

```

## 四、API 文档
#### 1. 可视化配置读取

> /api/configuration_file/observer

- 请求方法： GET
- 请求示例：`http://127.0.0.1:8080/api/configuration_file/observer?current_state=MySQL`
- 请求参数：

| 参数名称        | 类型      | 说明                              |
| -------------- | ----------- | --------------------------------- |
| current_state | String         | 当前状态 要配置的目标软件的名称例如（MySQL,Redis...） |

- 返回参数：

| 参数名称      | 类型   | 说明                                                  |
| ------------- | ------ | ----------------------------------------------------- |
| current_state | String | 当前状态 要配置的目标软件的名称例如（MySQL,Redis...） |
| config        | String | 解析后的配置项                                        |
| success              |   String     |      数据是否正常（true or false）                                                 |

#### 2. 可视化配置写入

> /api/configuration_file/observer

- 请求方法： POST
- 请求示例：`http://127.0.0.1:8080/api/configuration_file/observer`
- 请求参数：

| 参数名称      | 类型   | 说明                                                  |
| ------------- | ------ | ----------------------------------------------------- |
|       配置项表单数据       |    JSON    |   包含一个 "current_state" 字段来表示当前配置状态                                                    |

- 返回参数：

| 参数名称      | 类型   | 说明                                                  |
| ------------- | ------ | ----------------------------------------------------- |
| success              |   String     |      数据是否正常（true or false）                                                 |

#### 3. 高级配置读取
> /api/configuration_file

- 请求方法： GET
- 请求示例：`http://127.0.0.1:8080/api/configuration_file?current_state=MySQL`
- 请求参数：

| 参数名称        | 类型      | 说明                              |
| -------------- | ----------- | --------------------------------- |
| current_state | String         | 当前状态 要配置的目标软件的名称例如（MySQL,Redis...） |

- 返回参数：

| 参数名称      | 类型   | 说明                                                  |
| ------------- | ------ | ----------------------------------------------------- |
| file_name | String | 配置文件名称（全路径） |
| content        | String | 配置文件的文本内容                                        |
| success              |   String     |      数据是否正常（true or false）

#### 4. 高级配置写入

> /api/configuration_file

- 请求方法： POST
- 请求示例：`http://127.0.0.1:8080/api/configuration_file`
- 请求参数：

| 参数名称      | 类型   | 说明                                                  |
| ------------- | ------ | ----------------------------------------------------- |
|       配置项表单数据       |    JSON    |   包含一个 "current_state" 字段来表示当前配置状态                                                    |

- 返回参数：

| 参数名称 | 类型   | 说明                          |
| -------- | ------ | ----------------------------- |
| success  | String | 数据是否正常（true or false） |
| updatedContent         |    String    |    更新后的配置文件内容                           |
