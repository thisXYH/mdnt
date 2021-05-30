package images

import "nt/internal"

// refStat 图片文件的引用统计
type refStat struct {
	path  string // path is full path for image
	count int    // count is refrenced times for image
}

// resolvedResult
// markdown 解析结果
type resolvedResult struct {
	path     string
	imageSet internal.Set // 引用的图片 set[RefImage]
}

type RefImageType int8

const (
	RelImage RefImageType = iota + 1
	AbsImage
	WebImage
)

type RefImage struct {
	Original     string       // ![comment](xxx/x.png)
	OriginalPath string       // xxx/x.png
	Name         string       // x.png
	Type         RefImageType //是否是网络图片
}
