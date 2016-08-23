package parameters

//A RepositoryParams repository for entities of the parameters
type RepositoryParams interface {
	Add(p Parameter) error
	Update(p Parameter) error
	FindByName(name string) (Parameter, error)
	Delete(p Parameter) error
	//Refresh(p *Parameter) error
	Close() error
	All() (CollectionRepositoryParams, error)
}

//A CollectionRepositoryParams returns collection of the parameters
type CollectionRepositoryParams interface {
	Get() []Parameter
}
