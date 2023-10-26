/*
The action package functions more like a helper rather than a service that handles business logic.
However, unlike typical helper functions, the action package has ties to models. For example,
functions like Login and Logout are tied to the model.User.

This package does not use dependency injection, which allows the functions within this package to be used
across various services, facilitating repeated usage.

Unlike services that handle business logic or helpers that are not tied to any models, the action
package provides utility functions that are tied to specific models and can be used throughout
the application. This makes it a versatile tool in the codebase.

Example usage:

    if err := action.Logout(c); err != nil {
		return err
	}
*/
package action
