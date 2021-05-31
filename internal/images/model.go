package images

import "github.com/thisXYH/NoteTools/nt/internal"

// refStat 图片文件的引用统计
type refStat struct {
	path  string // path is full path for image
	count int    // count is refrenced times for image
}

// resolvedResult
// markdown 解析结果
type resolvedResult struct {
	path     string
	imageSet internal.Seter // 引用的图片 set[RefImage]
}

type refImage struct {
	original     string   // ![comment](xxx/x.png)
	originalPath string   // xxx/x.png
	name         string   // x.png
	pt           pathType // 引用的路径类型
}

type pathType int8

const (
	relPath pathType = iota + 1
	absPath
	webPath
)
