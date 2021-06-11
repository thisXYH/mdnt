package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/thisXYH/mdnt/internal"
	"github.com/thisXYH/mdnt/internal/images"
)

const (
	// 环境变量key
	mdnt_img_i_env = "mdnt_img_i"
	mdnt_img_m_env = "mdnt_img_m"
)

var imagesOps *images.Options = &images.Options{}

var imagesCmd = &cobra.Command{
	Use:   "img",
	Short: "维护 markdown 文档的图片引用",
	Long:  "维护 markdown 文档的图片引用，以及对应的图片文件",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAndDealImagesOps(); err != nil {
			return err
		}

		// 执行 img 子命令
		if err := images.Execute(*imagesOps); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	imagesCmd.Flags().StringVarP(&imagesOps.ImgDir, "image-dir", "i", "", "图片目录路径,可从环境变量"+mdnt_img_i_env+"读取")
	//imagesCmd.MarkFlagRequired("image-dir")
	imagesCmd.Flags().StringVarP(&imagesOps.MdDir, "markdown-dir", "m", "", "文档目录路径,可从环境变量"+mdnt_img_m_env+"读取")
	//imagesCmd.MarkFlagRequired("markdown-dir")
	imagesCmd.Flags().BoolVarP(&imagesOps.DoImgDel, "delete-unref", "d", false, "删除没有引用的图片文件")
	imagesCmd.Flags().BoolVarP(&imagesOps.DoRelPathFix, "fix-ref", "f", false, "修复图片的相对路径引用")
	imagesCmd.Flags().BoolVarP(&imagesOps.DoWebImgDownload, "down-web", "w", false, "删除没有引用的图片文件")
	imagesCmd.Flags().BoolVarP(&imagesOps.Verbose, "verbose", "v", false, "详细日志")
}

// checkAndDealImagesOps 检查和润色输入值
func checkAndDealImagesOps() error {
	if imagesOps.ImgDir == "" {
		if imagesOps.ImgDir = os.Getenv(mdnt_img_i_env); imagesOps.ImgDir == "" {
			return fmt.Errorf("-i is empty")
		}
	}

	if !filepath.IsAbs(imagesOps.ImgDir) {
		imagesOps.ImgDir, _ = filepath.Abs(imagesOps.ImgDir)
	}

	if !internal.IsFileOrDirExist(imagesOps.ImgDir) {
		return fmt.Errorf("not found the image-dir path: %s", imagesOps.ImgDir)
	}

	if imagesOps.MdDir == "" {
		if imagesOps.MdDir = os.Getenv(mdnt_img_m_env); imagesOps.MdDir == "" {
			return fmt.Errorf("-m is empty")
		}
	}

	if !filepath.IsAbs(imagesOps.MdDir) {
		imagesOps.MdDir, _ = filepath.Abs(imagesOps.MdDir)
	}

	if !internal.IsFileOrDirExist(imagesOps.MdDir) {
		return fmt.Errorf("not found the markdown-dir path: %s", imagesOps.MdDir)
	}

	return nil
}
