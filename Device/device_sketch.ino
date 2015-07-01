/******************************************************************************
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
******************************************************************************/

int redLED = 13;
int greenLED = 9;
int blueLED = 10;

int xPin = 0;
int yPin = 1;
int zPin = 2;

int panicButton = 1;

struct s53Payload {
  byte xValue;
  byte yValue;
  byte zValue;
};

void setup()
{
  //Wait a bit of time to prevent reflashing issues
  delay(5000); 
  
  // initialize serial communication at 9600 bits per second:
  Serial.begin(9600);
  
  //Set out I/O params
  pinMode(redLED,OUTPUT);
  pinMode(greenLED,OUTPUT);
  pinMode(blueLED,OUTPUT);
  pinMode(panicButton,INPUT);
  
  //Say hi!
  Serial.println("+----------------------------------------------------------+");
  Serial.println("| s53 - Reasonable Doubt Engine against RIPA s.49 notices  |");
  Serial.println("|                                                          |");
  Serial.println("| https://BrassHornCommunications.uk                       |");
  Serial.println("+----------------------------------------------------------+");
}

// the loop routine runs over and over again forever:
void loop() 
{
  int panicButtonValue = digitalRead(panicButton);
  int xValue = analogRead(xPin);
  int yValue = analogRead(yPin);
  int zValue = analogRead(zPin);
  
  if(panicButtonValue == 1)
  {
    setLED(2);
  }
  else
  {
    setLED(0);
  }
  
  //Print our status
  Serial.print("s53");
  Serial.print("\t");
  Serial.print(millis());
  Serial.print("\t");
  
  //print our safety values
  Serial.print(xValue);
  Serial.print("\t");
  Serial.print(yValue);
  Serial.print("\t");
  Serial.print(zValue);
  Serial.print("\t");
  Serial.println(panicButtonValue);
  
  
  delay(250);        // delay in between reads for stability
}

/**
0 = All good (blue)
1 = Armed (green)
2 = Triggered (red)
**/
void setLED(int state)
{
  switch (state)
  {
    case 0:
      digitalWrite(redLED,LOW);
      digitalWrite(greenLED,LOW);
      digitalWrite(blueLED,HIGH);
      break;
    case 1:
      digitalWrite(redLED,LOW);
      digitalWrite(greenLED,HIGH);
      digitalWrite(blueLED,LOW);
      break;
    case 2:
      digitalWrite(redLED,HIGH);
      digitalWrite(greenLED,LOW);
      digitalWrite(blueLED,LOW);
      break;
  }
}
