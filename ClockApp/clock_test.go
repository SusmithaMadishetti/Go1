package main

import "testing"

var resultMap map[int]string

func Test_TickTockBong_Case1(t *testing.T) {
	resultMap = TickTockBong()
	if resultMap[60] != "tock" { //60 sec
		t.Error("Clock is not working as expected for every 60 seconds..!")
	} else {
		t.Log("Clock is working as expected for every 60 seconds..!")
	}
}

func Test_TickTockBong_Case2(t *testing.T) {
	if resultMap[3600] != "bong" { //3600 sec-->60*60
		t.Error("Clock is not working as expected for every 1 hour..!")
	} else {
		t.Log("Clock is working as expected for every 1 hour..!")
	}
}

func Test_TickTockBong_Case3(t *testing.T) {
	if resultMap[10800] != "bong" { //10800 sec -->3*60*60
		t.Error("Clock is not working as expected  3 hour..!")
	} else {
		t.Log("Clock is working as expected for  3 hour..!")
	}
}

func Test_TickTockBong_Case4(t *testing.T) {
	resultMap = TickTockBong()
	if resultMap[10801] != "" { //10801 sec -->3*60*60 +1 (after 3 hour sec)
		t.Error("Clock is not stopped as expected after 3 hour..!")
	} else {
		t.Log("Clock is stopped as expected after 3 hour..!")
	}
}
