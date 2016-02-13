package main

import (
	"encoding/xml"
)

type Location struct {
	Pta string `xml:"pta,attr"`
	Ptd string `xml:"ptd,attr"`
	Tpl string `xml:"tpl,attr"`
	Wta string `xml:"wta,attr"`
	Wtd string `xml:"wtd,attr"`
	Arr struct {
		Et  string `xml:"et,attr"`
		Src string `xml:"src,attr"`
		Wet string `xml:"wet,attr"`
	} `xml:"arr"`
	Dep struct {
		Et  string `xml:"et,attr"`
		Src string `xml:"src,att"r`
	} `xml:"dep"`
	Plat struct {
		CisPlatsup string `xml:"cisPlatsup,attr"`
		Platsup    bool   `xml:"platsup,attr"`
		Plat       string `xml:",chardata"`
	} `xml:"plat"`
}
type Pport struct {
	Ts      string `xml:"ts,attr"`
	Version string `xml:"version,attr"`
	UR      struct {
		updateOrigin string `xml:"updateOrigin,attr"`
		TS           struct {
			Rid        string `xml:"rid,attr"`
			Ssd        string `xml:"ssd,attr"`
			Uid        string `xml:"uid,attr"`
			LateReason struct {
				Reason int `xml:",chardata"`
			} `xml:"LateReason"`
			Locations []Location `xml:"Location"`
		} `xml:"TS"`
	} `xml:"uR"`
}

func XmlToStructs(xmlBytes []byte) *Pport {
	data := &Pport{}
	err := xml.Unmarshal(xmlBytes, data)
	failIf(err)
	return data
}
