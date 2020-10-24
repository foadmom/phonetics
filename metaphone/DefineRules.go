package metaphone;

import "fmt";
import "bytes";
import "strconv";

type mapFunctions struct {
	char      byte;
	transFunc func(pp *PhoneticPhrase) (rc int);
}

var vowels = []byte{'A', 'E', 'I', 'O', 'U'};

// 48-122;
var translationTable = []mapFunctions{
	mapFunctions{0, invalidMap},
	mapFunctions{1, invalidMap},
	mapFunctions{2, invalidMap},
	mapFunctions{3, invalidMap},
	mapFunctions{4, invalidMap},
	mapFunctions{5, invalidMap},
	mapFunctions{6, invalidMap},
	mapFunctions{7, invalidMap},
	mapFunctions{8, invalidMap},
	mapFunctions{9, invalidMap},
	mapFunctions{10, invalidMap},
	mapFunctions{11, invalidMap},
	mapFunctions{12, invalidMap},
	mapFunctions{13, invalidMap},
	mapFunctions{14, invalidMap},
	mapFunctions{15, invalidMap},
	mapFunctions{16, invalidMap},
	mapFunctions{17, invalidMap},
	mapFunctions{18, invalidMap},
	mapFunctions{19, invalidMap},
	mapFunctions{20, invalidMap},
	mapFunctions{21, invalidMap},
	mapFunctions{22, invalidMap},
	mapFunctions{23, invalidMap},
	mapFunctions{24, invalidMap},
	mapFunctions{25, invalidMap},
	mapFunctions{26, invalidMap},
	mapFunctions{27, invalidMap},
	mapFunctions{28, invalidMap},
	mapFunctions{29, invalidMap},
	mapFunctions{30, invalidMap},
	mapFunctions{31, invalidMap},
	mapFunctions{32, invalidMap},
	mapFunctions{33, invalidMap},
	mapFunctions{34, invalidMap},
	mapFunctions{35, invalidMap},
	mapFunctions{36, invalidMap},
	mapFunctions{37, invalidMap},
	mapFunctions{38, invalidMap},
	mapFunctions{39, invalidMap},
	mapFunctions{40, invalidMap},
	mapFunctions{41, invalidMap},
	mapFunctions{42, invalidMap},
	mapFunctions{43, invalidMap},
	mapFunctions{44, invalidMap},
	mapFunctions{45, invalidMap},
	mapFunctions{46, invalidMap},
	mapFunctions{47, invalidMap},
	mapFunctions{48, invalidMap},
	mapFunctions{49, invalidMap},
	mapFunctions{50, invalidMap},
	mapFunctions{51, invalidMap},
	mapFunctions{52, invalidMap},
	mapFunctions{53, invalidMap},
	mapFunctions{54, invalidMap},
	mapFunctions{55, invalidMap},
	mapFunctions{56, invalidMap},
	mapFunctions{57, invalidMap},
	mapFunctions{58, invalidMap},
	mapFunctions{59, invalidMap},
	mapFunctions{60, invalidMap},
	mapFunctions{61, invalidMap},
	mapFunctions{62, invalidMap},
	mapFunctions{63, invalidMap},
	mapFunctions{64, invalidMap},
	mapFunctions{'A', translate_vowels}, // A;
	mapFunctions{'B', translate_B},
	mapFunctions{'C', translate_C},
	mapFunctions{'D', translate_D},
	mapFunctions{'E', translate_vowels}, // E;
	mapFunctions{'F', mapToSameChar},
	mapFunctions{'G', translate_G},      // G;
	mapFunctions{'H', translate_H},      // H;
	mapFunctions{'I', translate_vowels}, // I;
	mapFunctions{'J', mapToSameChar},
	mapFunctions{'K', mapToSameChar},
	mapFunctions{'L', mapToSameChar},
	mapFunctions{'M', mapToSameChar},
	mapFunctions{'N', mapToSameChar},
	mapFunctions{'O', translate_vowels}, // O;
	mapFunctions{'P', translate_P},
	mapFunctions{'Q', translate_Q},
	mapFunctions{'R', mapToSameChar},
	mapFunctions{'S', translate_S},
	mapFunctions{'T', translate_T},
	mapFunctions{'U', translate_vowels}, // U;
	mapFunctions{'V', translate_V},
	mapFunctions{'W', translate_W},
	mapFunctions{'X', translate_X},
	mapFunctions{'Y', translate_Y},
	mapFunctions{'Z', translate_Z},
	mapFunctions{91, invalidMap},
	mapFunctions{92, invalidMap},
	mapFunctions{93, invalidMap},
	mapFunctions{94, invalidMap},
	mapFunctions{95, invalidMap},
	mapFunctions{96, invalidMap},
	mapFunctions{97, invalidMap},
	mapFunctions{98, invalidMap},
	mapFunctions{99, invalidMap},
	mapFunctions{100, invalidMap},
	mapFunctions{101, invalidMap},
	mapFunctions{102, invalidMap},
	mapFunctions{103, invalidMap},
	mapFunctions{104, invalidMap},
	mapFunctions{105, invalidMap},
	mapFunctions{106, invalidMap},
	mapFunctions{107, invalidMap},
	mapFunctions{108, invalidMap},
	mapFunctions{109, invalidMap},
	mapFunctions{110, invalidMap},
	mapFunctions{111, invalidMap},
	mapFunctions{112, invalidMap},
	mapFunctions{113, invalidMap},
	mapFunctions{114, invalidMap},
	mapFunctions{115, invalidMap},
	mapFunctions{116, invalidMap},
	mapFunctions{117, invalidMap},
	mapFunctions{118, invalidMap},
	mapFunctions{119, invalidMap},
	mapFunctions{120, invalidMap},
	mapFunctions{121, invalidMap},
	mapFunctions{122, invalidMap},
	mapFunctions{123, invalidMap},
	mapFunctions{124, invalidMap},
	mapFunctions{125, invalidMap},
	mapFunctions{126, invalidMap},
	mapFunctions{127, invalidMap},
	mapFunctions{128, invalidMap},
}

func RollCall() string {
	return ("present");
}

// ==========================================================;
// returns true if the char passed in is a vowel;
// ==========================================================;
func isVowel(char byte) bool {
	return (bytes.ContainsRune(vowels, rune(char)));
}

func invalidMap(pp *PhoneticPhrase) (rc int) {
	//    fmt.Printf ("Invalid character mapping. phrase=%s, index=%d\n", string(phrase), pp.pIndex);;
	fmt.Println("Invalid character mapping. index=" +
		strconv.Itoa(pp.pIndex) + "  char=" +
		string(pp.bytePhrase[pp.pIndex]) + " phrase=" + string(pp.bytePhrase));
	pp.kIndex = pp.kIndex + 1;
	pp.pIndex = pp.pIndex + 1;
	return -1;
}

func TranslateChar(pp *PhoneticPhrase) (rc int) {
	return (translationTable[pp.bytePhrase[pp.pIndex]].transFunc(pp));

}

// ==========================================================;
// there are certain combination of chars at the start of;
// a phrase that can be silent.;
// eg KNife  is nife, k is silent;
//    WRite  is rite, w is silent;
// ==========================================================;
var ignoreTheFirstCharArray = [][]byte{
	{'K', 'N'},
	{'G', 'N'},
	{'P', 'N'},
	{'A', 'E'},
	{'W', 'R'},
};
var wh = []byte{'W', 'H'}

// ==========================================================;
// Checks for the start of the phrase for silent chars;
// based on the rule in ignoreTheFirstCharArray;
// ==========================================================;
func silentCharsAtStart(pp *PhoneticPhrase) (rc int) {
	for _, b := range ignoreTheFirstCharArray {
		if bytes.Compare(pp.bytePhrase[:2], b) == 0 {
			pp.key[0] = pp.bytePhrase[1];
			pp.kIndex = 1;
			pp.pIndex = 2;
			return 0;
		}
	}
	if bytes.Compare(pp.bytePhrase[:2], wh) == 0 {
		pp.key[0] = 'W';
		pp.kIndex = 1;
		pp.pIndex = 2;
	}
	return 0;
}

func translateTheStart(pp *PhoneticPhrase) (rc int) {
	silentCharsAtStart(pp);
	if pp.bytePhrase[0] == 'X' {
		pp.key[0] = 'S';
		pp.kIndex = 1;
		pp.pIndex = 1;
	}
	return 1;
}

// ==========================================================;
// this one translate byte 0 because we introduced it;
// replacing the duplicate chars.;
// just ignore the char and move on to the next char;
// ==========================================================;
func translate_ignore(pp *PhoneticPhrase) (rc int) {
	pp.pIndex++;
	return 1;
}

// ==========================================================;
// how the vowels are treated;
// ==========================================================;
func translate_vowels(pp *PhoneticPhrase) (rc int) {
	return (translate_ignore(pp));
}

// ==========================================================;
// map the char to the same char;
// ==========================================================;
func mapToSameChar(pp *PhoneticPhrase) (rc int) {
	pp.key[pp.kIndex] = pp.bytePhrase[pp.pIndex];
	pp.kIndex++;
	pp.pIndex++;
	return 1;
}

// ==========================================================;
// translate to this char. eg p -> F;
// ==========================================================;
func mapToChar(pp *PhoneticPhrase, mappedChar byte) (rc int) {
	pp.key[pp.kIndex] = mappedChar;
	pp.kIndex++;
	pp.pIndex++;
	return 0;
}

func mapToA(pp *PhoneticPhrase) (rc int) {
	pp.key[pp.kIndex] = 'A';
	pp.kIndex++;
	pp.pIndex++;
	return 1;
}

func translate_B(pp *PhoneticPhrase) (rc int) {
	if (pp.pIndex == pp.stLength-1) && (pp.bytePhrase[pp.pIndex-1] == 'M') {
		// ignore B if it is at the end and after M like dumb;
	} else {
		pp.key[pp.kIndex] = 'B';
		pp.kIndex = pp.kIndex + 1;
	}
	pp.pIndex++;
	return 0;
}

func translate_C(pp *PhoneticPhrase) (rc int) {
	var _index int = pp.pIndex;
    if pp.bytePhrase[_index+1] == 'I' && pp.bytePhrase[_index+2] == 'A' { // replace cia with x
		pp.key[pp.kIndex] = 'X';
		pp.kIndex = pp.kIndex + 1;
		pp.pIndex = pp.pIndex + 3;
    } else if pp.bytePhrase[_index+1] == 'H' {
		pp.key[pp.kIndex] = 'X';
    	pp.kIndex = pp.kIndex + 1;
		pp.pIndex = pp.pIndex + 2;
	} else if pp.bytePhrase[_index+1] == 'I' || pp.bytePhrase[_index+1] == 'E' || pp.bytePhrase[_index+1] == 'Y' {
	    // if C is followed by I, E or Y then replace with S
			pp.key[pp.kIndex] = 'S';
			pp.kIndex = pp.kIndex + 1;
			pp.pIndex = pp.pIndex + 2;
	} else if pp.bytePhrase[_index+1] == 'K' {  // if CK then replace with K
		pp.key[pp.kIndex] = 'K';
		pp.kIndex++;
		pp.pIndex++;
		pp.pIndex++;
	} else {
		pp.key[pp.kIndex] = 'K';
		pp.kIndex++;
		pp.pIndex++;
	}

	return 0;
}

// ==========================================================;
// D turns to J if followed by G otherwise T;
// eg DG -> J  otherwise D -> T;
// ==========================================================;
func translate_D(pp *PhoneticPhrase) (rc int) {
	if pp.bytePhrase[pp.pIndex+1] == 'G' {
		// check to see if D is not the last char and next char is G;
		pp.key[pp.kIndex] = 'J';
		pp.kIndex++;
		pp.pIndex = pp.pIndex + 2;
	} else {
		pp.key[pp.kIndex] = 'T';
		pp.kIndex++;
		pp.pIndex++;
	}
	return 0;
}

// ==========================================================;
//   G ---->        SILENT if in "-gh-" and not at end or before a vowel;
//                            in "-gn" or "-gned";
//                            in "-dge-" etc., as in above rule;
//           J      if before "i", or "e", or "y" if not double "gg";
//           K      otherwise;
//;
// ==========================================================;
func translate_G(pp *PhoneticPhrase) (rc int) {
	var nextChar = pp.bytePhrase[pp.pIndex+1];
	var nextCharPlusOne = pp.bytePhrase[pp.pIndex+2];
	var pLen = pp.stLength;

	if nextChar == 'H' { // if g is followed by h;
		if pp.pIndex < pLen-2 { // and it is not at the end;
			if isVowel(nextCharPlusOne) == false { // and h is not followed by a vowel;
				pp.pIndex++ // then g is silent;
			}
		}
	} else if nextChar == 'N' {
		pp.pIndex++ // then g is silent;
	} else if nextChar == 'E' || nextChar == 'I' || nextChar == 'Y' {
		pp.key[pp.kIndex] = 'J';
		pp.pIndex++;
		pp.kIndex++;
	} else {
		pp.key[pp.kIndex] = 'K';
		pp.pIndex++;
		pp.kIndex++;
	}

	return 0;
}

// ==========================================================;
//   H ---->        SILENT if after vowel and no vowel follows;
//                         or after "-ch-", "-sh-", "-ph-", "-th-", "-gh-";
//           H      otherwise;
//;
// ==========================================================;
func translate_H(pp *PhoneticPhrase) (rc int) {
	var _nextChar = pp.bytePhrase[pp.pIndex+1];
	var _prevChar = pp.bytePhrase[pp.pIndex-1];

	if (isVowel(_prevChar)) && (!isVowel(_nextChar)) {
		pp.pIndex++;
	} else {
		mapToSameChar(pp);
	}

	return 0;
}

// ==========================================================;
//   P ----> F      if before "h";
//           P      otherwise;
// ==========================================================;
func translate_P(pp *PhoneticPhrase) (rc int) {
	var _nextChar = pp.bytePhrase[pp.pIndex+1];

	if _nextChar == 'H' {
		mapToChar(pp, 'F');
		pp.pIndex++;
	} else {
		mapToChar(pp, 'P');
	}

	return 0;
}

// ==========================================================;
// Q turns to K;
// ==========================================================;
func translate_Q(pp *PhoneticPhrase) (rc int) {
	mapToChar(pp, 'K');
	return 0;
}

// ==========================================================;
//   S ----> X      (sh) if before "h" or in "-sio-" or "-sia-";
//           S      otherwise;
// ==========================================================;
func translate_S(pp *PhoneticPhrase) (rc int) {
	var _nextChar = pp.bytePhrase[pp.pIndex+1];
	var _nextCharPlusOne = pp.bytePhrase[pp.pIndex+2];

	if (_nextChar == 'C') && (_nextCharPlusOne == 'H') {
		mapToChar(pp, 'S');
		mapToChar(pp, 'K');
		pp.pIndex++;
	} else if (_nextChar == 'I') && ((_nextCharPlusOne == 'O') || (_nextCharPlusOne == 'A')) {
		// sia or sio translates to X;
		mapToChar(pp, 'X');
		pp.pIndex = pp.pIndex + 2; // ignore the the IO or IA following S
	} else if (_nextChar == 'H') {
		mapToChar(pp, 'X');
		pp.pIndex++; // ignore the H following S
	} else {
	    mapToChar(pp, 'S');
	}

	return 0;
}

// ==========================================================;
//   T ----> X      (sh) if "-tia-" or "-tio-";
//           0      (th) if before "h";
//                  silent if in "-tch-";
//           T      otherwise;
// ==========================================================;
func translate_T(pp *PhoneticPhrase) (rc int) {
	var _nextChar = pp.bytePhrase[pp.pIndex+1];
	var _nextCharPlusOne = pp.bytePhrase[pp.pIndex+2];

	if (_nextChar == 'I') && ((_nextCharPlusOne == 'A') || (_nextCharPlusOne == 'O')) {
		// tia or tio translates to x;
		mapToChar(pp, 'X');
	} else if _nextChar == 'H' { // th translates to 0;
		mapToChar(pp, '0');
		pp.pIndex++;
	} else if _nextChar == 'C' && _nextCharPlusOne == 'H' { // t is silent in tch;
		pp.kIndex++;
		pp.pIndex++;
	} else {
		pp.key[pp.kIndex] = 'T';
		pp.kIndex++;
		pp.pIndex++;
	}

	return 0;
}

// ==========================================================;
// V turns to F;
// ==========================================================;
func translate_V(pp *PhoneticPhrase) (rc int) {
	pp.key[pp.kIndex] = 'F';
	pp.kIndex++;
	pp.pIndex++;
	return 0;
}

// ==========================================================;
//   W ---->        SILENT if not followed by a vowel;
//           W      if followed by a vowel;
// ==========================================================;
func translate_W(pp *PhoneticPhrase) (rc int) {
	if isVowel(pp.bytePhrase[pp.pIndex+1]) {
		mapToChar(pp, 'W');
	} else {
		pp.pIndex++ // silent so move on;
	}
	return 0;
}

// ==========================================================;
//   X ----> KS;
// ==========================================================;
func translate_X(pp *PhoneticPhrase) (rc int) {
	pp.key[pp.kIndex] = 'K';
	pp.key[pp.kIndex+1] = 'S';
	pp.kIndex = pp.kIndex + 2;
	pp.pIndex++;
	return 0;
}

// ==========================================================;
//   Y ---->        SILENT if not followed by a vowel;
//           Y      if followed by a vowel;
// ==========================================================;
func translate_Y(pp *PhoneticPhrase) (rc int) {
	if isVowel(pp.bytePhrase[pp.pIndex+1]) {
		mapToChar(pp, 'Y');
	} else {
		pp.pIndex++ // silent so move on;
	}
	return 0;
}

// ==========================================================;
// Z turns to S;
// ==========================================================;
func translate_Z(pp *PhoneticPhrase) (rc int) {
	pp.key[pp.kIndex] = 'S';
	pp.kIndex++;
	pp.pIndex++;
	return 1;
}
