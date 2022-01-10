package grecover

import "log"

// DefaultRecovery default recover func.
func DefaultRecovery() {
	if e := recover(); e != nil {
		log.Printf("wrapper exec recover:%v\n", e)
	}
}
