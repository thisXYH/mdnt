package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thisXYH/mdnt/internal/markdown"
)

var markdownOps = &markdown.Options{}
var markdownCmd = &cobra.Command{
	Use:   "md",
	Short: "维护 markdown 笔记的引用",
	Long:  "维护 markdown 笔记的引用，以及笔记Id",
	RunE: func(cmd *cobra.Command, args []string) error {
		// if err := checkAndDealImagesOps(); err != nil {
		// 	return err
		// }

		// 执行 img 子命令
		if err := markdown.Execute(*markdownOps); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	markdownCmd.Flags().StringVarP(&markdownOps.MdDir, "markdown-dir", "m", "", "文档目录路径,可从环境变量"+mdnt_img_m_env+"读取")
	markdownCmd.Flags().BoolVarP(&markdownOps.DoIdSet, "id-set", "s", false, "给未设置id的笔记添加id")
	markdownCmd.Flags().BoolVarP(&markdownOps.DoRelPathFix, "fix-ref", "f", false, "维护笔记的相对路径引用")
	markdownCmd.Flags().BoolVarP(&markdownOps.NewId, "id-new", "n", false, "生成一个新id")
}
