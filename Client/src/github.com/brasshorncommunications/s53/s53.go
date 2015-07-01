/*
Copyright (c) 2015, Brass Horn Communications
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package main

import (
        "log"
	"strings"
	"strconv"
        "github.com/tarm/serial"
	"code.google.com/p/gcfg"
	"os/exec"
)

type Config struct {
	Daemon struct {
		Debug bool
	}
	Serial struct {
		Port string
		Baud int
	}
	Deadhand struct {
		Command string
		Email string
		Misses int
		Triggerdeviation int
	}
}

var (
	//How many bytes to bite off the serial comms
	byteCount int = 256;

	//Our temp buffer from the serial port
        buf []byte;

	//Contains the (eventually) full status string from the arduino
        statusLine string = "";

	//For detecting and managing communications from the arduino
        //var carriageReturnChar byte = '\r';
        newLineChar byte = '\n';
        nullChar byte = '\x00';

	//Allows breaking out of the inner for loop
	foundFullLine bool = false;

	//The config options
	cfg Config;

	//Trigger threshold - high on purpose
	triggerThreshold int = 5000;

	//Slices for holding 2mins of past values for reference
	pastX []int;
	pastY []int;
	pastZ []int;
	
)

//func evals53Meta(meta string, pastX *[]int, pastY *[]int, pastZ *[]int) {
func evals53Meta(meta string) {

        var metaElements []string = strings.Split(meta, "\t");

        if(len(metaElements) == 6) {
		if(cfg.Daemon.Debug) {
                	log.Printf("Uptime: %v",metaElements[1]);
                	log.Printf("X 	 : %v",metaElements[2]);
                	log.Printf("Y 	 : %v",metaElements[3]);
                	log.Printf("Z 	 : %v",metaElements[4]);
			log.Printf("Panic: %v\n\n",metaElements[5]);
		}
        } else {
		if(cfg.Daemon.Debug) {
			log.Printf("A full s53 status string wasn't returned, skipping");
		}
		return;
	}

	var (
		deadHandTrigger bool	= false;
		readXVal       	int64 	= 0;
		readYVal       	int64 	= 0;
		readZVal       	int64 	= 0;
		panicVal   	int64 	= 0;
	)

	//Capture our values
	readXVal,_ = strconv.ParseInt(metaElements[2],10,0);
        readYVal,_ = strconv.ParseInt(metaElements[3],10,0);
        readZVal,_ = strconv.ParseInt(metaElements[4],10,0);
        panicVal,_ = strconv.ParseInt(metaElements[5],10,0);

	//Append our values	
	log.Printf("Adding %v to pastX(%v), %v to pastY(%v), %v to pastZ(%v)",readXVal,len(pastX),readYVal,len(pastY),readZVal,len(pastZ));

	pastX = append(pastX,int(readXVal));
	pastY = append(pastY,int(readYVal));
	pastZ = append(pastZ,int(readZVal));

	if(len(pastX) < 16) {
		if(cfg.Daemon.Debug) {
			log.Printf("pastX only %v large, skipping",len(pastX));
		}
		return
	}

	//Zero out all the x,y,z vals so we can use them as an average
	xVal := 0;
	yVal := 0;
	zVal := 0;
	minX := 65536;
	maxX := 0;
	minY := 65536;
	maxY := 0;
	minZ := 65536;
	maxZ := 0;

	for _,x := range pastX[len(pastX)-16:len(pastX)] {
		xVal += x;
		if(x < minX) {
			minX = x;
		}

		if(x > maxX) {
			maxX = x;
		}
	}

	xVal = (xVal / 16);

	for _,y := range pastY[len(pastY)-16:len(pastY)] {
		
                yVal += y;

		if(y < minY) {
                        minY = y;
                }

                if(y > maxY) {
                        maxY = y;
                }
        }

	yVal = (yVal / 16);

	for _,z := range pastZ[len(pastZ)-16:len(pastZ)] {
                zVal += z;

		if(z < minZ) {
                        minZ = z;
                }

                if(z > maxZ) {
                        maxZ = z;
                }
        }

	zVal = (zVal / 16);



	if((maxX - minX) > triggerThreshold) {
		deadHandTrigger = true;
		log.Printf("X: %v - %v = %v",maxX,minX,(maxX-minX));
	} else if((maxY - minY) > triggerThreshold) {
                deadHandTrigger = true;
		log.Printf("Y: %v - %v = %v",maxY,minY,(maxY-minY));
        } else if((maxZ - minZ) > triggerThreshold) {
                deadHandTrigger = true;
		log.Printf("X: %v - %v = %v",maxZ,minZ,(maxZ-minZ));
        }

	if(int(panicVal) != 0) {
                deadHandTrigger = true;
		log.Printf("Panic trigger active");
        }


	if(deadHandTrigger == true) {
		println("s53 Deadhand threshold reached, triggering.....");
		
		//cmd := exec.Command(cfg.Deadhand.Command);
		cmd := exec.Command("ls -l /tmp");
		stdout, err := cmd.Output();

    		if err != nil {
        		println(err.Error())
        		return
    		}

		if(cfg.Daemon.Debug) {
                	log.Printf("Exec returned: %s",stdout)
             	}

	}
}



func main() {

	buf := make([]byte, byteCount);

	err := gcfg.ReadFileInto(&cfg, "/etc/s53.conf")
	
	if err != nil {
		log.Printf("Could not load /etc/s53.conf, error: %s\n", err)
		return
	}

	//Record the threshold for later use
	triggerThreshold = cfg.Deadhand.Triggerdeviation;

	//Set up our averaging slices
	pastX = make([]int, 0);
	pastY = make([]int, 0);
	pastZ = make([]int, 0);


	log.Printf("Opening serial port %s",cfg.Serial.Port);	

	c := &serial.Config{Name: cfg.Serial.Port, Baud: cfg.Serial.Baud}
        s, err := serial.OpenPort(c)
        
	if err != nil {
                log.Fatal(err)
        }

	log.Printf("\n\n\ns53 - Because encryption isn't a crime\n\n\n");

	//Infinite loop of doom
	for {
		foundFullLine = false;

		// Within this loop we look to match the tab delimited chars sent by the
		// ardunio and once we find a new-line char we pass the values on
		for {
			n, err := s.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			
			if(cfg.Daemon.Debug) {
				log.Printf("Received: %s",buf[:n])
			}
		
			for i, s := range buf {
				if( s == newLineChar ) {
					if(cfg.Daemon.Debug) {
						log.Printf("%v",statusLine);
					}
					foundFullLine = true;
					break;
				} else {
					if( s != nullChar ) {
						statusLine += string(s);	
						//log.Printf("String now: %v",statusLine);
					} else {
						//log.Printf("Found null char");
					}
					/*length := len(statusLine);
					log.Printf("String length: %d",length);*/
				}
				if(cfg.Daemon.Debug) {
					log.Printf("\n-----------\n Char (%v): %q (%v)", i, s, s );
				}
			}


			//Reset the read buffer
                        for i := 0; i < byteCount; i++ {
                                buf[i] = nullChar;
                        }

			//If we've found a full status line lets quit and evaluate it
			if(foundFullLine) {
				break;
			}
		}

		//log.Printf("%v",statusLine);
		
		evals53Meta(statusLine);

		statusLine = "";
		
	}
}
