package main

 import "testing"

 func TestCalculator (t *testing.T) {
	 if res , err:=Calculate ("+",1,2); res !=3 || err !=nil{
		t.Errorf("Addition test failed: %s", err.Error())
	 }

	 if res , err:=Calculate ("-",1,2); res !=-1 || err !=nil{
		t.Errorf("Addition test failed: %s", err.Error())
	 }

	 if res , err:=Calculate ("*",1,2); res !=2 || err !=nil{
		t.Errorf("Addition test failed: %s", err.Error())
	 }

	 if res , err:=Calculate ("/",1,2); res !=0 || err !=nil{
		t.Errorf("Addition test failed: %s", err.Error())
	 }
	 
 }