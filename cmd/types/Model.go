package types

type ImgResult struct {
	MdFullName string
	ImageSet   Set
}

type RefImageType int

const (
	Unknow RefImageType = iota
	RelImage
	WebImage
	AbsImage
)

type RefImage struct {
	Original     string       // ![comment](xxx/x.png)
	OriginalPath string       // xxx/x.png
	Name         string       // x.png
	Type         RefImageType //是否是网络图片
}

type ImageDetail struct {
	FullName string
	RefCount int
}
