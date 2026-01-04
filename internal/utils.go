package internal

import (
	"os"
	"path/filepath"
	"strings"
)

// FileExists 检查文件是否存在
func FileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// IsPDFFile 检查文件是否为PDF文件
func IsPDFFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".pdf"
}

// FindExecutable 在PATH中查找可执行文件
func FindExecutable(name string) string {
	// 检查是否已经是完整路径
	if FileExists(name) {
		return name
	}

	// 在PATH中查找
	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		return ""
	}

	paths := filepath.SplitList(pathEnv)
	for _, dir := range paths {
		fullPath := filepath.Join(dir, name)
		if FileExists(fullPath) {
			return fullPath
		}
		// 尝试添加.exe扩展名（Windows）
		if !strings.HasSuffix(strings.ToLower(name), ".exe") {
			fullPath = filepath.Join(dir, name+".exe")
			if FileExists(fullPath) {
				return fullPath
			}
		}
	}

	return ""
}
