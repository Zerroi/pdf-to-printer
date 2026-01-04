package printer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/zerroi/pdf-to-printer/internal"
)

// 常见的SumatraPDF安装路径
var commonSumatraPaths = []string{
	`C:\Program Files\SumatraPDF\SumatraPDF.exe`,
	`C:\Program Files (x86)\SumatraPDF\SumatraPDF.exe`,
	`C:\Users\%USERNAME%\AppData\Local\SumatraPDF\SumatraPDF.exe`,
}

// FindSumatraPDF 查找SumatraPDF可执行文件
func FindSumatraPDF() (string, error) {
	// 首先检查配置中是否指定了路径
	config := GetConfig()
	if config.SumatraPath != "" {
		if internal.FileExists(config.SumatraPath) {
			return config.SumatraPath, nil
		}
		return "", ErrSumatraNotFound
	}

	// 检查常见安装路径
	for _, path := range commonSumatraPaths {
		// 展开环境变量
		expandedPath := os.ExpandEnv(path)
		if internal.FileExists(expandedPath) {
			return expandedPath, nil
		}
	}

	// 在PATH中查找
	if path := internal.FindExecutable("SumatraPDF"); path != "" {
		return path, nil
	}

	return "", ErrSumatraNotFound
}

// BuildPrintCommand 构建打印命令
func BuildPrintCommand(pdfPath, printerName string, options PrintOptions) ([]string, error) {
	// 验证PDF文件
	if !internal.FileExists(pdfPath) {
		return nil, ErrFileNotFound
	}

	if !internal.IsPDFFile(pdfPath) {
		return nil, ErrInvalidPDF
	}

	// 获取绝对路径
	absPDFPath, err := filepath.Abs(pdfPath)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrFileNotFound, err)
	}

	// 构建命令参数
	var args []string

	// 如果显示打印对话框
	if options.ShowDialog {
		args = append(args, "-print-dialog", "-exit-when-done")
	} else {
		// 使用指定打印机或默认打印机
		if printerName != "" {
			args = append(args, "-print-to", printerName)
		} else {
			args = append(args, "-print-to-default")
		}

		// 静默模式
		if options.Silent {
			args = append(args, "-silent")
		}

		// 构建打印设置
		var settings []string

		// 页面范围
		if options.PageRange != "" {
			settings = append(settings, options.PageRange)
		}

		// 纸张方向
		if options.Orientation == "portrait" || options.Orientation == "landscape" {
			settings = append(settings, options.Orientation)
		}

		// 缩放
		if options.Scaling == "noscale" || options.Scaling == "shrink" || options.Scaling == "fit" {
			settings = append(settings, options.Scaling)
		}

		// 颜色模式
		if options.ColorMode == "color" || options.ColorMode == "monochrome" {
			settings = append(settings, options.ColorMode)
		}

		// 双面打印
		if options.Duplex == "duplex" || options.Duplex == "duplexshort" || options.Duplex == "duplexlong" || options.Duplex == "simplex" {
			settings = append(settings, options.Duplex)
		}

		// 纸张托盘
		if options.Bin != "" {
			settings = append(settings, options.Bin)
		}

		// 纸张大小
		if options.PaperSize != "" {
			settings = append(settings, "paper="+options.PaperSize)
		}

		// 打印份数
		if options.Copies > 1 {
			settings = append(settings, fmt.Sprintf("%dx", options.Copies))
		}

		// 添加打印设置参数
		if len(settings) > 0 {
			args = append(args, "-print-settings", strings.Join(settings, ","))
		}
	}

	// 添加PDF文件路径
	args = append(args, absPDFPath)

	return args, nil
}
