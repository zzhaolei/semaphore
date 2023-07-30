package semaphore

/*
#include <stdlib.h>
#include <fcntl.h>
*/
import "C"

import (
	"fmt"
	"syscall"
	"unsafe"
)

type Semaphore struct {
	sd uintptr
}

func New() *Semaphore {
	return &Semaphore{}
}

func (sem *Semaphore) Open(name string, mode int, value int) error {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	r1, _, err := syscall.Syscall6(
		syscall.SYS_SEM_OPEN,
		uintptr(unsafe.Pointer(cs)),
		uintptr(C.O_CREAT),
		uintptr(mode),
		uintptr(value),
		0,
		0,
	)
	if err != 0 {
		sem.Unlink(name)
		return fmt.Errorf("create semaphore failed: %s", err)
	}
	sem.sd = r1
	return nil
}

func (sem *Semaphore) TryAcquire() error {
	_, _, err := syscall.Syscall(syscall.SYS_SEM_TRYWAIT, sem.sd, 0, 0)
	if err != 0 {
		return fmt.Errorf("try acquire failed: %s", err.Error())
	}
	return nil
}

func (sem *Semaphore) Acquire() {
	_, _, _ = syscall.Syscall(syscall.SYS_SEM_WAIT, sem.sd, 0, 0)
}

func (sem *Semaphore) Release() {
	_, _, _ = syscall.Syscall(syscall.SYS_SEM_POST, sem.sd, 0, 0)
}

func (sem *Semaphore) Unlink(name string) {
	cs := C.CString(name)
	defer C.free(unsafe.Pointer(cs))

	sem.close()
	_, _, _ = syscall.Syscall(syscall.SYS_SEM_UNLINK, uintptr(unsafe.Pointer(cs)), 0, 0)
}

func (sem *Semaphore) close() {
	_, _, _ = syscall.Syscall(syscall.SYS_SEM_CLOSE, sem.sd, 0, 0)
}
