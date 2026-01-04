package main

import (
	"fmt"
	"log"

	"github.com/zerroi/pdf-to-printer/printer"
)

func main() {
	// 示例1: 使用默认打印机打印PDF
	fmt.Println("示例1: 使用默认打印机打印PDF")
	err := printer.PrintToDefault("document.pdf")
	if err != nil {
		log.Printf("打印失败: %v\n", err)
	} else {
		fmt.Println("打印成功!")
	}

	// 示例2: 指定打印机打印PDF
	fmt.Println("\n示例2: 指定打印机打印PDF")
	options := printer.PrintOptions{
		PrinterName: "HP LaserJet",
		Copies:      2,
		Silent:      true,
	}
	err = printer.Print("document.pdf", options)
	if err != nil {
		log.Printf("打印失败: %v\n", err)
	} else {
		fmt.Println("打印成功!")
	}

	// 示例3: 获取可用打印机列表
	fmt.Println("\n示例3: 获取可用打印机列表")
	printers, err := printer.GetPrinters()
	if err != nil {
		log.Printf("获取打印机列表失败: %v\n", err)
	} else {
		fmt.Println("可用打印机:")
		for i, p := range printers {
			fmt.Printf("  %d. %s\n", i+1, p)
		}
	}

	// 示例4: 获取默认打印机
	fmt.Println("\n示例4: 获取默认打印机")
	defaultPrinter, err := printer.GetDefaultPrinter()
	if err != nil {
		log.Printf("获取默认打印机失败: %v\n", err)
	} else {
		fmt.Printf("默认打印机: %s\n", defaultPrinter)
	}

	// 示例5: 使用Printer接口
	fmt.Println("\n示例5: 使用Printer接口")
	p := printer.NewPrinter()
	err = p.PrintToDefault("document.pdf")
	if err != nil {
		log.Printf("打印失败: %v\n", err)
	} else {
		fmt.Println("打印成功!")
	}

	// 示例6: 检查SumatraPDF是否可用
	fmt.Println("\n示例6: 检查SumatraPDF是否可用")
	if p.IsSumatraAvailable() {
		fmt.Println("SumatraPDF可用")
	} else {
		fmt.Println("SumatraPDF未找到，请先安装")
	}

	// 示例7: 自定义SumatraPDF路径
	fmt.Println("\n示例7: 自定义SumatraPDF路径")
	err = printer.SetConfig(printer.Config{
		SumatraPath: "C:\\Program Files\\SumatraPDF\\SumatraPDF.exe",
	})
	if err != nil {
		log.Printf("设置配置失败: %v\n", err)
	} else {
		fmt.Println("配置设置成功!")
	}

	// 示例8: 使用高级打印选项
	fmt.Println("\n示例8: 使用高级打印选项")
	advancedOptions := printer.PrintOptions{
		PrinterName: "HP LaserJet",
		Copies:      1,
		Silent:      true,
		PageRange:   "1-3,5",      // 打印第1-3页和第5页
		Orientation: "portrait",   // 纵向
		Scaling:     "fit",        // 适应纸张大小
		ColorMode:   "color",      // 彩色打印
		Duplex:      "duplexlong", // 长边双面打印
		PaperSize:   "A4",         // A4纸张
	}
	err = printer.Print("document.pdf", advancedOptions)
	if err != nil {
		log.Printf("打印失败: %v\n", err)
	} else {
		fmt.Println("打印成功!")
	}

	// 示例9: 显示打印对话框
	fmt.Println("\n示例9: 显示打印对话框")
	dialogOptions := printer.PrintOptions{
		ShowDialog: true,
	}
	err = printer.Print("document.pdf", dialogOptions)
	if err != nil {
		log.Printf("打印失败: %v\n", err)
	} else {
		fmt.Println("打印成功!")
	}
}
