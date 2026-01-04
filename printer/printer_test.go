package printer

import (
	"errors"
	"testing"
)

func TestConfig(t *testing.T) {
	// 保存原始配置
	_ = GetConfig()
	defer resetConfig()

	// 测试设置配置
	t.Run("SetConfig", func(t *testing.T) {
		// 注意：SetConfig会验证文件是否存在，所以这里测试设置空路径
		err := SetConfig(Config{
			SumatraPath: "",
		})
		if err != nil {
			t.Errorf("SetConfig failed: %v", err)
		}

		config := GetConfig()
		if config.SumatraPath != "" {
			t.Errorf("SumatraPath = %q, want empty string", config.SumatraPath)
		}
	})

	// 测试重置配置
	t.Run("ResetConfig", func(t *testing.T) {
		resetConfig()
		config := GetConfig()
		if config.SumatraPath != "" {
			t.Errorf("SumatraPath = %q, want empty string", config.SumatraPath)
		}
	})
}

func TestPrintOptionsValidation(t *testing.T) {
	t.Run("DefaultOptions", func(t *testing.T) {
		options := PrintOptions{}
		// 默认Copies是0，但在Print函数中会自动修正为1
		// 这里只验证默认值是0
		if options.Copies != 0 {
			t.Errorf("Default copies = %d, want 0", options.Copies)
		}
	})
}

func TestErrorTypes(t *testing.T) {
	tests := []struct {
		name  string
		err   error
		check func(error) bool
	}{
		{"ErrSumatraNotFound", ErrSumatraNotFound, func(e error) bool { return errors.Is(e, ErrSumatraNotFound) }},
		{"ErrPrinterNotFound", ErrPrinterNotFound, func(e error) bool { return errors.Is(e, ErrPrinterNotFound) }},
		{"ErrFileNotFound", ErrFileNotFound, func(e error) bool { return errors.Is(e, ErrFileNotFound) }},
		{"ErrInvalidPDF", ErrInvalidPDF, func(e error) bool { return errors.Is(e, ErrInvalidPDF) }},
		{"ErrPrintFailed", ErrPrintFailed, func(e error) bool { return errors.Is(e, ErrPrintFailed) }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.check(tt.err) {
				t.Errorf("Error type mismatch")
			}
		})
	}
}

func TestNewPrinter(t *testing.T) {
	p := NewPrinter()
	if p == nil {
		t.Fatal("NewPrinter returned nil")
	}

	// 检查是否实现了Printer接口
	var _ Printer = p
}

func TestConvenienceFunctions(t *testing.T) {
	// 这些函数在没有实际SumatraPDF和打印机的情况下无法完全测试
	// 但我们可以检查它们不会panic
	t.Run("GetPrinters", func(t *testing.T) {
		printers, err := GetPrinters()
		// 在Windows上应该能获取打印机列表，在其他平台上可能失败
		if err != nil {
			t.Logf("GetPrinters failed (expected on non-Windows): %v", err)
		} else {
			t.Logf("Found %d printers", len(printers))
		}
	})

	t.Run("GetDefaultPrinter", func(t *testing.T) {
		printer, err := GetDefaultPrinter()
		if err != nil {
			t.Logf("GetDefaultPrinter failed (expected on non-Windows): %v", err)
		} else {
			t.Logf("Default printer: %s", printer)
		}
	})
}

func TestPrinterInterface(t *testing.T) {
	p := NewPrinter()

	// 检查所有接口方法是否存在
	t.Run("InterfaceMethods", func(t *testing.T) {
		_ = p.Print
		_ = p.PrintToDefault
		_ = p.GetPrinters
		_ = p.GetDefaultPrinter
	})
}
