# TOC

* Base Build on helloempty
* Porting on helloempty

# Base Build on helloempty

Notes, code change, and build sequence

```
~/emgo/src/github.com/ziutek/emgo/egpath/src/stm32/examples/f4-discovery/helloempty
$ cat readme-local-porting-1-empty

this helloempty folder is copied from ...stm32...f4-discovery/helloworld and modified to contains only an empty for-loop:


$ diff -r ../helloworld/ ./
diff -r ../helloworld/main.go ./main.go
8,10c8,10
< import (
<       "fmt"
< )
---
> //import (
> //    "fmt"
> //)
14c14
<               fmt.Println("Hello world!")
---
>               //fmt.Println("Hello world!")


$ ../build.sh -v 1
pkg start 1     .
pkg start 2     runtime
pkg start 3     internal
pkg end   3  1  internal
pkg start 4     sync/atomic
pkg end   4  2  sync/atomic
pkg start 5     runtime/noos
pkg start 6     mem
pkg end   6  3  mem
pkg start 7     sync/fence
pkg end   7  4  sync/fence
pkg start 8     syscall
pkg start 9     arch/cortexm
pkg end   9  5  arch/cortexm
pkg start 10    arch/cortexm/scb
pkg start 11    bits
pkg end   11 6  bits
pkg start 12    mmio
pkg end   12 7  mmio
pkg end   12 8  arch/cortexm/scb
pkg end   12 9  syscall
pkg start 13    arch/cortexm/acc
pkg end   13 10 arch/cortexm/acc
pkg start 14    arch/cortexm/cmt
pkg end   14 11 arch/cortexm/cmt
pkg start 15    arch/cortexm/fpu
pkg end   15 12 arch/cortexm/fpu
pkg start 16    arch/cortexm/pft
pkg end   16 13 arch/cortexm/pft
pkg start 17    arch/cortexm/debug/itm
pkg end   17 14 arch/cortexm/debug/itm
pkg start 18    arch/cortexm/mpu
pkg end   18 15 arch/cortexm/mpu
pkg start 19    arch/cortexm/nvic
pkg end   19 16 arch/cortexm/nvic
pkg start 20    math/rand
pkg end   20 17 math/rand
pkg end   20 18 runtime/noos
pkg end   20 19 runtime
pkg end   20 20 .

$ arm-none-eabi-size.exe cortexm4f.elf
   text    data     bss     dec     hex filename
   7956     180     104    8240    2030 cortexm4f.elf
```

# Porting on helloempty

Notes, code change, and build sequence

```
~/emgo/src/github.com/ziutek/emgo/egpath/src/stm32/examples/f4-discovery/helloempty
$ cat readme-local-porting-1-stm32-f4-discovery

change the arch to mblite from cortexm7f
    two more changes:
    [1] change egc/build.go and egc/buildtools.go to define alignment size and gcc command options.
    [2] copy arch-specific files from their cortexm4f to -mblite version and change +build too.
        the list is by fixing missing symbols in compile or link but should evaluate through packages in build log.


$ git diff
diff --git a/egc/build.go b/egc/build.go
index 7fc97390..ad1c5686 100755
--- a/egc/build.go
+++ b/egc/build.go
@@ -49,6 +49,7 @@ var (
                "cortexm7d": cortexmSizes,

                "amd64": &gotoc.StdSizes{8, 8},
+               "mblite": cortexmSizes, /* copy of cortexm4f */
        }
 )

diff --git a/egc/buildtools.go b/egc/buildtools.go
index 6d0fbb07..25285233 100755
--- a/egc/buildtools.go
+++ b/egc/buildtools.go
@@ -18,6 +18,7 @@ var archMap = map[string]string{
        "cortexm7f": "-mcpu=cortex-m7 -mthumb -mfloat-abi=hard -mfpu=fpv5-sp-d16",
        "cortexm7d": "-mcpu=cortex-m7 -mthumb -mfloat-abi=hard -mfpu=fpv5-dp-d16",
        "amd64":     "",
+       "mblite": "-mcpu=cortex-m4 -mthumb -mfloat-abi=hard -mfpu=fpv4-sp-d16", /* copy of cortexm4f */
 }

 var osMap = map[string]struct{ cc, ld string }{


$ git status
Changes not staged for commit:

        modified:   egc/build.go
        modified:   egc/buildtools.go

Untracked files:

        egroot/src/internal/copy-mblite.s
        egroot/src/internal/memcmp-mblite.s
        egroot/src/internal/memset-mblite.s
        egroot/src/internal/syscall-noos-mblite+.h
        egroot/src/runtime/noos/config-mblite.go
        egroot/src/runtime/noos/cpu-mblite.go
        egroot/src/runtime/noos/faults-mblite.go
        egroot/src/runtime/noos/faults-mblite.s
        egroot/src/runtime/noos/panic-mblite.go
        egroot/src/runtime/noos/pendsv-mblite.go
        egroot/src/runtime/noos/pendsv-mblitefpu+.c
        egroot/src/runtime/noos/reset-mblite+.c
        egroot/src/runtime/noos/stack-mblite+.c
        egroot/src/runtime/noos/stack-mblite.go
        egroot/src/runtime/noos/svc-mblite+.c
        egroot/src/runtime/noos/svc-mblite.go
        egroot/src/runtime/noos/tasker-mblite.go
        egroot/src/runtime/noos/vectors-mblite+.c
        egroot/src/runtime/noos/vectors-mblite.go
        egroot/src/sync/fence/fence-mblite+.h
        egroot/src/syscall/irq-noos-mblite.go
        egroot/src/syscall/schednext-noos-mblite.go
```
