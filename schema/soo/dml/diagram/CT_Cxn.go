// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/clearmann/gooxml"
	"github.com/clearmann/gooxml/schema/soo/dml"
)

type CT_Cxn struct {
	ModelIdAttr    ST_ModelId
	TypeAttr       ST_CxnType
	SrcIdAttr      ST_ModelId
	DestIdAttr     ST_ModelId
	SrcOrdAttr     uint32
	DestOrdAttr    uint32
	ParTransIdAttr *ST_ModelId
	SibTransIdAttr *ST_ModelId
	PresIdAttr     *string
	ExtLst         *dml.CT_OfficeArtExtensionList
}

func NewCT_Cxn() *CT_Cxn {
	ret := &CT_Cxn{}
	return ret
}

func (m *CT_Cxn) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "modelId"},
		Value: fmt.Sprintf("%v", m.ModelIdAttr)})
	if m.TypeAttr != ST_CxnTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "srcId"},
		Value: fmt.Sprintf("%v", m.SrcIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "destId"},
		Value: fmt.Sprintf("%v", m.DestIdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "srcOrd"},
		Value: fmt.Sprintf("%v", m.SrcOrdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "destOrd"},
		Value: fmt.Sprintf("%v", m.DestOrdAttr)})
	if m.ParTransIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "parTransId"},
			Value: fmt.Sprintf("%v", *m.ParTransIdAttr)})
	}
	if m.SibTransIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sibTransId"},
			Value: fmt.Sprintf("%v", *m.SibTransIdAttr)})
	}
	if m.PresIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "presId"},
			Value: fmt.Sprintf("%v", *m.PresIdAttr)})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Cxn) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "modelId" {
			parsed, err := ParseUnionST_ModelId(attr.Value)
			if err != nil {
				return err
			}
			m.ModelIdAttr = parsed
			continue
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "srcId" {
			parsed, err := ParseUnionST_ModelId(attr.Value)
			if err != nil {
				return err
			}
			m.SrcIdAttr = parsed
			continue
		}
		if attr.Name.Local == "destId" {
			parsed, err := ParseUnionST_ModelId(attr.Value)
			if err != nil {
				return err
			}
			m.DestIdAttr = parsed
			continue
		}
		if attr.Name.Local == "srcOrd" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SrcOrdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "destOrd" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.DestOrdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "parTransId" {
			parsed, err := ParseUnionST_ModelId(attr.Value)
			if err != nil {
				return err
			}
			m.ParTransIdAttr = &parsed
			continue
		}
		if attr.Name.Local == "sibTransId" {
			parsed, err := ParseUnionST_ModelId(attr.Value)
			if err != nil {
				return err
			}
			m.SibTransIdAttr = &parsed
			continue
		}
		if attr.Name.Local == "presId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PresIdAttr = &parsed
			continue
		}
	}
lCT_Cxn:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_Cxn %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Cxn
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Cxn and its children
func (m *CT_Cxn) Validate() error {
	return m.ValidateWithPath("CT_Cxn")
}

// ValidateWithPath validates the CT_Cxn and its children, prefixing error messages with path
func (m *CT_Cxn) ValidateWithPath(path string) error {
	if err := m.ModelIdAttr.ValidateWithPath(path + "/ModelIdAttr"); err != nil {
		return err
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.SrcIdAttr.ValidateWithPath(path + "/SrcIdAttr"); err != nil {
		return err
	}
	if err := m.DestIdAttr.ValidateWithPath(path + "/DestIdAttr"); err != nil {
		return err
	}
	if m.ParTransIdAttr != nil {
		if err := m.ParTransIdAttr.ValidateWithPath(path + "/ParTransIdAttr"); err != nil {
			return err
		}
	}
	if m.SibTransIdAttr != nil {
		if err := m.SibTransIdAttr.ValidateWithPath(path + "/SibTransIdAttr"); err != nil {
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
