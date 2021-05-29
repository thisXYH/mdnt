# 摘要
做 `MarkDown` 笔记的时候用到的辅助工具集

## 功能
* 引用图片处理
    * [X] ~~*删除没有引用的图片文件*~~ [2021-05-29]
    * [X] ~~*修复引用图片的相对路径引用*~~ [2021-05-29]
    * [ ] 引用网络图片转本地图片引用
    * [ ] 不删除指定命名格式图片文件（norm-xxx.png）

## 如何构建
* 依赖: Golang 1.16

1. 下载库： `git clone git@github.com:thisXYH/NoteTools.git && cd NoteTools`
1. 编译：`go install .`

## 如何使用
````
Usag: nt command [options]

Commands:
        img     图片管理

Detail: nt <command> -h 查看详情

Command img :
  -d    删除无引用的图片，否则只打印路径
  -f    修复引用图片的相对路径，否则只打印路径
  -h    显示帮助菜单
  -i string
        图片目录，不能为空
  -m string
        文档目录，不能为空
  -w    下载引用的网络图片，否则只打印路径
````
