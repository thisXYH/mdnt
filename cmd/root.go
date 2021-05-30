package cmd

import "github.com/spf13/cobra"

const VERSION = "v1.1.0"

var rootCmd = &cobra.Command{
	Use:   "nt [command] [flags]",
	Short: "维护 Markdown 笔记的工具集",
	Long: `nt 提供了一个维护 Markdown 笔记的工具集。
方便日常Markdown的维护，比如移动了笔记的位置，
导致引用的图片路径对不上的问题。

项目地址：https://github.com/thisXYH/NoteTools`,
	Version: VERSION,
}

func init() {
	rootCmd.AddCommand(imagesCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
