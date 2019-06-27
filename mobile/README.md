#Original files for mobile folder:
https://github.com/sergioabril/lnd/tree/mobile-build-tool

#Modified according Youtube talk: 
https://www.youtube.com/watch?v=IIxCYrRiP78

#How To build:
just cd github.com/lightningnetwork/lnd &&& make ios

#Notes, problems and questions:
- I'm not using build_mobile.sh
- looks like there is a double/triple circular reference during build process.

- If I don't empty the Lnd.main() call on mobile/bindings.go, an error is thrown on build:
```
# github.com/lightningnetwork/lnd/mobile
mobile/bindings.go:26:21: too many arguments in call to lnd.Main
	have (*bufconn.Listener, *bufconn.Listener)
	want ()
```

- On the resulting framework, there are *things missing*, like *protocol LndmobileRecvStream*, and four gRPC methods.

- Makefile make ios is clumsy btw.