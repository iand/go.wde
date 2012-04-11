/*
   Copyright 2012 John Asmuth

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package xgb

import (
	"code.google.com/p/jamslam-x-go-binding/xgb"
	"image/draw"
	"image"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"sync"
)

var xu *xgbutil.XUtil
var connLock sync.Mutex

func ensureConnection() (err error) {
	connLock.Lock()
	defer connLock.Unlock()

	if xu != nil {
		return
	}

	xu, err = xgbutil.Dial(":0.0")

	return
}

type Window struct {
	id xgb.Id
	conn *xgb.Conn
	buffer draw.Image
	width, height int
}

func NewWindow(width, height int) (w *Window, err error) {

	println(width, height)

	err = ensureConnection()
	if err != nil {
		return
	}

	w = new(Window)
	w.width, w.height = width, height
	w.buffer = image.NewNRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	w.conn = xu.Conn()
	screen := xu.Screen()

	w.id = w.conn.NewId()
	w.conn.CreateWindow(xgb.WindowClassCopyFromParent, w.id, screen.Root, 200, 200, uint16(width), uint16(height), 10, xgb.WindowClassInputOutput, screen.RootVisual, 0, []uint32{})

	return
}

func (w *Window) SetTitle(title string) {
	// cannot
	return
}

func (w *Window) SetSize(width, height int) {
	// cannot
	return
}

func (w *Window) Size() (width, height int)  {
	width, height = w.width, w.height
	return
}

func (w *Window) Show() {
	w.conn.MapWindow(w.id)

}

func (w *Window) Screen() (im draw.Image) {
	im = w.buffer
	return
}

func (w *Window) FlushImage() {
	xgraphics.PaintImg(xu, w.id, w.buffer)
}

func (w *Window) Close() (err error) {

	return
}

func (w *Window) EventChan() (events <-chan interface{}) {
	return make(chan interface{})
}