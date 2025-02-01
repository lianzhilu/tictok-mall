build-tiktok-mall:
	export MALL_CORE_CONFIG_DIR=$(shell pwd)/conf/
	go build -o output/ttmall ./cmd