# s53
s53 is a software client and hardware device combination designed to provide reasonable doubt as to whether one still has decryption keys for the computer to which the device is attached.

It does this by destroying or changing the keys of all encrypted disks on a protected computer if the device is removed, a panic button is triggered or X,Y,Z movement is detected.

This is a test of the [Regulation of Investigatory Powers Act 2000 Section 53 Paragraph 3 defence](http://www.legislation.gov.uk/ukpga/2000/23/section/53#section-53-3)

## Directory Structure
### Client
GoLang source and bash scripts for running on the computer with LUKS / BSD crypto RAID protected partitions.

### Device
Ardunio sketch for use with an Freetronics LeoStick *(although any arduino will work)* and an ADXL337 chip to detect motion, receive wireless signals from panic buttons or other s53 devices *(if enabled)* and communicate with the client software *(heartbeats etc)*

### 3D
Various 3D printer files for housing elements of the s53 hardware

### Website
A copy of the s53 website - this is for mirroring the s53 website if required.

## Legal Notice:
We are not a lawyers, solictiors, or in any way well versed in the law. Relying on this software to refuse to comply with a Regulation of Investigatory Powers Act 2000 Section 49 notice may result in you being found guilty of **Failure to comply with a notice** under the Regulation of Investigatory Powers Act 2000 and sentenced to up to 5 years in prison.

## Copyright
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
