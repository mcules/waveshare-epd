package main

import (
	"fmt"
	"os"

	"github.com/mcules/waveshare-epd"
	"github.com/mcules/waveshare-epd/epd_config"
	"github.com/mcules/waveshare-epd/imageutil"
)

func main() {

	e := epd.Epd{
		Config: epd_config.EpdConfig{},
	}
	e.Setup()
	e.Clear()

	img, err := imageutil.OpenImage("../imageutil/test/test_sb.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	e.Display(&img, epd.MODE_MONO_DITHER_ON)

	e.Sleep()

}
