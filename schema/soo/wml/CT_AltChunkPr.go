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

	"github.com/clearmann/gooxml"
)

type CT_AltChunkPr struct {
	// Keep Source Formatting on Import
	MatchSrc *CT_OnOff
}

func NewCT_AltChunkPr() *CT_AltChunkPr {
	ret := &CT_AltChunkPr{}
	return ret
}

func (m *CT_AltChunkPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.MatchSrc != nil {
		sematchSrc := xml.StartElement{Name: xml.Name{Local: "w:matchSrc"}}
		e.EncodeElement(m.MatchSrc, sematchSrc)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AltChunkPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_AltChunkPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "matchSrc"}:
				m.MatchSrc = NewCT_OnOff()
				if err := d.DecodeElement(m.MatchSrc, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_AltChunkPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AltChunkPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AltChunkPr and its children
func (m *CT_AltChunkPr) Validate() error {
	return m.ValidateWithPath("CT_AltChunkPr")
}

// ValidateWithPath validates the CT_AltChunkPr and its children, prefixing error messages with path
func (m *CT_AltChunkPr) ValidateWithPath(path string) error {
	if m.MatchSrc != nil {
		if err := m.MatchSrc.ValidateWithPath(path + "/MatchSrc"); err != nil {
			return err
		}
	}
	return nil
}
