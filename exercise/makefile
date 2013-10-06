GCCGO=gccgo
TARGET=find
OBJS=find.o

all:${TARGET}
${TARGET}:${OBJS}
	${GCCGO} -o $@ $<
clean:
	rm -rf *.o ${TARGET}
%.o:%.go
	${GCCGO} -c $<
