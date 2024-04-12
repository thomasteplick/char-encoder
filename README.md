<h3>Alpha-Numeric Encoder for Training Multilayer Perceptron Character Recognition Application </h3>

<p>
This program is a web application written in Go that uses the html/template package to create the web page.  From the src/encoder/ directory, issue "go run encode.go" to start the Character Encoder server.
In a web browser, enter URL http://127.0.0.1:8080/char-encoder to start the application.  Select the character from the drop-down list and enter its pattern by checking the boxes in the grid.  When satisfied with
the pattern, click the Encode button.  The character and its 91-integer encoding is saved to data/encoded_chars.csv.  Additional character entries are appended to the file.  The program converts the pattern entered
in the 91-checkbox grid to a sequence of 1 and -1; the checked boxes are given 1, and the unchecked boxes are given -1.  These patterns can be used to train a Neural Network to perform character recognition. 
</p>

<p>Shown below is an example for entering A in the checkbox grid.  The following is what is written to the disk file:
A,-1,-1,-1,-1,1,-1,-1,-1,-1,-1,-1,-1,1,-1,1,-1,-1,-1,-1,-1,1,-1,-1,-1,1,-1,-1,-1,-1,1,-1,-1,-1,1,-1,-1,-1,-1,1,1,1,1,1,-1,-1,-1,-1,1,-1,-1,-1,1,-1,-1,-1,1,-1,-1,-1,-1,-1,1,-1,-1,1,-1,-1,-1,-1,-1,1,-1,1,-1,-1,-1,-1,-1,-1,-1,1
</p>

![image](https://github.com/thomasteplick/char-encoder/assets/117768679/a22f862b-63bb-4aaa-9b3e-33a288fa5223)

