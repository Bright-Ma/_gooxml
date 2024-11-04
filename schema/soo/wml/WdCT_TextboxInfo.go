// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/clearmann/gooxml"
	"github.com/clearmann/gooxml/schema/soo/dml"
)

type WdCT_TextboxInfo struct {
	IdAttr      *uint16
	TxbxContent *WdCT_TxbxContent
	ExtLst      *dml.CT_OfficeArtExtensionList
}

func NewWdCT_TextboxInfo() *WdCT_TextboxInfo {
	ret := &WdCT_TextboxInfo{}
	ret.TxbxContent = NewWdCT_TxbxContent()
	return ret
}

func (m *WdCT_TextboxInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	setxbxContent := xml.StartElement{Name: xml.Name{Local: "wp:txbxContent"}}
	e.EncodeElement(m.TxbxContent, setxbxContent)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "wp:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_TextboxInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TxbxContent = NewWdCT_TxbxContent()
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 16)
			if err != nil {
				return err
			}
			pt := uint16(parsed)
			m.IdAttr = &pt
			continue
		}
	}
lWdCT_TextboxInfo:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "txbxContent"}:
				if err := d.DecodeElement(m.TxbxContent, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on WdCT_TextboxInfo %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_TextboxInfo
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_TextboxInfo and its children
func (m *WdCT_TextboxInfo) Validate() error {
	return m.ValidateWithPath("WdCT_TextboxInfo")
}

// ValidateWithPath validates the WdCT_TextboxInfo and its children, prefixing error messages with path
func (m *WdCT_TextboxInfo) ValidateWithPath(path string) error {
	if err := m.TxbxContent.ValidateWithPath(path + "/TxbxContent"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
