module github.com/fasthttp/session/v2

go 1.12

require (
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	github.com/go-redis/redis/v8 v8.3.4 // Don't upgrade to keep go1.12 compatibility
	github.com/go-sql-driver/mysql v1.5.0
	github.com/lib/pq v1.9.0
	github.com/mattn/go-sqlite3 v1.14.6
	github.com/savsgio/dictpool v0.0.0-20210104112344-3a2df09158c0
	github.com/savsgio/gotils v0.0.0-20210104112019-96a5e1e9898f
	github.com/valyala/bytebufferpool v1.0.0
	github.com/valyala/fasthttp v1.19.0
)
