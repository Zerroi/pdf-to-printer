# PDF to Printer

一个简单易用的Go语言库，用于在Windows系统上打印PDF文件。该库通过调用SumatraPDF实现打印功能，可以轻松集成到你的Go项目中。

## 功能特性

- ✅ 简单易用的API接口
- ✅ 支持指定打印机打印
- ✅ 支持使用默认打印机打印
- ✅ 支持打印份数设置
- ✅ 支持页面范围打印
- ✅ 支持纸张方向设置
- ✅ 支持缩放模式设置
- ✅ 支持颜色模式设置
- ✅ 支持双面打印设置
- ✅ 支持纸张大小设置
- ✅ 支持显示打印对话框
- ✅ 自动检测SumatraPDF安装路径
- ✅ 获取系统打印机列表
- ✅ 获取默认打印机名称
- ✅ 完整的错误处理

## 系统要求

- Windows 7 或更高版本
- Go 1.21 或更高版本
- [SumatraPDF](https://www.sumatrapdfreader.org/download-free-pdf-viewer.html)

## 安装

### 1. 安装SumatraPDF

从 [SumatraPDF官网](https://www.sumatrapdfreader.org/download-free-pdf-viewer.html) 下载并安装SumatraPDF。

### 2. 安装Go库

```bash
go get github.com/zerroi/pdf-to-printer
```

## 快速开始

### 基本使用

```go
package main

import (
    "log"
    "github.com/zerroi/pdf-to-printer/printer"
)

func main() {
    // 使用默认打印机打印PDF
    err := printer.PrintToDefault("document.pdf")
    if err != nil {
        log.Fatal(err)
    }
}
```

### 指定打印机打印

```go
package main

import (
    "log"
    "github.com/zerroi/pdf-to-printer/printer"
)

func main() {
    options := printer.PrintOptions{
        PrinterName: "HP LaserJet",
        Copies:      2,
        Silent:      true,
    }
    err := printer.Print("document.pdf", options)
    if err != nil {
        log.Fatal(err)
    }
}
```

### 高级打印选项

```go
package main

import (
    "log"
    "github.com/zerroi/pdf-to-printer/printer"
)

func main() {
    options := printer.PrintOptions{
        PrinterName: "HP LaserJet",
        Copies:      1,
        Silent:      true,
        PageRange:   "1-3,5",           // 打印第1-3页和第5页
        Orientation: "portrait",        // 纵向 (portrait/landscape)
        Scaling:     "fit",            // 适应纸张 (noscale/shrink/fit)
        ColorMode:   "color",          // 彩色打印 (color/monochrome)
        Duplex:      "duplexlong",     // 长边双面 (duplex/duplexshort/duplexlong/simplex)
        PaperSize:   "A4",             // 纸张大小 (A2/A3/A4/A5/A6/letter/legal/tabloid/statement)
    }
    err := printer.Print("document.pdf", options)
    if err != nil {
        log.Fatal(err)
    }
}
```

### 显示打印对话框

```go
package main

import (
    "log"
    "github.com/zerroi/pdf-to-printer/printer"
)

func main() {
    options := printer.PrintOptions{
        ShowDialog: true,  // 显示打印对话框，打印完成后自动退出
    }
    err := printer.Print("document.pdf", options)
    if err != nil {
        log.Fatal(err)
    }
}
```

### 获取打印机列表

```go
package main

import (
    "fmt"
    "log"
    "github.com/zerroi/pdf-to-printer/printer"
)

func main() {
    printers, err := printer.GetPrinters()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("可用打印机:")
    for _, p := range printers {
        fmt.Println("-", p)
    }
}
```

### 使用Printer接口

```go
package main

import (
    "log"
    "github.com/zerroi/pdf-to-printer/printer"
)

func main() {
    p := printer.NewPrinter()
    
    // 检查SumatraPDF是否可用
    if !p.IsSumatraAvailable() {
        log.Fatal("SumatraPDF未找到，请先安装")
    }
    
    // 打印PDF
    err := p.PrintToDefault("document.pdf")
    if err != nil {
        log.Fatal(err)
    }
}
```

### 自定义SumatraPDF路径

```go
package main

import (
    "log"
    "github.com/zerroi/pdf-to-printer/printer"
)

func main() {
    // 设置自定义SumatraPDF路径
    err := printer.SetConfig(printer.Config{
        SumatraPath: "C:\\Custom\\Path\\SumatraPDF.exe",
    })
    if err != nil {
        log.Fatal(err)
    }
    
    // 打印PDF
    err = printer.PrintToDefault("document.pdf")
    if err != nil {
        log.Fatal(err)
    }
}
```

## API文档

### PrintOptions

打印选项结构体：

```go
type PrintOptions struct {
    PrinterName string  // 打印机名称（为空则使用默认打印机）
    Copies      int     // 打印份数（默认1）
    Silent      bool    // 静默模式（默认true）
    PageRange   string  // 页面范围，如 "1-3,5,10-8"
    Orientation string  // 纸张方向: "portrait" 或 "landscape"
    Scaling     string  // 缩放: "noscale", "shrink", "fit"
    ColorMode   string  // 颜色模式: "color" 或 "monochrome"
    Duplex      string  // 双面打印: "duplex", "duplexshort", "duplexlong", "simplex"
    Bin         string  // 纸张托盘: bin=<num or name>
    PaperSize   string  // 纸张大小: A2, A3, A4, A5, A6, letter, legal, tabloid, statement
    ShowDialog  bool    // 显示打印对话框
}
```

**打印选项说明：**

- `PageRange`: 页面范围，支持多种格式：
  - `"1-3"`: 打印第1到第3页
  - `"1,3,5"`: 打印第1、3、5页
  - `"1-3,5,10-8"`: 打印第1-3页、第5页、第8-10页
  - `"odd"`: 打印奇数页
  - `"even"`: 打印偶数页

- `Orientation`: 纸张方向
  - `"portrait"`: 纵向（默认）
  - `"landscape"`: 横向

- `Scaling`: 缩放模式
  - `"noscale"`: 不缩放
  - `"shrink"`: 缩小以适应纸张
  - `"fit"`: 适应纸张大小

- `ColorMode`: 颜色模式
  - `"color"`: 彩色打印
  - `"monochrome"`: 单色打印

- `Duplex`: 双面打印
  - `"duplex"`: 双面打印
  - `"duplexshort"`: 短边双面
  - `"duplexlong"`: 长边双面
  - `"simplex"`: 单面打印

- `PaperSize`: 纸张大小
  - `"A2"`, `"A3"`, `"A4"`, `"A5"`, `"A6"`
  - `"letter"`, `"legal"`, `"tabloid"`, `"statement"`

- `ShowDialog`: 显示打印对话框
  - `true`: 显示打印对话框，打印完成后自动退出
  - `false`: 直接打印（默认）

### Printer接口

打印器接口定义：

```go
type Printer interface {
    // Print 打印PDF文件
    Print(pdfPath string, options PrintOptions) error
    
    // PrintToDefault 使用默认打印机打印
    PrintToDefault(pdfPath string) error
    
    // GetPrinters 获取可用打印机列表
    GetPrinters() ([]string, error)
    
    // GetDefaultPrinter 获取默认打印机名称
    GetDefaultPrinter() (string, error)
}
```

### 主要函数

#### `NewPrinter() *PDFPrinter`

创建新的打印器实例。

#### `Print(pdfPath string, options PrintOptions) error`

打印PDF文件。

#### `PrintToDefault(pdfPath string) error`

使用默认打印机打印PDF文件。

#### `GetPrinters() ([]string, error)`

获取系统可用打印机列表。

#### `GetDefaultPrinter() (string, error)`

获取系统默认打印机名称。

#### `SetConfig(config Config) error`

设置全局配置。

#### `GetConfig() Config`

获取当前配置。

### PDFPrinter方法

#### `IsSumatraAvailable() bool`

检查SumatraPDF是否可用。

## 错误处理

库定义了以下错误类型：

- `ErrSumatraNotFound` - SumatraPDF未找到
- `ErrPrinterNotFound` - 打印机未找到
- `ErrFileNotFound` - PDF文件未找到
- `ErrInvalidPDF` - 无效的PDF文件
- `ErrPrintFailed` - 打印操作失败
- `ErrInvalidOptions` - 无效的打印选项

## 项目结构

```
pdf-to-printer/
├── printer/               # 核心包
│   ├── printer.go         # 主要API接口
│   ├── sumatra.go         # SumatraPDF相关操作
│   ├── config.go          # 配置管理
│   └── errors.go          # 自定义错误类型
├── internal/              # 内部工具
│   ├── utils.go           # 工具函数
│   └── windows.go         # Windows特定功能
├── examples/              # 使用示例
│   └── basic.go           # 基本使用示例
└── plans/                 # 设计文档
    └── project-plan.md    # 项目计划
```

## 运行示例

```bash
# 进入示例目录
cd examples

# 运行基本示例（需要有一个名为document.pdf的PDF文件）
go run basic.go
```

## 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./printer
go test ./internal

# 查看测试覆盖率
go test -cover ./...
```

## 注意事项

1. 本库仅支持Windows系统，因为SumatraPDF是Windows应用程序。
2. 确保SumatraPDF已正确安装并可在系统中找到。
3. 打印机名称必须与系统中显示的名称完全一致（区分大小写）。
4. PDF文件路径可以是相对路径或绝对路径。

## 许可证

MIT License

## 贡献

欢迎提交问题和拉取请求！

## 相关链接

- [SumatraPDF官网](https://www.sumatrapdfreader.org/)
- [npm pdf-to-printer](https://www.npmjs.com/package/pdf-to-printer) - 参考的Node.js版本

## 更新日志

### v1.1.0 (2024-01-04)

- 根据SumatraPDF官方文档更新打印选项
- 支持页面范围打印
- 支持纸张方向、缩放、颜色模式设置
- 支持双面打印设置
- 支持纸张大小设置
- 支持显示打印对话框
- 移除不存在的-version参数

### v1.0.0 (2024-01-04)

- 初始版本发布
- 支持基本PDF打印功能
- 支持打印机列表查询
- 支持默认打印机获取
- 完整的错误处理
