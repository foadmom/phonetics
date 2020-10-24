package soundex

const VOWEL  int = 48;
const IGNORE int = 1;

type SoundexCharStruct struct {
    Char   	byte;
    CodeI	int;
    CodeC	byte;
}


// group defs for standard soundex
const S_GROUP_1     int = 49;
const S_GROUP_2     int = 50;
const S_GROUP_3     int = 51;
const S_GROUP_4     int = 52;
const S_GROUP_5     int = 53;
const S_GROUP_6     int = 54;

var charToStandardCode = [26] SoundexCharStruct {
    SoundexCharStruct {'A', VOWEL, byte(VOWEL)},
    SoundexCharStruct {'B', S_GROUP_1, '1'},
    SoundexCharStruct {'C', S_GROUP_2, '2'},
    SoundexCharStruct {'D', S_GROUP_3, '3'},
    SoundexCharStruct {'E', VOWEL, byte(VOWEL)},
    SoundexCharStruct {'F', S_GROUP_1, '1'},
    SoundexCharStruct {'G', S_GROUP_2, '2'},
    SoundexCharStruct {'H', IGNORE, byte(IGNORE)},
    SoundexCharStruct {'I', VOWEL, byte(VOWEL)},
    SoundexCharStruct {'J', S_GROUP_2, '2'},
    SoundexCharStruct {'K', S_GROUP_2, '2'},
    SoundexCharStruct {'L', S_GROUP_4, '4'},
    SoundexCharStruct {'M', S_GROUP_5, '5'},
    SoundexCharStruct {'N', S_GROUP_5, '5'},
    SoundexCharStruct {'O', VOWEL, byte(VOWEL)},
    SoundexCharStruct {'P', S_GROUP_1, '1'},
    SoundexCharStruct {'Q', S_GROUP_2, '2'},
    SoundexCharStruct {'R', S_GROUP_6, '6'},
    SoundexCharStruct {'S', S_GROUP_2, '2'},
    SoundexCharStruct {'T', S_GROUP_3, '3'},
    SoundexCharStruct {'U', VOWEL, byte(VOWEL)},
    SoundexCharStruct {'V', S_GROUP_1, '1'},
    SoundexCharStruct {'W', IGNORE, byte(IGNORE)},
    SoundexCharStruct {'X', S_GROUP_2, '2'},
    SoundexCharStruct {'Y', VOWEL, byte(VOWEL)},
    SoundexCharStruct {'Z', S_GROUP_2, '2'} };




// group defs for refined soundex
const R_GROUP_1     int = 49;
const R_GROUP_2     int = 50;
const R_GROUP_3     int = 51;
const R_GROUP_4     int = 52;
const R_GROUP_5     int = 53;
const R_GROUP_6     int = 54;
const R_GROUP_7     int = 55;
const R_GROUP_8     int = 56;
const R_GROUP_9     int = 57;


var charToRefindCode = [26] SoundexCharStruct {
    SoundexCharStruct {'A', VOWEL, byte(VOWEL)},
    SoundexCharStruct {'B', R_GROUP_1, '1'},
    SoundexCharStruct {'C', R_GROUP_3, '3'},
    SoundexCharStruct {'D', R_GROUP_6, '6'},
    SoundexCharStruct {'E', VOWEL, byte(VOWEL)},
    SoundexCharStruct {'F', R_GROUP_2, '2'},
    SoundexCharStruct {'G', R_GROUP_4, '4'},
    SoundexCharStruct {'H', IGNORE, byte(IGNORE)},
    SoundexCharStruct {'I', VOWEL, byte(VOWEL)},
    SoundexCharStruct {'J', R_GROUP_4, '4'},
    SoundexCharStruct {'K', R_GROUP_3, '3'},
    SoundexCharStruct {'L', R_GROUP_7, '7'},
    SoundexCharStruct {'M', R_GROUP_8, '8'},
    SoundexCharStruct {'N', R_GROUP_8, '8'},
    SoundexCharStruct {'O', VOWEL, byte(VOWEL)},
    SoundexCharStruct {'P', R_GROUP_1, '1'},
    SoundexCharStruct {'Q', R_GROUP_5, '5'},
    SoundexCharStruct {'R', R_GROUP_9, '9'},
    SoundexCharStruct {'S', R_GROUP_3, '3'},
    SoundexCharStruct {'T', R_GROUP_6, '6'},
    SoundexCharStruct {'U', VOWEL, byte(VOWEL)},
    SoundexCharStruct {'V', R_GROUP_2, '2'},
    SoundexCharStruct {'W', IGNORE, byte(IGNORE)},
    SoundexCharStruct {'X', R_GROUP_5, '5'},
    SoundexCharStruct {'Y', VOWEL, byte(VOWEL)},
    SoundexCharStruct {'Z', R_GROUP_5, '5'} };



func RollCall () string {
    return ("present");
}

func StandardSoundexCharCode (char byte) SoundexCharStruct {
    return charToStandardCode [char-'A'];
}

func RefinedSoundexCharCode (char byte) SoundexCharStruct {
    return charToRefindCode [char-'A'];
}
