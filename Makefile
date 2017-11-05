BUILD := build
PROGRAM := sahaayak

.PHONY: all clean

all:
	mkdir -p $(BUILD)
	go build -o $(BUILD)/$(PROGRAM)

clean:
	rm -rf $(BUILD)
