//Tab.g4
grammar Tab;

//// Tokens
//INTEGER: [0-9]+;
//STR: [1-9];
//DURMOD: [._]+;
//WHITESPACE: [ \r\n\t]+ -> skip;
//
//// Rules
//start: durationGroups EOF;
//
//durationGroups
//    : durationGroup durationGroups
//    | durationGroup
//    ;
//
//durationGroup
//    : duration '(' noteGroup ')'
//    ;
//
//duration
//    : INTEGER DURMOD INTEGER      #DurDescribed
//    | INTEGER                     #DurDefault
//    ;
//
//noteGroup
//    : note noteGroup
//    | '[' chord ']' noteGroup
//    | note
//    ;
//
//chord
//    : noteGroup #SimpleChord
//    ;
//
//note
//    : STR INTEGER               #PlayFret
//    | 'p'                       #Pause
//    | 'b(' note ',' INTEGER ')' #Bend
//    ;

// Tokens
INTEGER: [0-9]+;
DURMOD: [._]+;
NOTELETTER: [A-G]|[a-g]'#'?;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start: bpm tuning durationGroup* EOF;

bpm : 'BPM' INTEGER;

tuning: (NOTELETTER INTEGER){6};

durationGroup
    : duration '(' entry* ')'
    ;

duration
    : INTEGER DURMOD INTEGER*     #DurDescribed
    | INTEGER                     #DurDefault
    ;

entry
    : INTEGER                      #PlayFret
    | '[' INTEGER* ']'             #SimpleChord
    | 'p'                          #Pause
    | 'b(' INTEGER ',' INTEGER ')' #Bend
    ;
