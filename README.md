# 摘要 mdnt - Markdown Notebook Tool
[![Go](https://github.com/thisXYH/mdnt/actions/workflows/go.yml/badge.svg)](https://github.com/thisXYH/mdnt/actions/workflows/go.yml)

维护 `Markdown` 笔记的工具集,方便日常`Markdown`的维护。
比如: 移动了笔记的位置，导致引用的图片路径对不上的问题。

## 功能
* 引用图片处理
  * [X] 删除没有引用的图片文件
  * [X] 修复引用图片的相对路径引用
  * [X] 引用网络图片转本地图片引用
  * [X] 图片目录和笔记目录支持从环境变量读取，详情mdnt help img
* 加密/解密处理
  * [X] 加密/解密指定文件（敏感笔记）
* 引用笔记处理
  * [ ] 为笔记生产唯一id
    > 指定`id`为笔记的第二行
    > **eg:** `> id:e4405e49fa95b2122e5e0965cf1f3724`
  * [ ] 为所有未设置`id`的笔记，设置`id`
  * [ ] 修复引用笔记的相对路径
    > 通过`[alter](path/note.md?<id>)`,语法匹配到对应的笔记

## 如何构建
* 依赖: Golang 1.16

1. 下载库： `git clone git@github.com:thisXYH/mdnt.git && cd mdnt`
1. 编译：`go install .`

## 如何使用
````
Usage:
  mdnt [command]

Available Commands:
  enc         加解密指定笔记
  help        Help about any command
  img         管理 markdown 文档的图片引用

Flags:
  -h, --help      help for nt
  -v, --version   version for nt

Use "nt [command] --help" for more information about a command.
````