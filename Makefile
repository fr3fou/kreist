MODULE_NAME=car
OUTPUT_DIR=car

ifeq ($(OS), Windows_NT)
	EXECUTABLE=$(MODULE_NAME).exe
else
	EXECUTABLE=$(MODULE_NAME)
endif

ifneq ("$(wildcard $(OUTPUT_DIR))","")
    OUTPUT_EXISTS = 1
endif

all: clean static-build bundle

static-build:
	go build -ldflags '-extldflags "-static"' .

bundle:
	mkdir $(OUTPUT_DIR)
	mv $(EXECUTABLE) $(OUTPUT_DIR)
	cp assets $(OUTPUT_DIR) -r
	cp levels $(OUTPUT_DIR) -r

clean:
ifeq ($(OUTPUT_EXISTS), 1)
	rm -r $(OUTPUT_DIR)
endif
	
