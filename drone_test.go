package main

import "testing"

func TestDroneRotateLeft(t *testing.T) {
	d := Drone{0, 0, "N", 10}
	d.RotateLeft()
	e := Drone{0, 0, "W", 10}
	if d != e {
        t.Fatalf(`%q != %q`, e, d)
    }
}

func TestDroneRotateRight(t *testing.T) {
	d := Drone{0, 0, "N", 10}
	d.RotateRight()
	e := Drone{0, 0, "E", 10}
	if d != e {
        t.Fatalf(`%q != %q`, e, d)
    }
}

func TestDroneRotateRightAndGoFoward(t *testing.T) {
	d := Drone{0, 0, "N", 10}
	d.RotateRight()
	d.GoFowardSingle()
	e := Drone{1, 0, "E", 10}
	if d != e {
        t.Fatalf(`%q != %q`, e, d)
    }
}

func TestDroneRotateLeftAndGoFoward(t *testing.T) {
	d := Drone{0, 0, "N", 10}
	d.RotateLeft()
	d.GoFowardSingle()
	e := Drone{-1, 0, "W", 10}
	if d != e {
        t.Fatalf(`%q != %q`, e, d)
    }
}

func TestDrone180RotateAndFoward(t *testing.T) {
	d := Drone{0, 0, "N", 10}
	d.RotateRight()
	d.RotateRight()
	d.GoFowardSingle()
	e := Drone{0, 0, "N", 10}
	e.RotateLeft()
	e.RotateLeft()
	e.GoFowardSingle()
	if d != e {
        t.Fatalf(`%q != %q`, e, d)
    }
}

func TestDroneLimitExceeded(t *testing.T) {
	limit := 4
	d := Drone{0, 0, "N", limit}
	d.RotateLeft()
	for i := 0; i < 6; i++ {
		err := d.GoFowardSingle()
		if err == nil && i == limit {
			t.Fatalf(`error has not been raised`)
		}
	}
}

func TestDroneReport(t *testing.T) {
	d := Drone{0, 0, "N", 10}
	d.RotateLeft()
	d.GoFowardSingle()
	exp := d.Report()
	res := "(-1, 0) dirección Occidente\n"
	if exp != res {
        t.Fatalf(`%q != %q`, exp, res)
    }
}