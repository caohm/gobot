package sysfs

import (
	"os"
	"syscall"
	"testing"

	"gobot.io/x/gobot/gobottest"
)

func TestPwmPin(t *testing.T) {
	fs := NewMockFilesystem([]string{
		"/sys/class/pwm/pwmchip0/export",
		"/sys/class/pwm/pwmchip0/unexport",
		"/sys/class/pwm/pwmchip0/pwm10/enable",
		"/sys/class/pwm/pwmchip0/pwm10/period",
		"/sys/class/pwm/pwmchip0/pwm10/duty_cycle",
	})

	SetFilesystem(fs)

	pin := NewPwmPin(10)
	gobottest.Assert(t, pin.pin, "10")

	err := pin.Unexport()
	gobottest.Assert(t, err, nil)
	gobottest.Assert(t, fs.Files["/sys/class/pwm/pwmchip0/unexport"].Contents, "10")

	err = pin.Export()
	gobottest.Assert(t, err, nil)
	gobottest.Assert(t, fs.Files["/sys/class/pwm/pwmchip0/export"].Contents, "10")

	gobottest.Refute(t, fs.Files["/sys/class/pwm/pwmchip0/pwm10/enable"].Contents, "1")
	err = pin.Enable("1")
	gobottest.Assert(t, err, nil)
	gobottest.Assert(t, fs.Files["/sys/class/pwm/pwmchip0/pwm10/enable"].Contents, "1")

	fs.Files["/sys/class/pwm/pwmchip0/pwm10/period"].Contents = "6"
	data, _ := pin.Period()
	gobottest.Assert(t, data, "6")
}

func TestPwmPinExportError(t *testing.T) {
	fs := NewMockFilesystem([]string{
		"/sys/class/pwm/pwmchip0/export",
		"/sys/class/pwm/pwmchip0/unexport",
		"/sys/class/pwm/pwmchip0/pwm10/enable",
		"/sys/class/pwm/pwmchip0/pwm10/period",
		"/sys/class/pwm/pwmchip0/pwm10/duty_cycle",
	})

	SetFilesystem(fs)

	pin := NewPwmPin(10)
	pin.write = func(string, []byte) (int, error) {
		return 0, &os.PathError{Err: syscall.EBUSY}
	}

	gobottest.Refute(t, pin.Export(), nil)
}

func TestPwmPinUnxportError(t *testing.T) {
	fs := NewMockFilesystem([]string{
		"/sys/class/pwm/pwmchip0/export",
		"/sys/class/pwm/pwmchip0/unexport",
		"/sys/class/pwm/pwmchip0/pwm10/enable",
		"/sys/class/pwm/pwmchip0/pwm10/period",
		"/sys/class/pwm/pwmchip0/pwm10/duty_cycle",
	})

	SetFilesystem(fs)

	pin := NewPwmPin(10)
	pin.write = func(string, []byte) (int, error) {
		return 0, &os.PathError{Err: syscall.EBUSY}
	}

	gobottest.Refute(t, pin.Unexport(), nil)
}