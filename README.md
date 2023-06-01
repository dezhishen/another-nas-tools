# another-nas-tools
[![wakatime](https://wakatime.com/badge/user/a2c981ca-317d-4b34-8ed9-4264fbfdb775/project/d4580dbe-68b9-4644-b19e-1fab9a292d96.svg)](https://wakatime.com/badge/user/a2c981ca-317d-4b34-8ed9-4264fbfdb775/project/d4580dbe-68b9-4644-b19e-1fab9a292d96)
- 本服务将以golang开发,使用插件化的方式,提供nas的各种功能
- 根据功能定义interface
- 通过插件实现interface,提供具体的功能实现
  > 例如媒体管理是功能,本地文件管理是一种插件,jellyfin管理是一种插件,他们都实现媒体管理的interface
  
## 功能清单
- 用户\角色\菜单
- 媒体管理
    - 媒体服务器对接
    - tmdb检索
    - 自动下载
    - 自动整理
    - 字幕任务自动提交
- 站点管理
    - 站点对接
- 下载管理
    - 下载任务自动提交
- 消息通知
## 模块
- [ ] [http服务](./pkg/http)
    - [ ] [用户管理](./pkg/http/user)
    - [ ] [角色管理](./pkg/http/role)
    - [ ] [媒体管理](./pkg/http/media)
    - [ ] [任务管理](./pkg/http/task)
    - [ ] [日志管理](./pkg/http/log)
    - [ ] [下载管理](./pkg/http/download)
    - [ ] [插件管理](./pkg/http/plugin)
    - [ ] [健康检查](./pkg/http/health) 
- [x] [应用生命周期](./pkg/lifecycle)
- [x] [任务执行器](./pkg/task)
- [ ] [定时任务](./pkg/cron)
- [ ] [下载管理](./pkg/download)
- [ ] [日志记录](./pkg/log)
- [ ] [插件管理](./pkg/plugin)
