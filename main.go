package main
import (
	"fmt"
	"os"
	"log"
	"strings"
	"io/ioutil"
	"strconv"
	"github.com/bradhe/stopwatch"
)
func checkErr( err error ) {
	if err != nil {
		log.Fatal(err)
	}
}
func readFile( filePath string ) string {
	content, err := ioutil.ReadFile( filePath )
	checkErr(err)
	return string(content)
} 
func findAndReplace( str string, find string, replace string ) string {
	fmt.Println("Replaced " + find + " with " + replace)
	return strings.Replace(str, find, replace, -1)
}
func replaceInstancePrefixes( content string ) string {
	prefix := "AutoInstance"

	for i := 0; i < 5; i++ {
		content = findAndReplace(content, prefix + strconv.Itoa(i) + "-", "")
	}
	return content
}
func main() {
	watch := stopwatch.Start()

	vmfPath := os.Args[1]
	vmfContent := readFile(vmfPath)
	vmfContent = replaceInstancePrefixes( vmfContent )
	
	fmt.Println("Done.")

	fmt.Println("Writing changes...")
	err := ioutil.WriteFile( vmfPath, []byte(vmfContent), 777 )
	checkErr(err)
	fmt.Println("Changes written.")
	
	
	watch.Stop()
	fmt.Printf("Precompile successful in %dms\n", watch.Milliseconds())
	fmt.Println("Press any key to continue...")
	fmt.Scanln()
}