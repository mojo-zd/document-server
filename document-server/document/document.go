package document

import (
	"errors"
	"fmt"

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
	rows := []*xlsx.Row{}
	fileName := ""
	if len(passiveSheet.Rows) > len(initiativeSheet.Rows) {
		rows = CompareRow(passiveSheet, initiativeSheet, d.PassiveColumn-1, d.InitiativeColumn-1)
		fileName = fmt.Sprintf(constant.COMPARED_FILE,
			tools.RemoveSuffix(d.PassiveFile, []tools.Remove{{Split: "/"}, {Split: ".", Reverse: true}}),
			tools.RemoveSuffix(d.InitiativeFile, []tools.Remove{{Split: "/"}, {Split: ".", Reverse: true}}),
		)
	} else {
		rows = CompareRow(initiativeSheet, passiveSheet, d.InitiativeColumn-1, d.PassiveColumn-1)
	}

	err = GenerateXLSX(constant.UPLOAD_DIR+fileName, rows)
	return
}

func GenerateXLSX(fileName string, rows []*xlsx.Row) (err error) {
	xlsxFile := xlsx.NewFile()
	sheet, err := xlsxFile.AddSheet("sheet1")
	if err != nil {
		log.Fatal("生成xlsx文件失败")
		return
	}

	for _, row := range rows {
		row.Sheet = sheet
	}
	err = xlsxFile.Save(fileName)
	return
}

func CompareRow(initiativeSheet, passiveSheet *xlsx.Sheet, initiativeColumn, passiveColumn int) (rows []*xlsx.Row) {
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
