package main

import "fmt"
import "bufio"
import "os";
import "strings";
import "regexp";

import metaphone "phonetics/metaphone"
import soundex   "phonetics/soundex"
//import utilities "phonetics/utilities";


var testData = []string{"ackermann", "Fusedale", "Charon", "Lampart",
	"Noulton", "Graham", "grahvm", "braz", ""}


func main() {
	fmt.Printf ("testing phonetics\n");
//	testStringSplit ("seyed-foad momtazi, anari");
//	testStringSplitRegExp ();
//	testPostgres ();
	testPhonetics();
//    test_errorTrace ();
}


func testStringSplitRegExp () {
	for true == true {
    	reader := bufio.NewReader(os.Stdin);
    	fmt.Print("Enter text: ");
    	_text, _ := reader.ReadString('\n');
    	_text = strings.Trim (_text, "\n");
    
        s := strings.Split("ian macdonald", "mac");
        fmt.Println(s);
        re := regexp.MustCompile(`[ ,_-']`);
    
        // Split based on pattern.
        // ... Second argument means "no limit."
        result := re.Split(_text, -1);
    
        for i := range(result) {
            fmt.Println(result[i]);
        }
	}
}
    
func testStringSplit (s string) {
    result := strings.Split(s, " ");

    // Display all elements.
    for i := range result {
        fmt.Println(result[i]);
    }
}


func testStandardSoundex() {
	for i := 0; i < len(testData); i++ {
		fmt.Printf("soundex code for %s c= %s   - ", testData[i], soundex.CreateStandardCode(testData[i]))
		fmt.Printf("soundex code for %s = %s\n", testData[i], soundex.CreateRefinedCode(testData[i]))
	}
}

func testPhonetics () {
	for true == true {
//	    var _err utilities.ErrorTrace;
		c := [50]byte{}
		reader := bufio.NewReader(os.Stdin);
		fmt.Print("Enter text: ");
		_text, _ := reader.ReadString('\n');
		_text = strings.Trim (_text, "\n");
		metaphone.GenerateCode(_text, c[:]);
		fmt.Printf("GenerateCode: %s\n", c);
		fmt.Printf("soundex code for %s c= %s\n", _text, soundex.CreateStandardCode(_text));
		fmt.Printf("refined soundex code for %s = %s\n", _text, soundex.CreateRefinedCode(_text));
	}
}






