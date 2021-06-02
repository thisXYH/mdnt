package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/thisXYH/mdnt/internal/encrypto"
)

var encryptoOps *encrypto.Options = &encrypto.Options{}

var encyptoCmd = &cobra.Command{
	Use:   "enc",
	Short: "加解密指定笔记",
	Long:  "加解密指定笔记，防止敏感信息泄露(请牢记密码！！)",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkAndDealEncOps(); err != nil {
			return err
		}

		// 执行 enc 子命令
		if err := encrypto.Execute(*encryptoOps); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	encyptoCmd.Flags().StringVarP(&encryptoOps.FilePath, "file", "f", "", "文件路径")
	encyptoCmd.MarkFlagRequired("file")
	encyptoCmd.Flags().StringVarP(&encryptoOps.Password, "password", "p", "", "加解密密码")
	encyptoCmd.MarkFlagRequired("password")
}

func checkAndDealEncOps() error {
	if !filepath.IsAbs(encryptoOps.FilePath) {
		encryptoOps.FilePath, _ = filepath.Abs(encryptoOps.FilePath)
	}

	f, err := os.Stat(encryptoOps.FilePath)
	if err != nil || f.IsDir() {
		return fmt.Errorf("file not exist: %s", encryptoOps.FilePath)
	}

	return nil
}
