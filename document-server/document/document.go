package document

import (
	"errors"
	"fmt"
	"strconv"

	"math"

	"github.com/document/document-server/constant"
	"github.com/document/document-server/tools"
	"github.com/tealeg/xlsx"
	"github.com/wise-registry/src/common/utils/log"
)

type Document struct {
	PassiveFile      string
	InitiativeFile   string //基准文件
	PassiveColumn    int    //用于比较的列 从1开始
	InitiativeColumn int
}

var default_compare_column = 1

func NewDocument(passiveFile, initiativeFile string) *Document {
	return &Document{PassiveFile: passiveFile, InitiativeFile: initiativeFile, PassiveColumn: default_compare_column, InitiativeColumn: default_compare_column}
}

func NewDocumentWithColumn(passiveFile, initiativeFile string, passiveColumn, initiativeColumn int) *Document {
	return &Document{PassiveFile: passiveFile, InitiativeFile: initiativeFile, PassiveColumn: passiveColumn, InitiativeColumn: initiativeColumn}
}

func (d *Document) SetPassiveColumn(column int) *Document {
	d.PassiveColumn = column
	return d
}
func (d *Document) SetInitiativeColumn(column int) *Document {
	d.InitiativeColumn = column
	return d
}

func (d *Document) Compare() (err error) {
	passiveXls, err := xlsx.OpenFile(d.PassiveFile)
	if err != nil {
		return
	}

	initiativeXls, err := xlsx.OpenFile(d.InitiativeFile)
	if err != nil {
		return
	}

	if err = d.Validate(passiveXls, initiativeXls); err != nil {
		return
	}

	passiveSheet := passiveXls.Sheets[0]
	initiativeSheet := initiativeXls.Sheets[0]
	positiveRows, reverseRows := []*xlsx.Row{}, []*xlsx.Row{}
	fileName := ""
	if len(passiveSheet.Rows) > len(initiativeSheet.Rows) {
		positiveRows = ComparePositiveRow(passiveSheet, initiativeSheet, d.PassiveColumn-1, d.InitiativeColumn-1)
		reverseRows = CompareReverseRow(passiveSheet, initiativeSheet, d.PassiveColumn-1, d.InitiativeColumn-1)
		fileName = fmt.Sprintf(constant.COMPARED_FILE,
			tools.RemoveSuffix(d.PassiveFile, []tools.Remove{{Split: "/"}, {Split: ".", Reverse: true}}),
			tools.RemoveSuffix(d.InitiativeFile, []tools.Remove{{Split: "/"}, {Split: ".", Reverse: true}}),
		)

	} else {
		positiveRows = ComparePositiveRow(initiativeSheet, passiveSheet, d.InitiativeColumn-1, d.PassiveColumn-1)
		reverseRows = CompareReverseRow(initiativeSheet, passiveSheet, d.InitiativeColumn-1, d.PassiveColumn-1)
		fileName = fmt.Sprintf(constant.COMPARED_FILE,
			tools.RemoveSuffix(d.InitiativeFile, []tools.Remove{{Split: "/"}, {Split: ".", Reverse: true}}),
			tools.RemoveSuffix(d.PassiveFile, []tools.Remove{{Split: "/"}, {Split: ".", Reverse: true}}),
		)
	}

	err = GenerateXLSX(constant.UPLOAD_DIR+fileName, positiveRows, reverseRows)
	return
}

func GenerateXLSX(fileName string, positiveRows, reverseRows []*xlsx.Row) (err error) {
	xlsxFile := xlsx.NewFile()

	FillSheet(xlsxFile, positiveRows, "sheet1")
	if len(reverseRows) > 0 {
		FillSheet(xlsxFile, reverseRows, "sheet2")
	}
	err = xlsxFile.Save(fileName)
	return
}

func FillSheet(xlsxFile *xlsx.File, rows []*xlsx.Row, sheetName string) (err error) {
	sheet, err := xlsxFile.AddSheet(sheetName)
	if err != nil {
		log.Fatalf("生成xlsx Sheet %s 失败", sheetName)
		return
	}

	for _, row := range rows {
		r := sheet.AddRow()
		for _, cell := range row.Cells {
			c := r.AddCell()
			if v, err := strconv.Atoi(cell.Value); err == nil && v < math.MaxInt32 {
				c.SetValue(v)
			} else {
				c.SetValue(cell.Value)
			}
		}
	}
	return
}

func ComparePositiveRow(initiativeSheet, passiveSheet *xlsx.Sheet, initiativeColumn, passiveColumn int) (rows []*xlsx.Row) {
	rows = []*xlsx.Row{}
	passiveMap := map[string]*xlsx.Row{}
	for _, row := range passiveSheet.Rows {
		if _, ok := passiveMap[row.Cells[passiveColumn].Value]; !ok {
			passiveMap[row.Cells[passiveColumn].Value] = row
		}
	}

	for _, row := range initiativeSheet.Rows {
		if _, ok := passiveMap[row.Cells[initiativeColumn].Value]; !ok {
			rows = append(rows, row)
		}
	}

	return
}

func CompareReverseRow(initiativeSheet, passiveSheet *xlsx.Sheet, initiativeColumn, passiveColumn int) (rows []*xlsx.Row) {
	rows = []*xlsx.Row{}
	initiativeMap := map[string]*xlsx.Row{}

	for _, row := range initiativeSheet.Rows {
		if _, ok := initiativeMap[row.Cells[initiativeColumn].Value]; !ok {
			initiativeMap[row.Cells[initiativeColumn].Value] = row
		}
	}

	for _, row := range passiveSheet.Rows {
		if _, ok := initiativeMap[row.Cells[passiveColumn].Value]; !ok {
			rows = append(rows, row)
		}
	}
	return
}

func (d *Document) Validate(passXls, initiative *xlsx.File) error {
	if passXls == nil || initiative == nil {
		return errors.New("文件验证失败, 请确认文件格式正确")
	}

	if len(passXls.Sheets) == 0 || len(initiative.Sheets) == 0 {
		return errors.New("未获取到文件中可比较的Sheet")
	}

	if len(passXls.Sheets[0].Cols) < d.PassiveColumn || len(initiative.Sheets[0].Cols) < d.InitiativeColumn {
		return errors.New("用于比较的列号大于文档列,请确保比较的列号小于文档总列数")
	}
	return nil
}
