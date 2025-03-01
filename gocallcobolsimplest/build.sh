cob2 -q64 -vV -comprc_ok=8 -q"list,LP(64)" -qexportall -c ccc.cbl
cob2 -Wl,dll,lp64 -o XDDLL ccc.o
go build
