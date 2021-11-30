package helpers

import "log"

// HandleErrors TODO: Get an optional message
// HandleErrors is a utility function to handle errors
func HandleErrors(err error) {
	log.Fatalln(err)
	return
}
