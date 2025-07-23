# Copyright 2025 mjh &lt;694142812@qq.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://example.com/miniblog. The professional
# version of this repository is https://example.com/onex.

## 安装 air 工具
go install github.com/air-verse/air@latest
# 启动 air 工具
air
# air 工具默认读取 .air.toml 文件中的配置
# 如果没有该文件，则会使用默认配置
# 可以使用 -c 参数指定配置文件
# air -c ./scripts/air.toml

## 安装 license 工具
go install github.com/nishanths/license/v5@latest

## 安装 addlicense 工具
go install github.com/marmotedu/addlicense@latest

# 为文件添加许可证头
# 该脚本会在 cmd 目录下的所有 Go 文件中添加许可证头
addlicense -v -f ./scripts/boilerplate.txt --skip-dirs=third_party,_output cmd 