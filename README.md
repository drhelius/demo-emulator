demo-emulator [![Build Status](https://travis-ci.org/drhelius/demo-emulator.svg?branch=master)](https://travis-ci.org/drhelius/demo-emulator)
=======
<b>Copyright &copy; 2016 by Ignacio Sanchez</b>

Nintendo Game Boy emulator written in Go to be used in workshops about emulator programming.

Follow me on Twitter for updates: http://twitter.com/drhelius

![Screenshot](/screenshots/screenshot1.png "Screenshot") ![Screenshot](/screenshots/screenshot2.png "Screenshot")

Presentation
------------

https://speakerdeck.com/drhelius/8-bit-emulator-programming-with-go

Requirements
------------

Before you start, make sure you have Go installed and ready to build applications: https://golang.org/doc/install

Once you have a working Go environment you'll need to install the following dependecies:

#### Windows

- GCC 64 bit installed: http://tdm-gcc.tdragon.net/download

#### Linux

- Ubuntu: <code>sudo apt-get install build-essential libgl1-mesa-dev xorg-dev</code>
- Fedora: <code>sudo dnf install @development-tools libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel libXi-devel</code>

#### Mac OS X

- You need Xcode or Command Line Tools for Xcode (<code>xcode-select --install</code>) for required headers and libraries.

Building
--------
Run this command to let Go download and build the sources. You don't even need to clone this repo, Go will do it for you:

```
go get -u github.com/drhelius/demo-emulator
```

Running
-------
Once built you can find the emulator binary in <code>$GOPATH/bin</code>. Use it with the <code>-rom</code> argument in order to load a Game Boy ROM file:

```
$GOPATH/bin/demo-emulator -rom path/to/your_rom.gb
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
