package cmd

import "github.com/spf13/cobra"

const VERSION = "v1.3.2"

var rootCmd = &cobra.Command{
	Use:   "mdnt [command] [flags]",
	Short: "维护 Markdown 笔记的工具集",
	Long: `mdnt - Markdown Notebook Tool
提供了一个维护 Markdown 笔记的工具集。
方便日常Markdown的维护，比如移动了笔记的位置，
导致引用的图片路径对不上的问题。

项目地址：https://github.com/thisXYH/mdnt`,
	Version: VERSION,
}

func init() {
	rootCmd.AddCommand(imagesCmd)
	rootCmd.AddCommand(encyptoCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
