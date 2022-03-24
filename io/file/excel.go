// Package file by chenaiquan<like.aiquan@qq.com> create on 2021/11/4 17.42
package file

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
)

type ORA struct {
	corp  string //单位
	name  string //业务系统
	name2 string //进程名
	v1000 string //10点值
	v2000 string //20点值
	H     string // 最大值
	L     string // 最小值
	A     string // 平均值
	TIME  string // 时间
}

func ReadXlsx(filename string) []ORA {
	var listOra []ORA
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Printf("open failed: %s\n", err)
	}
	for _, sheet := range xlFile.Sheets {
		//fmt.Printf("Sheet Name: %s\n", sheet.Name)
		tmpOra := ORA{}
		// 获取标签页(时间)
		//tmpOra.TIME = sheet.Name
		for _, row := range sheet.Rows {
			var strs []string
			for _, cell := range row.Cells {
				text := cell.String()
				strs = append(strs, text)
			}
			// 获取标签页(时间)
			tmpOra.TIME = sheet.Name
			tmpOra.corp = strs[0]
			tmpOra.name = strs[1]
			tmpOra.name2 = strs[2]
			tmpOra.v1000 = strs[3]
			tmpOra.v2000 = strs[4]
			tmpOra.H = strs[5]
			tmpOra.L = strs[6]
			tmpOra.A = strs[7]
			listOra = append(listOra, tmpOra)
		}
	}
	return listOra
}

func WritingXlsx(oraList []ORA) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	row.SetHeightCM(0.5)
	cell = row.AddCell()
	cell.Value = "单位"
	cell = row.AddCell()
	cell.Value = "业务系统"
	cell = row.AddCell()
	cell.Value = "进程名"
	cell = row.AddCell()
	cell.Value = "V1000"
	cell = row.AddCell()
	cell.Value = "V2000"
	cell = row.AddCell()
	cell.Value = "H"
	cell = row.AddCell()
	cell.Value = "L"
	cell = row.AddCell()
	cell.Value = "A"
	cell = row.AddCell()
	cell.Value = "TIME"

	for _, i := range oraList {
		if i.corp == "单位" {
			continue
		}

		// 判断是否为-9999，是的变为0.0
		var row1 *xlsx.Row
		if i.v1000 == "-9999" {
			i.v1000 = "0.0"
		}
		if i.v2000 == "-9999" {
			i.v2000 = "0.0"
		}
		if i.H == "-9999" {
			i.H = "0.0"
		}
		if i.L == "-9999" {
			i.L = "0.0"
		}

		row1 = sheet.AddRow()
		row1.SetHeightCM(0.5)

		cell = row1.AddCell()
		cell.Value = i.corp
		cell = row1.AddCell()
		cell.Value = i.name
		cell = row1.AddCell()
		cell.Value = i.name2

		// 判断值是大于7200，大于变成红色
		v1, _ := strconv.ParseFloat(i.v1000, 64)
		if v1 > 7200 {
			cell = row1.AddCell()
			cell.Value = i.v1000
			cell.GetStyle().Font.Color = "00FF0000"
		} else {
			cell = row1.AddCell()
			cell.Value = i.v1000
		}

		//v2, _ := strconv.Atoi(i.v2000)
		v2, _ := strconv.ParseFloat(i.v2000, 64)
		if v2 > 7200 {
			cell = row1.AddCell()
			cell.Value = i.v2000
			cell.GetStyle().Font.Color = "00FF0000"
		} else {
			cell = row1.AddCell()
			cell.Value = i.v2000
		}

		//vH, _ := strconv.Atoi(i.H)
		vH, _ := strconv.ParseFloat(i.H, 64)
		if vH > 7200 {
			cell = row1.AddCell()
			cell.Value = i.H
			cell.GetStyle().Font.Color = "00FF0000"
		} else {
			cell = row1.AddCell()
			cell.Value = i.H
		}

		//vL, _ := strconv.Atoi(i.L)
		vL, _ := strconv.ParseFloat(i.L, 64)
		if vL > 7200 {
			cell = row1.AddCell()
			cell.Value = i.L
			cell.GetStyle().Font.Color = "00FF0000"
		} else {
			cell = row1.AddCell()
			cell.Value = i.L

		}

		//vA, _ := strconv.Atoi(i.A)
		vA, _ := strconv.ParseFloat(i.A, 64)
		if vA > 7200 {
			cell = row1.AddCell()
			cell.Value = i.A
			cell.GetStyle().Font.Color = "00FF0000"
		} else {
			cell = row1.AddCell()
			cell.Value = i.A
		}

		// 打印时间
		cell = row1.AddCell()
		cell.Value = i.TIME
	}

	err = file.Save("after file name")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
