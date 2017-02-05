package server

import "net/url"

type (
	//GetNotSupported implicates that the GET method cannot be used
	GetNotSupported struct{}
	//PostNotSupported implicates that the POST method cannot be used
	PostNotSupported struct{}
	//PutNotSupported implicates that the PUT method cannot be used
	PutNotSupported struct{}
	//DeleteNotSupported implicates that the DELETE method cannot be used
	DeleteNotSupported struct{}
)

//Get defines an endpoint with a not supported GET method
func (GetNotSupported) Get(values url.Values) (int, interface{}) {
	return 405, ""
}

//Post defines an endpoint with a not supported POST method
func (PostNotSupported) Post(values url.Values) (int, interface{}) {
	return 405, ""
}

//Put defines an endpoint with a not supported PUT method
func (PutNotSupported) Put(values url.Values) (int, interface{}) {
	return 405, ""
}

//Delete defines an endpoint with a not supported DELETE method
func (DeleteNotSupported) Delete(values url.Values) (int, interface{}) {
	return 405, ""
}
