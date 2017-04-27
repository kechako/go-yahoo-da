package da

import (
	"encoding/xml"
	"io"
)

// A ResultSet represents a result set from the API.
type ResultSet struct {
	XMLName xml.Name `xml:"ResultSet"`
	Results []Result `xml:"Result"`
}

// A Result represents a result of analysis for each clausea phrase.
type Result struct {
	Chunks []Chunk `xml:"ChunkList>Chunk"`
}

// A Chunk represents information of a clausea phrase.
type Chunk struct {
	ID         int        `xml:"Id"`
	Dependency int        `xml:"Dependency"`
	Morphemes  []Morpheme `xml:"MorphemList>Morphem"`
}

// String returns a string of chunk.
func (c Chunk) String() string {
	s := ""
	for _, m := range c.Morphemes {
		s += m.Surface
	}
	return s
}

// A Morpheme represents information of a Morpheme.
type Morpheme struct {
	Surface  string `xml:"Surface"`
	Reading  string `xml:"Reading"`
	Baseform string `xml:"Baseform"`
	POS      string `xml:"POS"`
	Feature  string `xml:"Feature"`
}

func decodeResultSet(r io.Reader) (res ResultSet, err error) {
	err = xml.NewDecoder(r).Decode(&res)
	return
}
