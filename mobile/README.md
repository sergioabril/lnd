# STATUS
It's fully working. Bindings have been generated, and a simple `./build_mobile.go` should produce a fully functional *framework* for iOS.

# HOW I MADE IT
#1 Pulled Master `lightningnetwork/lnd`

#2 Forked into my own branch mobile-build, to start modifying it.
`https://github.com/sergioabril/lnd/tree/mobile-build-tool`

#3 Modified *bindings.go* according Youtube talk.

https://www.youtube.com/watch?v=IIxCYrRiP78
What is not said there, is that you need to edit the `lightningnetwork/lnd/lnd.go` to include a two args Main() method that is called from *bindings.go*. Otherwise, it will fail to build an iOS module.

#4 This is how I Regenerated bindings to have the latest lnd calls:
1. Downloaded and installed https://github.com/halseth/promobile
2. *shame* > Manually modified *promobile* to set listeners inside for every service, because I wasn't able to feed them thru args:
```
	listeners["walletunlocker"] = "bufWalletUnlockerLis"
	listeners["lightning"] = "bufLightningLis"
```
3. `go install` (inside promobile folder) to rebuild binaries of *promobile* with my change.
4. ran *./gen_bindings.sh*
5. Made sure *bufWalletUnlockerLis* and *bufLightningLis* was declared and used on bindings.go, lightning_api_generated.go, and walletunlocker_api_generated.go.
5. It should be ready for step 5.

#5 How I Built it
just cd github.com/lightningnetwork/lnd && make ios
(A dep init & dep ensure could be needed beforehand)

#6 Notes, problems and questions:

1. Even after this, XCODE throws errors for these grabbed from lightning-app mobile repo (I had to remove them):
```
                    @"Status" : ^(NSData* bytes, NativeCallback* cb) { LndmobileStatus(bytes, cb); },
                       @"SetScores" : ^(NSData* bytes, NativeCallback* cb) { LndmobileSetScores(bytes, cb); },
                       @"QueryScores" : ^(NSData* bytes, NativeCallback* cb) { LndmobileQueryScores(bytes, cb); },
                       @"ModifyStatus" : ^(NSData* bytes, NativeCallback* cb) { LndmobileModifyStatus(bytes, cb); },
```