package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/thisXYH/mdnt/internal"
	"github.com/thisXYH/mdnt/internal/markdown"
)

var markdownOps = &markdown.Options{}
var markdownCmd = &cobra.Command{
	Use:   "md",
	Short: "维护 markdown 文档的引用",
	Long:  "维护 markdown 文档的引用，以及文档Id",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAndDealMarkdownOps(); err != nil {
			return err
		}

		// 执行 img 子命令
		if err := markdown.Execute(*markdownOps); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	markdownCmd.Flags().StringVarP(&markdownOps.MdDir, "markdown-dir", "m", "", "文档目录路径,可从环境变量"+mdnt_img_m_env+"读取")
	markdownCmd.Flags().BoolVarP(&markdownOps.DoIdSet, "id-set", "s", false, "给未设置id的文档添加id")
	markdownCmd.Flags().BoolVarP(&markdownOps.DoRelPathFix, "fix-ref", "f", false, "维护文档的相对路径引用")
	markdownCmd.Flags().BoolVarP(&markdownOps.NewId, "id-new", "n", false, "生成一个新id")
}

func checkAndDealMarkdownOps() error {
	// 两个需要目录的选项
	if markdownOps.DoIdSet || markdownOps.DoRelPathFix {
		if markdownOps.MdDir == "" {
			if markdownOps.MdDir = os.Getenv(mdnt_img_m_env); markdownOps.MdDir == "" {
				return fmt.Errorf("-m is empty")
			}
		}

		if !filepath.IsAbs(markdownOps.MdDir) {
			markdownOps.MdDir, _ = filepath.Abs(markdownOps.MdDir)
		}

		if !internal.IsFileOrDirExist(markdownOps.MdDir) {
			return fmt.Errorf("not found the markdown-dir path: %s", markdownOps.MdDir)
		}
	}

	return nil
}
