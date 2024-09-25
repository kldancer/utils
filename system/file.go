package system

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func readAndWrite() {
	if _, err := os.Stat("/root/boc-2.conflist"); os.IsNotExist(err) {
		// If the config file does not exist, it will be ignored
		log.Errorf("err : %s", err)
	}
	data, err := os.ReadFile("/root/boc-2.conflist")
	if err != nil {
		log.Errorf("err : %s", err)
	}

	err = os.WriteFile("/root/boc.conflist", data, 0644)
	if err != nil {
		log.Errorf("err : %s", err)
	}

}
