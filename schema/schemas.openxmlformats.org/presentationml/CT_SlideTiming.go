// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"
)

type CT_SlideTiming struct {
	TnLst *CT_TimeNodeList
	// Build List
	BldLst *CT_BuildList
	ExtLst *CT_ExtensionListModify
}

func NewCT_SlideTiming() *CT_SlideTiming {
	ret := &CT_SlideTiming{}
	return ret
}
func (m *CT_SlideTiming) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.TnLst != nil {
		setnLst := xml.StartElement{Name: xml.Name{Local: "p:tnLst"}}
		e.EncodeElement(m.TnLst, setnLst)
	}
	if m.BldLst != nil {
		sebldLst := xml.StartElement{Name: xml.Name{Local: "p:bldLst"}}
		e.EncodeElement(m.BldLst, sebldLst)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_SlideTiming) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SlideTiming:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "tnLst":
				m.TnLst = NewCT_TimeNodeList()
				if err := d.DecodeElement(m.TnLst, &el); err != nil {
					return err
				}
			case "bldLst":
				m.BldLst = NewCT_BuildList()
				if err := d.DecodeElement(m.BldLst, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionListModify()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideTiming
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_SlideTiming) Validate() error {
	return m.ValidateWithPath("CT_SlideTiming")
}
func (m *CT_SlideTiming) ValidateWithPath(path string) error {
	if m.TnLst != nil {
		if err := m.TnLst.ValidateWithPath(path + "/TnLst"); err != nil {
			return err
		}
	}
	if m.BldLst != nil {
		if err := m.BldLst.ValidateWithPath(path + "/BldLst"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}