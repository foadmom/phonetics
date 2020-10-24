package utilities

import (

)

//   x xx xxx
//   |  |  |
//   |  |  |_______ leaf code. details codes
//   |  |
//   |  |__________ branch code. 
//   |
//   |_____________ trunk code 1=system errors, 2=input errors
//
//
//
//

var	 OK                             ErrorStruct = ErrorStruct {0,       ""};
var	 NULL_STRING_NOT_ALLOWED		ErrorStruct = ErrorStruct {-201001, "input string is null"};
var  STRING_TOO_SHORT 				ErrorStruct = ErrorStruct {-201002, "input string less than minimum length"};

