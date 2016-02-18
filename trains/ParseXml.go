package main

import (
	"encoding/xml"
	"log"
)

type Location struct {
	Tpl string  `xml:"tpl,attr"`
	Pta *string `xml:"pta,attr"`
	Ptd *string `xml:"ptd,attr"`
	Wta *string `xml:"wta,attr"`
	Wtd *string `xml:"wtd,attr"`
	Wtp *string `xml:"wtp,attr"`
	Arr *struct {
		Et  string `xml:"et,attr"`
		Src string `xml:"src,attr"`
		Wet string `xml:"wet,attr"`
	} `xml:"arr"`
	Dep *struct {
		Et  string `xml:"et,attr"`
		Src string `xml:"src,att"r`
	} `xml:"dep"`
	Plat *struct {
		CisPlatsup string `xml:"cisPlatsup,attr"`
		Platsup    bool   `xml:"platsup,attr"`
		Plat       string `xml:",chardata"`
	} `xml:"plat"`
	Pass *struct {
		Et  string `xml:"et,attr"`
		Src string `xml:"src,attr"`
	}
}

type TS struct {
	Rid        string `xml:"rid,attr"`
	Ssd        string `xml:"ssd,attr"`
	Uid        string `xml:"uid,attr"`
	LateReason *struct {
		Reason int `xml:",chardata"`
	} `xml:"LateReason"`
	Locations []Location `xml:"Location"`
}

type Pport struct {
	Ts      string `xml:"ts,attr"`
	Version string `xml:"version,attr"`
	Ur      *struct {
		updateOrigin string `xml:"updateOrigin,attr"`
		Deactivated  *struct {
			Rid string `xml:"rid,attr"`
		} `xml:"deactivated"`
		Ts *TS `xml:"TS"`
	} `xml:"uR"`
}

func XmlToStructs(xmlBytes []byte) *Pport {
	data := &Pport{}
	err := xml.Unmarshal(xmlBytes, data)
	if err != nil {
		log.Printf("Failed to parse:\n%s\nError is %s ",
			string(xmlBytes), err)
	}
	return data
}
