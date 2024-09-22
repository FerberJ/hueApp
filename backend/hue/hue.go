package hue

import (
	"context"
	"fmt"
	"openHueApp/backend/models"

	"github.com/openhue/openhue-go"
)

type Hue struct {
	ctx    context.Context
	Home   *openhue.Home
	Rooms  []models.Room
	Lights []models.Light
}

func NewHue() *Hue {
	home, _ := GetHome()
	return &Hue{
		Home: home,
	}
}

func (hue *Hue) Startup(ctx context.Context) {
	hue.ctx = ctx
}

func GetHome() (*openhue.Home, error) {
	return openhue.NewHome(openhue.LoadConf())
}

func (hue *Hue) GetLights() ([]models.Light, error) {
	if hue.Home == nil {
		return nil, fmt.Errorf("home was not initilised")
	}

	lights, err := hue.Home.GetLights()
	if err != nil {
		return nil, err
	}

	hue.Lights = []models.Light{}

	for _, light := range lights {
		hue.Lights = append(hue.Lights, models.Light{
			Model: models.Model{
				ID: *light.Id,
			},
			Name: *light.Metadata.Name,
			On:   light.IsOn(),
		})
	}

	return hue.Lights, nil
}

func (hue *Hue) ToggleLight(light models.Light) error {
	if hue.Home == nil {
		return fmt.Errorf("home was not initilised")
	}

	hue.Home.UpdateLight(light.ID, openhue.LightPut{
		On: &openhue.On{
			On: func() *bool {
				toggled := !light.On
				return &toggled
			}(),
		},
	})

	for _, hueLight := range hue.Lights {
		if hueLight.ID == light.ID {
			hueLight.On = !hueLight.On
		}
	}

	return nil
}
