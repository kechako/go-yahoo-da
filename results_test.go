package da

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"
)

const testXML = `<?xml version="1.0" encoding="UTF-8" ?>
<ResultSet xsi:schemaLocation="urn:yahoo:jp:jlp:DAService https://jlp.yahooapis.jp/DAService/V1/parseResponse.xsd">
  <Result>
    <ChunkList>
      <Chunk>
        <Id>0</Id>
        <Dependency>1</Dependency>
        <MorphemList>
          <Morphem>
          <Surface>うち</Surface>
          <Reading>うち</Reading>
          <Baseform>うち</Baseform>
          <POS>名詞</POS>
          <Feature>名詞,地名町名,*,うち,うち,うち</Feature>
          </Morphem>
          <Morphem>
          <Surface>の</Surface>
          <Reading>の</Reading>
          <Baseform>の</Baseform>
          <POS>助詞</POS>
          <Feature>助詞,助詞連体化,*,の,の,の</Feature>
          </Morphem>
        </MorphemList>
      </Chunk>
      <Chunk>
        <Id>1</Id>
        <Dependency>3</Dependency>
        <MorphemList>
          <Morphem>
          <Surface>庭</Surface>
          <Reading>にわ</Reading>
          <Baseform>庭</Baseform>
          <POS>名詞</POS>
          <Feature>名詞,名詞場所,*,庭,にわ,庭</Feature>
          </Morphem>
          <Morphem>
          <Surface>に</Surface>
          <Reading>に</Reading>
          <Baseform>に</Baseform>
          <POS>助詞</POS>
          <Feature>助詞,格助詞,*,に,に,に</Feature>
          </Morphem>
          <Morphem>
          <Surface>は</Surface>
          <Reading>は</Reading>
          <Baseform>は</Baseform>
          <POS>助詞</POS>
          <Feature>助詞,係助詞,*,は,は,は</Feature>
          </Morphem>
        </MorphemList>
      </Chunk>
      <Chunk>
        <Id>2</Id>
        <Dependency>3</Dependency>
        <MorphemList>
          <Morphem>
          <Surface>二羽</Surface>
          <Reading>2わ</Reading>
          <Baseform>2羽</Baseform>
          <POS>接尾辞</POS>
          <Feature>接尾辞,助数,*,二羽,2わ,2羽</Feature>
          </Morphem>
          <Morphem>
          <Surface>鶏</Surface>
          <Reading>にわとり</Reading>
          <Baseform>鶏</Baseform>
          <POS>名詞</POS>
          <Feature>名詞,名詞,*,鶏,にわとり,鶏</Feature>
          </Morphem>
          <Morphem>
          <Surface>が</Surface>
          <Reading>が</Reading>
          <Baseform>が</Baseform>
          <POS>助詞</POS>
          <Feature>助詞,格助詞,*,が,が,が</Feature>
          </Morphem>
        </MorphemList>
      </Chunk>
      <Chunk>
        <Id>3</Id>
        <Dependency>-1</Dependency>
        <MorphemList>
          <Morphem>
          <Surface>い</Surface>
          <Reading>い</Reading>
          <Baseform>い</Baseform>
          <POS>動詞</POS>
          <Feature>動詞,一段,未然ウ接続,い,い,い</Feature>
          </Morphem>
          <Morphem>
          <Surface>ます</Surface>
          <Reading>ま</Reading>
          <Baseform>ま</Baseform>
          <POS>助動詞</POS>
          <Feature>助動詞,助動詞ます,基本形,ます,ま,ま</Feature>
          </Morphem>
          <Morphem>
          <Surface>。</Surface>
          <Reading>。</Reading>
          <Baseform>。</Baseform>
          <POS>特殊</POS>
          <Feature>特殊,句点,*,。,。,。</Feature>
          </Morphem>
        </MorphemList>
      </Chunk>
    </ChunkList>
  </Result>
</ResultSet>`

var testResultSet = ResultSet{
	XMLName: xml.Name{
		Local: "ResultSet",
	},
	Results: []Result{
		{
			Chunks: []Chunk{
				{
					ID:         0,
					Dependency: 1,
					Morphemes: []Morpheme{
						{
							Surface:  "うち",
							Reading:  "うち",
							Baseform: "うち",
							POS:      "名詞",
							Feature:  "名詞,地名町名,*,うち,うち,うち",
						},
						{
							Surface:  "の",
							Reading:  "の",
							Baseform: "の",
							POS:      "助詞",
							Feature:  "助詞,助詞連体化,*,の,の,の",
						},
					},
				},
				{
					ID:         1,
					Dependency: 3,
					Morphemes: []Morpheme{
						{
							Surface:  "庭",
							Reading:  "にわ",
							Baseform: "庭",
							POS:      "名詞",
							Feature:  "名詞,名詞場所,*,庭,にわ,庭",
						},
						{
							Surface:  "に",
							Reading:  "に",
							Baseform: "に",
							POS:      "助詞",
							Feature:  "助詞,格助詞,*,に,に,に",
						},
						{
							Surface:  "は",
							Reading:  "は",
							Baseform: "は",
							POS:      "助詞",
							Feature:  "助詞,係助詞,*,は,は,は",
						},
					},
				},
				{
					ID:         2,
					Dependency: 3,
					Morphemes: []Morpheme{
						{
							Surface:  "二羽",
							Reading:  "2わ",
							Baseform: "2羽",
							POS:      "接尾辞",
							Feature:  "接尾辞,助数,*,二羽,2わ,2羽",
						},
						{
							Surface:  "鶏",
							Reading:  "にわとり",
							Baseform: "鶏",
							POS:      "名詞",
							Feature:  "名詞,名詞,*,鶏,にわとり,鶏",
						},
						{
							Surface:  "が",
							Reading:  "が",
							Baseform: "が",
							POS:      "助詞",
							Feature:  "助詞,格助詞,*,が,が,が",
						},
					},
				},
				{
					ID:         3,
					Dependency: -1,
					Morphemes: []Morpheme{
						{
							Surface:  "い",
							Reading:  "い",
							Baseform: "い",
							POS:      "動詞",
							Feature:  "動詞,一段,未然ウ接続,い,い,い",
						},
						{
							Surface:  "ます",
							Reading:  "ま",
							Baseform: "ま",
							POS:      "助動詞",
							Feature:  "助動詞,助動詞ます,基本形,ます,ま,ま",
						},
						{
							Surface:  "。",
							Reading:  "。",
							Baseform: "。",
							POS:      "特殊",
							Feature:  "特殊,句点,*,。,。,。",
						},
					},
				},
			},
		},
	},
}

func Test_decodeResultSet(t *testing.T) {
	var res ResultSet
	var err error

	res, err = decodeResultSet(strings.NewReader("Invalid XML"))
	if err == nil {
		t.Fatalf("should be fail: %v", err)
	}

	res, err = decodeResultSet(strings.NewReader(testXML))
	if err != nil {
		t.Fatalf("should not be fail: %v", err)
	}

	if !reflect.DeepEqual(res, testResultSet) {
		t.Errorf("want %#v\nbut %#v", testResultSet, res)
	}
}
