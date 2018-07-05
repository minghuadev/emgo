package main

import (
	"bufio"
	"delay"
	"fmt"
	"rtos"
	"text/linewriter"

	"sdcard"
	"sdcard/sdio"

	"stm32/hal/dma"
	"stm32/hal/exti"
	"stm32/hal/gpio"
	"stm32/hal/irq"
	"stm32/hal/sdmmc"
	"stm32/hal/system"
	"stm32/hal/system/timer/systick"
	"stm32/hal/usart"

	"stm32/hal/raw/pwr"
	"stm32/hal/raw/rcc"
)

var (
	led     gpio.Pin
	sd      *sdmmc.DriverDMA
	tts     *usart.Driver
	bcmRSTn gpio.Pin
	bcmD1   gpio.Pin
)

func init() {
	system.Setup96(26)
	systick.Setup(2e6)

	// GPIO

	gpio.A.EnableClock(false)
	//bcmIRQ := gpio.A.Pin(0)
	tx2 := gpio.A.Pin(2)
	rx2 := gpio.A.Pin(3)
	led = gpio.A.Pin(4)
	bcmCMD := gpio.A.Pin(6)
	//flashMOSI = gpio.A.Pin(7)
	bcmD1 = gpio.A.Pin(8) // Also LSE output (MCO1) to WLAN powersave clock.
	bcmD2 := gpio.A.Pin(9)
	//flashCSn := gpio.A.Pin(15)

	gpio.B.EnableClock(true)
	//flashSCK := gpio.B.Pin(3)
	//flashMISO := gpio.B.Pin(4)
	bcmD3 := gpio.B.Pin(5)
	bcmD0 := gpio.B.Pin(7)
	bcmRSTn = gpio.B.Pin(14)
	bcmCLK := gpio.B.Pin(15)

	// LED

	led.Set()
	led.Setup(&gpio.Config{
		Mode:   gpio.Out,
		Driver: gpio.OpenDrain,
		Speed:  gpio.Low,
	})

	// USART2

	tx2.Setup(&gpio.Config{Mode: gpio.Alt})
	rx2.Setup(&gpio.Config{Mode: gpio.AltIn, Pull: gpio.PullUp})
	tx2.SetAltFunc(gpio.USART2)
	rx2.SetAltFunc(gpio.USART2)
	d := dma.DMA1
	d.EnableClock(true) // DMA clock must remain enabled in sleep mode.
	tts = usart.NewDriver(
		usart.USART2, d.Channel(6, 4), d.Channel(5, 4), make([]byte, 88),
	)
	tts.Periph().EnableClock(true)
	tts.Periph().SetBaudRate(115200)
	tts.Periph().Enable()
	tts.EnableRx()
	tts.EnableTx()
	rtos.IRQ(irq.USART2).Enable()
	rtos.IRQ(irq.DMA1_Stream5).Enable()
	rtos.IRQ(irq.DMA1_Stream6).Enable()
	fmt.DefaultWriter = linewriter.New(
		bufio.NewWriterSize(tts, 88),
		linewriter.CRLF,
	)

	// WLAN (BCM43362: SDIO, reset, IRQ)

	bcmRSTn.Setup(&gpio.Config{Mode: gpio.Out, Speed: gpio.Low})

	cfg := &gpio.Config{Mode: gpio.Alt, Speed: gpio.VeryHigh, Pull: gpio.PullUp}
	for _, pin := range []gpio.Pin{bcmCLK, bcmCMD, bcmD0, bcmD1, bcmD2, bcmD3} {
		pin.Setup(cfg)
		pin.SetAltFunc(gpio.SDIO)
	}
	d = dma.DMA2
	d.EnableClock(true)
	sd = sdmmc.NewDriverDMA(sdmmc.SDIO, d.Channel(6, 4), bcmD0)
	sd.Periph().EnableClock(true)
	sd.Periph().Enable()
	rtos.IRQ(irq.SDIO).Enable()
	rtos.IRQ(irq.EXTI9_5).Enable()
}

func checkErr(what string, err error) {
	if err == nil {
		return
	}
	fmt.Printf("%s: %v\n", what, err)
	for {
		//led.Clear()
		delay.Millisec(100)
		//led.Set()
		delay.Millisec(100)
	}
}

func main() {
	fmt.Printf("Try to communicate with BCM43362:\n")

	// Initialize WLAN

	bcmRSTn.Store(0) // Set WLAN into reset state.

	sd.SetBusWidth(sdcard.Bus1)
	sd.SetClock(400e3, true)

	// Provide WLAN powersave clock on PA8 (SDIO_D1).
	RCC := rcc.RCC
	PWR := pwr.PWR
	RCC.PWREN().Set()
	PWR.DBP().Set()
	RCC.LSEON().Set()
	for RCC.LSERDY().Load() == 0 {
		led.Clear()
		delay.Millisec(50)
		led.Set()
		delay.Millisec(50)
	}
	RCC.MCO1().Store(1 << rcc.MCO1n) // LSE on MCO1.
	PWR.DBP().Clear()
	RCC.PWREN().Clear()
	bcmD1.SetAltFunc(gpio.MCO)

	delay.Millisec(1)
	bcmRSTn.Store(1)

	var (
		rca uint16
		cs  sdcard.CardStatus
	)
	led.Clear()
	for {
		delay.Millisec(1)
		sd.SendCmd(sdcard.CMD0())
		checkErr("CMD0", sd.Err(true))
		sd.SendCmd(sdcard.CMD5(0))
		sd.Err(true)
		rca, cs = sd.SendCmd(sdcard.CMD3()).R6()
		if sd.Err(true) == nil {
			break
		}
	}
	led.Set()
	fmt.Printf("CMD5: rca=%x cs=%s\n", rca, cs)

	cs = sd.SendCmd(sdcard.CMD7(rca)).R1()
	checkErr("CMD7", sd.Err(true))
	fmt.Printf("CMD7: cs=%s\n", cs)

	fmt.Printf("Enable FN1: ")
	for {
		ioen, st := sd.SendCmd(sdcard.CMD52(
			sdio.CIA, sdio.CCCR_IOEN, sdcard.Write|sdcard.RAW, sdio.FN1,
		)).R5()
		checkErr("CMD52", sd.Err(true))
		if st&^sdcard.IO_CURRENT_STATE != 0 {
			fmt.Println(st)
			return
		}
		if ioen&sdio.FN1 != 0 {
			break
		}
		delay.Millisec(1)
	}
	fmt.Printf("OK\n")
}

func ttsISR() {
	tts.ISR()
}

func ttsRxDMAISR() {
	tts.RxDMAISR()
}

func ttsTxDMAISR() {
	tts.TxDMAISR()
}

func sdioISR() {
	sd.ISR()
}

func exti9_5ISR() {
	pending := exti.Pending() & 0x3E0
	pending.ClearPending()
	if pending&sd.BusyLine() != 0 {
		sd.BusyISR()
	}
}

//c:__attribute__((section(".ISRs")))
var ISRs = [...]func(){
	irq.USART2:       ttsISR,
	irq.DMA1_Stream5: ttsRxDMAISR,
	irq.DMA1_Stream6: ttsTxDMAISR,

	irq.SDIO:    sdioISR,
	irq.EXTI9_5: exti9_5ISR,
}
