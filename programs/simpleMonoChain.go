package main

import (
	"epd"
	"epd/epd_config"
	"epd/imageutil"
	"fmt"
	"os"
	"time"
)

func main() {

	e := epd.Epd{
		Config: epd_config.EpdConfig{},
	}
	e.Setup()
	e.Clear()

	img, err := imageutil.OpenImage("../imageutil/test/test_portrait.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	e.Display(&img, epd.MODE_MONO_DITHER_OFF)

	time.Sleep(5 * time.Second)

	img, err = imageutil.OpenImage("../imageutil/test/test_portrait.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	e.Display(&img, epd.MODE_MONO_DITHER_ON)

	e.Sleep()

}
