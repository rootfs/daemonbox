# Daemonbox
Daemonbox is a daemon process with inspiration from busybox

## Build

```console
  $ make godeps
  $ make build
```
## Start daemonbox

```console
  SERVICE_PORT=5000 ./daemonbox
```

## Test daemonbox

On a different console:

```console
 $curl http://localhost:5000/cmd/mount/`echo -n "-t gluster localhost:/test /mnt" |base64 `
```