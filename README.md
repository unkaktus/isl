## isl

`isl` (Insecure Socket Layer) is a tool to allow talking with devices that do not support modern TLS versions and are stuck with SSLv3. For example, that's the case for older pre-POODLE DELL servers and their iDRAC interface.

### Usage
Install:
```shell
go install github.com/unkaktus/isl@latest
```

Launch:
```shell
isl -t 192.168.0.20
```

Go ahead to https://localhost and have fun upgrading firmware to something modern.