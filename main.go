//go:build windows
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/devlights/goxcel"
)

func runPict() ([]byte, error) {
	var (
		cmd *exec.Cmd
		out []byte
		err error
	)

	cmd = exec.Command("./pict.exe", "./model.txt")

	out, err = cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("cmd.Output() failed: %w", err)
	}

	return out, nil
}

func run() error {
	//
	// pict.exeを起動して結果を取得
	//
	var (
		cmdOut []byte
		err    error
	)

	cmdOut, err = runPict()
	if err != nil {
		return err
	}

	//
	// Excelを起動
	//
	var (
		quit    func()
		excel   *goxcel.Goxcel
		release goxcel.ReleaseFunc
		visible = false
	)

	quit = goxcel.MustInitGoxcel()
	defer quit()

	excel, release = goxcel.MustNewGoxcel()
	defer release()

	excel.MustSilent(visible)

	//
	// 新規ブックを開く
	//
	var (
		wbs       *goxcel.Workbooks
		wb        *goxcel.Workbook
		wbRelease goxcel.ReleaseFunc
		ws        *goxcel.Worksheet
	)

	wbs = excel.MustWorkbooks()
	wb, wbRelease = wbs.MustAdd()
	defer wbRelease()

	ws = wb.MustSheets(1)

	//
	// pictの結果をExcelに書き込み
	//
	var (
		buf     = bytes.NewBuffer(cmdOut)
		scanner = bufio.NewScanner(buf)
		row     = 1
	)

	for scanner.Scan() {
		tsv := scanner.Text()
		parts := strings.Split(tsv, "\t")

		for col := range len(parts) {
			cell := ws.MustCells(row, (col + 1))
			cell.MustSetValue(parts[col])
		}

		row++
	}

	//
	// ファイル出力
	//
	var (
		cwd     string
		outPath string
	)

	cwd, err = os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd() failed: %w", err)
	}

	outPath = filepath.Join(cwd, "result.xlsx")

	err = wb.SaveAs(outPath)
	if err != nil {
		return fmt.Errorf("wb.SaveAs() failed: %w", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
