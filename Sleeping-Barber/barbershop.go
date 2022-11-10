package main

import (
	"github.com/fatih/color"
	"time"
)

type Barbershop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *Barbershop) AddBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for the clients.", barber)

		for {
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do %s takes a nap.", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up and starts cutting %s's hair.", client, barber)
					isSleeping = false
				}

				// cut hair
				shop.CutHair(barber, client)

			} else {
				// shop is closed, so send the barber home and close this goroutine
				shop.SendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *Barbershop) CutHair(barber, client string) {
	color.Green("%s is getting a haircut from %s.", client, barber)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is done getting a haircut from %s.", client, barber)
}

func (shop *Barbershop) SendBarberHome(barber string) {
	color.Cyan("%s is going home for the day.", barber)
	shop.BarbersDoneChan <- true
}

func (shop *Barbershop) CloseShopForTheDay() {
	color.Cyan("The shop is closing for the day.")
	close(shop.ClientsChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}
	close(shop.BarbersDoneChan)

	color.Green("-------------------------------")
	color.Green("The shop is closed for the day.")
}

func (shop *Barbershop) AddClient(client string) {
	color.Green("%s arrives!", client)
	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room.", client)
		default:
			color.Red("%s leaves because the shop is full.", client)
		}
	} else {
		color.Red("The shop is already closed, so % leaves", client)
	}
	shop.ClientsChan <- client
}
