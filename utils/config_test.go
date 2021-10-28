package utils

import (
	"fmt"
	"testing"
)
func TestParseConfig(t *testing.T) {
	config:=InitAndGetConf("../config.yaml")
	fmt.Printf("%+v\n",config)
	fmt.Printf("%+v\n",config.Consensus)
}
