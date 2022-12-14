// Code generated by xgen. DO NOT EDIT.

package schema

// CvReferenceList ...
type CvReferenceList struct {
	CvSourceVersionAttr interface{}    `xml:"cvSourceVersion,attr,omitempty"`
	CvReference         []*CvReference `xml:"CvReference"`
}

// CvMappingRuleList ...
type CvMappingRuleList struct {
	CvMappingRule []string `xml:"CvMappingRule"`
}

// CvMapping ...
type CvMapping struct {
	ModelNameAttr     string             `xml:"modelName,attr"`
	ModelURIAttr      string             `xml:"modelURI,attr"`
	ModelVersionAttr  string             `xml:"modelVersion,attr"`
	CvReferenceList   *CvReferenceList   `xml:"CvReferenceList"`
	CvMappingRuleList *CvMappingRuleList `xml:"CvMappingRuleList"`
}

// CvMappingRule ...
type CvMappingRule struct {
	ScopePathAttr               string `xml:"scopePath,attr"`
	CvElementPathAttr           string `xml:"cvElementPath,attr"`
	CvTermsCombinationLogicAttr string `xml:"cvTermsCombinationLogic,attr"`
	RequirementLevelAttr        string `xml:"requirementLevel,attr"`
	NameAttr                    string `xml:"name,attr,omitempty"`
	IdAttr                      string `xml:"id,attr"`
	CvMappingRule               string `xml:"CvMappingRule"`
}

// CvTerm ...
type CvTerm struct {
	CvIdentifierRefAttr string `xml:"cvIdentifierRef,attr"`
	TermAccessionAttr   string `xml:"termAccession,attr"`
	TermNameAttr        string `xml:"termName,attr"`
	UseTermNameAttr     bool   `xml:"useTermName,attr,omitempty"`
	UseTermAttr         bool   `xml:"useTerm,attr"`
	AllowChildrenAttr   bool   `xml:"allowChildren,attr"`
	IsRepeatableAttr    bool   `xml:"isRepeatable,attr,omitempty"`
}

// CvReference ...
type CvReference struct {
	CvIdentifierAttr string `xml:"cvIdentifier,attr"`
	CvNameAttr       string `xml:"cvName,attr"`
}
