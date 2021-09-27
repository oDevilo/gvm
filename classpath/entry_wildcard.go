package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// 其实也是组合entry

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // remove *
	compositeEntry := []Entry{}
	// 遍历 baseDir 下所有的文件，并执行回调
	_ = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	})
	return compositeEntry
}
