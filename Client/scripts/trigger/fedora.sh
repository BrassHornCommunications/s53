#!/bin/bash

# Removing all keys is dangerous but hey
cryptsetup luksErase /dev/sda
cryptsetup luksErase /dev/mapper/fedora-root

# Do an rm -rf so that the SSD (assuming you have one) will also start doing
# it's reclaimation and data destruction routines
rm -rf /

# Now lets shutdown so the RAM degradation starts
#shutdown -h now s53 has detected tampering - all LUKS headers have been destroyed and the file system erased

#Or just send a wall for all to see
/usr/bin/wall s53 has detected tampering - all LUKS headers have been destroyed and the file system erased
