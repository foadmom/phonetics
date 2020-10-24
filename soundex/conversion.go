package soundex

import "strings"

const maxStandardCodeLen int = 4;
const maxRefinedCodeLen  int = 20;

// ==============================================================
//
// ==============================================================
func CreateStandardCode (_phrase string) string {
    var phrase []byte = []byte (strings.ToUpper (_phrase));
    
    var soundexGeneratedCode = [maxStandardCodeLen] byte {'0', '0', '0', '0'};
    var soundexIndex         int 	 = 0;
    var previousCode         int	 = 0;
    
    soundexIndex++;

    if (len (phrase)>1) {
        soundexGeneratedCode[0] = phrase[0];
        for i:=1; i<len(phrase) && soundexIndex < maxStandardCodeLen; i++ {
            var sc SoundexCharStruct = StandardSoundexCharCode (phrase[i]);

            if (sc.CodeI > VOWEL) && (sc.CodeI != previousCode) {
                soundexGeneratedCode [soundexIndex] = byte(sc.CodeI);
                soundexIndex++;
                previousCode = sc.CodeI;
            }
             
        }
    }
    return string(soundexGeneratedCode[:maxStandardCodeLen]);
}


// ==============================================================
//
// ==============================================================
func CreateRefinedCode (_phrase string) string {
    var phrase []byte = []byte (strings.ToUpper (_phrase));
    
    var soundexGeneratedCode = [maxRefinedCodeLen] byte {};
    var soundexIndex         int 	 = 0;
    var previousCode         int	 = 0;
    
    soundexIndex++;

    if (len (phrase)>1) {
        soundexGeneratedCode[0] = phrase[0];
        for i:=0; i<len(phrase) && soundexIndex < maxRefinedCodeLen; i++ {
            var sc SoundexCharStruct = RefinedSoundexCharCode (phrase[i]);

            if (sc.CodeI > IGNORE) && (sc.CodeI != previousCode) {
                soundexGeneratedCode [soundexIndex] = byte(sc.CodeI);
                soundexIndex++;
                previousCode = sc.CodeI;
            }
             
        }
    }
    return string(soundexGeneratedCode[:soundexIndex]);
}





