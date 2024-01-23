// Package classification PetStore.
//
// Documentation of PetStore API.
//
//		Version: 1.0.0
//		Schemes:
//		- http
//		BasePath: /
//
//		Consumes:
//		- application/json
//		- multipart/form-data
//
//		Produces:
//		- application/json
//
//		Security:
//		- basic
//
//		SecurityDefinitions:
//		 Bearer:
//	   type: apiKey
//	   name: Authorization
//	   in: header
//
// swagger:meta
package docs

//go:generate swagger generate spec -o ../../public/swagger.json --scan-models
