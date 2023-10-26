/*
The action package functions more like a helper rather than a service that handles business logic.
It does not use dependency injection, which allows the functions within this package to be used
across various services, facilitating repeated usage.

Unlike services that handle business logic, the action package provides utility functions that can
be used throughout the application. This makes it a versatile tool in the codebase.

One such function within this package is Logout, which is used in the auth service's SignOut function
and also in the user service's Delete function. This is an example of how functions in the
action package can be used across services.

Example usage:

    if err := action.Logout(c); err != nil {
		return err
	}
*/
package action
