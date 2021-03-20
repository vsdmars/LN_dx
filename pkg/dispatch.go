package pkg

import (
	"errors"
	"fmt"
)

const (
	errMP            = "mp: [%s] does not exist"
	errAction        = "action: [%s] not defined"
	errMpHasNoAction = "mp: [%s] has no action: [%s] defined"
)

func dispatch(mp string, action string) error {
	imp, ok := mpToMp[mp]
	if !ok {
		return errors.New(fmt.Sprint(errMP, mp))
	}

	iaction, ok := actionToAction[action]
	if !ok {
		return errors.New(fmt.Sprint(errAction, action))
	}

	found := false
	for _, act := range mpActionType[imp] {
		if act == iaction {
			found = true
			break
		}
	}

	if !found {
		return errors.New(fmt.Sprint(errMpHasNoAction, mp, action))
	}
	return nil
}
