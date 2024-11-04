// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"

	"github.com/clearmann/gooxml"
)

type CT_AccPr struct {
	Chr    *CT_Char
	CtrlPr *CT_CtrlPr
}

func NewCT_AccPr() *CT_AccPr {
	ret := &CT_AccPr{}
	return ret
}

func (m *CT_AccPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Chr != nil {
		sechr := xml.StartElement{Name: xml.Name{Local: "m:chr"}}
		e.EncodeElement(m.Chr, sechr)
	}
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AccPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_AccPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "chr"}:
				m.Chr = NewCT_Char()
				if err := d.DecodeElement(m.Chr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "ctrlPr"}:
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_AccPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AccPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AccPr and its children
func (m *CT_AccPr) Validate() error {
	return m.ValidateWithPath("CT_AccPr")
}

// ValidateWithPath validates the CT_AccPr and its children, prefixing error messages with path
func (m *CT_AccPr) ValidateWithPath(path string) error {
	if m.Chr != nil {
		if err := m.Chr.ValidateWithPath(path + "/Chr"); err != nil {
			return err
		}
	}
	if m.CtrlPr != nil {
		if err := m.CtrlPr.ValidateWithPath(path + "/CtrlPr"); err != nil {
			return err
		}
	}
	return nil
}
