package main

import (
  "bufio"
  "log"
  "os"

  "github.com/algoGuy/EasyMIDI/smf"
  "github.com/algoGuy/EasyMIDI/smfio"
)

func main() {

  // Create division
  division, err := smf.NewDivision(960, smf.NOSMTPE)
  checkErr(err)

  // Create new midi struct    创建midi文件
  midi, err := smf.NewSMF(smf.Format0, *division)
  checkErr(err)

  // Create track struct     创建track对象.// 运算符优先级.后缀	()[]->.++ --	左到右  //一元	+ -!~++ --(type)*&sizeof	右到左   所以下面的运算是先计算.然后再取地址.   记忆:从使用频率上理解,  .属性这些 .方法这些是所有语言最最常用的都没有之一,所以他们优先级最高,对于编程来说最合乎情理.
  track := &smf.Track{} //表示构造一个空的结构体Track

  // Add track to new midi struct
  err = midi.AddTrack(track)
  checkErr(err)

  // Create some midi and meta events
  midiEventOne, err := smf.NewMIDIEvent(0, smf.NoteOnStatus, 0x00, 0x30, 0x50)
  checkErr(err)
  midiEventTwo, err := smf.NewMIDIEvent(10000, smf.NoteOnStatus, 0x00, 0x30, 0x00)
  checkErr(err)
  metaEventOne, err := smf.NewMetaEvent(0, smf.MetaEndOfTrack, []byte{})
  checkErr(err)

  // Add created events to track
  err = track.AddEvent(midiEventOne)
  checkErr(err)
  err = track.AddEvent(midiEventTwo)
  checkErr(err)
  err = track.AddEvent(metaEventOne)
  checkErr(err)

  // Save to new midi source file
  outputMidi, err := os.Create("outputMidi.mid")
  checkErr(err)
  defer outputMidi.Close()

  // Create buffering stream
  writer := bufio.NewWriter(outputMidi)
  smfio.Write(writer, midi)
  writer.Flush()
}

func checkErr(err error) {
  if err != nil {
    log.Fatalln(err)
  }
}
