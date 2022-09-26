TEST_CASE ?= 03
MODE := unordered

CLEAN_TARGET ?= \
	input/input02.txt input/input03.txt \
	output/output02.txt output/output02.txt

.PHONY: download
download:
	wget -cO input/input02.txt 'https://raw.githubusercontent.com/benhoyt/countwords/master/kjvbible.txt'
	for _ in $$(seq 10); do cat input/input02.txt; done > input/input03.txt
	wget -cO output/output03.txt 'https://raw.githubusercontent.com/benhoyt/countwords/master/output.txt'
	sed "s/0$$//" output/output03.txt > output/output02.txt
