# 摘要 mdnt - Markdown Notebook Tool
[![Go](https://github.com/thisXYH/mdnt/actions/workflows/go.yml/badge.svg)](https://github.com/thisXYH/mdnt/actions/workflows/go.yml)

维护 `Markdown` 笔记的工具集,方便日常`Markdown`的维护。
比如: 移动了笔记的位置，导致引用的图片路径对不上的问题。

## 功能
* 引用图片处理
  * [X] 删除没有引用的图片文件  [2021-05-29]
  * [X] 修复引用图片的相对路径引用  [2021-05-29]
  * [X] 引用网络图片转本地图片引用  [2021-06-01]
  * [X] 图片目录和笔记目录支持从环境变量读取，详情mdnt help img  [2021-06-01]
* 加密/解密处理
  * [ ] 加密/解密指定文件

## 如何构建
* 依赖: Golang 1.16

1. 下载库： `git clone git@github.com:thisXYH/mdnt.git && cd mdnt`
1. 编译：`go install .`

## 如何使用
````
Usage:
  mdnt [command]

Available Commands:
  help        Help about any command
  img         管理 markdown 文档的图片引用

Flags:
  -h, --help      help for nt
  -v, --version   version for nt

Use "nt [command] --help" for more information about a command.
````