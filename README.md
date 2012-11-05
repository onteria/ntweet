ntweet
======

This is the source code of GoNitori, a Twitter client I'm writing out of boredom and the desire to pickup some Google Go while I'm at it.

Requirements
============

This is written with Google Go 1.0.3. I update the go hg release branch daily so it will keep up with whatever the stable version is and any changes along with it. For oauth it requires the [go-oauth](http://github.com/garyburd/go-oauth/oauth) library to handle oAuth authentication. You can easily grab it by setting $GOPATH and running:

```
go get github.com/garyburd/go-oauth/oauth
```

License
=======

GoNitori is licensed under the [WTFPL license](http://sam.zoy.org/wtfpl/). 

Support
=======

As this is mostly a hobby project, don't expect me to do much as far as pull requests or overal support is concerned. Up front and hones if it's not something I want this program to do I won't pull it in. It's more of a code dump for others to look at and see how to work with the Twitter API using Google Go and a more recent version of the API with actual oAuth. 