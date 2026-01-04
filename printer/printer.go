package printer

import (
	"fmt"
	"os/exec"

	"github.com/zerroi/pdf-to-printer/internal"
)

// PrintOptions 打印选项
type PrintOptions struct {
	PrinterName string // 打印机名称（为空则使用默认打印机）
	Copies      int    // 打印份数（默认1）
	Silent      bool   // 静默模式（默认true）
	PageRange   string // 页面范围，如 "1-3,5,10-8"
	Orientation string // 纸张方向: "portrait" 或 "landscape"
	Scaling     string // 缩放: "noscale", "shrink", "fit"
	ColorMode   string // 颜色模式: "color" 或 "monochrome"
	Duplex      string // 双面打印: "duplex", "duplexshort", "duplexlong", "simplex"
	Bin         string // 纸张托盘: bin=<num or name>
	PaperSize   string // 纸张大小: A2, A3, A4, A5, A6, letter, legal, tabloid, statement
	ShowDialog  bool   // 显示打印对话框
}

// PDFPrinter PDF打印器实现
type PDFPrinter struct {
	sumatraPath string
}

// NewPrinter 创建新的打印器实例
func NewPrinter() *PDFPrinter {
	sumatraPath, _ := FindSumatraPDF()
	return &PDFPrinter{
		sumatraPath: sumatraPath,
	}
}

// Print 打印PDF文件
func (p *PDFPrinter) Print(pdfPath string, options PrintOptions) error {
	// 验证选项
	if options.Copies < 1 {
		options.Copies = 1
	}

	// 查找SumatraPDF
	if p.sumatraPath == "" {
		sumatraPath, err := FindSumatraPDF()
		if err != nil {
			return err
		}
		p.sumatraPath = sumatraPath
	}

	// 确定打印机名称
	printerName := options.PrinterName
	if printerName == "" {
		var err error
		printerName, err = internal.GetDefaultPrinter()
		if err != nil {
			return fmt.Errorf("failed to get default printer: %w", err)
		}
	}

	// 验证打印机是否存在
	if !internal.PrinterExists(printerName) {
		return fmt.Errorf("%w: %s", ErrPrinterNotFound, printerName)
	}

	// 构建打印命令
	args, err := BuildPrintCommand(pdfPath, printerName, options)
	if err != nil {
		return err
	}

	// 执行打印命令
	cmd := exec.Command(p.sumatraPath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%w: %v, output: %s", ErrPrintFailed, err, string(output))
	}

	return nil
}

// PrintToDefault 使用默认打印机打印
func (p *PDFPrinter) PrintToDefault(pdfPath string) error {
	return p.Print(pdfPath, PrintOptions{
		Silent: true,
	})
}

// GetPrinters 获取可用打印机列表
func (p *PDFPrinter) GetPrinters() ([]string, error) {
	return internal.GetPrinters()
}

// GetDefaultPrinter 获取默认打印机名称
func (p *PDFPrinter) GetDefaultPrinter() (string, error) {
	return internal.GetDefaultPrinter()
}

// Printer 打印器接口
type Printer interface {
	Print(pdfPath string, options PrintOptions) error
	PrintToDefault(pdfPath string) error
	GetPrinters() ([]string, error)
	GetDefaultPrinter() (string, error)
}

// 确保PDFPrinter实现了Printer接口
var _ Printer = (*PDFPrinter)(nil)

// Print 便捷函数：使用默认打印器打印PDF
func Print(pdfPath string, options PrintOptions) error {
	p := NewPrinter()
	return p.Print(pdfPath, options)
}

// PrintToDefault 便捷函数：使用默认打印机打印PDF
func PrintToDefault(pdfPath string) error {
	p := NewPrinter()
	return p.PrintToDefault(pdfPath)
}

// GetPrinters 便捷函数：获取可用打印机列表
func GetPrinters() ([]string, error) {
	p := NewPrinter()
	return p.GetPrinters()
}

// GetDefaultPrinter 便捷函数：获取默认打印机名称
func GetDefaultPrinter() (string, error) {
	p := NewPrinter()
	return p.GetDefaultPrinter()
}

// IsSumatraAvailable 检查SumatraPDF是否可用
func (p *PDFPrinter) IsSumatraAvailable() bool {
	if p.sumatraPath == "" {
		_, err := FindSumatraPDF()
		return err == nil
	}
	return internal.FileExists(p.sumatraPath)
}
