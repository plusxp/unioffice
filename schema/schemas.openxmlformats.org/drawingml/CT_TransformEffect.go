// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_TransformEffect struct {
	SxAttr *ST_Percentage
	SyAttr *ST_Percentage
	KxAttr *int32
	KyAttr *int32
	TxAttr *ST_Coordinate
	TyAttr *ST_Coordinate
}

func NewCT_TransformEffect() *CT_TransformEffect {
	ret := &CT_TransformEffect{}
	return ret
}
func (m *CT_TransformEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.SxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sx"},
			Value: fmt.Sprintf("%v", *m.SxAttr)})
	}
	if m.SyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sy"},
			Value: fmt.Sprintf("%v", *m.SyAttr)})
	}
	if m.KxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "kx"},
			Value: fmt.Sprintf("%v", *m.KxAttr)})
	}
	if m.KyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ky"},
			Value: fmt.Sprintf("%v", *m.KyAttr)})
	}
	if m.TxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tx"},
			Value: fmt.Sprintf("%v", *m.TxAttr)})
	}
	if m.TyAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ty"},
			Value: fmt.Sprintf("%v", *m.TyAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TransformEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "sx" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.SxAttr = &parsed
		}
		if attr.Name.Local == "sy" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.SyAttr = &parsed
		}
		if attr.Name.Local == "kx" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.KxAttr = &pt
		}
		if attr.Name.Local == "ky" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.KyAttr = &pt
		}
		if attr.Name.Local == "tx" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.TxAttr = &parsed
		}
		if attr.Name.Local == "ty" {
			parsed, err := ParseUnionST_Coordinate(attr.Value)
			if err != nil {
				return err
			}
			m.TyAttr = &parsed
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TransformEffect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_TransformEffect) Validate() error {
	return m.ValidateWithPath("CT_TransformEffect")
}
func (m *CT_TransformEffect) ValidateWithPath(path string) error {
	if m.SxAttr != nil {
		if err := m.SxAttr.ValidateWithPath(path + "/SxAttr"); err != nil {
			return err
		}
	}
	if m.SyAttr != nil {
		if err := m.SyAttr.ValidateWithPath(path + "/SyAttr"); err != nil {
			return err
		}
	}
	if m.KxAttr != nil {
		if *m.KxAttr <= -5400000 {
			return fmt.Errorf("%s/m.KxAttr must be > -5400000 (have %v)", path, *m.KxAttr)
		}
		if *m.KxAttr >= 5400000 {
			return fmt.Errorf("%s/m.KxAttr must be < 5400000 (have %v)", path, *m.KxAttr)
		}
	}
	if m.KyAttr != nil {
		if *m.KyAttr <= -5400000 {
			return fmt.Errorf("%s/m.KyAttr must be > -5400000 (have %v)", path, *m.KyAttr)
		}
		if *m.KyAttr >= 5400000 {
			return fmt.Errorf("%s/m.KyAttr must be < 5400000 (have %v)", path, *m.KyAttr)
		}
	}
	if m.TxAttr != nil {
		if err := m.TxAttr.ValidateWithPath(path + "/TxAttr"); err != nil {
			return err
		}
	}
	if m.TyAttr != nil {
		if err := m.TyAttr.ValidateWithPath(path + "/TyAttr"); err != nil {
			return err
		}
	}
	return nil
}