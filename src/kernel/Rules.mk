#ifeq ($(strip $(CIRCLEHOME)),)
CIRCLEHOME = ../../
#endif

include $(CIRCLEHOME)/Rules.mk
