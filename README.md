# notifissue

## Usage
### Install
```
$ go get -u github.com/nemusou/notifissue
```

### for Fish
add `config.fish` or your config file.
```
# notifissue
set -x NOTIFISSUE_PATH "YOUR_NOTIFISSUE_PATH"
function _notifissue
  for i in ($NOTIFISSUE_PATH -u=nemusou)
    echo $i
  end
end
_notifissue
```
