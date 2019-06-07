package controller

import (
	"github.com/olegtrautvein/testcrd/operatorsdk/kublr-operator/pkg/controller/logging"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, logging.Add)
}
