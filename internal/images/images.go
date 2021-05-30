package images

import (
	"fmt"
	"io/fs"
	"nt/internal"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

type Options struct {
	ImgDir           string //ImgDir 图片目录绝对路径
	MdDir            string //MdDir 文档目录绝对路径
	DoImgDel         bool
	DoRelPathFix     bool
	DoWebImgDownload bool
}

var imagesOps Options

func Execute(p Options) error {
	imagesOps = p
	// future, 类似 c# 中的 task.run
	refStatCh := getRefStatMap(p.ImgDir)

	// 解析 markdown
	ch := markdownResolved(p.MdDir)

	// future, 类似 c# 中的 task.result
	refStatMap := <-refStatCh

	// 用自己实现的sync 限流同步
	ms := internal.NewMySync(5)
	for i := 0; i < 3; i++ {
		go resolvedResultHandler(refStatMap, ch, ms)
	}
	ms.Wait()

	ImageFileHandler(refStatMap)

	return nil
}

func ImageFileHandler(imageMap map[string]*refStat) {
	sb := strings.Builder{}
	sb.WriteString("无引用图片：\n")
	for k, stat := range imageMap {
		if stat.count > 0 {
			continue
		}

		if imagesOps.DoImgDel {
			os.Remove(stat.path)
			sb.WriteString(k + "\t Deleted \n")
		} else {
			sb.WriteString(k + "\n")
		}

	}
	fmt.Println(sb.String())
}

// markdownResolved 解析 markdown 的图片引用
func markdownResolved(mdDir string) <-chan resolvedResult {
	wg := &sync.WaitGroup{}
	ch := make(chan resolvedResult, 5)

	filepath.WalkDir(mdDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() || filepath.Ext(d.Name()) != ".md" {
			return nil
		}
		go markdownResolvedCore(path, ch, wg)
		return nil
	})

	// close chan
	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

// markdownResolvedCore 解析 markdown 的图片引用核心逻辑
func markdownResolvedCore(fullAddr string, ch chan resolvedResult, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	result := resolvedResult{path: fullAddr, imageSet: internal.NewMySet()}

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
		refImag := refImage{}
		refImag.original = v[0]
		refImag.originalPath = v[1]
		refImag.name = v[2]
		refImag.pt = getRefImageType(refImag.originalPath)

		result.imageSet.Add(refImag)
	}

	ch <- result
}

// 处理 解析结果
func resolvedResultHandler(refStatMap map[string]*refStat, ch <-chan resolvedResult, ms *internal.MySync) {
	ms.Add(1)
	defer ms.Done(1)

	for result := range ch {
		if result.imageSet.Len() <= 0 {
			continue
		}
		var show strings.Builder
		show.WriteString("=================================\n")
		show.WriteString(fmt.Sprintf("%s  count:%d\n", result.path, result.imageSet.Len()))

		// 两个需要修改的选项
		var content string
		if imagesOps.DoRelPathFix || imagesOps.DoWebImgDownload {
			buf, _ := os.ReadFile(result.path)
			content = string(buf)
		}

		for _, v := range result.imageSet.Iterator() {
			temp := v.(refImage)

			switch temp.pt {
			case absPath, relPath:
				refStatItem, ok := refStatMap[temp.name]
				if !ok {
					show.WriteString("没找到图片文件：" + temp.originalPath + "\n")
					continue
				}

				// 指针
				refStatItem.count++

				if imagesOps.DoRelPathFix {
					relPath, _ := filepath.Rel(filepath.Dir(result.path), refStatItem.path)
					relPath = filepath.ToSlash(relPath)
					content = strings.ReplaceAll(content, temp.originalPath, relPath)
					show.WriteString(temp.originalPath + "  =>  " + relPath + "\n")
				} else {
					show.WriteString(temp.originalPath + "\n")
				}
			case webPath:
				if imagesOps.DoWebImgDownload {
					//TODO:
				} else {
					show.WriteString(temp.originalPath + "\n")
				}
			}
		}

		fmt.Println(show.String())
		if content != "" {
			internal.WriteToFile(result.path, content)
		}
	}
}

// getRefStatMap 获取所有图片文件的map
// 使用 Future 模式
func getRefStatMap(imageDir string) <-chan map[string]*refStat {
	ch := make(chan map[string]*refStat, 1)

	go func() {
		// 遍历图片目录，图片名称要求唯一
		imageMap := make(map[string]*refStat)
		filepath.WalkDir(imageDir, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}

			switch ext := filepath.Ext(d.Name()); ext {
			case ".png", ".jpg":
				imageMap[d.Name()] = &refStat{path: path, count: 0}
			}

			return nil
		})

		ch <- imageMap
	}()

	return ch
}

func getRefImageType(path string) pathType {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return webPath
	}

	if filepath.IsAbs(path) {
		return absPath
	} else {
		return relPath
	}
}
