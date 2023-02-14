package cloner
import (
	"os"
	"testing"
	"fmt"
    "path/filepath"
)
//"os/exec"
func TestMain(m *testing.M) {
	
    // Create the "results" directory in the current working directory
    dir := "results"
    err := os.Mkdir(dir, 0755)
    if err != nil {
        fmt.Println("Error creating directory:", err)
        return
    }

    // Get the absolute path to the "results" directory
    absPath, err := filepath.Abs(dir)
    if err != nil {
        fmt.Println("Error getting absolute path:", err)
        return
    }

    // Call the cloneRepo function with the repository URL and "results" directory
    err = cloneRepo("https://github.com/KevinMi2023p/ECE461_team33", absPath)
    if err != nil {
        fmt.Println("Error cloning repository:", err)
    } else {
        fmt.Println("Repository cloned successfully!")
    }
}