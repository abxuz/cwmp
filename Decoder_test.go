package cwmp

import (
	"encoding/xml"
	"strings"
	"testing"
)

const data = `<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/"
    xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/"
    xmlns:xsd="http://www.w3.org/2001/XMLSchema"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:cwmp="urn:dslforum-org:cwmp-1-0">
    <SOAP-ENV:Header>
        <cwmp:ID SOAP-ENV:mustUnderstand="1">19665046</cwmp:ID>
    </SOAP-ENV:Header>
    <SOAP-ENV:Body>
	<cwmp:Inform>
            <DeviceId>
                <Manufacturer>SCTY</Manufacturer>
                <OUI>40F420</OUI>
                <ProductClass>TEWA-1000G</ProductClass>
                <SerialNumber>4685040F420F87E30</SerialNumber>
            </DeviceId>
            <Event SOAP-ENC:arrayType="cwmp:EventStruct[3]">
                <EventStruct>
                    <EventCode>0 BOOTSTRAP</EventCode>
                    <CommandKey></CommandKey>
                </EventStruct>
                <EventStruct>
                    <EventCode>1 BOOT</EventCode>
                    <CommandKey></CommandKey>
                </EventStruct>
                <EventStruct>
                    <EventCode>X CT-COM BIND</EventCode>
                    <CommandKey></CommandKey>
                </EventStruct>
            </Event>
            <MaxEnvelopes>1</MaxEnvelopes>
            <CurrentTime>2025-04-28T20:30:20+08:00</CurrentTime>
            <RetryCount>5</RetryCount>
            <ParameterList SOAP-ENC:arrayType="cwmp:ParameterValueStruct[0010]">
                <ParameterValueStruct>
                    <Name>InternetGatewayDevice.DeviceSummary</Name>
                    <Value xsi:type="xsd:string">InternetGatewayDevice:1.5[](Baseline:1,
                        EthernetLAN:1, USBLAN:1, Time:1, IPPing:1, DeviceAssociation:1, QoS:1,
                        WiFiLAN:1) , VoiceService:1.0[1](Endpoint:1, SIPEndpoint:1)</Value>
                </ParameterValueStruct>
                <ParameterValueStruct>
                    <Name>InternetGatewayDevice.DeviceInfo.SpecVersion</Name>
                    <Value xsi:type="xsd:string">1.0</Value>
                </ParameterValueStruct>
                <ParameterValueStruct>
                    <Name>InternetGatewayDevice.DeviceInfo.HardwareVersion</Name>
                    <Value xsi:type="xsd:string">V1.1</Value>
                </ParameterValueStruct>
                <ParameterValueStruct>
                    <Name>InternetGatewayDevice.DeviceInfo.SoftwareVersion</Name>
                    <Value xsi:type="xsd:string">Tianyi_V1.0.1</Value>
                </ParameterValueStruct>
                <ParameterValueStruct>
                    <Name>InternetGatewayDevice.DeviceInfo.ProvisioningCode</Name>
                    <Value xsi:type="xsd:string"></Value>
                </ParameterValueStruct>
                <ParameterValueStruct>
                    <Name>InternetGatewayDevice.ManagementServer.ConnectionRequestURL</Name>
                    <Value xsi:type="xsd:string">http://10.99.0.2:30010/</Value>
                </ParameterValueStruct>
                <ParameterValueStruct>
                    <Name>InternetGatewayDevice.ManagementServer.ParameterKey</Name>
                    <Value xsi:type="xsd:string"></Value>
                </ParameterValueStruct>
                <ParameterValueStruct>
                    <Name>
                        InternetGatewayDevice.WANDevice.1.WANConnectionDevice.1.WANIPConnection.1.ExternalIPAddress</Name>
                    <Value xsi:type="xsd:string">10.99.0.2</Value>
                </ParameterValueStruct>
                <ParameterValueStruct>
                    <Name>InternetGatewayDevice.WANDevice.1.WANCommonInterfaceConfig.WANAccessType</Name>
                    <Value xsi:type="xsd:string">GPON</Value>
                </ParameterValueStruct>
                <ParameterValueStruct>
                    <Name>InternetGatewayDevice.X_CT-COM_UserInfo.UserName</Name>
                    <Value xsi:type="xsd:string">LG2078428108</Value>
                </ParameterValueStruct>
                <ParameterValueStruct>
                    <Name>InternetGatewayDevice.X_CT-COM_UserInfo.UserId</Name>
                    <Value xsi:type="xsd:string"></Value>
                </ParameterValueStruct>
            </ParameterList>
        </cwmp:Inform>
    </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`

func TestDecoder(t *testing.T) {
	decoder := xml.NewDecoder(strings.NewReader(data))

	enveloper := new(EnvelopeDecoder)
	err := decoder.Decode(enveloper)
	if err != nil {
		t.Fatal(err)
	}
}
