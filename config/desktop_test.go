package config

import (
	"log"
	"os"
	"testing"
)

// Test function for the GetDesktopPath function
func TestGetDesktopPath(t *testing.T) {
	folderInfo, err := os.Stat(GetDesktopPath("standard"))
// If the directory does not exist, an error is reported in the testing framework
	if os.IsNotExist(err) {
		t.Errorf("Failed, expecting directory to exist: %v", err)
	}
// Logs the folder information to the console
	log.Println(folderInfo)
}
