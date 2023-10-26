/*
The helper package provides a set of utility functions that can be used throughout the application
without any ties to any models. These functions typically perform tasks such as data processing,
validation, or other operations that do not require interaction with the database or models.

Functions in the helper package are designed to be used repeatedly and can help simplify code by
extracting common logic into reusable functions.

Example usage:

    b, err := helper.BindRequest[request.SignUpRequest](c)
	if err != nil {
		return err
	}
*/
package helper
