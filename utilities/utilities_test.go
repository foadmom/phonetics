package utilities

import (
    "testing"
    "fmt"
    "strings"
)

func Test_utilities (t *testing.T) {
    _input := "FOAD HAS GONE to dinner";
    _output := LowerCaseCapitalFirstChar (_input);
    if (strings.Compare (string (_output), "Foad has gone to dinner") != 0) {
       t.Errorf("LowerCaseCapitalFirstChar did not work for %s it returned %s.", _input, _output);
    } else {
        fmt.Println ("tested OK"); 
    }
}

func Test_InputStringValid (t *testing.T) {
    _input := "AB";
    if (InputStringValid("AB") == true) {
       t.Errorf("InputStringValid (%s) returned true", _input);
    }
 }