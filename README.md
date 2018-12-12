# gocron
gocron is wrapper of launchd.
Support target is crond.

# command
## display cron for current user 
```bash
$ gocron -l
PID     Status  Label
-1      0       com.trendmicro.itis.uninstaller

```

## display cron for all user 
```bash
$ gocron -l -a
PID     Status  Label
-1      0       com.apple.safaridavclient
-1      0       com.apple.securityuploadd
-1      0       com.apple.AddressBook.abd
...more
```