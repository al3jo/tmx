package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	tmx := unmarshall()
	sort.Sort(tmx.Body)
	markup := marshall(&tmx)
	fmt.Print(markup)
}

func test() {
	s := "   En la Unión Europea (UE) para el manejo de pacientes con fibrosis"
	t := strings.TrimSpace(string(re.ReplaceAll([]byte(s), void)))
	fmt.Printf("%s\n%s\n", s, t)
}

func unmarshall() (tmx Tmx) {
	xmlFile, err := os.Open("test/tm.tmx")
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		return
	}
	defer xmlFile.Close()

	data, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(data, &tmx)
	return tmx
}

func print(t Tmx) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d)%s\n", i+1, t.Body.Tus[i].Tuvs[0])
	}
}

func marshall(t *Tmx) (markup string) {
	start := `<?xml version="1.0" encoding="utf-8"?>`
	d, err := xml.Marshal(t)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return ""
	}
	return fmt.Sprintf("%s\n%s", start, string(d))
}

// Implements the TMX schema https://www.gala-global.org/tmx-14b
type Tmx struct {
	Header  Header
	Body    Body
	XMLName xml.Name `xml:"tmx"`
}

func (t Tmx) String() string {
	return fmt.Sprintf("%s [%s %s]", t.XMLName, t.Header, t.Body)
}

type Header struct {
	AdminLang           string     `xml:"adminlang,attr"`
	ChangeDate          *string    `xml:"changedate,attr"`
	ChangeId            *string    `xml:"changeid,attr"`
	CreationDate        *string    `xml:"creationdate,attr"`
	CreationId          *string    `xml:"creationid,attr"`
	CreationTool        string     `xml:"creationtool,attr"`
	CreationToolVersion string     `xml:"creationtoolversion,attr"`
	DataType            string     `xml:"datatype,attr"`
	OEncoding           *string    `xml:"o-encoding,attr"`
	OTmf                string     `xml:"o-tmf,attr"`
	Props               []Property `xml:"prop"`
	SegType             string     `xml:"segtype,attr"`
	SrcLang             string     `xml:"srclang,attr"`
	XMLName             xml.Name   `xml:"header"`
}

func (h Header) String() string {
	return fmt.Sprintf("%s [Props: %d]", h.XMLName, len(h.Props))
}

type Property struct {
	XMLName xml.Name `xml:"prop"`
	Name    string   `xml:"type,attr"`
	Prop    string   `xml:",chardata"`
}

func (p Property) String() string {
	return fmt.Sprintf("%s [%s=%s]", p.XMLName, p.Name, p.Prop)
}

type Body struct {
	Tus     []TranslationUnit `xml:"tu"`
	XMLName xml.Name          `xml:"body"`
}

var void = []byte("")
var re = regexp.MustCompile(`^[–%# &"]*`)

func clean(s string) string {
	return strings.TrimSpace(string(re.ReplaceAll([]byte(s), void)))
}

func (b Body) Len() int { return len(b.Tus) }
func (b Body) Less(i, j int) bool {
	l := clean(b.Tus[i].String())
	r := clean(b.Tus[j].String())
	return strings.Compare(l, r) < 0
}
func (b Body) String() string { return fmt.Sprintf("%s [Tus: %d]", b.XMLName, len(b.Tus)) }
func (b Body) Swap(i, j int)  { b.Tus[i], b.Tus[j] = b.Tus[j], b.Tus[i] }

type TranslationUnit struct {
	ChangeDate          *string                  `xml:"changedate,attr"`
	ChangeId            *string                  `xml:"changeid,attr"`
	CreationDate        *string                  `xml:"creationdate,attr"`
	CreationId          *string                  `xml:"creationid,attr"`
	CreationTool        *string                  `xml:"creationtool,attr"`
	CreationToolVersion *string                  `xml:"creationtoolversion,attr"`
	DataType            *string                  `xml:"datatype,attr"`
	LastUsageDate       *string                  `xml:"lastusagedate,attr"`
	OEncoding           *string                  `xml:"o-encoding,attr"`
	OTmf                *string                  `xml:"o-tmf,attr"`
	Props               *[]Property              `xml:"prop"`
	SegType             *string                  `xml:"segtype,attr"`
	SrcLang             *string                  `xml:"srclang,attr"`
	TuId                *string                  `xml:"tuid,attr"`
	Tuvs                []TranslationUnitVariant `xml:"tuv"`
	UsageCount          *string                  `xml:"usagecount,attr"`
	XMLName             xml.Name                 `xml:"tu"`
}

func (t TranslationUnit) String() string {
	return fmt.Sprintf("%s", clean(t.Tuvs[0].Seg))
}

type TranslationUnitVariant struct {
	ChangeDate          *string     `xml:"changedate,attr"`
	ChangeId            *string     `xml:"changeid,attr"`
	CreationDate        *string     `xml:"creationdate,attr"`
	CreationId          *string     `xml:"creationid,attr"`
	CreationTool        *string     `xml:"creationtool,attr"`
	CreationToolVersion *string     `xml:"creationtoolversion,attr"`
	DataType            *string     `xml:"datatype,attr"`
	Language            string      `xml:"lang,attr"`
	LastUsageDate       *string     `xml:"lastusagedate,attr"`
	OEncoding           *string     `xml:"o-encoding,attr"`
	OTmf                *string     `xml:"o-tmf,attr"`
	Props               *[]Property `xml:"prop"`
	Seg                 string      `xml:"seg"`
	XMLName             xml.Name    `xml:"tuv"`
}

func (t TranslationUnitVariant) String() string {
	return fmt.Sprintf("%s", t.Seg)
}
