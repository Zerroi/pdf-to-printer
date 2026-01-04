package internal

import (
	"bytes"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"
)

// Windows API 相关定义
var (
	modwin32              = syscall.NewLazyDLL("winspool.drv")
	procEnumPrinters      = modwin32.NewProc("EnumPrintersA")
	procGetDefaultPrinter = modwin32.NewProc("GetDefaultPrinterA")
)

const (
	PRINTER_ENUM_LOCAL = 0x00000002
	PRINTER_ENUM_NAME  = 0x00000008
)

// GetPrinters 获取系统打印机列表
func GetPrinters() ([]string, error) {
	// 使用wmic命令获取打印机列表
	cmd := exec.Command("wmic", "printer", "get", "name")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	// 解析输出
	lines := strings.Split(out.String(), "\n")
	var printers []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && line != "Name" {
			printers = append(printers, line)
		}
	}

	return printers, nil
}

// GetDefaultPrinter 获取默认打印机名称
func GetDefaultPrinter() (string, error) {
	var buf [256]byte
	var size uint32 = 256

	// 使用Windows API获取默认打印机
	ret, _, err := procGetDefaultPrinter.Call(
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(&size)),
	)
	if ret == 0 {
		return "", err
	}

	return strings.TrimSpace(string(buf[:size])), nil
}

// PrinterExists 检查打印机是否存在
func PrinterExists(printerName string) bool {
	printers, err := GetPrinters()
	if err != nil {
		return false
	}

	for _, p := range printers {
		if strings.EqualFold(p, printerName) {
			return true
		}
	}
	return false
}
