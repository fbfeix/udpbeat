package udpbeat

import (
	"github.com/gravitational/trace"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&trace.TextFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

//func main() {
//	log.Infof("starting udpbeat, check trace.yml for details on config")
//	err := beat.Run(ELKBeatName, ELKBeatVersion, NewELK())
//	if err != nil {
//		log.Fatalf("err: %v", err)
//	}
//}
