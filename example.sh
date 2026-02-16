#!/bin/sh

ijson(){
	jq -c -n '{
		helo: "wrld",
		wrld: 42,
		noti: "42",
		real: 42.195,
		imag: false,
		fake: true,
		even: [2,4,6,8],
		prompt: "helo\nwrld",
		dict: {
			store: 0,
			deflate: 8,
		},
	}'
}

ijson |
	wasmtime \
		run \
		./jsonmap2yaml.wasm \
		-indent 8 \
		-indent-sequence \
		-auto-int \
		-literal-multiline \
		-single-quote |
	bat --language=yaml
