# GoAutoLoginZzuliStudent
使用 Go 语言编写的用于自动登录郑州轻工业大学校园网的小工具，通过后台服务的方式，实现校园网自动登录验证，断线自动重连。


## 安装方法

### 1. 从 [github release](https://github.com/allwaysLove/GoAutoLoginZzuliStudent/releases) 中下载最新版本
### 2. 通过源码安装
   ```shell
   go build GoAutoLoginZzuliStudent\main -o bin\GoAutoLoginZzuliStudent.exe
   ```

## 使用方法

1. 打开终端，使用 `cd` 命令，将工作目录切换至程序所在目录，执行 `autoLoginWifi install` 命令，按提示输入 **用户名**、**密码** 与 **账户类型**，其中账户类型取值包括：*校园网*、*校园移动*、*校园联通*、*校园单宽*
2. 服务安装完成后，执行 `autoLoginWifi start` 命令即可启动
3. 其他参数详见下方表格

    | 参数名称      | 参数功能   |
    |-----------|--------|
    | install   | 安装服务   |
    | uninstall | 卸载服务   |
    | start     | 启动服务   |
    | stop      | 停止服务   |
    | restart   | 重启服务   |
    | status    | 查看服务状态 |

## 备注

服务开启后，会自动在 *D:* 盘根目录创建一个名为 *GoAutoLoginZzuliStudent.txt* 的文件，用于记录本次运行的日志，关闭服务时会将该日志文件清空。


## 待办事项

- [ ] 添加任务栏托盘图标，用于对服务进行更方便、快捷的控制
