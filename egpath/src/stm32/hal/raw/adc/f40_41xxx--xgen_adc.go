// +build f40_41xxx

package adc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type ADC_Periph struct {
	SR    SR
	CR1   CR1
	CR2   CR2
	SMPR1 SMPR1
	SMPR2 SMPR2
	JOFR1 JOFR1
	JOFR2 JOFR2
	JOFR3 JOFR3
	JOFR4 JOFR4
	HTR   HTR
	LTR   LTR
	SQR1  SQR1
	SQR2  SQR2
	SQR3  SQR3
	JSQR  JSQR
	JDR   [4]JDR
	DR    DR
}

func (p *ADC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var ADC1 = (*ADC_Periph)(unsafe.Pointer(uintptr(mmap.ADC1_BASE)))

var ADC2 = (*ADC_Periph)(unsafe.Pointer(uintptr(mmap.ADC2_BASE)))

var ADC3 = (*ADC_Periph)(unsafe.Pointer(uintptr(mmap.ADC3_BASE)))

type SR_Bits uint32

func (b SR_Bits) Field(mask SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR_Bits) J(v int) SR_Bits {
	return SR_Bits(bits.Make32(v, uint32(mask)))
}

type SR struct{ mmio.U32 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U32.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U32.Store(uint32(b)) }

type SR_Mask struct{ mmio.UM32 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM32.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) AWD() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(AWD)}}
}

func (p *ADC_Periph) EOC() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(EOC)}}
}

func (p *ADC_Periph) JEOC() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(JEOC)}}
}

func (p *ADC_Periph) JSTRT() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(JSTRT)}}
}

func (p *ADC_Periph) STRT() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(STRT)}}
}

func (p *ADC_Periph) OVR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(OVR)}}
}

type CR1_Bits uint32

func (b CR1_Bits) Field(mask CR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1_Bits) J(v int) CR1_Bits {
	return CR1_Bits(bits.Make32(v, uint32(mask)))
}

type CR1 struct{ mmio.U32 }

func (r *CR1) Bits(mask CR1_Bits) CR1_Bits { return CR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR1) StoreBits(mask, b CR1_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR1) SetBits(mask CR1_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CR1) ClearBits(mask CR1_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CR1) Load() CR1_Bits              { return CR1_Bits(r.U32.Load()) }
func (r *CR1) Store(b CR1_Bits)            { r.U32.Store(uint32(b)) }

type CR1_Mask struct{ mmio.UM32 }

func (rm CR1_Mask) Load() CR1_Bits   { return CR1_Bits(rm.UM32.Load()) }
func (rm CR1_Mask) Store(b CR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) AWDCH() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(AWDCH)}}
}

func (p *ADC_Periph) EOCIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(EOCIE)}}
}

func (p *ADC_Periph) AWDIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(AWDIE)}}
}

func (p *ADC_Periph) JEOCIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(JEOCIE)}}
}

func (p *ADC_Periph) SCAN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(SCAN)}}
}

func (p *ADC_Periph) AWDSGL() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(AWDSGL)}}
}

func (p *ADC_Periph) JAUTO() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(JAUTO)}}
}

func (p *ADC_Periph) DISCEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(DISCEN)}}
}

func (p *ADC_Periph) JDISCEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(JDISCEN)}}
}

func (p *ADC_Periph) DISCNUM() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(DISCNUM)}}
}

func (p *ADC_Periph) JAWDEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(JAWDEN)}}
}

func (p *ADC_Periph) AWDEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(AWDEN)}}
}

func (p *ADC_Periph) RES() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(RES)}}
}

func (p *ADC_Periph) OVRIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(OVRIE)}}
}

type CR2_Bits uint32

func (b CR2_Bits) Field(mask CR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2_Bits) J(v int) CR2_Bits {
	return CR2_Bits(bits.Make32(v, uint32(mask)))
}

type CR2 struct{ mmio.U32 }

func (r *CR2) Bits(mask CR2_Bits) CR2_Bits { return CR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR2) StoreBits(mask, b CR2_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR2) SetBits(mask CR2_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CR2) ClearBits(mask CR2_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CR2) Load() CR2_Bits              { return CR2_Bits(r.U32.Load()) }
func (r *CR2) Store(b CR2_Bits)            { r.U32.Store(uint32(b)) }

type CR2_Mask struct{ mmio.UM32 }

func (rm CR2_Mask) Load() CR2_Bits   { return CR2_Bits(rm.UM32.Load()) }
func (rm CR2_Mask) Store(b CR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) ADON() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(ADON)}}
}

func (p *ADC_Periph) CONT() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(CONT)}}
}

func (p *ADC_Periph) DMA() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(DMA)}}
}

func (p *ADC_Periph) DDS() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(DDS)}}
}

func (p *ADC_Periph) EOCS() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(EOCS)}}
}

func (p *ADC_Periph) ALIGN() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(ALIGN)}}
}

func (p *ADC_Periph) JEXTSEL() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(JEXTSEL)}}
}

func (p *ADC_Periph) JEXTEN() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(JEXTEN)}}
}

func (p *ADC_Periph) JSWSTART() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(JSWSTART)}}
}

func (p *ADC_Periph) EXTSEL() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(EXTSEL)}}
}

func (p *ADC_Periph) EXTEN() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(EXTEN)}}
}

func (p *ADC_Periph) SWSTART() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(SWSTART)}}
}

type SMPR1_Bits uint32

func (b SMPR1_Bits) Field(mask SMPR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SMPR1_Bits) J(v int) SMPR1_Bits {
	return SMPR1_Bits(bits.Make32(v, uint32(mask)))
}

type SMPR1 struct{ mmio.U32 }

func (r *SMPR1) Bits(mask SMPR1_Bits) SMPR1_Bits { return SMPR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *SMPR1) StoreBits(mask, b SMPR1_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SMPR1) SetBits(mask SMPR1_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *SMPR1) ClearBits(mask SMPR1_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *SMPR1) Load() SMPR1_Bits                { return SMPR1_Bits(r.U32.Load()) }
func (r *SMPR1) Store(b SMPR1_Bits)              { r.U32.Store(uint32(b)) }

type SMPR1_Mask struct{ mmio.UM32 }

func (rm SMPR1_Mask) Load() SMPR1_Bits   { return SMPR1_Bits(rm.UM32.Load()) }
func (rm SMPR1_Mask) Store(b SMPR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SMP10() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP10)}}
}

func (p *ADC_Periph) SMP11() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP11)}}
}

func (p *ADC_Periph) SMP12() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP12)}}
}

func (p *ADC_Periph) SMP13() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP13)}}
}

func (p *ADC_Periph) SMP14() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP14)}}
}

func (p *ADC_Periph) SMP15() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP15)}}
}

func (p *ADC_Periph) SMP16() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP16)}}
}

func (p *ADC_Periph) SMP17() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP17)}}
}

func (p *ADC_Periph) SMP18() SMPR1_Mask {
	return SMPR1_Mask{mmio.UM32{&p.SMPR1.U32, uint32(SMP18)}}
}

type SMPR2_Bits uint32

func (b SMPR2_Bits) Field(mask SMPR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SMPR2_Bits) J(v int) SMPR2_Bits {
	return SMPR2_Bits(bits.Make32(v, uint32(mask)))
}

type SMPR2 struct{ mmio.U32 }

func (r *SMPR2) Bits(mask SMPR2_Bits) SMPR2_Bits { return SMPR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *SMPR2) StoreBits(mask, b SMPR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SMPR2) SetBits(mask SMPR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *SMPR2) ClearBits(mask SMPR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *SMPR2) Load() SMPR2_Bits                { return SMPR2_Bits(r.U32.Load()) }
func (r *SMPR2) Store(b SMPR2_Bits)              { r.U32.Store(uint32(b)) }

type SMPR2_Mask struct{ mmio.UM32 }

func (rm SMPR2_Mask) Load() SMPR2_Bits   { return SMPR2_Bits(rm.UM32.Load()) }
func (rm SMPR2_Mask) Store(b SMPR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SMP0() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP0)}}
}

func (p *ADC_Periph) SMP1() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP1)}}
}

func (p *ADC_Periph) SMP2() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP2)}}
}

func (p *ADC_Periph) SMP3() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP3)}}
}

func (p *ADC_Periph) SMP4() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP4)}}
}

func (p *ADC_Periph) SMP5() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP5)}}
}

func (p *ADC_Periph) SMP6() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP6)}}
}

func (p *ADC_Periph) SMP7() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP7)}}
}

func (p *ADC_Periph) SMP8() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP8)}}
}

func (p *ADC_Periph) SMP9() SMPR2_Mask {
	return SMPR2_Mask{mmio.UM32{&p.SMPR2.U32, uint32(SMP9)}}
}

type JOFR1_Bits uint32

func (b JOFR1_Bits) Field(mask JOFR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JOFR1_Bits) J(v int) JOFR1_Bits {
	return JOFR1_Bits(bits.Make32(v, uint32(mask)))
}

type JOFR1 struct{ mmio.U32 }

func (r *JOFR1) Bits(mask JOFR1_Bits) JOFR1_Bits { return JOFR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *JOFR1) StoreBits(mask, b JOFR1_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JOFR1) SetBits(mask JOFR1_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *JOFR1) ClearBits(mask JOFR1_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *JOFR1) Load() JOFR1_Bits                { return JOFR1_Bits(r.U32.Load()) }
func (r *JOFR1) Store(b JOFR1_Bits)              { r.U32.Store(uint32(b)) }

type JOFR1_Mask struct{ mmio.UM32 }

func (rm JOFR1_Mask) Load() JOFR1_Bits   { return JOFR1_Bits(rm.UM32.Load()) }
func (rm JOFR1_Mask) Store(b JOFR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JOFFSET1() JOFR1_Mask {
	return JOFR1_Mask{mmio.UM32{&p.JOFR1.U32, uint32(JOFFSET1)}}
}

type JOFR2_Bits uint32

func (b JOFR2_Bits) Field(mask JOFR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JOFR2_Bits) J(v int) JOFR2_Bits {
	return JOFR2_Bits(bits.Make32(v, uint32(mask)))
}

type JOFR2 struct{ mmio.U32 }

func (r *JOFR2) Bits(mask JOFR2_Bits) JOFR2_Bits { return JOFR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *JOFR2) StoreBits(mask, b JOFR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JOFR2) SetBits(mask JOFR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *JOFR2) ClearBits(mask JOFR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *JOFR2) Load() JOFR2_Bits                { return JOFR2_Bits(r.U32.Load()) }
func (r *JOFR2) Store(b JOFR2_Bits)              { r.U32.Store(uint32(b)) }

type JOFR2_Mask struct{ mmio.UM32 }

func (rm JOFR2_Mask) Load() JOFR2_Bits   { return JOFR2_Bits(rm.UM32.Load()) }
func (rm JOFR2_Mask) Store(b JOFR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JOFFSET2() JOFR2_Mask {
	return JOFR2_Mask{mmio.UM32{&p.JOFR2.U32, uint32(JOFFSET2)}}
}

type JOFR3_Bits uint32

func (b JOFR3_Bits) Field(mask JOFR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JOFR3_Bits) J(v int) JOFR3_Bits {
	return JOFR3_Bits(bits.Make32(v, uint32(mask)))
}

type JOFR3 struct{ mmio.U32 }

func (r *JOFR3) Bits(mask JOFR3_Bits) JOFR3_Bits { return JOFR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *JOFR3) StoreBits(mask, b JOFR3_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JOFR3) SetBits(mask JOFR3_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *JOFR3) ClearBits(mask JOFR3_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *JOFR3) Load() JOFR3_Bits                { return JOFR3_Bits(r.U32.Load()) }
func (r *JOFR3) Store(b JOFR3_Bits)              { r.U32.Store(uint32(b)) }

type JOFR3_Mask struct{ mmio.UM32 }

func (rm JOFR3_Mask) Load() JOFR3_Bits   { return JOFR3_Bits(rm.UM32.Load()) }
func (rm JOFR3_Mask) Store(b JOFR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JOFFSET3() JOFR3_Mask {
	return JOFR3_Mask{mmio.UM32{&p.JOFR3.U32, uint32(JOFFSET3)}}
}

type JOFR4_Bits uint32

func (b JOFR4_Bits) Field(mask JOFR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JOFR4_Bits) J(v int) JOFR4_Bits {
	return JOFR4_Bits(bits.Make32(v, uint32(mask)))
}

type JOFR4 struct{ mmio.U32 }

func (r *JOFR4) Bits(mask JOFR4_Bits) JOFR4_Bits { return JOFR4_Bits(r.U32.Bits(uint32(mask))) }
func (r *JOFR4) StoreBits(mask, b JOFR4_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JOFR4) SetBits(mask JOFR4_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *JOFR4) ClearBits(mask JOFR4_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *JOFR4) Load() JOFR4_Bits                { return JOFR4_Bits(r.U32.Load()) }
func (r *JOFR4) Store(b JOFR4_Bits)              { r.U32.Store(uint32(b)) }

type JOFR4_Mask struct{ mmio.UM32 }

func (rm JOFR4_Mask) Load() JOFR4_Bits   { return JOFR4_Bits(rm.UM32.Load()) }
func (rm JOFR4_Mask) Store(b JOFR4_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JOFFSET4() JOFR4_Mask {
	return JOFR4_Mask{mmio.UM32{&p.JOFR4.U32, uint32(JOFFSET4)}}
}

type HTR_Bits uint32

func (b HTR_Bits) Field(mask HTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HTR_Bits) J(v int) HTR_Bits {
	return HTR_Bits(bits.Make32(v, uint32(mask)))
}

type HTR struct{ mmio.U32 }

func (r *HTR) Bits(mask HTR_Bits) HTR_Bits { return HTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *HTR) StoreBits(mask, b HTR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HTR) SetBits(mask HTR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *HTR) ClearBits(mask HTR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *HTR) Load() HTR_Bits              { return HTR_Bits(r.U32.Load()) }
func (r *HTR) Store(b HTR_Bits)            { r.U32.Store(uint32(b)) }

type HTR_Mask struct{ mmio.UM32 }

func (rm HTR_Mask) Load() HTR_Bits   { return HTR_Bits(rm.UM32.Load()) }
func (rm HTR_Mask) Store(b HTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) HT() HTR_Mask {
	return HTR_Mask{mmio.UM32{&p.HTR.U32, uint32(HT)}}
}

type LTR_Bits uint32

func (b LTR_Bits) Field(mask LTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LTR_Bits) J(v int) LTR_Bits {
	return LTR_Bits(bits.Make32(v, uint32(mask)))
}

type LTR struct{ mmio.U32 }

func (r *LTR) Bits(mask LTR_Bits) LTR_Bits { return LTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *LTR) StoreBits(mask, b LTR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *LTR) SetBits(mask LTR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *LTR) ClearBits(mask LTR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *LTR) Load() LTR_Bits              { return LTR_Bits(r.U32.Load()) }
func (r *LTR) Store(b LTR_Bits)            { r.U32.Store(uint32(b)) }

type LTR_Mask struct{ mmio.UM32 }

func (rm LTR_Mask) Load() LTR_Bits   { return LTR_Bits(rm.UM32.Load()) }
func (rm LTR_Mask) Store(b LTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) LT() LTR_Mask {
	return LTR_Mask{mmio.UM32{&p.LTR.U32, uint32(LT)}}
}

type SQR1_Bits uint32

func (b SQR1_Bits) Field(mask SQR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR1_Bits) J(v int) SQR1_Bits {
	return SQR1_Bits(bits.Make32(v, uint32(mask)))
}

type SQR1 struct{ mmio.U32 }

func (r *SQR1) Bits(mask SQR1_Bits) SQR1_Bits { return SQR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *SQR1) StoreBits(mask, b SQR1_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SQR1) SetBits(mask SQR1_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *SQR1) ClearBits(mask SQR1_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *SQR1) Load() SQR1_Bits               { return SQR1_Bits(r.U32.Load()) }
func (r *SQR1) Store(b SQR1_Bits)             { r.U32.Store(uint32(b)) }

type SQR1_Mask struct{ mmio.UM32 }

func (rm SQR1_Mask) Load() SQR1_Bits   { return SQR1_Bits(rm.UM32.Load()) }
func (rm SQR1_Mask) Store(b SQR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ13() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{&p.SQR1.U32, uint32(SQ13)}}
}

func (p *ADC_Periph) SQ14() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{&p.SQR1.U32, uint32(SQ14)}}
}

func (p *ADC_Periph) SQ15() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{&p.SQR1.U32, uint32(SQ15)}}
}

func (p *ADC_Periph) SQ16() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{&p.SQR1.U32, uint32(SQ16)}}
}

func (p *ADC_Periph) L() SQR1_Mask {
	return SQR1_Mask{mmio.UM32{&p.SQR1.U32, uint32(L)}}
}

type SQR2_Bits uint32

func (b SQR2_Bits) Field(mask SQR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR2_Bits) J(v int) SQR2_Bits {
	return SQR2_Bits(bits.Make32(v, uint32(mask)))
}

type SQR2 struct{ mmio.U32 }

func (r *SQR2) Bits(mask SQR2_Bits) SQR2_Bits { return SQR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *SQR2) StoreBits(mask, b SQR2_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SQR2) SetBits(mask SQR2_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *SQR2) ClearBits(mask SQR2_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *SQR2) Load() SQR2_Bits               { return SQR2_Bits(r.U32.Load()) }
func (r *SQR2) Store(b SQR2_Bits)             { r.U32.Store(uint32(b)) }

type SQR2_Mask struct{ mmio.UM32 }

func (rm SQR2_Mask) Load() SQR2_Bits   { return SQR2_Bits(rm.UM32.Load()) }
func (rm SQR2_Mask) Store(b SQR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ7() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{&p.SQR2.U32, uint32(SQ7)}}
}

func (p *ADC_Periph) SQ8() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{&p.SQR2.U32, uint32(SQ8)}}
}

func (p *ADC_Periph) SQ9() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{&p.SQR2.U32, uint32(SQ9)}}
}

func (p *ADC_Periph) SQ10() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{&p.SQR2.U32, uint32(SQ10)}}
}

func (p *ADC_Periph) SQ11() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{&p.SQR2.U32, uint32(SQ11)}}
}

func (p *ADC_Periph) SQ12() SQR2_Mask {
	return SQR2_Mask{mmio.UM32{&p.SQR2.U32, uint32(SQ12)}}
}

type SQR3_Bits uint32

func (b SQR3_Bits) Field(mask SQR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SQR3_Bits) J(v int) SQR3_Bits {
	return SQR3_Bits(bits.Make32(v, uint32(mask)))
}

type SQR3 struct{ mmio.U32 }

func (r *SQR3) Bits(mask SQR3_Bits) SQR3_Bits { return SQR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *SQR3) StoreBits(mask, b SQR3_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SQR3) SetBits(mask SQR3_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *SQR3) ClearBits(mask SQR3_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *SQR3) Load() SQR3_Bits               { return SQR3_Bits(r.U32.Load()) }
func (r *SQR3) Store(b SQR3_Bits)             { r.U32.Store(uint32(b)) }

type SQR3_Mask struct{ mmio.UM32 }

func (rm SQR3_Mask) Load() SQR3_Bits   { return SQR3_Bits(rm.UM32.Load()) }
func (rm SQR3_Mask) Store(b SQR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) SQ1() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{&p.SQR3.U32, uint32(SQ1)}}
}

func (p *ADC_Periph) SQ2() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{&p.SQR3.U32, uint32(SQ2)}}
}

func (p *ADC_Periph) SQ3() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{&p.SQR3.U32, uint32(SQ3)}}
}

func (p *ADC_Periph) SQ4() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{&p.SQR3.U32, uint32(SQ4)}}
}

func (p *ADC_Periph) SQ5() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{&p.SQR3.U32, uint32(SQ5)}}
}

func (p *ADC_Periph) SQ6() SQR3_Mask {
	return SQR3_Mask{mmio.UM32{&p.SQR3.U32, uint32(SQ6)}}
}

type JSQR_Bits uint32

func (b JSQR_Bits) Field(mask JSQR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JSQR_Bits) J(v int) JSQR_Bits {
	return JSQR_Bits(bits.Make32(v, uint32(mask)))
}

type JSQR struct{ mmio.U32 }

func (r *JSQR) Bits(mask JSQR_Bits) JSQR_Bits { return JSQR_Bits(r.U32.Bits(uint32(mask))) }
func (r *JSQR) StoreBits(mask, b JSQR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JSQR) SetBits(mask JSQR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *JSQR) ClearBits(mask JSQR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *JSQR) Load() JSQR_Bits               { return JSQR_Bits(r.U32.Load()) }
func (r *JSQR) Store(b JSQR_Bits)             { r.U32.Store(uint32(b)) }

type JSQR_Mask struct{ mmio.UM32 }

func (rm JSQR_Mask) Load() JSQR_Bits   { return JSQR_Bits(rm.UM32.Load()) }
func (rm JSQR_Mask) Store(b JSQR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JSQ1() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{&p.JSQR.U32, uint32(JSQ1)}}
}

func (p *ADC_Periph) JSQ2() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{&p.JSQR.U32, uint32(JSQ2)}}
}

func (p *ADC_Periph) JSQ3() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{&p.JSQR.U32, uint32(JSQ3)}}
}

func (p *ADC_Periph) JSQ4() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{&p.JSQR.U32, uint32(JSQ4)}}
}

func (p *ADC_Periph) JL() JSQR_Mask {
	return JSQR_Mask{mmio.UM32{&p.JSQR.U32, uint32(JL)}}
}

type JDR_Bits uint32

func (b JDR_Bits) Field(mask JDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask JDR_Bits) J(v int) JDR_Bits {
	return JDR_Bits(bits.Make32(v, uint32(mask)))
}

type JDR struct{ mmio.U32 }

func (r *JDR) Bits(mask JDR_Bits) JDR_Bits { return JDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *JDR) StoreBits(mask, b JDR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *JDR) SetBits(mask JDR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *JDR) ClearBits(mask JDR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *JDR) Load() JDR_Bits              { return JDR_Bits(r.U32.Load()) }
func (r *JDR) Store(b JDR_Bits)            { r.U32.Store(uint32(b)) }

type JDR_Mask struct{ mmio.UM32 }

func (rm JDR_Mask) Load() JDR_Bits   { return JDR_Bits(rm.UM32.Load()) }
func (rm JDR_Mask) Store(b JDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) JDATA(n int) JDR_Mask {
	return JDR_Mask{mmio.UM32{&p.JDR[n].U32, uint32(JDATA)}}
}

type DR_Bits uint32

func (b DR_Bits) Field(mask DR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR_Bits) J(v int) DR_Bits {
	return DR_Bits(bits.Make32(v, uint32(mask)))
}

type DR struct{ mmio.U32 }

func (r *DR) Bits(mask DR_Bits) DR_Bits { return DR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DR) StoreBits(mask, b DR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DR) SetBits(mask DR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DR) ClearBits(mask DR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DR) Load() DR_Bits             { return DR_Bits(r.U32.Load()) }
func (r *DR) Store(b DR_Bits)           { r.U32.Store(uint32(b)) }

type DR_Mask struct{ mmio.UM32 }

func (rm DR_Mask) Load() DR_Bits   { return DR_Bits(rm.UM32.Load()) }
func (rm DR_Mask) Store(b DR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Periph) DATA() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(DATA)}}
}

func (p *ADC_Periph) ADC2DATA() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(ADC2DATA)}}
}
