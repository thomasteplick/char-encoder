/*
  Character encoder.  This is a web application that encodes alpha-numeric characters
  to 81 integers that represent their image in a 81-cell grid (9x9).  The integers are
  1 or -1, with 1 representing an "on" cell and -1 representing an "off" cell.  The
  user sketches the character in a 81-cell grid consisting of checkboxes.  A checkbox
  is checked to turn the cell on.  The user submits the form and the 81 integers are
  saved to a disk file along with the character it represents.  For example, each line
  in the file would look as shown below.
  A,-1,-1,-1,1,1,...,-1,-1

  These encoded characters could be used as training/testing samples for a MLP
  Neural Network for character recognition.
*/

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	rows         int = 9
	cols         int = 9
	tmpl             = "templates/encoder.html" // html template relative address
	addr             = "127.0.0.1:8080"         // http server listen address
	pattern          = "/char-encoder"          // http handler for character encoding
	encodedChars     = "encoded_chars.csv"      // file containing the encoded characters
	cells            = rows * cols              // number of checkboxes
	dataDir          = "data/"                  // directory for the encoded characters
)

// dynamic html data used in actions
type PlotT struct {
	Grid   []string // encoder grid for sketching the character
	Status string
}

type Encoder struct {
	output []string
	plot   PlotT
}

var (
	tmplEncode *template.Template
)

// init parses the html template file done only once
func init() {
	tmplEncode = template.Must(template.ParseFiles(tmpl))
}

// runEncoder encodes the sketch and saves it to disk
func (enc *Encoder) runEncoder(r *http.Request) error {
	// Determine if anything was entered yet
	r.ParseForm()
	charSelect := r.FormValue("charselect")
	charChecked := r.PostForm["charcheck"]
	if len(charSelect) == 0 || len(charChecked) == 0 {
		fmt.Printf("Select character and sketch it in the grid\n")
		// Set encoder status
		enc.plot.Status = "Select character and sketch it in grid"
		// not an error, but initial invocation, so return nil
		return nil
	}

	// Set the checked boxes to 1
	for _, val := range charChecked {
		index, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("int conversion %v error %v", val, err.Error())
		}
		enc.output[index] = "1"
	}

	// Set encoder status
	enc.plot.Status = fmt.Sprintf("Encoded character written to disk file %s", path.Join(dataDir, encodedChars))

	// Save the encoded char to disk
	f, err := os.OpenFile(
		path.Join(dataDir, encodedChars), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Open file %v error: %v", encodedChars, err.Error())
	}
	defer f.Close()
	// Each line looks like:
	// a,-1,-1,...,1,1,-1,...,-1, the character followed by 81 ones or minus ones
	fmt.Fprintf(f, "%s,%s\n", charSelect, strings.Join(enc.output, ","))

	return nil
}

// restoreOutputAndGridValues resets the output and Grid to default values
func (enc *Encoder) initOutputAndGridValues() error {
	// Set the unchecked boxes to -1
	// Write the values of the grid so the checkbox locations can be determined
	for i := range enc.output {
		enc.plot.Grid[i] = strconv.Itoa(i)
		enc.output[i] = "-1"
	}
	return nil
}

// handleCharEncoding performs character encoding
func handleCharEncoding(w http.ResponseWriter, r *http.Request) {

	encoder := Encoder{
		output: make([]string, cells),
		plot:   PlotT{Grid: make([]string, cells), Status: ""},
	}

	err := encoder.initOutputAndGridValues()
	if err != nil {
		encoder.plot.Status = "restoreOutputAndGridValues error.  See log."
		fmt.Printf("restoreOutputAndGridValues error: %v\n", err.Error())
		if err := tmplEncode.Execute(w, encoder.plot); err != nil {
			log.Fatalf("Write to HTTP output using template with grid error: %v\n", err)
		}
		return
	}

	err = encoder.runEncoder(r)
	if err != nil {
		encoder.plot.Status = "runEncoder error.  See log."
		fmt.Printf("runEncoder error: %v\n", err.Error())
		if err := tmplEncode.Execute(w, encoder.plot); err != nil {
			log.Fatalf("Write to HTTP output using template with grid error: %v\n", err)
		}
		return
	}

	// Write to HTTP output using template and grid
	if err := tmplEncode.Execute(w, encoder.plot); err != nil {
		log.Fatalf("Write to HTTP output using template with grid error: %v\n", err)
	}
}

func main() {
	// Setup http server with handler for character encoding
	http.HandleFunc(pattern, handleCharEncoding)
	fmt.Printf("Character Encoder server listening on %v.\n", addr)
	http.ListenAndServe(addr, nil)
}
