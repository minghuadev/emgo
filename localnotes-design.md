
# https://github.com/ziutek/emgo/issues/14

### ziutek commented on Jun 20, 2018 

This project evolved from proof of concept Go to C translator.

The "compiler" related code is in the following directories

**gotoc**

Contains Go to C translator. It's one of the most horrible and convoluted code I've written in my life. Much worse than the code of my first big Go project from 2010: mymysql.

I was started without any solid knowledge of Go type system and AST. At this time there was no go/type package yet. It evolved rapidly to handle more and more corner cases and still don't handle all well. This is a transitional solution and I prepare to write a real front-end to GCC or LLVM.

**egc**

The egc (Emgo compiler) uses gotoc to ranslate Go sources to C and uses gcc, ar, ld to compile, archive, link them into packages or executables.

**egroot/egc**

Contains definitions of basic Go types.

That's all of the strict compiler related code.

There is **egroot/src/internal** package that is in the middle between compiler and the standard library. It implements basic and architecture dependent things that compiler wants like memcpy, memset, typeinfo, ...

All of the above things are subject to change when real compiler will be created because they are specific to Go to C translation.

There is **egroot/runtime** that contains things that are architecture and OS dependent. Currently it supports only noos OS. There is also not finished but partially working support for linux.

The code base contains also:

The standard library located in **egroot** and other packages, examples, etc. located in **egpath**.

About porting to other targets. I have surprisingly a lot of such proposals from people interested in Emgo and almost any question of using of the current code on supported architectures. It looks a bit like everyone who want to learn and use the real Go will start from porting it to new architecture/OS...

If you are interested in Emgo start to use it, ask concrete question about using it, write something about it and after that you will be more familiar with this project, you will know more its good and weak points and you can start to consider to port it to other targets.

You are absolutely right about the needs of the community. I started to collect things to publish as a wiki page.

I'm waiting for any experience from use of Emgo (these bad and good ones). The link to the article / blog post will be published on Emgo blog. I'm open to articles for the wiki page.

Please analyze the code and ask more questions about Emgo internals. This issue can be used to create the first article about Emgo internals.

See also issue #13



# https://github.com/ziutek/emgo/issues/13

### maxekman commented on Jun 17, 2018

Are there any plans for generating board/arch definitions using SVD files? The Rust embedded effort uses a tool called svd2rust to generate type safe memory mappings (I’m not familiar with all the details in Rust as I’m a Go developer).

### ziutek commented on Jun 18, 2018

I don't know anything about SVD. Quick check in Google and I known that there is something like CMSIS-SVD and It looks interesting.

Currently packages in stm32/hal/raw are automatically generated from ST Cube MX C header files but it isn't a perfect solution.

### maxekman commented on Jun 19, 2018

Yes, CMSIS-SVD is the format they are parsing, apart from that I don't know much about it.

If I get some time to experiment with it and start on a parser I'll let you know.

### ziutek commented on Jun 20, 2018

Currently Emgo has something called xgen.

See for example this file that describes SPI peripheral in STM32L476:

https://github.com/ziutek/emgo/blob/master/egpath/src/stm32/hal/raw/spi/l476xx--spi.go

The top comment describes the peripheral, its instances and registers:

```
// Peripheral: SPI_Periph  Serial Peripheral Interface.
// Instances:
//  SPI2  mmap.SPI2_BASE
//  SPI3  mmap.SPI3_BASE
//  SPI1  mmap.SPI1_BASE
// Registers:
//  0x00 32  CR1    Control register 1.
//  0x04 32  CR2    Control register 2.
//  0x08 32  SR     Status register.
//  0x0C 32  DR     Data register.
//  0x10 32  CRCPR  CRC polynomial register.
//  0x14 32  RXCRCR Rx CRC register.
//  0x18 32  TXCRCR Tx CRC register.
// Import:
//  stm32/o/l476xx/mmap
package spi
```

This is easy to write and read. Next the file contains constants that describe all bits/bitfields of these registers:

```
const (
	CPHA     CR1 = 0x01 << 0  //+ Clock Phase.
	CPOL     CR1 = 0x01 << 1  //+ Clock Polarity.
	MSTR     CR1 = 0x01 << 2  //+ Master Selection.
	BR       CR1 = 0x07 << 3  //+ BR[2:0] bits (Baud Rate Control).
	BR_0     CR1 = 0x01 << 3  //  Bit 0.
	BR_1     CR1 = 0x02 << 3  //  Bit 1.
	BR_2     CR1 = 0x04 << 3  //  Bit 2.
	SPE      CR1 = 0x01 << 6  //+ SPI Enable.
	LSBFIRST CR1 = 0x01 << 7  //+ Frame Format.
	SSI      CR1 = 0x01 << 8  //+ Internal slave select.
	SSM      CR1 = 0x01 << 9  //+ Software slave management.
	RXONLY   CR1 = 0x01 << 10 //+ Receive only.
	CRCL     CR1 = 0x01 << 11 //+ CRC Length.
	CRCNEXT  CR1 = 0x01 << 12 //+ Transmit CRC next.
	CRCEN    CR1 = 0x01 << 13 //+ Hardware CRC calculation enable.
	BIDIOE   CR1 = 0x01 << 14 //+ Output enable in bidirectional mode.
	BIDIMODE CR1 = 0x01 << 15 //+ Bidirectional data mode enable.
)

const (
	CPHAn     = 0
	CPOLn     = 1
	MSTRn     = 2
	BRn       = 3
	SPEn      = 6
	LSBFIRSTn = 7
	SSIn      = 8
	SSMn      = 9
	RXONLYn   = 10
	CRCLn     = 11
	CRCNEXTn  = 12
	CRCENn    = 13
	BIDIOEn   = 14
	BIDIMODEn = 15
)
```

The xgen parses this peripheral file and generates the remaining code that allows you to write:

```
if spi.SPI2.CR1.Load() & spi.LSBFIRST != 0 {
//...
}
```

or even

```
if spi.SPI2.LSBFIRST().Load() != 0 {
//...
}
```

See generated code:

https://github.com/ziutek/emgo/blob/master/egpath/src/stm32/hal/raw/spi/f746xx--xgen_spi.go

In this case even the peripheral file isn't written by hand but the stm32xgen generates it from STM32 Cube MX C header file.

Because the peripherals are not identical, the stm32/hal/spi package try to provide unified interface to the SPI peripheral and default driver that uses DMA and IRQ.

Summarizing:

The stm32/hal/raw and stm32/hal/irq contain automatically generated packages that allow raw access the peripherals and names all IRQs. Other packages in stm32/hal are written by hand to provide unified interface and drivers.

 maxekman commented on Jun 20, 2018

Thanks for the description, looks pretty neat.

I was just reading an issue on embedded Rust regarding how they plan to handle boards and targets in the future, could be interesting as inspiration for emgo also: rust-embedded/wg#101



# https://github.com/ziutek/emgo/issues/15

### wumingyu12 commented on Jun 24, 2018

hello，first it is a wonderful project，i has hate to use c to write Embedded program。

i find you use xgen to create sdk .go file by register config .go file。
now i want to add something in register config .go file
how can i use xgen to update sdk .go file like xgen_ficr.go

### ziutek commented on Jun 25, 2018

Simly run:

```
xgen ficr.go
```

It regenerates xgen_ficr.go.

For STM32 see issue #13

### ziutek commented on Jun 25, 2018

There is one unobvious thing.

Bit field names must have a comment beginning with "//+". See the following example:

const (
    SPEED    = 3 << 5 //+
    VERYSLOW = 0 << 5
    SLOW     = 1 << 5
    NORMAL   = 2 << 5
    FAST     = 3 << 5
    DIR      = 1 << 7 //+
)

xgen uses it to distinguish between field name and its values. Of course, it is good to put a full description of any bit/bitfield.

