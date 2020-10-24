package utilities

import "unicode";
import "strings";

const  MINIMUM_LEN		int		= 4;

// ==========================================================
// convert string to lowercase with the first char in Upper
// ==========================================================
func LowerCaseCapitalFirstChar (s string) []byte {
    var phrase []byte = ConvertToLowerCase (s);
    var firstChar byte = byte(unicode.ToUpper(rune(phrase [0]))); 
    phrase[0] = firstChar;
    return phrase;
}


// ==========================================================
// convert string to uppercase
// ==========================================================
func ConvertToLowerCase (s string) []byte {
    var phrase []byte = []byte(strings.ToLower (s));
    return phrase;
}


// ==========================================================
// prepare string for phonetics processing
// ==========================================================
func PrepareString (s string) string {
    s = strings.TrimSpace(strings.ToUpper (s));
    return s;
}


// ==========================================================
// 
// ==========================================================
func InputStringValid (phrase string) *ErrorStruct {
    if (len(phrase) < MINIMUM_LEN) {
        return (&STRING_TOO_SHORT);
    }
    
    return nil;
}


