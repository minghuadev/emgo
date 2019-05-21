
# Windows Build and Use with cygwin


## Install packages

**Install the Golang Windows SDK**

**Install Cygwin**

**Place emgo in a working direcotry**

Create a top-level directory somewhere under the cygwin '/home/...' and cd into the directory. All further env-setting commands will be issued from this directory. 

Create a directory at 'libs/src/github.com/ziutek/' and checkout the emgo repo under it. 

**Install the 'gnu-none-eabi' toolchain**

Refer to the examples building section below, download and install a pre-built toolchain. 


## Cygwin Environment Setup for Building egc

Assume the cygwin '/home' is mapped to windows 'D:\home' directory. Adjust accordingly if your directory mapping is different. 

Source this file so that the 'egc' from emgo can be built: 

```
$ cat envset-cygwin-compile-egc

# file: envset-cygwin-compile-egc

#
# cygwin /home maps to windows D:\home
# emgo repo checked out under libs/src/github.com/ziutek/emgo
#

# gopath for compiling egc using the windows go sdk:
GOPATH=$( pwd | sed -e 's/\/home\//D:\\\\home\\\\/' | sed -e 's/\//\\\\/')\\\\libs
export GOPATH

# path for using egc inside cygwin:
PATH=$( pwd )/libs/bin:$PATH
export PATH
```

Then as usual, cd into the emgo egc directory and issue the 'go install' command to build egc. 


## Cygwin Environment Setup for building emgo stm32 examples

After sourcing the above 'envset-cygwin-compile-emgo' file, source this file so as to be able to build emgo examples using the just built egc command: 

```
$ cat envset-cygwin-toolchain

# file: envset-cygwin-toolchain

#
# download the toolchain file gcc-arm-none-eabi-8-2018-14-major-win32.zip from
#    https://developer.arm.com/-/media/Files/downloads/gnu-rm/8-2018q4/gcc-arm-none-eabi-8-2018-q4-major-win32.zip?revision=ab9cb8f8-6a9d-4a6e-818a-295f5d1ca982?product=GNU%20Arm%20Embedded%20Toolchain,ZIP,,Windows,8-2018-q4-major
#    with md5: 9b1cfb7539af11b0badfaa960679ea6f
# then extract to cygwin $HOME/gcc-arm-none-eabi-8-2018-q4-major-win32/bin
#

# env for toolchain commands:
toolchain_bindir=$HOME/gcc-arm-none-eabi-8-2018-q4-major-win32/bin
export EGCC=$toolchain_bindir/arm-none-eabi-gcc
export EGLD=$toolchain_bindir/arm-none-eabi-ld
export EGAR=$toolchain_bindir/arm-none-eabi-ar

# env for emgo:
export EGROOT=$(pwd | xargs)/libs/src/github.com/ziutek/emgo/egroot
export EGPATH=$(pwd | xargs)/libs/src/github.com/ziutek/emgo/egpath
```

Then as usually, cd into the emgo egpath stm32 examples directory and issue the '../build.sh' command for one of the samples. 


