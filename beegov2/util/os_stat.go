package util

import (
	"fmt"
	"log"

	"github.com/mackerelio/go-osstat/cpu"
)

type CpuCheck struct {
}

func (c *CpuCheck) Check() error {
	before, err := cpu.Get()
	if err != nil {
		return fmt.Errorf("get cpu stat failed:%w", err)
	}
	after, err := cpu.Get()
	if err != nil {
		return fmt.Errorf("get cpu stat failed:%w", err)
	}
	total := float64(after.Total - before.Total)
	log.Println("cpu user:", float64(after.User-before.User)/total*100, "%")

	return nil
}
