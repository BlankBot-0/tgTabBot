//Tab.g4
grammar Tab;
// Tokens
INTEGER: [0-9]+;
DURMOD: [._]+;
NOTELETTER: [A-G]|[a-g]'#'?;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start: bpm tuning durationGroup* EOF;

bpm : 'BPM' INTEGER;

tuning: tuningImpl tuningImpl tuningImpl tuningImpl tuningImpl tuningImpl;
tuningImpl : NOTELETTER INTEGER;

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
