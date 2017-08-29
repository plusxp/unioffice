// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.
package spreadsheet_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"testing"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/spreadsheetml"

	"baliance.com/gooxml/spreadsheet"
	"baliance.com/gooxml/testhelper"
	"baliance.com/gooxml/zippkg"
)

func TestConstructor(t *testing.T) {
	wb := spreadsheet.New()
	if wb == nil {
		t.Errorf("expected a non-nil workbook")
	}
	if wb.X() == nil {
		t.Errorf("expected a non-nil workbook XML entity")
	}
	if wb.ContentTypes.X() == nil {
		t.Errorf("expected a non-nil content-types")
	}
	if wb.CoreProperties.X() == nil {
		t.Errorf("expected a non-nil core properties")
	}
	if wb.AppProperties.X() == nil {
		t.Errorf("expected a non-nil app properties")
	}
	if wb.Rels.X() == nil {
		t.Errorf("expected a non-nil relationship")
	}
	if wb.StyleSheet.X() == nil {
		t.Errorf("expected a non-nil stylesheet")
	}
}
func TestWorkbookUnmarshal(t *testing.T) {
	f, err := os.Open("testdata/workbook.xml")
	if err != nil {
		t.Fatalf("error reading content types file")
	}
	dec := xml.NewDecoder(f)
	r := spreadsheetml.NewWorkbook()
	if err := dec.Decode(r); err != nil {
		t.Errorf("error decoding content types: %s", err)
	}
	got := &bytes.Buffer{}
	fmt.Fprintf(got, zippkg.XMLHeader)
	enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
	if err := enc.Encode(r); err != nil {
		t.Errorf("error encoding content types: %s", err)
	}

	testhelper.CompareGoldenXML(t, "workbook.xml", got.Bytes())
}

func TestSimpleSheet(t *testing.T) {
	wb := spreadsheet.New()
	sheet := wb.AddSheet("foo")
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.SetString("testing 123")
	got := bytes.Buffer{}
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid spreadsheet: %s", err)
	}
	wb.Save(&got)
	testhelper.CompareGoldenZip(t, "simple-1.xlsx", got.Bytes())
}
func TestOpen(t *testing.T) {
	wb, err := spreadsheet.Open("testdata/simple-1.xlsx")
	if err != nil {
		t.Errorf("error opening workbook: %s", err)
	}

	got := bytes.Buffer{}
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid spreadsheet: %s", err)
	}
	wb.Save(&got)
	testhelper.CompareZip(t, "simple-1.xlsx", got.Bytes())
}

func TestOpenExcel2016(t *testing.T) {
	wb, err := spreadsheet.Open("../testdata/Office2016/Excel-Windows.xlsx")
	if err != nil {
		t.Fatalf("error opening workbook: %s", err)
	}

	got := bytes.Buffer{}
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid spreadsheet: %s", err)
	}
	wb.Save(&got)
	//testhelper.CompareZip(t, "../../testdata/Office2016/Excel-Windows.xlsx", got.Bytes())
}

func TestSheetCount(t *testing.T) {
	wb := spreadsheet.New()
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid spreadsheet: %s", err)
	}
	if wb.SheetCount() != 0 {
		t.Errorf("expected 0 sheets, got %d", wb.SheetCount())
	}
	wb.AddSheet("Sheet 1")
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid spreadsheet: %s", err)
	}
	if wb.SheetCount() != 1 {
		t.Errorf("expected 1 sheets, got %d", wb.SheetCount())
	}
	wb.AddSheet("Sheet 2")
	if err := wb.Validate(); err != nil {
		t.Errorf("created an invalid spreadsheet: %s", err)
	}
	if wb.SheetCount() != 2 {
		t.Errorf("expected 2 sheets, got %d", wb.SheetCount())
	}
}

func TestPreserveSpace(t *testing.T) {
	ss := spreadsheet.New()
	sheet := ss.AddSheet("Sheet 1")
	row := sheet.AddRow()
	values := []string{"  foo  ", " bar \t", "foo\r\nbar", "\t\r\nfoo\t123\r\n"}
	for i, s := range values {
		cell := row.AddCell()
		if i%2 == 0 {
			cell.SetInlineString(s)
		} else {
			cell.SetString(s)
		}

	}

	buf := bytes.Buffer{}
	if err := ss.Save(&buf); err != nil {
		log.Fatalf("error saving spreadsheet: %s", err)
	}

	op, err := spreadsheet.Read(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		log.Fatalf("error loading saved spreadsheet: %s", err)
	}
	sheets := op.Sheets()
	if len(sheets) != 1 {
		t.Errorf("expected 1 sheet, got %d", len(sheets))
	}
	rows := sheets[0].Rows()
	if len(rows) != 1 {
		t.Errorf("expected 1 row, got %d", len(rows))
	}
	cells := rows[0].Cells()
	if len(cells) != 4 {
		t.Errorf("expected 4 cells, got %d", len(cells))
	}
	for i, c := range cells {
		got, err := c.GetValue()
		if err != nil {
			t.Errorf("error retrieving cell value: %s", err)
		}
		if got != values[i] {
			t.Errorf("expected cell value = %q, got %q", values[i], got)
		}
	}
}