demo-emulator
=======
<b>Copyright &copy; 2016 by Ignacio Sanchez</b>

----------
[![Build Status](https://travis-ci.org/drhelius/demo-emulator.svg?branch=master)](https://travis-ci.org/drhelius/demo-emulator)

This is a Nintendo Game Boy emulator written in Go for demonstration purposes only.

Follow me on Twitter for updates: http://twitter.com/drhelius

----------
Requirements
------------

You need a Go working environment. In addition you'll need the following dependecies:

### Windows

- GCC 64 bit installed: http://tdm-gcc.tdragon.net/download

### Linux

- Ubuntu: <code>sudo apt-get install build-essential libgl1-mesa-dev xorg-dev</code>.

### Mac OS X

- You need Xcode or Command Line Tools for Xcode (<code>xcode-select --install</code>) for required headers and libraries.

Building
--------
```
go get github.com/drhelius/demo-emulator
```

Running
-------
```
$GOPATH/bin/demo-emulator -rom your_rom.gb
```

Controls
--------
```
START = Enter
SELECT = Space
A = S
B = A
Pad = Cursors
```

License
-------

<i>This program is free software: you can redistribute it and/or modify</i>
<i>it under the terms of the GNU General Public License as published by</i>
<i>the Free Software Foundation, either version 3 of the License, or</i>
<i>any later version.</i>

<i>This program is distributed in the hope that it will be useful,</i>
<i>but WITHOUT ANY WARRANTY; without even the implied warranty of</i>
<i>MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the</i>
<i>GNU General Public License for more details.</i>

<i>You should have received a copy of the GNU General Public License</i>
<i>along with this program.  If not, see http://www.gnu.org/licenses/</i>
