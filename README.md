# notifissue

![notifissue run](https://user-images.githubusercontent.com/18340344/70629082-b3fd7780-1c6c-11ea-9b2b-f6773e8a2844.png)

## Usage
### Install
```
$ go get -u github.com/nemusou/notifissue
```

### for Fish
add `config.fish` or any config file.
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
