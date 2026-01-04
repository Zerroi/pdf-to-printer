package printer

import "errors"

// 自定义错误类型
var (
	// ErrSumatraNotFound SumatraPDF未找到
	ErrSumatraNotFound = errors.New("SumatraPDF not found")
	// ErrPrinterNotFound 打印机未找到
	ErrPrinterNotFound = errors.New("printer not found")
	// ErrFileNotFound PDF文件未找到
	ErrFileNotFound = errors.New("PDF file not found")
	// ErrInvalidPDF 无效的PDF文件
	ErrInvalidPDF = errors.New("invalid PDF file")
	// ErrPrintFailed 打印操作失败
	ErrPrintFailed = errors.New("print operation failed")
	// ErrInvalidOptions 无效的打印选项
	ErrInvalidOptions = errors.New("invalid print options")
)
