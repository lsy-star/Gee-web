module day7-panic-recover

go 1.17

require gee v0.0.0

require middlewares v0.0.0

replace (
	gee => ./gee
	middlewares => ./middlewares
)
