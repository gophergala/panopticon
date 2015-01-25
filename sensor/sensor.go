package main

// vim:sw=4:ts=4

import (
	"time"

	"github.com/gophergala/panopticon/entry"
)

var prevMousePos Point
var lastMouseMovement = time.Now()

func MakeEntry() (*entry.Entry, error) {
	mousePos, err := GetCursorPos()
	if err != nil {
		return nil, err
	}
	kbdLastActive, err := GetLastInputInfo()
	if err != nil {
		return nil, err
	}
	var mouseIdleTime time.Duration
	if prevMousePos == *mousePos {
		mouseIdleTime = time.Now().Sub(lastMouseMovement) / time.Millisecond
	} else {
		lastMouseMovement = time.Now()
		prevMousePos = *mousePos
	}
	kbdIdleTime := time.Duration(int64(GetTickCount()-kbdLastActive) * int64(time.Millisecond))
	idle := mouseIdleTime
	if kbdIdleTime < mouseIdleTime {
		idle = kbdIdleTime
	}
	e := entry.Entry{
		Time:    time.Now(),
		WasIdle: false,
		Idle:    time.Duration(idle * time.Millisecond),
		Title:   WindowTitle()}
	return &e, nil
}
