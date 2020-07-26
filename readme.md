# C Code Hash

Generate function Hash for C

> https://feilong.home.blog/2020/07/26/c%e8%af%ad%e8%a8%80%e5%87%bd%e6%95%b0%e4%bd%93%e5%8f%98%e5%8c%96%e7%82%b9%e6%8a%bd%e5%87%ba/

## example

```
> make test
pushd src && go build -o ../build/CCodeHash .
~/workspace/workspace/workspace/c-code-hash/src ~/workspace/workspace/workspace/c-code-hash
========== TEST START =================
./build/CCodeHash ./test/bt.c
voidinsert(node**tree,node*item)        b48f74e6680a2e10cc17a7d36b6bb668
voidprintout(node*tree) 9b6ccf989b6aed52cb359e877a274a6f
voidmain()      d0d6ad0c4abbcf1502233c5656c69583
./build/CCodeHash ./test/bt_comment.c
voidinsert(node**tree,node*item)        b48f74e6680a2e10cc17a7d36b6bb668
voidprintout(node*tree) 9b6ccf989b6aed52cb359e877a274a6f
voidmain()      d0d6ad0c4abbcf1502233c5656c69583
./build/CCodeHash ./test/bt_codechange.c
voidinsert(node**tree,node*item)        e75391829bc8c2ab8bb63e33f9cc6d34
voidprintout(node*tree) 9b6ccf989b6aed52cb359e877a274a6f
voidmain()      d0d6ad0c4abbcf1502233c5656c69583
```
