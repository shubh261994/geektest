# Gotham

This is an app written in Go

Just run the following

## SETUP

1) Install Golang version >= 1.14

```
apt install golang-go
```

2) Verify version is >= 1.14

```
go version
```

3) Add the following to profile file ( the path where go is installed)

```
vi ~/.bash_profile
export PATH=$PATH:/usr/local/go/bin
```

4) Make a go folder inside your home directory. Then run the following commands:

```
cd go
mkdir bin src
cd src
git clone {repo-url}
export GOPATH={path of go directory that you created}
go get
go install
```

5) Update the configuration file for go -

```
cp config/config-sample config/config.go
```

7) To build and install the app (every time changes are made to the code inside the geektest repo):

```
go install && go run main.go
```

8) Run the server using the following command

```
go run main.go
```

This will start the Service on port 8080.