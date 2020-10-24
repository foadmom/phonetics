package utilities;

// errors Wrapf function

type ErrorStruct struct {
    Code	int;
    Message	string;
}

type ErrorTrace struct {
    Error	*ErrorStruct;
    Trace   [] ErrorStruct;
}

func AddErrorToTrace (tree *ErrorTrace, err *ErrorStruct) {
    if (tree.Error == nil) {
        tree.Error = err;
    } else {
        tree.Trace = append (tree.Trace, *err);
    }
     
}