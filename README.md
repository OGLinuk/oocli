# (O)nion (O)mega2 (CLI)

```
The Onion Omega2 runs OpenWRT, which allows the device to act as an access
        point and bridge WiFi to connected devices. Since it runs BusyBox (Linux), it can 
        run a small Golang cross-compiled fileserver. Any device connected to the Onion 
        Omega2 will have access to the fileserver via 192.168.3.1:9001. This CLI is used 
        to easily interact with that fileserver. The media files are also available to be 
        viewed, and all files to be downloaded through firefox.

Usage:
  oocli [command]

Available Commands:
  delete      Delete arg[0] from the Onion Omega2 fileserver
  download    Download arg[0] to arg[1] from the Onion Omega2 fileserver
  help        Help about any command
  initfs      Initialize the fileserver on the Onion Omega2
  killall     Execute killall arg[0]
  move        Move arg[0] to arg[1] with arg[2] as the name on the Onion Omega2 fileserver
  ping        Ping the Onion Omega2
  psgrep      Execute 'ps | grep arg[0]' on the Onion Omega2
  restart     Restart the Onion Omega2
  udug        Update the Onion Omega2
  upload      Upload arg[0] to arg[1] on the Onion Omega2 fileserver
  view        View arg[0] from the Onion Omega2 fileserver on the default browser
```