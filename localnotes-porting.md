
# TOC

* Debug Change to Print Running Sequence
* Running Sequence for stm32/examples/f4-explorer/helloworld

# Debug Change to Print Running Sequence

Apply the diff that follows and use "-v 1" to print the order of running steps. Or use "-v 2" to also print the gcc commands. 

```
diff --git a/egc/build.go b/egc/build.go
index c4a2e31e..7fc97390
--- a/egc/build.go
+++ b/egc/build.go
@@ -52,6 +52,8 @@ var (
        }
 )

+var compileStartCnt uint = 0
+var compileEndCnt uint = 0
 func compile(bp *build.Package) error {
        if ok, err := checkPkg(bp); err != nil {
                return err
@@ -59,7 +61,13 @@ func compile(bp *build.Package) error {
                return nil
        }
        if verbosity > 0 {
-               defer fmt.Println(bp.ImportPath)
+               compileStartCnt ++
+               fmt.Printf("pkg start %-2d %2s %s\n",
+                       compileStartCnt, "", bp.ImportPath)
+               defer func(){
+                       compileEndCnt ++;
+                       fmt.Printf("pkg end   %-2d %-2d %s\n",
+                               compileStartCnt, compileEndCnt, bp.ImportPath)}()
        }

        // Parse
diff --git a/egc/main.go b/egc/main.go
index 5c2c6b28..454fdb31
--- a/egc/main.go
+++ b/egc/main.go
@@ -101,7 +101,13 @@ func main() {

        var err error

-       tmpDir, err = ioutil.TempDir("", "eg-build")
+       tmpDirTop := ""
+       tmpDirPrefix := "eg-build"
+       if verbosity > 0 {
+               tmpDirTop = "."
+               tmpDirPrefix = "tmp-eg-build"
+       }
+       tmpDir, err = ioutil.TempDir(tmpDirTop, tmpDirPrefix)
        if err != nil {
                logErr(err)
                os.Exit(1)

```

# Running Sequence for stm32/examples/f4-explorer/helloworld

```
pkg start 1     stm32/examples/f4-discovery/helloworld
pkg start 2     fmt
pkg start 3     rtos
pkg start 4     syscall
pkg start 5     internal
pkg end   5  1  internal
pkg start 6     sync/atomic
pkg end   6  2  sync/atomic
pkg start 7     sync/fence
pkg end   7  3  sync/fence
pkg start 8     arch/cortexm
pkg end   8  4  arch/cortexm
pkg start 9     arch/cortexm/scb
pkg start 10    bits
pkg end   10 5  bits
pkg start 11    mmio
pkg end   11 6  mmio
pkg end   11 7  arch/cortexm/scb
pkg end   11 8  syscall
pkg end   11 9  rtos
pkg start 12    io
pkg start 13    errors
pkg end   13 10 errors
pkg end   13 11 io
pkg start 14    reflect
pkg start 15    mem
pkg end   15 12 mem
pkg end   15 13 reflect
pkg start 16    strconv
pkg start 17    math
pkg end   17 14 math
pkg end   17 15 strconv
pkg start 18    strings
pkg end   18 16 strings
pkg end   18 17 fmt
pkg start 19    runtime
pkg start 20    runtime/noos
pkg start 21    arch/cortexm/acc
pkg end   21 18 arch/cortexm/acc
pkg start 22    arch/cortexm/cmt
pkg end   22 19 arch/cortexm/cmt
pkg start 23    arch/cortexm/fpu
pkg end   23 20 arch/cortexm/fpu
pkg start 24    arch/cortexm/pft
pkg end   24 21 arch/cortexm/pft
pkg start 25    arch/cortexm/debug/itm
pkg end   25 22 arch/cortexm/debug/itm
pkg start 26    arch/cortexm/mpu
pkg end   26 23 arch/cortexm/mpu
pkg start 27    arch/cortexm/nvic
pkg end   27 24 arch/cortexm/nvic
pkg start 28    math/rand
pkg end   28 25 math/rand
pkg end   28 26 runtime/noos
pkg end   28 27 runtime
pkg end   28 28 stm32/examples/f4-discovery/helloworld
```
