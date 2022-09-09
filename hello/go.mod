module hello

go 1.15

replace example.com/greetings => ../greetings

require (
	bou.ke/monkey v1.0.2
	github.com/namsral/flag v1.7.4-pre
)
