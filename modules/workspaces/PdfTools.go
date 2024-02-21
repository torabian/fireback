package workspaces

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/alexeyco/simpletable"
	"github.com/schollz/progressbar/v3"
	"github.com/signintech/gopdf"
	middleeasttools "pixelplux.com/fireback/modules/workspaces/thirdparty/arabictools"
	fonts "pixelplux.com/fireback/modules/workspaces/ttf"
)

type Color struct {
	R uint8
	G uint8
	B uint8
}

type BaseStyle struct {
	BackgroundColor Color
}

type TableStyle struct {
	Th BaseStyle
	Td BaseStyle
}

func WriteBodyRow(pdf *gopdf.GoPdf, row []string, catalog *PdfWriterMeta, y *float64, style TableStyle) {
	WriteRowIntoPdf(pdf, row, catalog, y, style, "th")
}
func WriteHeadRow(pdf *gopdf.GoPdf, row []string, catalog *PdfWriterMeta, y *float64, style TableStyle) {
	WriteRowIntoPdf(pdf, row, catalog, y, style, "td")
}

func WriteRowIntoPdf(pdf *gopdf.GoPdf, row []string, catalog *PdfWriterMeta, y *float64, style TableStyle, area string) {
	pdf.SetLineWidth(0.1)
	h := float64(30)

	var cell BaseStyle
	if area == "th" {
		cell = style.Th
	}
	if area == "td" {
		cell = style.Td
	}

	x := float64(catalog.GutterLeft)
	for i := 0; i < len(catalog.CellWidth); i++ {

		pdf.SetFillColor(cell.BackgroundColor.R, cell.BackgroundColor.G, cell.BackgroundColor.B)
		pdf.RectFromUpperLeftWithStyle(x, *y, catalog.CellWidth[i], h, "FD")

		// x = float64(float64(i)*w) + float64(paddingLeft)
		// pdf.RectFromUpperLeftWithStyle(x, y, w, h, "FD")
		pdf.SetXY(x+catalog.CellPadding, *y+catalog.CellPadding)
		pdf.SetFontSize(6)
		txt := row[i]
		texts, _ := pdf.SplitText(txt, catalog.CellWidth[i])
		for _, text := range texts {
			// pdf.Text(text)
			pdf.SetX(x + catalog.CellPadding)
			pdf.SetFillColor(0, 0, 0)
			pdf.CellWithOption(&gopdf.Rect{W: catalog.CellWidth[i], H: 10}, text, gopdf.CellOption{Align: gopdf.Center})
			pdf.SetY(pdf.GetY() + 10)
		}

		x += catalog.CellWidth[i]

		// pdf.Text(fmt.Sprint(row) + " > " + fmt.Sprint(i) + "ajhsdj ksahdj ahskhakfhaskjfasf")
	}

	// We right the new height there, because this function can decide about the height of
	// the row only.
	*y += h

}

type PdfWriterMeta struct {
	GutterLeft  float64
	GutterRight float64
	CellPadding float64
	Header      string
	SubHeader   string
	Padding     float64
	CellCount   int64
	CellWidth   []float64
}

func BeginPdfExport(pdf *gopdf.GoPdf, catalog *PdfWriterMeta) {
	pdf.SetX(10)
	pdf.SetY(30)
	pdf.Text(catalog.Header)

	pdf.SetX(10)
	pdf.SetY(50)
	pdf.Text(catalog.SubHeader)

	color := &Color{120, 120, 120}
	pdf.SetLineWidth(0.1)
	pdf.SetFillColor(color.R, color.G, color.B) //setup fill color

	pdf.SetY(80)
}

func HandlePagination() {

}

var DefaultTableStyle TableStyle = TableStyle{
	Th: BaseStyle{
		BackgroundColor: Color{R: 210, G: 210, B: 210},
	},
	Td: BaseStyle{
		BackgroundColor: Color{R: 230, G: 230, B: 230},
	},
}

func BeginPdf() (gopdf.Config, *gopdf.GoPdf, error) {

	pdf := gopdf.GoPdf{}
	cfg := gopdf.Config{PageSize: *gopdf.PageSizeA4Landscape}
	pdf.Start(cfg)

	pdf.AddPage()

	fontData, _ := fonts.FontsFs.ReadFile("IRANSansWeb.ttf")
	err := pdf.AddTTFFontData("wts11", fontData)
	if err != nil {
		return cfg, nil, err
	}
	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		return cfg, nil, err
	}

	return cfg, &pdf, nil
}

func CreatePdfCatalog[T any](refl reflect.Value, data *PdfExportData) (gopdf.Config, *gopdf.GoPdf, *PdfWriterMeta) {
	gutter := 15.0
	cfg, pdf, _ := BeginPdf()
	cells := GetColumnsFromReflect[T](refl)
	cellWidth := []float64{}

	w := float64((cfg.PageSize.W - gutter) / float64(len(cells)))
	for _, cell := range cells {
		_ = cell
		cellWidth = append(cellWidth, w)
	}

	pdfCatalog := &PdfWriterMeta{
		CellPadding: 2,
		Header:      middleeasttools.Shape(data.Name),
		SubHeader:   middleeasttools.Shape(data.Description),
		GutterLeft:  gutter / 2,
		CellWidth:   cellWidth,
	}

	BeginPdfExport(pdf, pdfCatalog)

	y := pdf.GetY()

	for index, cell := range cells {

		if v, okay := data.FieldsMap[cell]; okay {
			cells[index] = middleeasttools.Shape(v)
		} else {
			cells[index] = middleeasttools.Shape(cell)
		}

	}

	WriteHeadRow(pdf, cells, pdfCatalog, &y, DefaultTableStyle)

	return cfg, pdf, pdfCatalog
}

func OperateOnRecord[T any](y *float64, cfg gopdf.Config, pdfCatalog *PdfWriterMeta, row *T, refl reflect.Value, pdf *gopdf.GoPdf) {
	data := ExtractRowStringValues[T](row, refl, false)
	WriteBodyRow(pdf, data, pdfCatalog, y, DefaultTableStyle)

	if *y > cfg.PageSize.H {
		pdf.AddPage()
		*y = float64(pdfCatalog.Padding)
		pdf.SetLineWidth(0.1)
		color := &Color{120, 120, 120}
		pdf.SetFillColor(color.R, color.G, color.B) //setup fill color

		pdf.SetY(0)
	}
}

func DatabaseScanner[T any](
	query QueryDSL,
	fn func(query QueryDSL) ([]*T, *QueryResultMeta, error),
) chan *T {

	stream := make(chan *T)

	readSize := int64(10)

	_, count, _ := fn(query)

	var index int64 = 0

	go func() {

		for ; index <= count.TotalItems; index += readSize {

			query.ItemsPerPage = int(readSize)
			query.StartIndex = int(index)
			items, _, _ := fn(query)

			if len(items) > 0 {
				for _, row := range items {
					stream <- row
				}
			}
		}
		close(stream)

	}()

	return stream
}

func DatabaseScannerNativeQuery[T any](
	q string,
	qc string,
	query QueryDSL,
) chan *T {

	stream := make(chan *T)

	readSize := int64(10)

	_, count, _ := UnsafeQuerySql[T](q, qc, query)

	var index int64 = 0

	go func() {

		for ; index <= count.TotalItems; index += readSize {

			query.ItemsPerPage = int(readSize)
			query.StartIndex = int(index)
			items, _, _ := UnsafeQuerySql[T](q, qc, query)

			if len(items) > 0 {
				for _, row := range items {
					stream <- row
				}
			}
		}
		close(stream)

	}()

	return stream
}

type PdfExportData struct {
	Name        string
	Description string
	FieldsMap   map[string]string
}

func PdfExporter[T any](
	path string,
	query QueryDSL,
	fn func(query QueryDSL) ([]*T, *QueryResultMeta, error),
	refl reflect.Value,
	bar *progressbar.ProgressBar,
	data *PdfExportData,
) *IError {

	cfg, pdf, pdfCatalog := CreatePdfCatalog[T](
		refl,
		data,
	)

	y := pdf.GetY()

	stream := DatabaseScanner(query, fn)
	for {
		row, more := <-stream

		if !more {
			break
		}
		OperateOnRecord[T](&y, cfg, pdfCatalog, row, refl, pdf)
		bar.Add(1)

	}

	pdf.WritePdf(path)
	return nil

}

func PdfExporterNativeQuery[T any](
	path string,
	query QueryDSL,
	report *Report,
	refl reflect.Value,
) *IError {

	data := &PdfExportData{
		Name:        report.Title,
		Description: report.Description,
	}

	cfg, pdf, pdfCatalog := CreatePdfCatalog[T](refl, data)
	y := pdf.GetY()

	stream := DatabaseScannerNativeQuery[T](report.Query, report.QueryCounter, query)
	for {
		row, more := <-stream
		if !more {
			break
		}
		OperateOnRecord[T](&y, cfg, pdfCatalog, row, refl, pdf)
	}

	pdf.WritePdf(path)
	return nil

}

func StructsToTerminalTable[T any](items []*T, v reflect.Value) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
		},
	}

	for _, n := range GetColumnsFromReflect[T](v) {
		table.Header.Cells = append(table.Header.Cells,
			&simpletable.Cell{Align: simpletable.AlignLeft, Text: n},
		)
	}

	for counter, device := range items {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", counter)},
		}

		for _, cellValue := range ExtractRowStringValues[T](device, v, false) {
			r = append(r, &simpletable.Cell{
				Align: simpletable.AlignRight, Text: cellValue,
			})
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}
	table.Println()
}

func GetReportById(id string, reports []Report) *Report {
	var report *Report

	for _, r := range reports {

		if r.UniqueId == id {
			report = &r
			break
		}
	}

	return report
}

func GetReport(reports []Report) *Report {

	items, _ := GetAppReportsString(reports)
	id := AskForSelect("Select report:", items)

	if id == "" {
		return nil
	}
	index := strings.Index(id, ">>>")
	if index <= 0 {
		return nil
	}
	id = strings.Trim(id[0:index], " ")
	return GetReportById(id, reports)

}
