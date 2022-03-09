module github.com/eyedeekay/goSam

require (
	github.com/getlantern/fdcount v0.0.0-20210503151800-5decd65b3731 // indirect
	github.com/getlantern/go-socks5 v0.0.0-20171114193258-79d4dd3e2db5
	github.com/getlantern/golog v0.0.0-20201105130739-9586b8bde3a9 // indirect
	github.com/getlantern/mockconn v0.0.0-20200818071412-cb30d065a848 // indirect
	github.com/getlantern/mtime v0.0.0-20200417132445-23682092d1f7 // indirect
	github.com/getlantern/netx v0.0.0-20190110220209-9912de6f94fd // indirect
	github.com/getlantern/ops v0.0.0-20200403153110-8476b16edcd6 // indirect
)

//replace github.com/eyedeekay/gosam v0.1.1-0.20190814195658-27e786578944 => github.com/eyedeekay/goSam ./

replace github.com/eyedeekay/gosam => ./

replace github.com/eyedeekay/goSam => ./

go 1.13
