#
# Makefile
#

CIRCLEHOME = ../../

OBJS	= main.o kernel.o

LIBS	= $(CIRCLEHOME)/lib/libcircle.a

INCLUDES = -I$(CIRCLEHOME)/include

SOURCES += \
	main.cpp \
	kernel.cpp

include Rules.mk

-include $(DEPS)
