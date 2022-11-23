//  User Api:
//   version: 0.0.1
//   title: Userapi
//  Schemes: http
//  Host: 192.168.33.50:8200
//  BasePath: /
//  Consumes:
//    - application/json

//  Produces:
//    - application/json
//
// swagger:meta
package docs

import "github.com/SzymekN/CRUD/pkg/model"

// Data structure representing single user
// swagger:response userResponse
type userResponseWrapper struct {
	// User data read from database
	// in: body
	Body model.User
}

// Data structure representing slice of users
// swagger:response usersResponse
type usersResponseWrapper struct {
	// User data read from database
	// in: body
	Body []model.User
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponeWrapper struct {
	// description of an error
	// in: body
	Body model.GenericError
}

// Generic message returned as a string
// swagger:response messageResponse
type messageResponeWrapper struct {
	// description of an error
	// in: body
	Body model.GenericMessage
}

// swagger:parameters  getUserV1 deleteUserV1 putUserV1 getUserV2 deleteUserV2 putUserV2
type userIDParameterWrapper struct {
	// ID of the user that will be deleted or read
	// in: path
	// required: true
	// example: 1
	ID int `json:"id"`
}

// swagger:parameters putUserV1 postUserV1 putUserV2 postUserV2
type userParameterWrapper struct {
	// User structure to update or create
	// in: body
	// required: true
	Body model.User
}
