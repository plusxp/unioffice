// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"log"
)

type CT_Headers struct {
	// Header Cell Reference
	Header *CT_String
}

func NewCT_Headers() *CT_Headers {
	ret := &CT_Headers{}
	ret.Header = NewCT_String()
	return ret
}
func (m *CT_Headers) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	seheader := xml.StartElement{Name: xml.Name{Local: "w:header"}}
	e.EncodeElement(m.Header, seheader)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_Headers) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Header = NewCT_String()
lCT_Headers:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "header":
				if err := d.DecodeElement(m.Header, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Headers
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_Headers) Validate() error {
	return m.ValidateWithPath("CT_Headers")
}
func (m *CT_Headers) ValidateWithPath(path string) error {
	if err := m.Header.ValidateWithPath(path + "/Header"); err != nil {
		return err
	}
	return nil
}