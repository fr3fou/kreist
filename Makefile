MODULE_NAME=car
OUTPUT_DIR=car
ifeq ($(OS), Windows_NT)
	EXECUTABLE=$(MODULE_NAME).exe
else
	EXECUTABLE=$(MODULE_NAME)
endif

all: clean static-build bundle

static-build:
	go build -ldflags '-extldflags "-static"' .

bundle:
	mkdir $(OUTPUT_DIR)
	mv $(EXECUTABLE) $(OUTPUT_DIR)
	cp assets $(OUTPUT_DIR) -r

clean:
	rm -r $(OUTPUT_DIR)