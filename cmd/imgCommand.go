package cmd

import (
	"flag"
	"fmt"
	"io/fs"
	"nt/cmd/types"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

const ImgCommand = "img"

var imgOptions *ImgOptions

// 路径使用绝对路径
type ImgOptions struct {
	ImgDir           string
	MdDir            string
	DoImgDel         bool
	DoRelPathFix     bool
	DoWebImgDownload bool
}

func EntryImgCommand(args []string) {
	imgOptions = getImgOptions(args)

	// 遍历图片目录，图片名称要求唯一
	imageMap := make(map[string]*types.ImageDetail)
	filepath.WalkDir(imgOptions.ImgDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		switch ext := filepath.Ext(d.Name()); ext {
		case ".png", ".jpg":
			imageMap[d.Name()] = &types.ImageDetail{FullName: path, RefCount: 0}
		}

		return nil
	})

	// 遍历文档目录
	ch := imgWalkMdDir(imgOptions.MdDir)

	// 用自己实现的sync同步
	ms := types.NewMySync()
	for i := 0; i < 3; i++ {
		go imgHandler(imageMap, ch, ms)
	}
	ms.Wait()

	ImageFileHandler(imageMap)
}

func ImageFileHandler(imageMap map[string]*types.ImageDetail) {
	sb := strings.Builder{}
	sb.WriteString("无引用图片：\n")
	for k, v := range imageMap {
		if v.RefCount > 0 {
			continue
		}

		if imgOptions.DoImgDel {
			os.Remove(v.FullName)
			sb.WriteString(k + "\t Deleted \n")
		} else {
			sb.WriteString(k + "\n")
		}

	}
	fmt.Println(sb.String())
}

// 遍历文档目录，使用sync结束ch
func imgWalkMdDir(mdDir string) <-chan types.ImgResult {
	wg := &sync.WaitGroup{}
	ch := make(chan types.ImgResult, 5)

	filepath.WalkDir(imgOptions.MdDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() || filepath.Ext(d.Name()) != ".md" {
			return nil
		}
		go imgResolved(path, ch, wg)
		return nil
	})

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

// 解析 markdown 图片引用
func imgResolved(fullAddr string, ch chan types.ImgResult, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	result := types.ImgResult{MdFullName: fullAddr, ImageSet: types.NewMySet()}

	buf, _ := os.ReadFile(fullAddr)
	content := string(buf)
	/*
		图片引用方式：
		1. 绝对路径
		2. 相对路径
		3. 网络路径
	*/
	re := regexp.MustCompile(`!\[.*\]\((.+(?:\\|\/)(.+(?:\.png|\.jpg)))\)`)
	for _, v := range re.FindAllStringSubmatch(content, -1) {
		refImag := types.RefImage{}
		refImag.Original = v[0]
		refImag.OriginalPath = v[1]
		refImag.Name = v[2]
		refImag.Type = getRefImageType(refImag.OriginalPath)

		result.ImageSet.Add(refImag)
	}

	ch <- result
}

// 处理 imgResolved 的解析结果
func imgHandler(imageMap map[string]*types.ImageDetail, ch <-chan types.ImgResult, ms *types.MySync) {
	ms.Add(1)
	defer ms.Done(1)

	for result := range ch {
		if result.ImageSet.Len() <= 0 {
			continue
		}
		var show strings.Builder
		show.WriteString("=================================\n")
		show.WriteString(fmt.Sprintf("%s  count:%d\n", result.MdFullName, result.ImageSet.Len()))

		// 两个需要修改的选项
		var content string
		if imgOptions.DoRelPathFix || imgOptions.DoWebImgDownload {
			buf, _ := os.ReadFile(result.MdFullName)
			content = string(buf)
		}

		for _, v := range result.ImageSet.Iterator() {
			temp := v.(types.RefImage)

			switch temp.Type {
			case types.AbsImage, types.RelImage:
				imageDetail, ok := imageMap[temp.Name]
				if !ok {
					show.WriteString("没找到图片文件：" + temp.OriginalPath + "\n")
					continue
				}

				// 指针
				imageDetail.RefCount++

				if imgOptions.DoRelPathFix {
					relPath, _ := filepath.Rel(filepath.Dir(result.MdFullName), imageDetail.FullName)
					relPath = filepath.ToSlash(relPath)
					content = strings.ReplaceAll(content, temp.OriginalPath, relPath)
					show.WriteString(temp.OriginalPath + "  =>  " + relPath + "\n")
				} else {
					show.WriteString(temp.OriginalPath + "\n")
				}
			case types.WebImage:
				if imgOptions.DoWebImgDownload {
					//TODO:
				} else {
					show.WriteString(temp.OriginalPath + "\n")
				}
			}
		}

		fmt.Println(show.String())
		if content != "" {
			writeToFile(result.MdFullName, content)
		}
	}
}

func getImgOptions(args []string) *ImgOptions {
	help := false
	imgOptions := &ImgOptions{}
	imgFlag := flag.NewFlagSet(ImgCommand, flag.ExitOnError)

	imgFlag.Usage = commandUsage(imgFlag)
	imgFlag.BoolVar(&help, "h", false, "显示帮助菜单")
	imgFlag.StringVar(&imgOptions.ImgDir, "i", "", "图片目录，不能为空")
	imgFlag.StringVar(&imgOptions.MdDir, "m", "", "文档目录，不能为空")
	imgFlag.BoolVar(&imgOptions.DoImgDel, "d", false, "删除无引用的图片，否则只打印路径")
	imgFlag.BoolVar(&imgOptions.DoRelPathFix, "f", false, "修复引用图片的相对路径，否则只打印路径")
	imgFlag.BoolVar(&imgOptions.DoWebImgDownload, "w", false, "下载引用的网络图片，否则只打印路径")

	imgFlag.Parse(args)

	if help {
		imgFlag.Usage()
		os.Exit(1)
	}

	if imgOptions.ImgDir == "" || imgOptions.MdDir == "" {
		printErrorAndExit(fmt.Errorf("-i && -m: 不能为空"), imgFlag.Usage)
	}

	if !filepath.IsAbs(imgOptions.ImgDir) {
		imgOptions.ImgDir, _ = filepath.Abs(imgOptions.ImgDir)
	}

	if !isFileOrDirExist(imgOptions.ImgDir) {
		printErrorAndExit(fmt.Errorf("图片目录路径不存在: "+imgOptions.ImgDir), imgFlag.Usage)
	}

	if !filepath.IsAbs(imgOptions.MdDir) {
		imgOptions.MdDir, _ = filepath.Abs(imgOptions.MdDir)
	}

	if !isFileOrDirExist(imgOptions.MdDir) {
		printErrorAndExit(fmt.Errorf("文档目录路径不存在:"+imgOptions.MdDir), imgFlag.Usage)
	}

	return imgOptions
}
