package customerr

type RepositoryErrorInterface interface {
	Error() string
}

type DataNotFoundError struct {
	Msg string
	Err error
}

type TooManyResultsFoundError struct {
	Msg string
	Err error
}

func (de *DataNotFoundError) Error() string {
	return de.Msg
}

func (de *DataNotFoundError) Unwrap() error {
	return de.Err
}

func (de *TooManyResultsFoundError) Error() string {
	return de.Msg
}

func (de *TooManyResultsFoundError) Unwrap() error {
	return de.Err
}
