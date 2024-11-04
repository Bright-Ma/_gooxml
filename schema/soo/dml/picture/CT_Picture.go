// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package picture

import (
	"encoding/xml"

	"github.com/clearmann/gooxml"
	"github.com/clearmann/gooxml/schema/soo/dml"
)

type CT_Picture struct {
	NvPicPr  *CT_PictureNonVisual
	BlipFill *dml.CT_BlipFillProperties
	SpPr     *dml.CT_ShapeProperties
}

func NewCT_Picture() *CT_Picture {
	ret := &CT_Picture{}
	ret.NvPicPr = NewCT_PictureNonVisual()
	ret.BlipFill = dml.NewCT_BlipFillProperties()
	ret.SpPr = dml.NewCT_ShapeProperties()
	return ret
}

func (m *CT_Picture) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	senvPicPr := xml.StartElement{Name: xml.Name{Local: "pic:nvPicPr"}}
	e.EncodeElement(m.NvPicPr, senvPicPr)
	seblipFill := xml.StartElement{Name: xml.Name{Local: "pic:blipFill"}}
	e.EncodeElement(m.BlipFill, seblipFill)
	sespPr := xml.StartElement{Name: xml.Name{Local: "pic:spPr"}}
	e.EncodeElement(m.SpPr, sespPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Picture) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvPicPr = NewCT_PictureNonVisual()
	m.BlipFill = dml.NewCT_BlipFillProperties()
	m.SpPr = dml.NewCT_ShapeProperties()
lCT_Picture:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/picture", Local: "nvPicPr"}:
				if err := d.DecodeElement(m.NvPicPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/picture", Local: "blipFill"}:
				if err := d.DecodeElement(m.BlipFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/picture", Local: "spPr"}:
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Picture %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Picture
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Picture and its children
func (m *CT_Picture) Validate() error {
	return m.ValidateWithPath("CT_Picture")
}

// ValidateWithPath validates the CT_Picture and its children, prefixing error messages with path
func (m *CT_Picture) ValidateWithPath(path string) error {
	if err := m.NvPicPr.ValidateWithPath(path + "/NvPicPr"); err != nil {
		return err
	}
	if err := m.BlipFill.ValidateWithPath(path + "/BlipFill"); err != nil {
		return err
	}
	if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
		return err
	}
	return nil
}
