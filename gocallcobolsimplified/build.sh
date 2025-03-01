cob2 -q64 -vV -comprc_ok=8 -q"list,LP(64)" -qexportall -c ccc2.cbl
cob2 -q64 -vV -comprc_ok=8 -q"list,LP(64)" -qexportall -c ccc1.cbl
cob2 -Wl,dll,lp64 -o XDDLL ccc1.o ccc2.o
go build
