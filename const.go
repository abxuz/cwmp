package cwmp

import "errors"

const (
	XsdTypeString string = "xsd:string"
	XsdTypeInt    string = "xsd:int"
	XsdTypeUint   string = "xsd:unsignedInt"
	XsdTypeBool   string = "xsd:boolean"
)

const (
	// SoapArray array type
	SoapArray string = "SOAP-ENC:Array"
)

const (
	// EventBootStrap first connection
	EventBootStrap string = "0 BOOTSTRAP"
	// EventBoot reset or power on
	EventBoot string = "1 BOOT"
	// EventPeriodic periodic inform
	EventPeriodic string = "2 PERIODIC"
	// EventScheduled scheduled infrorm
	EventScheduled string = "3 SCHEDULED"
	// EventValueChange value change event
	EventValueChange string = "4 VALUE CHANGE"
	// EventKicked acs notify cpe
	EventKicked string = "5 KICKED"
	// EventConnectionRequest cpe request connection
	EventConnectionRequest string = "6 CONNECTION REQUEST"
	// EventTransferComplete download complete
	EventTransferComplete string = "7 TRANSFER COMPLETE"
	// EventClientChange custom event client online/offline
	EventClientChange string = "8 CLIENT CHANGE"
)

var (
	ErrInvalidCWMPXML = errors.New("invalid cwmp xml data")
	ErrNotImplemented = errors.New("not implemented")
)
