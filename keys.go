// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package evdev

import "unsafe"

// KeymapEntry.Flags values.
// They specify how the kernel should handle a keymap request.
const (
	// Kernel should perform lookup in keymap by @index instead of @scancode
	InputKeymapByIndex = 1 << iota
)

// KeymapEntry is used to retrieve and modify keymap data. Users have
// option of performing lookup either by @scancode itself or by @index
// in a keymap entry. Device.KeyMap() will also return scancode or index
// (depending on which element was used to perform lookup).
type KeymapEntry struct {
	Flags    uint8     // They specify how the kernel should handle a keymap request.
	Len      uint8     // Length of the scancode that resides in Scancode buffer.
	Index    uint16    // Index in the keymap, may be used instead of scancode
	Keycode  uint32    // Key code assigned to this scancode
	Scancode [32]uint8 // Scancode represented in machine-endian form.
}

// Keystate returns the current,global key- and button- states.
//
// This is only applicable to devices with EvKey event support.
func (d *Device) KeyState() Bitset {
	bs := NewBitset(KeyMax)
	buf := bs.Bytes()
	ioctl(d.fd.Fd(), _EVIOCGKEY(len(buf)), unsafe.Pointer(&buf[0]))
	return bs
}

// KeyMap fills the key mapping for the given key.
// E.g.: Pressing M, will input N into the input system.
// This allows us to rewire physical keys.
//
// Refer to `Device.SetKeyMap()` for information on what
// this means.
//
// Be aware that the KeyMap functions may not work on every keyboard.
// This is only applicable to devices with EvKey event support.
func (d *Device) KeyMap(keycode int) KeymapEntry {
	var entry KeymapEntry
	entry.Keycode = uint32(keycode)
	ioctl(d.fd.Fd(), _EVIOCGKEYCODE, unsafe.Pointer(&entry))
	return entry
}

// SetKeyMap sets the given key to the specified mapping.
// E.g.: Pressing M, will input N into the input system.
// This allows us to rewire physical keys.
//
// Some input drivers support variable mappings between the keys
// held down (which are interpreted by the keyboard scan and reported
// as scancodes) and the events sent to the input layer.
//
// You can change which key is associated with each scancode
// using this call. The value of the scancode is the first element
// in the integer array (list[n][0]), and the resulting input
// event key number (keycode) is the second element in the array.
// (list[n][1]).
//
// Be aware that the KeyMap functions may not work on every keyboard.
// This is only applicable to devices with EvKey event support.
func (d *Device) SetKeyMap(entry KeymapEntry) bool {
	return ioctl(d.fd.Fd(), _EVIOCSKEYCODE, unsafe.Pointer(&entry)) == nil
}

/*
	Keys and buttons

Events take the form Key<name> or Btn<name>. For example, KeyA is used
to represent the 'A' key on a keyboard. When a key is depressed, an event with
the key's code is emitted with value 1. When the key is released, an event is
emitted with value 0. Some hardware send events when a key is repeated. These
events have a value of 2. In general, Key<name> is used for keyboard keys, and
Btn<name> is used for other types of momentary switch events.

A few codes have special meanings:

  - BtnTool<name>:

  - These codes are used in conjunction with input trackpads, tablets, and
    touchscreens. These devices may be used with fingers, pens, or other tools.
    When an event occurs and a tool is used, the corresponding BtnTool<name>
    code should be set to a value of 1. When the tool is no longer interacting
    with the input device, the BtnTool<name> code should be reset to 0. All
    trackpads, tablets, and touchscreens should use at least one BtnTool<name>
    code when events are generated.

  - BtnTouch:
    BtnTouch is used for touch contact. While an input tool is determined to be
    within meaningful physical contact, the value of this property must be set
    to 1. Meaningful physical contact may mean any contact, or it may mean
    contact conditioned by an implementation defined property. For example, a
    touchpad may set the value to 1 only when the touch pressure rises above a
    certain value. BtnTouch may be combined with BtnTool<name> codes. For
    example, a pen tablet may set BtnToolPen to 1 and BtnTouch to 0 while the
    pen is hovering over but not touching the tablet surface.

  - BtnToolFinger, BtnToolDoubleTap, BtnToolTrippleTap, BtnToolQuadTap:

  - These codes denote one, two, three, and four finger interaction on a
    trackpad or touchscreen. For example, if the user uses two fingers and moves
    them on the touchpad in an effort to scroll content on screen,
    BtnToolDoubleTap should be set to value 1 for the duration of the motion.
    Note that all BtnTool<name> codes and the BtnTouch code are orthogonal in
    purpose. A trackpad event generated by finger touches should generate events
    for one code from each group. At most only one of these BtnTool<name>
    codes should have a value of 1 during any synchronization frame.
*/
const (
	KeyReserved         = 0
	KeyEscape           = 1
	Key1                = 2
	Key2                = 3
	Key3                = 4
	Key4                = 5
	Key5                = 6
	Key6                = 7
	Key7                = 8
	Key8                = 9
	Key9                = 10
	Key0                = 11
	KeyMinus            = 12
	KeyEqual            = 13
	KeyBackSpace        = 14
	KeyTab              = 15
	KeyQ                = 16
	KeyW                = 17
	KeyE                = 18
	KeyR                = 19
	KeyT                = 20
	KeyY                = 21
	KeyU                = 22
	KeyI                = 23
	KeyO                = 24
	KeyP                = 25
	KeyLeftBrace        = 26
	KeyRightBrace       = 27
	KeyEnter            = 28
	KeyLeftCtrl         = 29
	KeyA                = 30
	KeyS                = 31
	KeyD                = 32
	KeyF                = 33
	KeyG                = 34
	KeyH                = 35
	KeyJ                = 36
	KeyK                = 37
	KeyL                = 38
	KeySemiColon        = 39
	KeyApostrophe       = 40
	KeyGrave            = 41
	KeyLeftShift        = 42
	KeyBackSlash        = 43
	KeyZ                = 44
	KeyX                = 45
	KeyC                = 46
	KeyV                = 47
	KeyB                = 48
	KeyN                = 49
	KeyM                = 50
	KeyComma            = 51
	KeyDot              = 52
	KeySlash            = 53
	KeyRightShift       = 54
	KeyKPAsterisk       = 55
	KeyLeftAlt          = 56
	KeySpace            = 57
	KeyCapsLock         = 58
	KeyF1               = 59
	KeyF2               = 60
	KeyF3               = 61
	KeyF4               = 62
	KeyF5               = 63
	KeyF6               = 64
	KeyF7               = 65
	KeyF8               = 66
	KeyF9               = 67
	KeyF10              = 68
	KeyNumLock          = 69
	KeyScrollLock       = 70
	KeyKP7              = 71
	KeyKP8              = 72
	KeyKP9              = 73
	KeyKPMinus          = 74
	KeyKP4              = 75
	KeyKP5              = 76
	KeyKP6              = 77
	KeyKPPlus           = 78
	KeyKP1              = 79
	KeyKP2              = 80
	KeyKP3              = 81
	KeyKP0              = 82
	KeyKPDot            = 83
	KeyZenkakuhankaku   = 85
	Key102ND            = 86
	KeyF11              = 87
	KeyF12              = 88
	KeyRO               = 89
	KeyKatakana         = 90
	KeyHiragana         = 91
	KeyHenkan           = 92
	KeyKatakanaHiragana = 93
	KeyMuhenkan         = 94
	KeyKPJPComma        = 95
	KeyKPEnter          = 96
	KeyRightCtrl        = 97
	KeyKPSlash          = 98
	KeySysRQ            = 99
	KeyRightAlt         = 100
	KeyLineFeed         = 101
	KeyHome             = 102
	KeyUp               = 103
	KeyPageUp           = 104
	KeyLeft             = 105
	KeyRight            = 106
	KeyEnd              = 107
	KeyDown             = 108
	KeyPageDown         = 109
	KeyInsert           = 110
	KeyDelete           = 111
	KeyMacro            = 112
	KeyMute             = 113
	KeyVolumeDown       = 114
	KeyVolumeUp         = 115
	KeyPower            = 116 // SC System Power Down
	KeyKPEqual          = 117
	KeyKPPlusMinus      = 118
	KeyPause            = 119
	KeyScale            = 120 // AL Compiz Scale (Expose)
	KeyKPComma          = 121
	KeyHangeul          = 122
	KeyHanguel          = KeyHangeul
	KeyHanja            = 123
	KeyYen              = 124
	KeyLeftMeta         = 125
	KeyRightMeta        = 126
	KeyCompose          = 127
	KeyStop             = 128 // AC Stop
	KeyAgain            = 129
	KeyProps            = 130 // AC Properties
	KeyUndo             = 131 // AC Undo
	KeyFront            = 132
	KeyCopy             = 133 // AC Copy
	KeyOpen             = 134 // AC Open
	KeyPaste            = 135 // AC Paste
	KeyFind             = 136 // AC Search
	KeyCut              = 137 // AC Cut
	KeyHelp             = 138 // AL Integrated Help Center
	KeyMenu             = 139 // Menu (show menu)
	KeyCalc             = 140 // AL Calculator
	KeySetup            = 141
	KeySleep            = 142 // SC System Sleep
	KeyWakeup           = 143 // System Wake Up
	KeyFile             = 144 // AL Local Machine Browser
	KeySendFile         = 145
	KeyDeleteFile       = 146
	KeyXFer             = 147
	KeyProg1            = 148
	KeyProg2            = 149
	KeyWWW              = 150 // AL Internet Browser
	KeyMSDOS            = 151
	KeyCoffee           = 152 // AL Terminal Lock/Screensaver
	KeyScreenlock       = KeyCoffee
	KeyDirection        = 153
	KeyCycleWindows     = 154
	KeyMail             = 155
	KeyBookmarks        = 156 // AC Bookmarks
	KeyComputer         = 157
	KeyBack             = 158 // AC Back
	KeyForward          = 159 // AC Forward
	KeyCloseCD          = 160
	KeyEjectCD          = 161
	KeyEjectCloseCD     = 162
	KeyNextSong         = 163
	KeyPlayPause        = 164
	KeyPreviousSong     = 165
	KeyStopCD           = 166
	KeyRecord           = 167
	KeyRewind           = 168
	KeyPhone            = 169 // Media Select Telephone
	KeyISO              = 170
	KeyConfig           = 171 // AL Consumer Control Configuration
	KeyHomepage         = 172 // AC Home
	KeyRefresh          = 173 // AC Refresh
	KeyExit             = 174 // AC Exit
	KeyMove             = 175
	KeyEdit             = 176
	KeyScrollUp         = 177
	KeyScrollDown       = 178
	KeyKPLeftParen      = 179
	KeyKPRightParen     = 180
	KeyNew              = 181 // AC New
	KeyRedo             = 182 // AC Redo/Repeat
	KeyF13              = 183
	KeyF14              = 184
	KeyF15              = 185
	KeyF16              = 186
	KeyF17              = 187
	KeyF18              = 188
	KeyF19              = 189
	KeyF20              = 190
	KeyF21              = 191
	KeyF22              = 192
	KeyF23              = 193
	KeyF24              = 194
	KeyPlayCD           = 200
	KeyPauseCD          = 201
	KeyProg3            = 202
	KeyProg4            = 203
	KeyDashboard        = 204 // AL Dashboard
	KeySuspend          = 205
	KeyClose            = 206 // AC Close
	KeyPlay             = 207
	KeyFastForward      = 208
	KeyBassBoost        = 209
	KeyPrint            = 210 // AC Print
	KeyHP               = 211
	KeyCanera           = 212
	KeySound            = 213
	KeyQuestion         = 214
	KeyEmail            = 215
	KeyChat             = 216
	KeySearch           = 217
	KeyConnect          = 218
	KeyFinance          = 219 // AL Checkbook/Finance
	KeySport            = 220
	KeyShop             = 221
	KeyAltErase         = 222
	KeyCancel           = 223 // AC Cancel
	KeyBrightnessDown   = 224
	KeyBrightnessUp     = 225
	KeyMedia            = 226
	KeySwitchVideoMode  = 227 // Cycle between available video  outputs (Monitor/LCD/TV-out/etc)
	KeyKBDIllumToggle   = 228
	KeyKBDIllumDown     = 229
	KeyKBDIllumUp       = 230
	KeySend             = 231 // AC Send
	KeyReply            = 232 // AC Reply
	KeyForwardMail      = 233 // AC Forward Msg
	KeySave             = 234 // AC Save
	KeyDocuments        = 235
	KeyBattery          = 236
	KeyBluetooth        = 237
	KeyWLAN             = 238
	KeyUWB              = 239
	KeyUnknown          = 240
	KeyVideoNext        = 241 // drive next video source
	KeyVideoPrevious    = 242 // drive previous video source
	KeyBrightnessCycle  = 243 // brightness up, after max is min
	KeyBrightnessZero   = 244 // brightness off, use ambient
	KeyDisplayOff       = 245 // display device to off state
	KeyWIMax            = 246
	KeyRFKill           = 247 // Key that controls all radios
	KeyMicMute          = 248 // Mute / unmute the microphone
	KeyOk               = 0x160
	KeySelect           = 0x161
	KeyGoto             = 0x162
	KeyClear            = 0x163
	KeyPower2           = 0x164
	KeyOption           = 0x165
	KeyInfo             = 0x166 // AL OEM Features/Tips/Tutorial
	KeyTime             = 0x167
	KeyVendor           = 0x168
	KeyArchive          = 0x169
	KeyProgram          = 0x16a // Media Select Program Guide
	KeyChannel          = 0x16b
	KeyFavorites        = 0x16c
	KeyEPG              = 0x16d
	KeyPVR              = 0x16e // Media Select Home
	KeyMHP              = 0x16f
	KeyLanguage         = 0x170
	KeyTitle            = 0x171
	KeySubtitle         = 0x172
	KeyAngle            = 0x173
	KeyZoom             = 0x174
	KeyMode             = 0x175
	KeyKeyboard         = 0x176
	KeyScreen           = 0x177
	KeyPC               = 0x178 // Media Select Computer
	KeyTV               = 0x179 // Media Select TV
	KeyTV2              = 0x17a // Media Select Cable
	KeyVCR              = 0x17b // Media Select VCR
	KeyVCR2             = 0x17c // VCR Plus
	KeySAT              = 0x17d // Media Select Satellite
	KeySAT2             = 0x17e
	KeyCD               = 0x17f // Media Select CD
	KeyTape             = 0x180 // Media Select Tape
	KeyRadio            = 0x181
	KeyTuner            = 0x182 // Media Select Tuner
	KeyPlayer           = 0x183
	KeyText             = 0x184
	KeyDVD              = 0x185 // Media Select DVD
	KeyAUX              = 0x186
	KeyMP3              = 0x187
	KeyAudio            = 0x188 // AL Audio Browser
	KeyVideo            = 0x189 // AL Movie Browser
	KeyDirectory        = 0x18a
	KeyList             = 0x18b
	KeyMemo             = 0x18c // Media Select Messages
	KeyCalender         = 0x18d
	KeyRed              = 0x18e
	KeyGreen            = 0x18f
	KeyYellow           = 0x190
	KeyBlue             = 0x191
	KeyChannelUp        = 0x192 // Channel Increment
	KeyChannelDown      = 0x193 // Channel Decrement
	KeyFirst            = 0x194
	KeyLast             = 0x195 // Recall Last
	KeyAB               = 0x196
	KeyNext             = 0x197
	KeyRestart          = 0x198
	KeySlow             = 0x199
	KeyShuffle          = 0x19a
	KeyBreak            = 0x19b
	KeyPrevious         = 0x19c
	KeyDigits           = 0x19d
	KeyTeen             = 0x19e
	KeyTwen             = 0x19f
	KeyVideoPhone       = 0x1a0 // Media Select Video Phone
	KeyGames            = 0x1a1 // Media Select Games
	KeyZoomIn           = 0x1a2 // AC Zoom In
	KeyZoomOut          = 0x1a3 // AC Zoom Out
	KeyZoomReset        = 0x1a4 // AC Zoom
	KeyWordProcessor    = 0x1a5 // AL Word Processor
	KeyEditor           = 0x1a6 // AL Text Editor
	KeySpreadsheet      = 0x1a7 // AL Spreadsheet
	KeyGraphicsEditor   = 0x1a8 // AL Graphics Editor
	KeyPresentation     = 0x1a9 // AL Presentation App
	KeyDatabase         = 0x1aa // AL Database App
	KeyNews             = 0x1ab // AL Newsreader
	KeyVoiceMail        = 0x1ac // AL Voicemail
	KeyAddressBook      = 0x1ad // AL Contacts/Address Book
	KeyMessenger        = 0x1ae // AL Instant Messaging
	KeyDisplayToggle    = 0x1af // Turn display (LCD) on and off
	KeySpellCheck       = 0x1b0 // AL Spell Check
	KeyLogoff           = 0x1b1 // AL Logoff
	KeyDollar           = 0x1b2
	KeyEuro             = 0x1b3
	KeyFrameBack        = 0x1b4 // Consumer - transport controls
	KeyframeForward     = 0x1b5
	KeyContextMenu      = 0x1b6 // GenDesc - system context menu
	KeyMediaRepeat      = 0x1b7 // Consumer - transport control
	Key10ChannelsUp     = 0x1b8 // 10 channels up (10+)
	Key10ChannelsDown   = 0x1b9 // 10 channels down (10-)
	KeyImages           = 0x1ba // AL Image Browser
	KeyDelEOL           = 0x1c0
	KeyDelEOS           = 0x1c1
	KeyInsLine          = 0x1c2
	KeyDelLine          = 0x1c3
	KeyFN               = 0x1d0
	KeyFNEsc            = 0x1d1
	KeyFNF1             = 0x1d2
	KeyFNF2             = 0x1d3
	KeyFNF3             = 0x1d4
	KeyFNF4             = 0x1d5
	KeyFNF5             = 0x1d6
	KeyFNF6             = 0x1d7
	KeyFNF7             = 0x1d8
	KeyFNF8             = 0x1d9
	KeyFNF9             = 0x1da
	KeyFNF10            = 0x1db
	KeyFNF11            = 0x1dc
	KeyFNF12            = 0x1dd
	KeyFN1              = 0x1de
	KeyFN2              = 0x1df
	KeyFND              = 0x1e0
	KeyFNE              = 0x1e1
	KeyFNF              = 0x1e2
	KeyFNS              = 0x1e3
	KeyFNB              = 0x1e4
	KeyBRLDot1          = 0x1f1
	KeyBRLDot2          = 0x1f2
	KeyBRLDot3          = 0x1f3
	KeyBRLDot4          = 0x1f4
	KeyBRLDot5          = 0x1f5
	KeyBRLDot6          = 0x1f6
	KeyBRLDot7          = 0x1f7
	KeyBRLDot8          = 0x1f8
	KeyBRLDot9          = 0x1f9
	KeyBRLDot10         = 0x1fa
	KeyNumeric0         = 0x200 // used by phones, remote controls,
	KeyNumeric1         = 0x201 // and other keypads
	KeyNumeric2         = 0x202
	KeyNumeric3         = 0x203
	KeyNumeric4         = 0x204
	KeyNumeric5         = 0x205
	KeyNumeric6         = 0x206
	KeyNumeric7         = 0x207
	KeyNumeric8         = 0x208
	KeyNumeric9         = 0x209
	KeyNumericStar      = 0x20a
	KeyNumericPound     = 0x20b
	KeyCameraFocus      = 0x210
	KeyWPSButton        = 0x211 // WiFi Protected Setup key
	KeyTouchpadToggle   = 0x212 // Request switch touchpad on or off
	KeyTouchpadOn       = 0x213
	KeyTouchpadOff      = 0x214
	KeyCameraZoomIn     = 0x215
	KeyCameraZoomOut    = 0x216
	KeyCameraUp         = 0x217
	KeyCameraDown       = 0x218
	KeyCameraLeft       = 0x219
	KeyCameraRight      = 0x21a
	KeyAttendantOn      = 0x21b
	KeyAttendantOff     = 0x21c
	KeyAttendantToggle  = 0x21d // Attendant call on or off
	KeyLightsToggle     = 0x21e // Reading light on or off

	// We avoid low common keys in module aliases so they don't get huge.
	KeyMinInteresting = KeyMute
	KeyMax            = 0x2ff
	KeyCount          = KeyMax + 1
)

// Button codes for mice and other devices.
const (
	BtnMisc           = 0x100
	Btn0              = 0x100
	Btn1              = 0x101
	Btn2              = 0x102
	Btn3              = 0x103
	Btn4              = 0x104
	Btn5              = 0x105
	Btn6              = 0x106
	Btn7              = 0x107
	Btn8              = 0x108
	Btn9              = 0x109
	BtnMouse          = 0x110
	BtnLeft           = 0x110
	BtnRight          = 0x111
	BtnMiddle         = 0x112
	BtnSide           = 0x113
	BtnExtra          = 0x114
	BtnForward        = 0x115
	BtnBack           = 0x116
	BtnTask           = 0x117
	BtnJoystick       = 0x120
	BtnTrigger        = 0x120
	BtnThumb          = 0x121
	BtnThumb2         = 0x122
	BtnTop            = 0x123
	BtnTop2           = 0x124
	BtnPinkie         = 0x125
	BtnBase           = 0x126
	BtnBase2          = 0x127
	BtnBase3          = 0x128
	BtnBase4          = 0x129
	BtnBase5          = 0x12a
	BtnBase6          = 0x12b
	BtnDead           = 0x12f
	BtnGamepad        = 0x130
	BtnA              = 0x130
	BtnB              = 0x131
	BtnC              = 0x132
	BtnX              = 0x133
	BtnY              = 0x134
	BtnZ              = 0x135
	BtnTL             = 0x136
	BtnTR             = 0x137
	BtnTL2            = 0x138
	BtnTR2            = 0x139
	BtnSelect         = 0x13a
	BtnStart          = 0x13b
	BtnMode           = 0x13c
	BtnThumbL         = 0x13d
	BtnThumbR         = 0x13e
	BtnDigi           = 0x140
	BtnToolPen        = 0x140
	BtnTooLRubber     = 0x141
	BtnToolBrush      = 0x142
	BtnToolPencil     = 0x143
	BtnToolAirbrush   = 0x144
	BtnToolFinger     = 0x145
	BtnToolMouse      = 0x146
	BtnToolLens       = 0x147
	BtnToolQuintTap   = 0x148 // Five fingers on trackpad
	BtnTouch          = 0x14a
	BtnStylus         = 0x14b
	BtnStylus2        = 0x14c
	BtnToolDoubleTap  = 0x14d
	BtnToolTrippleTap = 0x14e
	BtnToolQuadTap    = 0x14f // Four fingers on trackpad
	BtnWheel          = 0x150
	BtnGearDown       = 0x150
	BtnGearUp         = 0x151
	BtnTriggerHappy   = 0x2c0
	BtnTriggerHappy1  = 0x2c0
	BtnTriggerHappy2  = 0x2c1
	BtnTriggerHappy3  = 0x2c2
	BtnTriggerHappy4  = 0x2c3
	BtnTriggerHappy5  = 0x2c4
	BtnTriggerHappy6  = 0x2c5
	BtnTriggerHappy7  = 0x2c6
	BtnTriggerHappy8  = 0x2c7
	BtnTriggerHappy9  = 0x2c8
	BtnTriggerHappy10 = 0x2c9
	BtnTriggerHappy11 = 0x2ca
	BtnTriggerHappy12 = 0x2cb
	BtnTriggerHappy13 = 0x2cc
	BtnTriggerHappy14 = 0x2cd
	BtnTriggerHappy15 = 0x2ce
	BtnTriggerHappy16 = 0x2cf
	BtnTriggerHappy17 = 0x2d0
	BtnTriggerHappy18 = 0x2d1
	BtnTriggerHappy19 = 0x2d2
	BtnTriggerHappy20 = 0x2d3
	BtnTriggerHappy21 = 0x2d4
	BtnTriggerHappy22 = 0x2d5
	BtnTriggerHappy23 = 0x2d6
	BtnTriggerHappy24 = 0x2d7
	BtnTriggerHappy25 = 0x2d8
	BtnTriggerHappy26 = 0x2d9
	BtnTriggerHappy27 = 0x2da
	BtnTriggerHappy28 = 0x2db
	BtnTriggerHappy29 = 0x2dc
	BtnTriggerHappy30 = 0x2dd
	BtnTriggerHappy31 = 0x2de
	BtnTriggerHappy32 = 0x2df
	BtnTriggerHappy33 = 0x2e0
	BtnTriggerHappy34 = 0x2e1
	BtnTriggerHappy35 = 0x2e2
	BtnTriggerHappy36 = 0x2e3
	BtnTriggerHappy37 = 0x2e4
	BtnTriggerHappy38 = 0x2e5
	BtnTriggerHappy39 = 0x2e6
	BtnTriggerHappy40 = 0x2e7
)
