package parametergetter

//A ParameterGetter use for fetch parameter from repository and etc.
type ParameterGetter interface {
	//IsSet(name string) bool
	Get(name string) interface{}
	GetDef(name string, def interface{}) interface{}
}
