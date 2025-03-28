// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/clearmann/gooxml"
	"github.com/clearmann/gooxml/common"
	"github.com/clearmann/gooxml/measurement"
	"github.com/clearmann/gooxml/schema/soo/dml"
	pic "github.com/clearmann/gooxml/schema/soo/dml/picture"
	"github.com/clearmann/gooxml/schema/soo/ofc/sharedTypes"
	"github.com/clearmann/gooxml/schema/soo/wml"
)

// Run is a run of text within a paragraph that shares the same formatting.
type Run struct {
	d *Document
	x *wml.CT_R
}

// X returns the inner wrapped XML type.
func (r Run) X() *wml.CT_R {
	return r.x
}

// Text returns the underlying tet in the run.
func (r Run) Text() string {
	if len(r.x.EG_RunInnerContent) == 0 {
		return ""
	}
	buf := bytes.Buffer{}
	for _, ic := range r.x.EG_RunInnerContent {
		if ic.T != nil {
			buf.WriteString(ic.T.Content)
		}
		if ic.Tab != nil {
			buf.WriteByte('\t')
		}
	}
	return buf.String()
}

// ClearContent clears any content in the run (text, tabs, breaks, etc.)
func (r Run) ClearContent() {
	r.x.EG_RunInnerContent = nil
}

// AddLatexText 添加包含 Latex 公式的文本
func (r Run) AddLatexText(s string) {
	if strings.Contains(s, "$") {
		re := regexp.MustCompile(`\$(.*?)\$`)
		matches := re.FindAllStringSubmatchIndex(s, -1)

		lastEnd := 0
		for _, match := range matches {
			if match[0] > lastEnd {
				r.AddText(s[lastEnd:match[0]])
			}

			latex := s[match[2]:match[3]]
			if omml, err := LatexToOMML(latex); err == nil {
				ic := r.newIC()
				ic.T = wml.NewCT_Text()
				ic.T.Content = fmt.Sprintf(`<p xmlns:m="http://schemas.openxmlformats.org/officeDocument/2006/math">
					<m:oMathPara>
						<m:oMath>
							%s
						</m:oMath>
					</m:oMathPara>
				</p>`, omml)
				r.Properties().SetMath(true)
			}

			lastEnd = match[1]
		}

		if lastEnd < len(s) {
			r.AddText(s[lastEnd:])
		}
	} else {
		r.AddText(s)
	}

}

// AddText adds text to a run.
func (r Run) AddText(s string) {
	ic := wml.NewEG_RunInnerContent()
	r.x.EG_RunInnerContent = append(r.x.EG_RunInnerContent, ic)
	ic.T = wml.NewCT_Text()
	if gooxml.NeedsSpacePreserve(s) {
		p := "preserve"
		ic.T.SpaceAttr = &p
	}
	ic.T.Content = s
}

// 自带缩进, 遇到\n自动换行,设置字体(样式大小)
func AddIndentedMultilineText(doc *Document, font string, text string, size int) {
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if line != "" {
			_para := doc.AddParagraph()
			// 设置每个新段落的左缩进
			indent := measurement.Distance(size) * measurement.Point * 2
			_para.Properties().SetFirstLineIndent(indent)
			run := _para.AddRun()
			run.Properties().SetSize(measurement.Distance(size) * measurement.Point)
			run.Properties().SetFontFamily(font)
			run.AddText(line)
		}
	}
}

// AddMultiLineText adds multi line text to a run.
func (r Run) AddMultiLineText(s string) {
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		r.AddText(line)
		r.AddBreak()
	}
}

func (r Run) newIC() *wml.EG_RunInnerContent {
	ic := wml.NewEG_RunInnerContent()
	r.x.EG_RunInnerContent = append(r.x.EG_RunInnerContent, ic)
	return ic
}

// Clear removes all of the content from within a run.
func (r Run) Clear() {
	r.x.EG_RunInnerContent = nil
}

// AddTab adds tab to a run and can be used with the the Paragraph's tab stops.
func (r Run) AddTab() {
	ic := r.newIC()
	ic.Tab = wml.NewCT_Empty()
}

// AddFieldWithFormatting adds a field (automatically computed text) to the
// document with field specifc formatting.
func (r Run) AddFieldWithFormatting(code string, fmt string, isDirty bool) {
	ic := r.newIC()
	ic.FldChar = wml.NewCT_FldChar()
	ic.FldChar.FldCharTypeAttr = wml.ST_FldCharTypeBegin
	if isDirty {
		ic.FldChar.DirtyAttr = &sharedTypes.ST_OnOff{}
		ic.FldChar.DirtyAttr.Bool = gooxml.Bool(true)
	}

	ic = r.newIC()
	ic.InstrText = wml.NewCT_Text()
	if fmt != "" {
		ic.InstrText.Content = code + " " + fmt
	} else {
		ic.InstrText.Content = code
	}

	ic = r.newIC()
	ic.FldChar = wml.NewCT_FldChar()
	ic.FldChar.FldCharTypeAttr = wml.ST_FldCharTypeEnd
}

// AddField adds a field (automatically computed text) to the document.
func (r Run) AddField(code string) {
	r.AddFieldWithFormatting(code, "", true)
}

// SetMath 设置运行块为数学公式模式
func (r RunProperties) SetMath(hasMath bool) {
	if hasMath {
		if r.x.OMath == nil {
			r.x.OMath = wml.NewCT_OnOff()
		}
		r.x.OMath.ValAttr = &sharedTypes.ST_OnOff{Bool: gooxml.Bool(true)}
	} else {
		r.x.OMath = nil
	}
}

// Properties returns the run properties.
func (r Run) Properties() RunProperties {
	if r.x.RPr == nil {
		r.x.RPr = wml.NewCT_RPr()
	}
	return RunProperties{r.x.RPr}
}

// AddBreak adds a line break to a run.
func (r Run) AddBreak() {
	ic := r.newIC()
	ic.Br = wml.NewCT_Br()
}

// AddPageBreak adds a page break to a run.
func (r Run) AddPageBreak() {
	ic := r.newIC()
	ic.Br = wml.NewCT_Br()
	ic.Br.TypeAttr = wml.ST_BrTypePage
}

// DrawingAnchored returns a slice of AnchoredDrawings.
func (r Run) DrawingAnchored() []AnchoredDrawing {
	ret := []AnchoredDrawing{}
	for _, ic := range r.x.EG_RunInnerContent {
		if ic.Drawing == nil {
			continue
		}
		for _, anc := range ic.Drawing.Anchor {
			ret = append(ret, AnchoredDrawing{r.d, anc})
		}
	}
	return ret
}

// AddDrawingAnchored adds an anchored (floating) drawing from an ImageRef.
func (r Run) AddDrawingAnchored(img common.ImageRef) (AnchoredDrawing, error) {
	ic := r.newIC()
	ic.Drawing = wml.NewCT_Drawing()
	anchor := wml.NewWdAnchor()

	ad := AnchoredDrawing{r.d, anchor}

	// required by Word on OSX for the file to open
	anchor.SimplePosAttr = gooxml.Bool(false)

	anchor.AllowOverlapAttr = true
	anchor.CNvGraphicFramePr = dml.NewCT_NonVisualGraphicFrameProperties()

	ic.Drawing.Anchor = append(ic.Drawing.Anchor, anchor)
	anchor.Graphic = dml.NewGraphic()
	anchor.Graphic.GraphicData = dml.NewCT_GraphicalObjectData()
	anchor.Graphic.GraphicData.UriAttr = "http://schemas.openxmlformats.org/drawingml/2006/picture"
	anchor.SimplePos.XAttr.ST_CoordinateUnqualified = gooxml.Int64(0)
	anchor.SimplePos.YAttr.ST_CoordinateUnqualified = gooxml.Int64(0)
	anchor.PositionH.RelativeFromAttr = wml.WdST_RelFromHPage
	anchor.PositionH.Choice = &wml.WdCT_PosHChoice{}
	anchor.PositionH.Choice.PosOffset = gooxml.Int32(0)

	anchor.PositionV.RelativeFromAttr = wml.WdST_RelFromVPage
	anchor.PositionV.Choice = &wml.WdCT_PosVChoice{}
	anchor.PositionV.Choice.PosOffset = gooxml.Int32(0)

	anchor.Extent.CxAttr = int64(float64(img.Size().X*measurement.Pixel72) / measurement.EMU)
	anchor.Extent.CyAttr = int64(float64(img.Size().Y*measurement.Pixel72) / measurement.EMU)
	anchor.Choice = &wml.WdEG_WrapTypeChoice{}
	anchor.Choice.WrapSquare = wml.NewWdCT_WrapSquare()
	anchor.Choice.WrapSquare.WrapTextAttr = wml.WdST_WrapTextBothSides

	// Mac Word chokes if the ID is greater than an int32, even though the field is a
	// uint32 in the XSD
	randID := 0x7FFFFFFF & rand.Uint32()
	anchor.DocPr.IdAttr = randID
	p := pic.NewPic()
	p.NvPicPr.CNvPr.IdAttr = randID

	// find the reference to the actual image file in the document relationships
	// so we can embed via the relationship ID
	imgID := img.RelID()
	if imgID == "" {
		return ad, errors.New("couldn't find reference to image within document relations")
	}

	anchor.Graphic.GraphicData.Any = append(anchor.Graphic.GraphicData.Any, p)
	p.BlipFill = dml.NewCT_BlipFillProperties()
	p.BlipFill.Blip = dml.NewCT_Blip()
	p.BlipFill.Blip.EmbedAttr = &imgID
	p.BlipFill.Stretch = dml.NewCT_StretchInfoProperties()
	p.BlipFill.Stretch.FillRect = dml.NewCT_RelativeRect()

	p.SpPr = dml.NewCT_ShapeProperties()
	// Required to allow resizing
	p.SpPr.Xfrm = dml.NewCT_Transform2D()
	p.SpPr.Xfrm.Off = dml.NewCT_Point2D()
	p.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = gooxml.Int64(0)
	p.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = gooxml.Int64(0)
	p.SpPr.Xfrm.Ext = dml.NewCT_PositiveSize2D()
	p.SpPr.Xfrm.Ext.CxAttr = int64(img.Size().X * measurement.Point)
	p.SpPr.Xfrm.Ext.CyAttr = int64(img.Size().Y * measurement.Point)
	// required by Word on OSX for the image to display
	p.SpPr.PrstGeom = dml.NewCT_PresetGeometry2D()
	p.SpPr.PrstGeom.PrstAttr = dml.ST_ShapeTypeRect

	return ad, nil
}

// AddDrawingInline adds an inline drawing from an ImageRef.
func (r Run) AddDrawingInline(img common.ImageRef) (InlineDrawing, error) {
	ic := r.newIC()
	ic.Drawing = wml.NewCT_Drawing()

	inl := wml.NewWdInline()
	inline := InlineDrawing{r.d, inl}

	// required by Word on OSX for the file to open
	//anchor.SimplePosAttr = gooxml.Bool(false)

	//anchor.AllowOverlapAttr = true
	inl.CNvGraphicFramePr = dml.NewCT_NonVisualGraphicFrameProperties()

	ic.Drawing.Inline = append(ic.Drawing.Inline, inl)
	inl.Graphic = dml.NewGraphic()
	inl.Graphic.GraphicData = dml.NewCT_GraphicalObjectData()
	inl.Graphic.GraphicData.UriAttr = "http://schemas.openxmlformats.org/drawingml/2006/picture"

	inl.DistTAttr = gooxml.Uint32(0)
	inl.DistLAttr = gooxml.Uint32(0)
	inl.DistBAttr = gooxml.Uint32(0)
	inl.DistRAttr = gooxml.Uint32(0)

	inl.Extent.CxAttr = int64(float64(img.Size().X*measurement.Pixel72) / measurement.EMU)
	inl.Extent.CyAttr = int64(float64(img.Size().Y*measurement.Pixel72) / measurement.EMU)

	// Mac Word chokes if the ID is greater than an int32, even though the field is a
	// uint32 in the XSD
	randID := 0x7FFFFFFF & rand.Uint32()
	inl.DocPr.IdAttr = randID
	p := pic.NewPic()
	p.NvPicPr.CNvPr.IdAttr = randID

	// find the reference to the actual image file in the document relationships
	// so we can embed via the relationship ID
	imgID := img.RelID()
	if imgID == "" {
		return inline, errors.New("couldn't find reference to image within document relations")
	}

	inl.Graphic.GraphicData.Any = append(inl.Graphic.GraphicData.Any, p)
	p.BlipFill = dml.NewCT_BlipFillProperties()
	p.BlipFill.Blip = dml.NewCT_Blip()
	p.BlipFill.Blip.EmbedAttr = &imgID
	p.BlipFill.Stretch = dml.NewCT_StretchInfoProperties()
	p.BlipFill.Stretch.FillRect = dml.NewCT_RelativeRect()

	p.SpPr = dml.NewCT_ShapeProperties()
	// Required to allow resizing
	p.SpPr.Xfrm = dml.NewCT_Transform2D()
	p.SpPr.Xfrm.Off = dml.NewCT_Point2D()
	p.SpPr.Xfrm.Off.XAttr.ST_CoordinateUnqualified = gooxml.Int64(0)
	p.SpPr.Xfrm.Off.YAttr.ST_CoordinateUnqualified = gooxml.Int64(0)
	p.SpPr.Xfrm.Ext = dml.NewCT_PositiveSize2D()
	p.SpPr.Xfrm.Ext.CxAttr = int64(img.Size().X * measurement.Point)
	p.SpPr.Xfrm.Ext.CyAttr = int64(img.Size().Y * measurement.Point)
	// required by Word on OSX for the image to display
	p.SpPr.PrstGeom = dml.NewCT_PresetGeometry2D()
	p.SpPr.PrstGeom.PrstAttr = dml.ST_ShapeTypeRect

	return inline, nil
}
