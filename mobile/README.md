#1 Pull Master lightningnetwork/lnd

#2 Forked into my own branch mobile-build
`https://github.com/sergioabril/lnd/tree/mobile-build-tool`

#3 Modified *bindings.go* according Youtube talk: 
https://www.youtube.com/watch?v=IIxCYrRiP78

#4 Regenerated bindings:
1. Downloaded and installed https://github.com/halseth/promobile
2. Modified it to set listeners inside for every service, because I wasn't able to feed them:
```
	listeners["walletunlocker"] = "bufWalletUnlockerLis"
	listeners["lightning"] = "bufLightningLis"
```
3. go install (inside promobile folder)
4. ran *./gen_bindings.sh*
5. Manually added vars to binding files: *bufLightningLis* and *bufWallet...*

#5 How I Built it
just cd github.com/lightningnetwork/lnd &&& make ios, or ./build.sh
(A dep init & dep ensure could be needed)

#6 Notes, problems and questions:

1. I can't build it with `//if err := lnd.Main(walletUnlockerLis, lightningLis); err != nil {`
I need to remove the args inside lnd.Main()

2. Even after this, XCODE throws errors for this (I have to remove them):
```
                    @"Status" : ^(NSData* bytes, NativeCallback* cb) { LndmobileStatus(bytes, cb); },
                       @"SetScores" : ^(NSData* bytes, NativeCallback* cb) { LndmobileSetScores(bytes, cb); },
                       @"QueryScores" : ^(NSData* bytes, NativeCallback* cb) { LndmobileQueryScores(bytes, cb); },
                       @"ModifyStatus" : ^(NSData* bytes, NativeCallback* cb) { LndmobileModifyStatus(bytes, cb); },
```
