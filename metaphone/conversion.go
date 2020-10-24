package metaphone

//import "strings";
import util "phonetics/utilities";
//import "errors";


type PhoneticPhrase struct {
	stPhrase   string
	bytePhrase []byte
	stLength   int    // original string length
	pIndex     int    // current index of the bytePhrase in the processing
	key        []byte // this is the generated metaphone code
	kIndex     int    // current index of the key being generated
//    errTrace   *util.ErrorTrace;
}

func GenerateCode(phrase string, code []byte) {
   	var _currentStruct PhoneticPhrase;
   	
	phrase = util.PrepareString (phrase);
	
	_err := util.InputStringValid (phrase);
	if (_err == nil ) {
    	var _stLen int = len (phrase);
    	if _stLen > 1 {
    	    // setup PhoneticPhrase structure 
    	    _currentStruct.stPhrase = phrase;
    		_currentStruct.bytePhrase = make([]byte, len(phrase)+5);
    		removeSuccessiveDuplicates([]byte(_currentStruct.stPhrase), &_currentStruct);
    		_currentStruct.key = make([]byte, len(phrase)+50);
    		_currentStruct.pIndex = 0;
    		_currentStruct.kIndex = 0;
//    		_currentStruct.errTrace = errTrace;
    
    		translateTheStart(&_currentStruct)
    		for _currentStruct.pIndex < _currentStruct.stLength {
    			TranslateChar(&_currentStruct)
    		}
    		copy(code, _currentStruct.key);
    	}
	} else {
//	    util.AddErrorToTrace (errTrace, _err);
//        errors.Wrapf (errors.New (util.STRING_TOO_SHORT.Message), "%d: ", util.STRING_TOO_SHORT.Code);
	}
}

// ==========================================================
// Func removeSuccessiveDuplicates Remove successive repeated chars unless it is C
// ==========================================================
func removeSuccessiveDuplicates(phrase []byte, pp *PhoneticPhrase) (length int) {
	var _pLen int
	_pLen = len(phrase)
	_nIndex := 1
	pp.bytePhrase[0] = phrase[0] // copy the first char
	for i := 1; i < _pLen; i++ {
		// starting from 2nd char, if it is the same as previous
		// then do not copy to the nPhrase
		if phrase[i] != phrase[i-1] || phrase[i] == 'C' {
			pp.bytePhrase[_nIndex] = phrase[i]
			_nIndex++
		}
		pp.stLength = _nIndex
	}
	return _nIndex
}
