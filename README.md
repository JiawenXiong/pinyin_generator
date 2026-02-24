# 拼音生成器

**[English Version](README_EN.md) | [中文版本](README.md)**

一个简洁的中文拼音生成工具，可以将中文文本转换为拼音并自动复制到剪贴板。

## 功能特性

- 将中文文本转换为拼音（不带声调）
- 生成全拼和首字母组合
- 自动将结果复制到系统剪贴板

## 安装

### 从源码编译

```bash
go build -o pinyingen.exe main.go
```

## 使用方法

### 基本用法

```bash
pinyingen.exe "中文文本"
```

### 示例

```bash
pinyingen.exe "今日头条"
```

**输出：**
```
已复制到剪贴板: jinritoutiaojrtt
```

生成的结果包含：
- `jinritoutiao` - 全拼
- `jrtt` - 首字母组合

## 命令行参数

```bash
pinyingen.exe [options] "中文文本"
```

使用 `pinyingen.exe -h` 查看帮助信息。

## 依赖

- Go 1.25.0+
- [github.com/atotto/clipboard](https://github.com/atotto/clipboard) - 剪贴板操作
- [github.com/mozillazg/go-pinyin](https://github.com/mozillazg/go-pinyin) - 拼音转换

## 许可证

本项目采用 MIT 许可证。