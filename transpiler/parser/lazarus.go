package parser

import "encoding/xml"

type LazarusInformation struct {
	XMLName         xml.Name        `xml:"CONFIG"`
	ProjectOptions  ProjectOptions  `xml:"ProjectOptions"`
	CompilerOptions CompilerOptions `xml:"CompilerOptions"`
	Debugging       Debugging       `xml:"Debugging"`
}

type ProjectOptions struct {
	XMLName        xml.Name       `xml:"ProjectOptions"`
	Version        int            `xml:"Version"`
	General        General        `xml:"General"`
	BuildModes     BuildModes     `xml:"BuildModes"`
	PublishOptions PublishOptions `xml:"PublishOptions"`
	RunParams      RunParams      `xml:"RunParams"`
	Units          Units          `xml:"Units"`
}

type General struct {
	XMLName        xml.Name `xml:"General"`
	Flags          Flags    `xml:"Flags"`
	SessionStorage string   `xml:"SessionStorage"`
	Title          string   `xml:"Title"`
	UseAppBundle   bool     `xml:"UseAppBundle"`
	ResourceType   string   `xml:"ResourceType"`
}

type Flags struct {
	XMLName                         xml.Name `xml:"Flags"`
	MainUnitHasCreateFormStatements bool     `xml:"MainUnitHasCreateFormStatements"`
	MainUnitHasTitleStatement       bool     `xml:"MainUnitHasTitleStatement"`
	Runnable                        bool     `xml:"Runnable"`
}

type BuildModes struct {
	XMLName xml.Name        `xml:"BuildModes"`
	Items   []BuildModeItem `xml:"Item"`
}

type BuildModeItem struct {
	XMLName xml.Name `xml:"Item"`
	Name    string   `xml:"Name,attr"`
	Default bool     `xml:"Default,attr"`
}

type PublishOptions struct {
	XMLName        xml.Name `xml:"PublishOptions"`
	Version        int      `xml:"Version"`
	UseFileFilters bool     `xml:"UseFileFilters"`
}

type RunParams struct {
	XMLName       xml.Name `xml:"RunParams"`
	FormatVersion int      `xml:"FormatVersion"`
}

type Units struct {
	XMLName xml.Name   `xml:"Units"`
	Unit    []UnitItem `xml:"Unit"`
}

type UnitItem struct {
	XMLName         xml.Name `xml:"Unit"`
	Filename        string   `xml:"Filename"`
	IsPartOfProject bool     `xml:"IsPartOfProject"`
}

type CompilerOptions struct {
	XMLName        xml.Name       `xml:"CompilerOptions"`
	Version        int            `xml:"Version"`
	Target         Target         `xml:"Target"`
	SearchPaths    SearchPaths    `xml:"SearchPaths"`
	CodeGeneration CodeGeneration `xml:"CodeGeneration"`
}

type Target struct {
	XMLName  xml.Name `xml:"Target"`
	Filename string   `xml:"Filename"`
}

type SearchPaths struct {
	XMLName             xml.Name `xml:"SearchPaths"`
	IncludeFiles        string   `xml:"IncludeFiles"`
	UnitOutputDirectory string   `xml:"UnitOutputDirectory"`
}

type CodeGeneration struct {
	XMLName          xml.Name `xml:"CodeGeneration"`
	TargetProcessor  string   `xml:"TargetProcessor"`
	TargetController string   `xml:"TargetController"`
	TargetCPU        string   `xml:"TargetCPU"`
	TargetOS         string   `xml:"TargetOS"`
}

type Debugging struct {
	XMLName    xml.Name   `xml:"Debugging"`
	Exceptions Exceptions `xml:"Exceptions"`
}

type Exceptions struct {
	XMLName xml.Name        `xml:"Exceptions"`
	Items   []ExceptionItem `xml:"Item"`
}

type ExceptionItem struct {
	XMLName xml.Name `xml:"Item"`
	Name    string   `xml:"Name"`
}
