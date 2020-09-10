Spotify Link Listener
===

A tiny piece of glue to allow easy association of Spotify tracks, artists or playlists with RFID tags on a [phoniebox](https://github.com/MiczFlor/RPi-Jukebox-RFID). Runs a simple web server that, when hit with a POST containing a single form param `link`, assigns the most recently seen RFID tag on the phoniebox with the spotify item belonging to the link.

Meant to be used with [this](https://www.icloud.com/shortcuts/fadc7f510eef480d90deb0ad93f6aa12) iOS Shortcut, which allows you to use the share button in the Spotify app to quickly assign anything without needing to use the phoniebox web UI.

Use the included systemd unit definition to run this as a daemon alongside the phoniebox web service:

```
sudo cp spotify-link-listener.service /etc/systemd/system/spotify-link-listener.service
sudo systemctl start spotify-link-listener.service
sudo systemctl enable spotify-link-listener.service
```

Building
---

You'll need to install golang on your pi if you don't already have it. Here's a script copied from a blog post:

```
export GOLANG="$(curl https://golang.org/dl/|grep armv6l|grep -v beta|head -1|awk -F\> {'print $3'}|awk -F\< {'print $1'})"
wget https://golang.org/dl/$GOLANG
sudo tar -C /usr/local -xzf $GOLANG
rm $GOLANG
unset GOLANG
```

Then run `go build .` in the root of this repo, which should create a binary called `spotify-link-listener`.