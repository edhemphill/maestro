
LDFLAGS ?= -lpthread -pthread 
# TWSOVERSION is the compiler version...
# see http://rute.2038bug.com/node26.html.gz

CXX ?= g++ -g -O0 -fPIC -std=c++11
CC ?= gcc -g -O0 -fPIC
AR ?= ar

CXXFLAGS= -fPIC -std=c++11

ARCH ?=x86 
#ARCH=armel
SYSCALLS= syscalls-$(ARCH).c

ALLOBJS= $($<:%.cpp=%.o)

DEBUG_OPTIONS=-rdynamic -D_TW_TASK_DEBUG_THREADS_ -DLOGGER_HEAVY_DEBUG
#-D_TW_BUFBLK_DEBUG_STACK_
CFLAGS= $(DEBUG_OPTIONS) $(GLIBCFLAG) -I./include  -D__DEBUG   -fPIC 

DEBUG_CFLAGS= -g -DERRCMN_DEBUG_BUILD

ROOT_DIR=.
OUTPUT_DIR=.


EXTRA_TARGET=

CFLAGS+= -fPIC $(DEBUG_CFLAGS) 

GLIBCFLAG=-D_USING_GLIBC_
LD_TEST_FLAGS= -lgtest


OBJS= $(SRCS_CPP:%.cc=$(OUTPUT_DIR)/%.o) $(SRCS_C:%.c=$(OUTPUT_DIR)/%.o)
OBJS_NAMES= $(SRCS_CPP:%.cc=$%.o) $(SRCS_C:%.c=%.o)

##tw_sparsehash.h

## The -fPIC option tells gcc to create position 
## independant code which is necessary for shared libraries. Note also, 
## that the object file created for the static library will be 
## overwritten. That's not bad, however, because we have a static 
## library that already contains the needed object file.

$(OUTPUT_DIR)/%.o: %.cc
	$(CXX) $(CXXFLAGS) $(CFLAGS) -c $< -o $@

$(OUTPUT_DIR)/%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

process_utils.o: processes/process_utils.c
	$(CC) $(CFLAGS) -c $< -o $@

native.a-debug: CFLAGS += -DDEBUG_BINDINGS
native.a-debug: process_utils.o
	$(AR) rcs bindings.a $^

native.a: process_utils.o	
	$(AR) rcs $@ $^ 

clean: 
	-rm -rf $(OUTPUT_DIR)/*.o $(OUTPUT_DIR)/*.obj $(OUTPUT_DIR)/*.rpo $(OUTPUT_DIR)/*.idb $(OUTPUT_DIR)/*.lib $(OUTPUT_DIR)/*.exe $(OUTPUT_DIR)/*.a $(OUTPUT_DIR)/*~ $(OUTPUT_DIR)/core
	-rm -rf Debug
	-rm -f $(TWSOLIBNAME) $(TWSONAME) $(TWSOVERSION)
	-rm -f bindings.a
# DO NOT DELETE
