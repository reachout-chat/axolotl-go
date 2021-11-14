package tests

import (
	"fmt"
	"testing"

	"github.com/reachout-chat/axolotl-go/util/keyhelper"
)

func TestRegistrationID(t *testing.T) {
	regID := keyhelper.GenerateRegistrationID()
	fmt.Println(regID)
}
