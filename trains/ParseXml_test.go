package main

import (
	"encoding/xml"
	"log"
	"testing"
)

func Test_XmlConversion(t *testing.T) {
	xml := []byte(
		`<Pport xmlns="http://www.thalesgroup.com/rtti/PushPort/v12" 
				xmlns:ns3="http://www.thalesgroup.com/rtti/PushPort/Forecasts/v2" 
				ts="2016-02-10T18:42:57.8582793Z" version="12.0">
				<uR updateOrigin="TD">
					<TS rid="201602102520591" ssd="2016-02-10" uid="W11208">
						<ns3:Location pta="18:53" ptd="18:55" 
								tpl="BSNGSTK" wta="18:53" wtd="18:55">
							<ns3:arr et="18:53" src="TD" wet="18:52"/>
							<ns3:dep et="18:55" src="Darwin"/>
							<ns3:plat cisPlatsup="true" platsup="true">2</ns3:plat>
						</ns3:Location>
					 </TS>
				 </uR> 
			 </Pport>`)
	pport := XmlToStructs(xml)
	t.Log(pport)
	if pport.Ts != "2016-02-10T18:42:57.8582793Z" {
		t.Error("Bad time: ", pport.Ts)
	}
	if pport.Ur.Deactivated != nil {
		t.Error("Deactivated should be nil but is", pport.Ur.Deactivated)
	}
	if pport.Ur.Ts == nil {
		t.Error("Ts in nil")
	}

	if pport.Ur.Ts.Locations[0].Tpl != "BSNGSTK" {
		t.Error("Expecing Basingstoke")
	}

	if *pport.Ur.Ts.Locations[0].Pta != "18:53" {
		t.Error("Expecing 18:53")
	}
}

func Test_Deactivated(t *testing.T) {

	xml := []byte(
		`<?xml version="1.0" encoding="UTF-8"?>
		<Pport xmlns="http://www.thalesgroup.com/rtti/PushPort/v12" ts="2016-02-14T14:17:00.7083652Z" version="12.0">
			<uR updateOrigin="Darwin">
				<deactivated rid="201602142630790"/>
			</uR>
		</Pport>`)
	pport := XmlToStructs(xml)
	if pport.Ur.Deactivated.Rid != "201602142630790" {
		t.Error("Unable to pass deactivated", pport)
	}
	if pport.Ur.Ts != nil {
		t.Error("This is a deactivated update there should be no TS")
	}
}

func Test_StructToXml(t *testing.T) {
	p := Pport{}
	p.Version = "12.1"
	log.Println("Pport:", p)
	b, err := xml.Marshal(p)
	if err != nil {
		t.Error("Failed to marshal:", err)
	}
	log.Println("Pport:", string(b))
}

func Test_ManyLocations(t *testing.T) {
	xml := []byte(
		`<Pport xmlns="http://www.thalesgroup.com/rtti/PushPort/v12" 
				xmlns:ns3="http://www.thalesgroup.com/rtti/PushPort/Forecasts/v2" ts="2016-02-10T18:43:01.0717969Z" version="12.0">
			<uR updateOrigin="Darwin">
				<TS rid="201602102497749" ssd="2016-02-10" uid="W11502">
					<ns3:LateReason>-1</ns3:LateReason>
					<ns3:Location tpl="WRTINGJ" wtp="18:38:30">
						<ns3:pass et="18:43" src="Darwin"/>
					</ns3:Location>
					<ns3:Location pta="18:41" ptd="18:43" tpl="BSNGSTK" wta="18:41" wtd="18:43:30">
						<ns3:arr et="18:46" src="Darwin"/>
						<ns3:dep et="18:47" src="Darwin"/>
						<ns3:plat cisPlatsup="true" platsup="true">3</ns3:plat>
					</ns3:Location>
					<ns3:Location pta="18:54" ptd="18:54" tpl="FLEET" wta="18:53:30" wtd="18:54:30">
						<ns3:arr et="18:57" src="Darwin"/>
						<ns3:dep et="18:57" src="Darwin"/>
						<ns3:plat cisPlatsup="true" platsup="true">1</ns3:plat>
					</ns3:Location>
					<ns3:Location pta="18:59" ptd="19:00" tpl="FRBRMN" wta="18:59" wtd="19:00">
						<ns3:arr et="19:01" src="Darwin"/>
						<ns3:dep et="19:02" src="Darwin"/>
						<ns3:plat cisPlatsup="true" platsup="true">1</ns3:plat>
					</ns3:Location>
					<ns3:Location tpl="WOKINGJ" wtp="19:07:30">
						<ns3:pass et="19:10" src="Darwin"/>
					</ns3:Location>
					<ns3:Location tpl="WOKING" wtp="19:08">
						<ns3:pass et="19:10" src="Darwin"/>
						<ns3:plat cisPlatsup="true" platsup="true">2</ns3:plat>
					</ns3:Location>
					<ns3:Location tpl="HCRTJN" wtp="19:16">
						<ns3:pass et="19:18" src="Darwin"/>
					</ns3:Location>
					<ns3:Location tpl="SURBITN" wtp="19:17">
						<ns3:pass et="19:19" src="Darwin"/>
						<ns3:plat cisPlatsup="true" platsup="true">2</ns3:plat>
					</ns3:Location>
					<ns3:Location tpl="NEWMLDN" wtp="19:19">
						<ns3:pass et="19:21" src="Darwin"/>
					</ns3:Location>
					<ns3:Location tpl="WDON" wtp="19:21">
						<ns3:pass et="19:23" src="Darwin"/>
						<ns3:plat cisPlatsup="true" platsup="true">6</ns3:plat>
					</ns3:Location>
					<ns3:Location pta="19:26" ptd="19:27" tpl="CLPHMJM" wta="19:25:30" wtd="19:27">
						<ns3:arr et="19:26" src="Darwin" wet="19:27"/>
						<ns3:dep et="19:28" src="Darwin"/>
						<ns3:plat cisPlatsup="true" platsup="true">7</ns3:plat>
					</ns3:Location>
					<ns3:Location pta="19:34" tpl="WATRLMN" wta="19:34">
						<ns3:arr et="19:34" src="Darwin" wet="19:35"/>
						<ns3:plat cisPlatsup="true" platsup="true">7</ns3:plat>
					</ns3:Location>
				</TS>
			</uR>
		</Pport>`)
	pport := XmlToStructs(xml)
	numLocs := len(pport.Ur.Ts.Locations)
	if numLocs != 12 {
		t.Error("Not the right number of locations: ", numLocs)
	}
}
