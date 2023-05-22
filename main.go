package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func usage() {
	fmt.Println(
		`pdfencrypter [the path for input file or directory] [optional: owner password]`,
	)
}

func main() {

	var help bool
	flag.BoolVar(&help, "h", false, "help")
	flag.Parse()
	if help {
		usage()
		return
	}
	args := os.Args[1:]

	inputPath := args[0]
	var password string
	if len(args) == 1 {
		password = mustReadTXTfile("pass.txt")
	} else {
		password = args[1]
	}

	// ファイルまたはフォルダが存在するかどうかを確認します。
	if _, err := os.Stat(inputPath); err != nil {
		// ファイルまたはフォルダが存在しない場合は、エラーを返します。
		fmt.Println(err)
		return
	}

	// ファイルまたはフォルダの種類を判別します。
	info, err := os.Stat(inputPath)
	if err != nil {
		// ファイルまたはフォルダの種類を判別できない場合は、エラーを返します。
		fmt.Println(err)
		return
	}

	// 第一引数がフォルダパスの場合
	if info.IsDir() {
		// フォルダ内の PDF ファイルの一覧を取得します。
		files, err := filepath.Glob(filepath.Join(inputPath, "*.pdf"))
		if err != nil {
			fmt.Println(err)
			return
		}

		// フォルダ内の PDF ファイルの一覧を出力します。
		for _, file := range files {
			fmt.Println(file)
			err = run(file, filepath.Base(file), password)
			if err != nil {
				log.Fatal(err)
			}
		}
		return
	}

	// 第一引数がファイルパスの場合
	err = run(inputPath, filepath.Base(inputPath), password)
	if err != nil {
		log.Fatal(err)
	}
}

func run(inputPath, outputPath, password string) error {
	// ファイルのロック
	// https://www.antenna.co.jp/ptl/cookbook/vol2/201805301211.html#t_2e201805301211_2eISO_2032000-1_20_E8_A1_A822_20_E3_83_A6_E3_83_BC_E3_82_B6_E3_83_BC_E3_82_A2_E3_82_AF_E3_82_BB_E3_82_B9_E8_A8_B1_E5_8F_AF_E2_80_95_E3_82_BF_E3_82_A4_E3_83_97_EF_BC_91
	conf := model.NewDefaultConfiguration()
	conf.OwnerPW = password
	conf.Permissions = model.PermissionsPrint
	fmt.Println(conf.String())
	fmt.Println("userPW:", conf.UserPW)
	fmt.Println("ownerPW:", conf.OwnerPW)
	fmt.Println()
	err := api.EncryptFile(inputPath, outputPath, conf)
	if err != nil {
		return err
	}

	return nil
}

func mustReadTXTfile(txtPath string) string {
	// ファイルを開きます。
	f, err := os.Open(txtPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// ファイルの内容を読み取ります。
	contents, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return string(contents)
}
