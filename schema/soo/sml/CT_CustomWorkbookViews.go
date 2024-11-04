// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/clearmann/gooxml"
)

type CT_CustomWorkbookViews struct {
	// Custom Workbook View
	CustomWorkbookView []*CT_CustomWorkbookView
}

func NewCT_CustomWorkbookViews() *CT_CustomWorkbookViews {
	ret := &CT_CustomWorkbookViews{}
	return ret
}

func (m *CT_CustomWorkbookViews) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secustomWorkbookView := xml.StartElement{Name: xml.Name{Local: "ma:customWorkbookView"}}
	for _, c := range m.CustomWorkbookView {
		e.EncodeElement(c, secustomWorkbookView)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CustomWorkbookViews) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CustomWorkbookViews:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "customWorkbookView"}:
				tmp := NewCT_CustomWorkbookView()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CustomWorkbookView = append(m.CustomWorkbookView, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_CustomWorkbookViews %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CustomWorkbookViews
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CustomWorkbookViews and its children
func (m *CT_CustomWorkbookViews) Validate() error {
	return m.ValidateWithPath("CT_CustomWorkbookViews")
}

// ValidateWithPath validates the CT_CustomWorkbookViews and its children, prefixing error messages with path
func (m *CT_CustomWorkbookViews) ValidateWithPath(path string) error {
	for i, v := range m.CustomWorkbookView {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CustomWorkbookView[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
