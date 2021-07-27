package main

import (
    "bufio"
    "log"
    "os"
	"errors"
	"fmt"
)

const LeftSymbol = "I" 
const RigthSymbol = "D" 
const FwdSymbol = "A" 

func parseLine(line string, limit int) (Drone, error) {
	d := Drone{0, 0, "N", limit}
	for _, char := range line {
		switch c := string(char); c {
			case LeftSymbol: // Rotate Left
				d.RotateLeft()
			case RigthSymbol: // Rotate Rigth
				d.RotateRight()
			case FwdSymbol:
				err := d.GoFowardSingle()
				if err != nil {
					return Drone{}, err
				}
			default:
				return Drone{}, errors.New(fmt.Sprintf("Symbol '%s' not recognized in '%s'", c, line))
		}
	}
	return d, nil
}

func processDroneRoute(inPath string, outPath string, limit int) {
	file, err := os.Open(inPath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	out, err := os.Create(outPath)
    if err != nil {
        log.Fatal(err)
    }
	defer out.Close()

    scanner := bufio.NewScanner(file)
	PrintOkMsg(fmt.Sprintf("== Reporte de Entregas - (%s) ==", outPath))
    for scanner.Scan() {
        d, err := parseLine(scanner.Text(), limit)
		if err != nil {
			log.Fatal(err)
		}
		rep := d.Report()
		fmt.Print(rep)
		_, err = out.WriteString(rep)
		if err != nil {
			log.Fatal(err)
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
