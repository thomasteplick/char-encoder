<h3>Alpha-Numeric Encoder for Training Multilayer Perceptron Character Recognition Application </h3>

<p>
This program is a web application written in Go that uses the html/template package to create the web page.  From the src/encoder/ directory, issue "go run encode.go" to start the Character Encoder server.
In a web browser, enter URL http://127.0.0.1:8080/char-encoder to start the application.  Select the character from the drop-down list and enter its pattern by checking the boxes in the grid.  When satisfied with
the pattern, click the Encode button.  The character and its 91-integer encoding is saved to data/encoded_chars.csv.  Additional character entries are appended to the file.  The program converts the pattern entered
in the 91-checkbox grid to a sequence of 1 and -1; the checked boxes are given 1, and the unchecked boxes are given -1.
</p>
