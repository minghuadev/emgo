// Peripheral: ADC_Periph  Analog to Digital Converter.
// Instances:
//  ADC1  mmap.ADC1_BASE
//  ADC2  mmap.ADC2_BASE
//  ADC3  mmap.ADC3_BASE
//  ADC4  mmap.ADC4_BASE
// Registers:
//  0x00 32  ISR     Interrupt and Status Register.
//  0x04 32  IER     Interrupt Enable Register.
//  0x08 32  CR      Control register.
//  0x0C 32  CFGR    Configuration register.
//  0x14 32  SMPR1   Sample time register 1.
//  0x18 32  SMPR2   Sample time register 2.
//  0x20 32  TR1     Watchdog threshold register 1.
//  0x24 32  TR2     Watchdog threshold register 2.
//  0x28 32  TR3     Watchdog threshold register 3.
//  0x30 32  SQR1    Regular sequence register 1.
//  0x34 32  SQR2    Regular sequence register 2.
//  0x38 32  SQR3    Regular sequence register 3.
//  0x3C 32  SQR4    Regular sequence register 4.
//  0x40 32  DR      Regular data register.
//  0x4C 32  JSQR    Injected sequence register.
//  0x60 32  OFR[4]  Offset registers.
//  0x80 32  JDR[4]  Injected data registers.
//  0xA0 32  AWD2CR  Analog Watchdog 2 Configuration Register.
//  0xA4 32  AWD3CR  Analog Watchdog 3 Configuration Register.
//  0xB0 32  DIFSEL  Differential Mode Selection Register.
//  0xB4 32  CALFACT Calibration Factors.
// Import:
//  stm32/o/f303xe/mmap
package adc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	RDY   ISR_Bits = 0x01 << 0  //+ ADC Ready (ADRDY) flag.
	EOSMP ISR_Bits = 0x01 << 1  //+ ADC End of Sampling flag.
	EOC   ISR_Bits = 0x01 << 2  //+ ADC End of Regular Conversion flag.
	EOS   ISR_Bits = 0x01 << 3  //+ ADC End of Regular sequence of Conversions flag.
	OVR   ISR_Bits = 0x01 << 4  //+ ADC overrun flag.
	JEOC  ISR_Bits = 0x01 << 5  //+ ADC End of Injected Conversion flag.
	JEOS  ISR_Bits = 0x01 << 6  //+ ADC End of Injected sequence of Conversions flag.
	AWD1  ISR_Bits = 0x01 << 7  //+ ADC Analog watchdog 1 flag.
	AWD2  ISR_Bits = 0x01 << 8  //+ ADC Analog watchdog 2 flag.
	AWD3  ISR_Bits = 0x01 << 9  //+ ADC Analog watchdog 3 flag.
	JQOVF ISR_Bits = 0x01 << 10 //+ ADC Injected Context Queue Overflow flag.
)

const (
	RDYn   = 0
	EOSMPn = 1
	EOCn   = 2
	EOSn   = 3
	OVRn   = 4
	JEOCn  = 5
	JEOSn  = 6
	AWD1n  = 7
	AWD2n  = 8
	AWD3n  = 9
	JQOVFn = 10
)

const (
	RDYIE   IER_Bits = 0x01 << 0  //+ ADC Ready (ADRDY) interrupt source.
	EOSMPIE IER_Bits = 0x01 << 1  //+ ADC End of Sampling interrupt source.
	EOCIE   IER_Bits = 0x01 << 2  //+ ADC End of Regular Conversion interrupt source.
	EOSIE   IER_Bits = 0x01 << 3  //+ ADC End of Regular sequence of Conversions interrupt source.
	OVRIE   IER_Bits = 0x01 << 4  //+ ADC overrun interrupt source.
	JEOCIE  IER_Bits = 0x01 << 5  //+ ADC End of Injected Conversion interrupt source.
	JEOSIE  IER_Bits = 0x01 << 6  //+ ADC End of Injected sequence of Conversions interrupt source.
	AWD1IE  IER_Bits = 0x01 << 7  //+ ADC Analog watchdog 1 interrupt source.
	AWD2IE  IER_Bits = 0x01 << 8  //+ ADC Analog watchdog 2 interrupt source.
	AWD3IE  IER_Bits = 0x01 << 9  //+ ADC Analog watchdog 3 interrupt source.
	JQOVFIE IER_Bits = 0x01 << 10 //+ ADC Injected Context Queue Overflow interrupt source.
)

const (
	RDYIEn   = 0
	EOSMPIEn = 1
	EOCIEn   = 2
	EOSIEn   = 3
	OVRIEn   = 4
	JEOCIEn  = 5
	JEOSIEn  = 6
	AWD1IEn  = 7
	AWD2IEn  = 8
	AWD3IEn  = 9
	JQOVFIEn = 10
)

const (
	ADEN       CR_Bits = 0x01 << 0  //+ ADC Enable control.
	ADDIS      CR_Bits = 0x01 << 1  //+ ADC Disable command.
	ADSTART    CR_Bits = 0x01 << 2  //+ ADC Start of Regular conversion.
	JADSTART   CR_Bits = 0x01 << 3  //+ ADC Start of injected conversion.
	ADSTP      CR_Bits = 0x01 << 4  //+ ADC Stop of Regular conversion.
	JADSTP     CR_Bits = 0x01 << 5  //+ ADC Stop of injected conversion.
	ADVREGEN   CR_Bits = 0x03 << 28 //+ ADC Voltage regulator Enable.
	ADVREGEN_0 CR_Bits = 0x01 << 28 //  ADC ADVREGEN bit 0.
	ADVREGEN_1 CR_Bits = 0x02 << 28 //  ADC ADVREGEN bit 1.
	ADCALDIF   CR_Bits = 0x01 << 30 //+ ADC Differential Mode for calibration.
	ADCAL      CR_Bits = 0x01 << 31 //+ ADC Calibration.
)

const (
	ADENn     = 0
	ADDISn    = 1
	ADSTARTn  = 2
	JADSTARTn = 3
	ADSTPn    = 4
	JADSTPn   = 5
	ADVREGENn = 28
	ADCALDIFn = 30
	ADCALn    = 31
)

const (
	DMAEN     CFGR_Bits = 0x01 << 0  //+ ADC DMA Enable.
	DMACFG    CFGR_Bits = 0x01 << 1  //+ ADC DMA configuration.
	RES       CFGR_Bits = 0x03 << 3  //+ ADC Data resolution.
	RES_0     CFGR_Bits = 0x01 << 3  //  ADC RES bit 0.
	RES_1     CFGR_Bits = 0x02 << 3  //  ADC RES bit 1.
	ALIGN     CFGR_Bits = 0x01 << 5  //+ ADC Data Alignment.
	EXTSEL    CFGR_Bits = 0x0F << 6  //+ ADC External trigger selection for regular group.
	EXTSEL_0  CFGR_Bits = 0x01 << 6  //  ADC EXTSEL bit 0.
	EXTSEL_1  CFGR_Bits = 0x02 << 6  //  ADC EXTSEL bit 1.
	EXTSEL_2  CFGR_Bits = 0x04 << 6  //  ADC EXTSEL bit 2.
	EXTSEL_3  CFGR_Bits = 0x08 << 6  //  ADC EXTSEL bit 3.
	EXTEN     CFGR_Bits = 0x03 << 10 //+ ADC External trigger enable and polarity selection for regular channels.
	EXTEN_0   CFGR_Bits = 0x01 << 10 //  ADC EXTEN bit 0.
	EXTEN_1   CFGR_Bits = 0x02 << 10 //  ADC EXTEN bit 1.
	OVRMOD    CFGR_Bits = 0x01 << 12 //+ ADC overrun mode.
	CONT      CFGR_Bits = 0x01 << 13 //+ ADC Single/continuous conversion mode for regular conversion.
	AUTDLY    CFGR_Bits = 0x01 << 14 //+ ADC Delayed conversion mode.
	DISCEN    CFGR_Bits = 0x01 << 16 //+ ADC Discontinuous mode for regular channels.
	DISCNUM   CFGR_Bits = 0x07 << 17 //+ ADC Discontinuous mode channel count.
	DISCNUM_0 CFGR_Bits = 0x01 << 17 //  ADC DISCNUM bit 0.
	DISCNUM_1 CFGR_Bits = 0x02 << 17 //  ADC DISCNUM bit 1.
	DISCNUM_2 CFGR_Bits = 0x04 << 17 //  ADC DISCNUM bit 2.
	JDISCEN   CFGR_Bits = 0x01 << 20 //+ ADC Discontinuous mode on injected channels.
	JQM       CFGR_Bits = 0x01 << 21 //+ ADC JSQR Queue mode.
	AWD1SGL   CFGR_Bits = 0x01 << 22 //+ Enable the watchdog 1 on a single channel or on all channels.
	AWD1EN    CFGR_Bits = 0x01 << 23 //+ ADC Analog watchdog 1 enable on regular Channels.
	JAWD1EN   CFGR_Bits = 0x01 << 24 //+ ADC Analog watchdog 1 enable on injected Channels.
	JAUTO     CFGR_Bits = 0x01 << 25 //+ ADC Automatic injected group conversion.
	AWD1CH    CFGR_Bits = 0x1F << 26 //+ ADC Analog watchdog 1 Channel selection.
	AWD1CH_0  CFGR_Bits = 0x01 << 26 //  ADC AWD1CH bit 0.
	AWD1CH_1  CFGR_Bits = 0x02 << 26 //  ADC AWD1CH bit 1.
	AWD1CH_2  CFGR_Bits = 0x04 << 26 //  ADC AWD1CH bit 2.
	AWD1CH_3  CFGR_Bits = 0x08 << 26 //  ADC AWD1CH bit 3.
	AWD1CH_4  CFGR_Bits = 0x10 << 26 //  ADC AWD1CH bit 4.
)

const (
	DMAENn   = 0
	DMACFGn  = 1
	RESn     = 3
	ALIGNn   = 5
	EXTSELn  = 6
	EXTENn   = 10
	OVRMODn  = 12
	CONTn    = 13
	AUTDLYn  = 14
	DISCENn  = 16
	DISCNUMn = 17
	JDISCENn = 20
	JQMn     = 21
	AWD1SGLn = 22
	AWD1ENn  = 23
	JAWD1ENn = 24
	JAUTOn   = 25
	AWD1CHn  = 26
)

const (
	SMP0   SMPR1_Bits = 0x07 << 0  //+ ADC Channel 0 Sampling time selection.
	SMP0_0 SMPR1_Bits = 0x01 << 0  //  ADC SMP0 bit 0.
	SMP0_1 SMPR1_Bits = 0x02 << 0  //  ADC SMP0 bit 1.
	SMP0_2 SMPR1_Bits = 0x04 << 0  //  ADC SMP0 bit 2.
	SMP1   SMPR1_Bits = 0x07 << 3  //+ ADC Channel 1 Sampling time selection.
	SMP1_0 SMPR1_Bits = 0x01 << 3  //  ADC SMP1 bit 0.
	SMP1_1 SMPR1_Bits = 0x02 << 3  //  ADC SMP1 bit 1.
	SMP1_2 SMPR1_Bits = 0x04 << 3  //  ADC SMP1 bit 2.
	SMP2   SMPR1_Bits = 0x07 << 6  //+ ADC Channel 2 Sampling time selection.
	SMP2_0 SMPR1_Bits = 0x01 << 6  //  ADC SMP2 bit 0.
	SMP2_1 SMPR1_Bits = 0x02 << 6  //  ADC SMP2 bit 1.
	SMP2_2 SMPR1_Bits = 0x04 << 6  //  ADC SMP2 bit 2.
	SMP3   SMPR1_Bits = 0x07 << 9  //+ ADC Channel 3 Sampling time selection.
	SMP3_0 SMPR1_Bits = 0x01 << 9  //  ADC SMP3 bit 0.
	SMP3_1 SMPR1_Bits = 0x02 << 9  //  ADC SMP3 bit 1.
	SMP3_2 SMPR1_Bits = 0x04 << 9  //  ADC SMP3 bit 2.
	SMP4   SMPR1_Bits = 0x07 << 12 //+ ADC Channel 4 Sampling time selection.
	SMP4_0 SMPR1_Bits = 0x01 << 12 //  ADC SMP4 bit 0.
	SMP4_1 SMPR1_Bits = 0x02 << 12 //  ADC SMP4 bit 1.
	SMP4_2 SMPR1_Bits = 0x04 << 12 //  ADC SMP4 bit 2.
	SMP5   SMPR1_Bits = 0x07 << 15 //+ ADC Channel 5 Sampling time selection.
	SMP5_0 SMPR1_Bits = 0x01 << 15 //  ADC SMP5 bit 0.
	SMP5_1 SMPR1_Bits = 0x02 << 15 //  ADC SMP5 bit 1.
	SMP5_2 SMPR1_Bits = 0x04 << 15 //  ADC SMP5 bit 2.
	SMP6   SMPR1_Bits = 0x07 << 18 //+ ADC Channel 6 Sampling time selection.
	SMP6_0 SMPR1_Bits = 0x01 << 18 //  ADC SMP6 bit 0.
	SMP6_1 SMPR1_Bits = 0x02 << 18 //  ADC SMP6 bit 1.
	SMP6_2 SMPR1_Bits = 0x04 << 18 //  ADC SMP6 bit 2.
	SMP7   SMPR1_Bits = 0x07 << 21 //+ ADC Channel 7 Sampling time selection.
	SMP7_0 SMPR1_Bits = 0x01 << 21 //  ADC SMP7 bit 0.
	SMP7_1 SMPR1_Bits = 0x02 << 21 //  ADC SMP7 bit 1.
	SMP7_2 SMPR1_Bits = 0x04 << 21 //  ADC SMP7 bit 2.
	SMP8   SMPR1_Bits = 0x07 << 24 //+ ADC Channel 8 Sampling time selection.
	SMP8_0 SMPR1_Bits = 0x01 << 24 //  ADC SMP8 bit 0.
	SMP8_1 SMPR1_Bits = 0x02 << 24 //  ADC SMP8 bit 1.
	SMP8_2 SMPR1_Bits = 0x04 << 24 //  ADC SMP8 bit 2.
	SMP9   SMPR1_Bits = 0x07 << 27 //+ ADC Channel 9 Sampling time selection.
	SMP9_0 SMPR1_Bits = 0x01 << 27 //  ADC SMP9 bit 0.
	SMP9_1 SMPR1_Bits = 0x02 << 27 //  ADC SMP9 bit 1.
	SMP9_2 SMPR1_Bits = 0x04 << 27 //  ADC SMP9 bit 2.
)

const (
	SMP0n = 0
	SMP1n = 3
	SMP2n = 6
	SMP3n = 9
	SMP4n = 12
	SMP5n = 15
	SMP6n = 18
	SMP7n = 21
	SMP8n = 24
	SMP9n = 27
)

const (
	SMP10   SMPR2_Bits = 0x07 << 0  //+ ADC Channel 10 Sampling time selection.
	SMP10_0 SMPR2_Bits = 0x01 << 0  //  ADC SMP10 bit 0.
	SMP10_1 SMPR2_Bits = 0x02 << 0  //  ADC SMP10 bit 1.
	SMP10_2 SMPR2_Bits = 0x04 << 0  //  ADC SMP10 bit 2.
	SMP11   SMPR2_Bits = 0x07 << 3  //+ ADC Channel 11 Sampling time selection.
	SMP11_0 SMPR2_Bits = 0x01 << 3  //  ADC SMP11 bit 0.
	SMP11_1 SMPR2_Bits = 0x02 << 3  //  ADC SMP11 bit 1.
	SMP11_2 SMPR2_Bits = 0x04 << 3  //  ADC SMP11 bit 2.
	SMP12   SMPR2_Bits = 0x07 << 6  //+ ADC Channel 12 Sampling time selection.
	SMP12_0 SMPR2_Bits = 0x01 << 6  //  ADC SMP12 bit 0.
	SMP12_1 SMPR2_Bits = 0x02 << 6  //  ADC SMP12 bit 1.
	SMP12_2 SMPR2_Bits = 0x04 << 6  //  ADC SMP12 bit 2.
	SMP13   SMPR2_Bits = 0x07 << 9  //+ ADC Channel 13 Sampling time selection.
	SMP13_0 SMPR2_Bits = 0x01 << 9  //  ADC SMP13 bit 0.
	SMP13_1 SMPR2_Bits = 0x02 << 9  //  ADC SMP13 bit 1.
	SMP13_2 SMPR2_Bits = 0x04 << 9  //  ADC SMP13 bit 2.
	SMP14   SMPR2_Bits = 0x07 << 12 //+ ADC Channel 14 Sampling time selection.
	SMP14_0 SMPR2_Bits = 0x01 << 12 //  ADC SMP14 bit 0.
	SMP14_1 SMPR2_Bits = 0x02 << 12 //  ADC SMP14 bit 1.
	SMP14_2 SMPR2_Bits = 0x04 << 12 //  ADC SMP14 bit 2.
	SMP15   SMPR2_Bits = 0x07 << 15 //+ ADC Channel 15 Sampling time selection.
	SMP15_0 SMPR2_Bits = 0x01 << 15 //  ADC SMP15 bit 0.
	SMP15_1 SMPR2_Bits = 0x02 << 15 //  ADC SMP15 bit 1.
	SMP15_2 SMPR2_Bits = 0x04 << 15 //  ADC SMP15 bit 2.
	SMP16   SMPR2_Bits = 0x07 << 18 //+ ADC Channel 16 Sampling time selection.
	SMP16_0 SMPR2_Bits = 0x01 << 18 //  ADC SMP16 bit 0.
	SMP16_1 SMPR2_Bits = 0x02 << 18 //  ADC SMP16 bit 1.
	SMP16_2 SMPR2_Bits = 0x04 << 18 //  ADC SMP16 bit 2.
	SMP17   SMPR2_Bits = 0x07 << 21 //+ ADC Channel 17 Sampling time selection.
	SMP17_0 SMPR2_Bits = 0x01 << 21 //  ADC SMP17 bit 0.
	SMP17_1 SMPR2_Bits = 0x02 << 21 //  ADC SMP17 bit 1.
	SMP17_2 SMPR2_Bits = 0x04 << 21 //  ADC SMP17 bit 2.
	SMP18   SMPR2_Bits = 0x07 << 24 //+ ADC Channel 18 Sampling time selection.
	SMP18_0 SMPR2_Bits = 0x01 << 24 //  ADC SMP18 bit 0.
	SMP18_1 SMPR2_Bits = 0x02 << 24 //  ADC SMP18 bit 1.
	SMP18_2 SMPR2_Bits = 0x04 << 24 //  ADC SMP18 bit 2.
)

const (
	SMP10n = 0
	SMP11n = 3
	SMP12n = 6
	SMP13n = 9
	SMP14n = 12
	SMP15n = 15
	SMP16n = 18
	SMP17n = 21
	SMP18n = 24
)

const (
	LT1    TR1_Bits = 0xFFF << 0  //+ ADC Analog watchdog 1 lower threshold.
	LT1_0  TR1_Bits = 0x01 << 0   //  ADC LT1 bit 0.
	LT1_1  TR1_Bits = 0x02 << 0   //  ADC LT1 bit 1.
	LT1_2  TR1_Bits = 0x04 << 0   //  ADC LT1 bit 2.
	LT1_3  TR1_Bits = 0x08 << 0   //  ADC LT1 bit 3.
	LT1_4  TR1_Bits = 0x10 << 0   //  ADC LT1 bit 4.
	LT1_5  TR1_Bits = 0x20 << 0   //  ADC LT1 bit 5.
	LT1_6  TR1_Bits = 0x40 << 0   //  ADC LT1 bit 6.
	LT1_7  TR1_Bits = 0x80 << 0   //  ADC LT1 bit 7.
	LT1_8  TR1_Bits = 0x100 << 0  //  ADC LT1 bit 8.
	LT1_9  TR1_Bits = 0x200 << 0  //  ADC LT1 bit 9.
	LT1_10 TR1_Bits = 0x400 << 0  //  ADC LT1 bit 10.
	LT1_11 TR1_Bits = 0x800 << 0  //  ADC LT1 bit 11.
	HT1    TR1_Bits = 0xFFF << 16 //+ ADC Analog watchdog 1 higher threshold.
	HT1_0  TR1_Bits = 0x01 << 16  //  ADC HT1 bit 0.
	HT1_1  TR1_Bits = 0x02 << 16  //  ADC HT1 bit 1.
	HT1_2  TR1_Bits = 0x04 << 16  //  ADC HT1 bit 2.
	HT1_3  TR1_Bits = 0x08 << 16  //  ADC HT1 bit 3.
	HT1_4  TR1_Bits = 0x10 << 16  //  ADC HT1 bit 4.
	HT1_5  TR1_Bits = 0x20 << 16  //  ADC HT1 bit 5.
	HT1_6  TR1_Bits = 0x40 << 16  //  ADC HT1 bit 6.
	HT1_7  TR1_Bits = 0x80 << 16  //  ADC HT1 bit 7.
	HT1_8  TR1_Bits = 0x100 << 16 //  ADC HT1 bit 8.
	HT1_9  TR1_Bits = 0x200 << 16 //  ADC HT1 bit 9.
	HT1_10 TR1_Bits = 0x400 << 16 //  ADC HT1 bit 10.
	HT1_11 TR1_Bits = 0x800 << 16 //  ADC HT1 bit 11.
)

const (
	LT1n = 0
	HT1n = 16
)

const (
	LT2   TR2_Bits = 0xFF << 0  //+ ADC Analog watchdog 2 lower threshold.
	LT2_0 TR2_Bits = 0x01 << 0  //  ADC LT2 bit 0.
	LT2_1 TR2_Bits = 0x02 << 0  //  ADC LT2 bit 1.
	LT2_2 TR2_Bits = 0x04 << 0  //  ADC LT2 bit 2.
	LT2_3 TR2_Bits = 0x08 << 0  //  ADC LT2 bit 3.
	LT2_4 TR2_Bits = 0x10 << 0  //  ADC LT2 bit 4.
	LT2_5 TR2_Bits = 0x20 << 0  //  ADC LT2 bit 5.
	LT2_6 TR2_Bits = 0x40 << 0  //  ADC LT2 bit 6.
	LT2_7 TR2_Bits = 0x80 << 0  //  ADC LT2 bit 7.
	HT2   TR2_Bits = 0xFF << 16 //+ ADC Analog watchdog 2 higher threshold.
	HT2_0 TR2_Bits = 0x01 << 16 //  ADC HT2 bit 0.
	HT2_1 TR2_Bits = 0x02 << 16 //  ADC HT2 bit 1.
	HT2_2 TR2_Bits = 0x04 << 16 //  ADC HT2 bit 2.
	HT2_3 TR2_Bits = 0x08 << 16 //  ADC HT2 bit 3.
	HT2_4 TR2_Bits = 0x10 << 16 //  ADC HT2 bit 4.
	HT2_5 TR2_Bits = 0x20 << 16 //  ADC HT2 bit 5.
	HT2_6 TR2_Bits = 0x40 << 16 //  ADC HT2 bit 6.
	HT2_7 TR2_Bits = 0x80 << 16 //  ADC HT2 bit 7.
)

const (
	LT2n = 0
	HT2n = 16
)

const (
	LT3   TR3_Bits = 0xFF << 0  //+ ADC Analog watchdog 3 lower threshold.
	LT3_0 TR3_Bits = 0x01 << 0  //  ADC LT3 bit 0.
	LT3_1 TR3_Bits = 0x02 << 0  //  ADC LT3 bit 1.
	LT3_2 TR3_Bits = 0x04 << 0  //  ADC LT3 bit 2.
	LT3_3 TR3_Bits = 0x08 << 0  //  ADC LT3 bit 3.
	LT3_4 TR3_Bits = 0x10 << 0  //  ADC LT3 bit 4.
	LT3_5 TR3_Bits = 0x20 << 0  //  ADC LT3 bit 5.
	LT3_6 TR3_Bits = 0x40 << 0  //  ADC LT3 bit 6.
	LT3_7 TR3_Bits = 0x80 << 0  //  ADC LT3 bit 7.
	HT3   TR3_Bits = 0xFF << 16 //+ ADC Analog watchdog 3 higher threshold.
	HT3_0 TR3_Bits = 0x01 << 16 //  ADC HT3 bit 0.
	HT3_1 TR3_Bits = 0x02 << 16 //  ADC HT3 bit 1.
	HT3_2 TR3_Bits = 0x04 << 16 //  ADC HT3 bit 2.
	HT3_3 TR3_Bits = 0x08 << 16 //  ADC HT3 bit 3.
	HT3_4 TR3_Bits = 0x10 << 16 //  ADC HT3 bit 4.
	HT3_5 TR3_Bits = 0x20 << 16 //  ADC HT3 bit 5.
	HT3_6 TR3_Bits = 0x40 << 16 //  ADC HT3 bit 6.
	HT3_7 TR3_Bits = 0x80 << 16 //  ADC HT3 bit 7.
)

const (
	LT3n = 0
	HT3n = 16
)

const (
	L     SQR1_Bits = 0x0F << 0  //+ ADC regular channel sequence length.
	L_0   SQR1_Bits = 0x01 << 0  //  ADC L bit 0.
	L_1   SQR1_Bits = 0x02 << 0  //  ADC L bit 1.
	L_2   SQR1_Bits = 0x04 << 0  //  ADC L bit 2.
	L_3   SQR1_Bits = 0x08 << 0  //  ADC L bit 3.
	SQ1   SQR1_Bits = 0x1F << 6  //+ ADC 1st conversion in regular sequence.
	SQ1_0 SQR1_Bits = 0x01 << 6  //  ADC SQ1 bit 0.
	SQ1_1 SQR1_Bits = 0x02 << 6  //  ADC SQ1 bit 1.
	SQ1_2 SQR1_Bits = 0x04 << 6  //  ADC SQ1 bit 2.
	SQ1_3 SQR1_Bits = 0x08 << 6  //  ADC SQ1 bit 3.
	SQ1_4 SQR1_Bits = 0x10 << 6  //  ADC SQ1 bit 4.
	SQ2   SQR1_Bits = 0x1F << 12 //+ ADC 2nd conversion in regular sequence.
	SQ2_0 SQR1_Bits = 0x01 << 12 //  ADC SQ2 bit 0.
	SQ2_1 SQR1_Bits = 0x02 << 12 //  ADC SQ2 bit 1.
	SQ2_2 SQR1_Bits = 0x04 << 12 //  ADC SQ2 bit 2.
	SQ2_3 SQR1_Bits = 0x08 << 12 //  ADC SQ2 bit 3.
	SQ2_4 SQR1_Bits = 0x10 << 12 //  ADC SQ2 bit 4.
	SQ3   SQR1_Bits = 0x1F << 18 //+ ADC 3rd conversion in regular sequence.
	SQ3_0 SQR1_Bits = 0x01 << 18 //  ADC SQ3 bit 0.
	SQ3_1 SQR1_Bits = 0x02 << 18 //  ADC SQ3 bit 1.
	SQ3_2 SQR1_Bits = 0x04 << 18 //  ADC SQ3 bit 2.
	SQ3_3 SQR1_Bits = 0x08 << 18 //  ADC SQ3 bit 3.
	SQ3_4 SQR1_Bits = 0x10 << 18 //  ADC SQ3 bit 4.
	SQ4   SQR1_Bits = 0x1F << 24 //+ ADC 4th conversion in regular sequence.
	SQ4_0 SQR1_Bits = 0x01 << 24 //  ADC SQ4 bit 0.
	SQ4_1 SQR1_Bits = 0x02 << 24 //  ADC SQ4 bit 1.
	SQ4_2 SQR1_Bits = 0x04 << 24 //  ADC SQ4 bit 2.
	SQ4_3 SQR1_Bits = 0x08 << 24 //  ADC SQ4 bit 3.
	SQ4_4 SQR1_Bits = 0x10 << 24 //  ADC SQ4 bit 4.
)

const (
	Ln   = 0
	SQ1n = 6
	SQ2n = 12
	SQ3n = 18
	SQ4n = 24
)

const (
	SQ5   SQR2_Bits = 0x1F << 0  //+ ADC 5th conversion in regular sequence.
	SQ5_0 SQR2_Bits = 0x01 << 0  //  ADC SQ5 bit 0.
	SQ5_1 SQR2_Bits = 0x02 << 0  //  ADC SQ5 bit 1.
	SQ5_2 SQR2_Bits = 0x04 << 0  //  ADC SQ5 bit 2.
	SQ5_3 SQR2_Bits = 0x08 << 0  //  ADC SQ5 bit 3.
	SQ5_4 SQR2_Bits = 0x10 << 0  //  ADC SQ5 bit 4.
	SQ6   SQR2_Bits = 0x1F << 6  //+ ADC 6th conversion in regular sequence.
	SQ6_0 SQR2_Bits = 0x01 << 6  //  ADC SQ6 bit 0.
	SQ6_1 SQR2_Bits = 0x02 << 6  //  ADC SQ6 bit 1.
	SQ6_2 SQR2_Bits = 0x04 << 6  //  ADC SQ6 bit 2.
	SQ6_3 SQR2_Bits = 0x08 << 6  //  ADC SQ6 bit 3.
	SQ6_4 SQR2_Bits = 0x10 << 6  //  ADC SQ6 bit 4.
	SQ7   SQR2_Bits = 0x1F << 12 //+ ADC 7th conversion in regular sequence.
	SQ7_0 SQR2_Bits = 0x01 << 12 //  ADC SQ7 bit 0.
	SQ7_1 SQR2_Bits = 0x02 << 12 //  ADC SQ7 bit 1.
	SQ7_2 SQR2_Bits = 0x04 << 12 //  ADC SQ7 bit 2.
	SQ7_3 SQR2_Bits = 0x08 << 12 //  ADC SQ7 bit 3.
	SQ7_4 SQR2_Bits = 0x10 << 12 //  ADC SQ7 bit 4.
	SQ8   SQR2_Bits = 0x1F << 18 //+ ADC 8th conversion in regular sequence.
	SQ8_0 SQR2_Bits = 0x01 << 18 //  ADC SQ8 bit 0.
	SQ8_1 SQR2_Bits = 0x02 << 18 //  ADC SQ8 bit 1.
	SQ8_2 SQR2_Bits = 0x04 << 18 //  ADC SQ8 bit 2.
	SQ8_3 SQR2_Bits = 0x08 << 18 //  ADC SQ8 bit 3.
	SQ8_4 SQR2_Bits = 0x10 << 18 //  ADC SQ8 bit 4.
	SQ9   SQR2_Bits = 0x1F << 24 //+ ADC 9th conversion in regular sequence.
	SQ9_0 SQR2_Bits = 0x01 << 24 //  ADC SQ9 bit 0.
	SQ9_1 SQR2_Bits = 0x02 << 24 //  ADC SQ9 bit 1.
	SQ9_2 SQR2_Bits = 0x04 << 24 //  ADC SQ9 bit 2.
	SQ9_3 SQR2_Bits = 0x08 << 24 //  ADC SQ9 bit 3.
	SQ9_4 SQR2_Bits = 0x10 << 24 //  ADC SQ9 bit 4.
)

const (
	SQ5n = 0
	SQ6n = 6
	SQ7n = 12
	SQ8n = 18
	SQ9n = 24
)

const (
	SQ10   SQR3_Bits = 0x1F << 0  //+ ADC 10th conversion in regular sequence.
	SQ10_0 SQR3_Bits = 0x01 << 0  //  ADC SQ10 bit 0.
	SQ10_1 SQR3_Bits = 0x02 << 0  //  ADC SQ10 bit 1.
	SQ10_2 SQR3_Bits = 0x04 << 0  //  ADC SQ10 bit 2.
	SQ10_3 SQR3_Bits = 0x08 << 0  //  ADC SQ10 bit 3.
	SQ10_4 SQR3_Bits = 0x10 << 0  //  ADC SQ10 bit 4.
	SQ11   SQR3_Bits = 0x1F << 6  //+ ADC 11th conversion in regular sequence.
	SQ11_0 SQR3_Bits = 0x01 << 6  //  ADC SQ11 bit 0.
	SQ11_1 SQR3_Bits = 0x02 << 6  //  ADC SQ11 bit 1.
	SQ11_2 SQR3_Bits = 0x04 << 6  //  ADC SQ11 bit 2.
	SQ11_3 SQR3_Bits = 0x08 << 6  //  ADC SQ11 bit 3.
	SQ11_4 SQR3_Bits = 0x10 << 6  //  ADC SQ11 bit 4.
	SQ12   SQR3_Bits = 0x1F << 12 //+ ADC 12th conversion in regular sequence.
	SQ12_0 SQR3_Bits = 0x01 << 12 //  ADC SQ12 bit 0.
	SQ12_1 SQR3_Bits = 0x02 << 12 //  ADC SQ12 bit 1.
	SQ12_2 SQR3_Bits = 0x04 << 12 //  ADC SQ12 bit 2.
	SQ12_3 SQR3_Bits = 0x08 << 12 //  ADC SQ12 bit 3.
	SQ12_4 SQR3_Bits = 0x10 << 12 //  ADC SQ12 bit 4.
	SQ13   SQR3_Bits = 0x1F << 18 //+ ADC 13th conversion in regular sequence.
	SQ13_0 SQR3_Bits = 0x01 << 18 //  ADC SQ13 bit 0.
	SQ13_1 SQR3_Bits = 0x02 << 18 //  ADC SQ13 bit 1.
	SQ13_2 SQR3_Bits = 0x04 << 18 //  ADC SQ13 bit 2.
	SQ13_3 SQR3_Bits = 0x08 << 18 //  ADC SQ13 bit 3.
	SQ13_4 SQR3_Bits = 0x10 << 18 //  ADC SQ13 bit 4.
	SQ14   SQR3_Bits = 0x1F << 24 //+ ADC 14th conversion in regular sequence.
	SQ14_0 SQR3_Bits = 0x01 << 24 //  ADC SQ14 bit 0.
	SQ14_1 SQR3_Bits = 0x02 << 24 //  ADC SQ14 bit 1.
	SQ14_2 SQR3_Bits = 0x04 << 24 //  ADC SQ14 bit 2.
	SQ14_3 SQR3_Bits = 0x08 << 24 //  ADC SQ14 bit 3.
	SQ14_4 SQR3_Bits = 0x10 << 24 //  ADC SQ14 bit 4.
)

const (
	SQ10n = 0
	SQ11n = 6
	SQ12n = 12
	SQ13n = 18
	SQ14n = 24
)

const (
	SQ15   SQR4_Bits = 0x1F << 0 //+ ADC 15th conversion in regular sequence.
	SQ15_0 SQR4_Bits = 0x01 << 0 //  ADC SQ15 bit 0.
	SQ15_1 SQR4_Bits = 0x02 << 0 //  ADC SQ15 bit 1.
	SQ15_2 SQR4_Bits = 0x04 << 0 //  ADC SQ15 bit 2.
	SQ15_3 SQR4_Bits = 0x08 << 0 //  ADC SQ15 bit 3.
	SQ15_4 SQR4_Bits = 0x10 << 0 //  ADC SQ105 bit 4.
	SQ16   SQR4_Bits = 0x1F << 6 //+ ADC 16th conversion in regular sequence.
	SQ16_0 SQR4_Bits = 0x01 << 6 //  ADC SQ16 bit 0.
	SQ16_1 SQR4_Bits = 0x02 << 6 //  ADC SQ16 bit 1.
	SQ16_2 SQR4_Bits = 0x04 << 6 //  ADC SQ16 bit 2.
	SQ16_3 SQR4_Bits = 0x08 << 6 //  ADC SQ16 bit 3.
	SQ16_4 SQR4_Bits = 0x10 << 6 //  ADC SQ16 bit 4.
)

const (
	SQ15n = 0
	SQ16n = 6
)

const (
	RDATA    DR_Bits = 0xFFFF << 0 //+ ADC regular Data converted.
	RDATA_0  DR_Bits = 0x01 << 0   //  ADC RDATA bit 0.
	RDATA_1  DR_Bits = 0x02 << 0   //  ADC RDATA bit 1.
	RDATA_2  DR_Bits = 0x04 << 0   //  ADC RDATA bit 2.
	RDATA_3  DR_Bits = 0x08 << 0   //  ADC RDATA bit 3.
	RDATA_4  DR_Bits = 0x10 << 0   //  ADC RDATA bit 4.
	RDATA_5  DR_Bits = 0x20 << 0   //  ADC RDATA bit 5.
	RDATA_6  DR_Bits = 0x40 << 0   //  ADC RDATA bit 6.
	RDATA_7  DR_Bits = 0x80 << 0   //  ADC RDATA bit 7.
	RDATA_8  DR_Bits = 0x100 << 0  //  ADC RDATA bit 8.
	RDATA_9  DR_Bits = 0x200 << 0  //  ADC RDATA bit 9.
	RDATA_10 DR_Bits = 0x400 << 0  //  ADC RDATA bit 10.
	RDATA_11 DR_Bits = 0x800 << 0  //  ADC RDATA bit 11.
	RDATA_12 DR_Bits = 0x1000 << 0 //  ADC RDATA bit 12.
	RDATA_13 DR_Bits = 0x2000 << 0 //  ADC RDATA bit 13.
	RDATA_14 DR_Bits = 0x4000 << 0 //  ADC RDATA bit 14.
	RDATA_15 DR_Bits = 0x8000 << 0 //  ADC RDATA bit 15.
)

const (
	RDATAn = 0
)

const (
	JL        JSQR_Bits = 0x03 << 0  //+ ADC injected channel sequence length.
	JL_0      JSQR_Bits = 0x01 << 0  //  ADC JL bit 0.
	JL_1      JSQR_Bits = 0x02 << 0  //  ADC JL bit 1.
	JEXTSEL   JSQR_Bits = 0x0F << 2  //+ ADC external trigger selection for injected group.
	JEXTSEL_0 JSQR_Bits = 0x01 << 2  //  ADC JEXTSEL bit 0.
	JEXTSEL_1 JSQR_Bits = 0x02 << 2  //  ADC JEXTSEL bit 1.
	JEXTSEL_2 JSQR_Bits = 0x04 << 2  //  ADC JEXTSEL bit 2.
	JEXTSEL_3 JSQR_Bits = 0x08 << 2  //  ADC JEXTSEL bit 3.
	JEXTEN    JSQR_Bits = 0x03 << 6  //+ ADC external trigger enable and polarity selection for injected channels.
	JEXTEN_0  JSQR_Bits = 0x01 << 6  //  ADC JEXTEN bit 0.
	JEXTEN_1  JSQR_Bits = 0x02 << 6  //  ADC JEXTEN bit 1.
	JSQ1      JSQR_Bits = 0x1F << 8  //+ ADC 1st conversion in injected sequence.
	JSQ1_0    JSQR_Bits = 0x01 << 8  //  ADC JSQ1 bit 0.
	JSQ1_1    JSQR_Bits = 0x02 << 8  //  ADC JSQ1 bit 1.
	JSQ1_2    JSQR_Bits = 0x04 << 8  //  ADC JSQ1 bit 2.
	JSQ1_3    JSQR_Bits = 0x08 << 8  //  ADC JSQ1 bit 3.
	JSQ1_4    JSQR_Bits = 0x10 << 8  //  ADC JSQ1 bit 4.
	JSQ2      JSQR_Bits = 0x1F << 14 //+ ADC 2nd conversion in injected sequence.
	JSQ2_0    JSQR_Bits = 0x01 << 14 //  ADC JSQ2 bit 0.
	JSQ2_1    JSQR_Bits = 0x02 << 14 //  ADC JSQ2 bit 1.
	JSQ2_2    JSQR_Bits = 0x04 << 14 //  ADC JSQ2 bit 2.
	JSQ2_3    JSQR_Bits = 0x08 << 14 //  ADC JSQ2 bit 3.
	JSQ2_4    JSQR_Bits = 0x10 << 14 //  ADC JSQ2 bit 4.
	JSQ3      JSQR_Bits = 0x1F << 20 //+ ADC 3rd conversion in injected sequence.
	JSQ3_0    JSQR_Bits = 0x01 << 20 //  ADC JSQ3 bit 0.
	JSQ3_1    JSQR_Bits = 0x02 << 20 //  ADC JSQ3 bit 1.
	JSQ3_2    JSQR_Bits = 0x04 << 20 //  ADC JSQ3 bit 2.
	JSQ3_3    JSQR_Bits = 0x08 << 20 //  ADC JSQ3 bit 3.
	JSQ3_4    JSQR_Bits = 0x10 << 20 //  ADC JSQ3 bit 4.
	JSQ4      JSQR_Bits = 0x1F << 26 //+ ADC 4th conversion in injected sequence.
	JSQ4_0    JSQR_Bits = 0x01 << 26 //  ADC JSQ4 bit 0.
	JSQ4_1    JSQR_Bits = 0x02 << 26 //  ADC JSQ4 bit 1.
	JSQ4_2    JSQR_Bits = 0x04 << 26 //  ADC JSQ4 bit 2.
	JSQ4_3    JSQR_Bits = 0x08 << 26 //  ADC JSQ4 bit 3.
	JSQ4_4    JSQR_Bits = 0x10 << 26 //  ADC JSQ4 bit 4.
)

const (
	JLn      = 0
	JEXTSELn = 2
	JEXTENn  = 6
	JSQ1n    = 8
	JSQ2n    = 14
	JSQ3n    = 20
	JSQ4n    = 26
)

const (
	OFFSET1      OFR_Bits = 0xFFF << 0 //+ ADC data offset 1 for channel programmed into bits OFFSET1_CH[4:0].
	OFFSET1_0    OFR_Bits = 0x01 << 0  //  ADC OFFSET1 bit 0.
	OFFSET1_1    OFR_Bits = 0x02 << 0  //  ADC OFFSET1 bit 1.
	OFFSET1_2    OFR_Bits = 0x04 << 0  //  ADC OFFSET1 bit 2.
	OFFSET1_3    OFR_Bits = 0x08 << 0  //  ADC OFFSET1 bit 3.
	OFFSET1_4    OFR_Bits = 0x10 << 0  //  ADC OFFSET1 bit 4.
	OFFSET1_5    OFR_Bits = 0x20 << 0  //  ADC OFFSET1 bit 5.
	OFFSET1_6    OFR_Bits = 0x40 << 0  //  ADC OFFSET1 bit 6.
	OFFSET1_7    OFR_Bits = 0x80 << 0  //  ADC OFFSET1 bit 7.
	OFFSET1_8    OFR_Bits = 0x100 << 0 //  ADC OFFSET1 bit 8.
	OFFSET1_9    OFR_Bits = 0x200 << 0 //  ADC OFFSET1 bit 9.
	OFFSET1_10   OFR_Bits = 0x400 << 0 //  ADC OFFSET1 bit 10.
	OFFSET1_11   OFR_Bits = 0x800 << 0 //  ADC OFFSET1 bit 11.
	OFFSET1_CH   OFR_Bits = 0x1F << 26 //+ ADC Channel selection for the data offset 1.
	OFFSET1_CH_0 OFR_Bits = 0x01 << 26 //  ADC OFFSET1_CH bit 0.
	OFFSET1_CH_1 OFR_Bits = 0x02 << 26 //  ADC OFFSET1_CH bit 1.
	OFFSET1_CH_2 OFR_Bits = 0x04 << 26 //  ADC OFFSET1_CH bit 2.
	OFFSET1_CH_3 OFR_Bits = 0x08 << 26 //  ADC OFFSET1_CH bit 3.
	OFFSET1_CH_4 OFR_Bits = 0x10 << 26 //  ADC OFFSET1_CH bit 4.
	OFFSET1_EN   OFR_Bits = 0x01 << 31 //+ ADC offset 1 enable.
)

const (
	OFFSET1n    = 0
	OFFSET1_CHn = 26
	OFFSET1_ENn = 31
)

const (
	JDATA    JDR_Bits = 0xFFFF << 0 //+ ADC Injected DATA.
	JDATA_0  JDR_Bits = 0x01 << 0   //  ADC JDATA bit 0.
	JDATA_1  JDR_Bits = 0x02 << 0   //  ADC JDATA bit 1.
	JDATA_2  JDR_Bits = 0x04 << 0   //  ADC JDATA bit 2.
	JDATA_3  JDR_Bits = 0x08 << 0   //  ADC JDATA bit 3.
	JDATA_4  JDR_Bits = 0x10 << 0   //  ADC JDATA bit 4.
	JDATA_5  JDR_Bits = 0x20 << 0   //  ADC JDATA bit 5.
	JDATA_6  JDR_Bits = 0x40 << 0   //  ADC JDATA bit 6.
	JDATA_7  JDR_Bits = 0x80 << 0   //  ADC JDATA bit 7.
	JDATA_8  JDR_Bits = 0x100 << 0  //  ADC JDATA bit 8.
	JDATA_9  JDR_Bits = 0x200 << 0  //  ADC JDATA bit 9.
	JDATA_10 JDR_Bits = 0x400 << 0  //  ADC JDATA bit 10.
	JDATA_11 JDR_Bits = 0x800 << 0  //  ADC JDATA bit 11.
	JDATA_12 JDR_Bits = 0x1000 << 0 //  ADC JDATA bit 12.
	JDATA_13 JDR_Bits = 0x2000 << 0 //  ADC JDATA bit 13.
	JDATA_14 JDR_Bits = 0x4000 << 0 //  ADC JDATA bit 14.
	JDATA_15 JDR_Bits = 0x8000 << 0 //  ADC JDATA bit 15.
)

const (
	JDATAn = 0
)

const (
	AWD2CH    AWD2CR_Bits = 0x3FFFF << 1 //+ ADC Analog watchdog 2 channel selection.
	AWD2CH_0  AWD2CR_Bits = 0x01 << 1    //  ADC AWD2CH bit 0.
	AWD2CH_1  AWD2CR_Bits = 0x02 << 1    //  ADC AWD2CH bit 1.
	AWD2CH_2  AWD2CR_Bits = 0x04 << 1    //  ADC AWD2CH bit 2.
	AWD2CH_3  AWD2CR_Bits = 0x08 << 1    //  ADC AWD2CH bit 3.
	AWD2CH_4  AWD2CR_Bits = 0x10 << 1    //  ADC AWD2CH bit 4.
	AWD2CH_5  AWD2CR_Bits = 0x20 << 1    //  ADC AWD2CH bit 5.
	AWD2CH_6  AWD2CR_Bits = 0x40 << 1    //  ADC AWD2CH bit 6.
	AWD2CH_7  AWD2CR_Bits = 0x80 << 1    //  ADC AWD2CH bit 7.
	AWD2CH_8  AWD2CR_Bits = 0x100 << 1   //  ADC AWD2CH bit 8.
	AWD2CH_9  AWD2CR_Bits = 0x200 << 1   //  ADC AWD2CH bit 9.
	AWD2CH_10 AWD2CR_Bits = 0x400 << 1   //  ADC AWD2CH bit 10.
	AWD2CH_11 AWD2CR_Bits = 0x800 << 1   //  ADC AWD2CH bit 11.
	AWD2CH_12 AWD2CR_Bits = 0x1000 << 1  //  ADC AWD2CH bit 12.
	AWD2CH_13 AWD2CR_Bits = 0x2000 << 1  //  ADC AWD2CH bit 13.
	AWD2CH_14 AWD2CR_Bits = 0x4000 << 1  //  ADC AWD2CH bit 14.
	AWD2CH_15 AWD2CR_Bits = 0x8000 << 1  //  ADC AWD2CH bit 15.
	AWD2CH_16 AWD2CR_Bits = 0x10000 << 1 //  ADC AWD2CH bit 16.
	AWD2CH_17 AWD2CR_Bits = 0x18000 << 1 //  ADC AWD2CH bit 17.
)

const (
	AWD2CHn = 1
)

const (
	AWD3CH    AWD3CR_Bits = 0x3FFFF << 1 //+ ADC Analog watchdog 2 channel selection.
	AWD3CH_0  AWD3CR_Bits = 0x01 << 1    //  ADC AWD3CH bit 0.
	AWD3CH_1  AWD3CR_Bits = 0x02 << 1    //  ADC AWD3CH bit 1.
	AWD3CH_2  AWD3CR_Bits = 0x04 << 1    //  ADC AWD3CH bit 2.
	AWD3CH_3  AWD3CR_Bits = 0x08 << 1    //  ADC AWD3CH bit 3.
	AWD3CH_4  AWD3CR_Bits = 0x10 << 1    //  ADC AWD3CH bit 4.
	AWD3CH_5  AWD3CR_Bits = 0x20 << 1    //  ADC AWD3CH bit 5.
	AWD3CH_6  AWD3CR_Bits = 0x40 << 1    //  ADC AWD3CH bit 6.
	AWD3CH_7  AWD3CR_Bits = 0x80 << 1    //  ADC AWD3CH bit 7.
	AWD3CH_8  AWD3CR_Bits = 0x100 << 1   //  ADC AWD3CH bit 8.
	AWD3CH_9  AWD3CR_Bits = 0x200 << 1   //  ADC AWD3CH bit 9.
	AWD3CH_10 AWD3CR_Bits = 0x400 << 1   //  ADC AWD3CH bit 10.
	AWD3CH_11 AWD3CR_Bits = 0x800 << 1   //  ADC AWD3CH bit 11.
	AWD3CH_12 AWD3CR_Bits = 0x1000 << 1  //  ADC AWD3CH bit 12.
	AWD3CH_13 AWD3CR_Bits = 0x2000 << 1  //  ADC AWD3CH bit 13.
	AWD3CH_14 AWD3CR_Bits = 0x4000 << 1  //  ADC AWD3CH bit 14.
	AWD3CH_15 AWD3CR_Bits = 0x8000 << 1  //  ADC AWD3CH bit 15.
	AWD3CH_16 AWD3CR_Bits = 0x10000 << 1 //  ADC AWD3CH bit 16.
	AWD3CH_17 AWD3CR_Bits = 0x18000 << 1 //  ADC AWD3CH bit 17.
)

const (
	AWD3CHn = 1
)

const (
	CALFACT_S   CALFACT_Bits = 0x7F << 0  //+ ADC calibration factors in single-ended mode.
	CALFACT_S_0 CALFACT_Bits = 0x01 << 0  //  ADC CALFACT_S bit 0.
	CALFACT_S_1 CALFACT_Bits = 0x02 << 0  //  ADC CALFACT_S bit 1.
	CALFACT_S_2 CALFACT_Bits = 0x04 << 0  //  ADC CALFACT_S bit 2.
	CALFACT_S_3 CALFACT_Bits = 0x08 << 0  //  ADC CALFACT_S bit 3.
	CALFACT_S_4 CALFACT_Bits = 0x10 << 0  //  ADC CALFACT_S bit 4.
	CALFACT_S_5 CALFACT_Bits = 0x20 << 0  //  ADC CALFACT_S bit 5.
	CALFACT_S_6 CALFACT_Bits = 0x40 << 0  //  ADC CALFACT_S bit 6.
	CALFACT_D   CALFACT_Bits = 0x7F << 16 //+ ADC calibration factors in differential mode.
	CALFACT_D_0 CALFACT_Bits = 0x01 << 16 //  ADC CALFACT_D bit 0.
	CALFACT_D_1 CALFACT_Bits = 0x02 << 16 //  ADC CALFACT_D bit 1.
	CALFACT_D_2 CALFACT_Bits = 0x04 << 16 //  ADC CALFACT_D bit 2.
	CALFACT_D_3 CALFACT_Bits = 0x08 << 16 //  ADC CALFACT_D bit 3.
	CALFACT_D_4 CALFACT_Bits = 0x10 << 16 //  ADC CALFACT_D bit 4.
	CALFACT_D_5 CALFACT_Bits = 0x20 << 16 //  ADC CALFACT_D bit 5.
	CALFACT_D_6 CALFACT_Bits = 0x40 << 16 //  ADC CALFACT_D bit 6.
)

const (
	CALFACT_Sn = 0
	CALFACT_Dn = 16
)
