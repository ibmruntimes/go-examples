       IDENTIFICATION DIVISION.
       PROGRAM-ID. SUBPROG.
       DATA DIVISION.
       WORKING-STORAGE SECTION.
      * int TESTAPI(const char *astring,
      *     int *intp, int mycount) 
      
       01  ASTRING.
           05 FILLER PIC X(5) VALUE 'BILLO'.
           05 FILLER PIC X    VALUE X'00'.
       01  ULTIMATEANSWER  PIC S9(8) COMP VALUE 42.
       01  MYCOUNT  PIC S9(8) COMP VALUE 777.
       01  RESULT       PIC S9(8) COMP VALUE 0.

       LINKAGE SECTION.
       01 FP     USAGE IS FUNCTION-POINTER.

       PROCEDURE DIVISION  USING BY VALUE FP.
           DISPLAY "in PROCEDURE division".
           DISPLAY "about to call function pointer"
           CALL FP USING
           BY CONTENT   ASTRING,
           BY REFERENCE ULTIMATEANSWER ,
           BY VALUE     MYCOUNT
           RETURNING RESULT.
           DISPLAY "RESULT " RESULT.
              GOBACK.
       END PROGRAM 'SUBPROG'.
