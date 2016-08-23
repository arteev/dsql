package parameters

import "strings"

//A Parameter store value by name
type Parameter struct {
	ID          int
	Name        string
	Value       string
	Description string
}

//String this stringer
func (p Parameter) String() string {
	val := p.Value
	if p.Value == "" {
		val = "(empty)"
	}
	return p.Name + "=" + val
}

//ValueStr  get Parameter.Value, but return "(empty)" if value is null or empty
func (p Parameter) ValueStr() string {
	if p.Value == "" {
		return "(empty)"
	}
	return p.Value
}

//Bool returns value as boolean cast from string value
func (p Parameter) Bool() bool {
	return p.Value != "" && p.Value != "0" && strings.ToUpper(p.Value) != "FALSE"
}

//MustParameter returns a new parameter
func MustParameter(name, value, description string) Parameter {
	return Parameter{
		Name:        name,
		Value:       value,
		Description: description,
	}
}

//IsPreDefined returns true if the parameter a pre-defined
func (p Parameter) IsPreDefined() bool {
	for _, def := range definedParams {
		if strings.ToUpper(p.Name) == strings.ToUpper(def.Name) {
			return true
		}
	}
	return false
}
