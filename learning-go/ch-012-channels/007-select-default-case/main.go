package channels

import (
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	// ?
}

// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}
