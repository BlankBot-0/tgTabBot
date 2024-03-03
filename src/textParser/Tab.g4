//Tab.g4
grammar Tab;

// Tokens
STR: [1-9];
INTEGER: [0-9]+;
DURMOD: [._]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start: durationGroups EOF;

durationGroups
    : durationGroup durationGroups
    | durationGroup
    ;

durationGroup
    : duration '(' noteGroup ')'
    ;

duration
    : INTEGER DURMOD INTEGER      #DurDescribed
    | INTEGER                     #DurDefault
    ;

noteGroup
    : note noteGroup
    | '[' chord ']' noteGroup
    | note
    ;

chord
    : noteGroup #SimpleChord
    ;

note
    : STR INTEGER               #PlayFret
    | 'p'                       #Pause
    | 'b(' note ',' INTEGER ')' #Bend
    ;