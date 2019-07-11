package greasego


//
// Copyright (c) 2019 ARM Limited and affiliates.
//
// SPDX-License-Identifier: MIT
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

const GREASE_LIB_OK int = 0
const GREASE_LIB_NOT_FOUND int = 0x01E00000

// This interface is for providing a special callback to called when the
// greaseLib starts
type GreaseLibStartCB func()


// GreaseError is used for error reporting from greaseLib, and is analagous
// to the same structure in the C library. An errornum of 0 means 'no error'
type GreaseError struct {
	Str   string
	Errno int
}


//This generic interfaces represents places in the
// greaseLib where the C GreaseLibCallback(err,void *) is used, but no data is ever
// passed back with the void pointer
type GreaseLibCallbackNoData interface {
}

// Callback used for a callback which
type GreaseLibAddTargetCB func(err *GreaseError, optsId int, targId uint32)

// The GreaseLib structure represents a single instantiation of the
// library. For now, only one instantiation is supported.
type GreaseLib struct {
	_greaseLibStartCB GreaseLibStartCB
}

type GreaseLibCallbackEvent struct {
	data interface{}
	err  *GreaseError
}

// GetGreaseLibVersion returns the version of the underlying greaseLib.so
// shared library in use
func GetGreaseLibVersion() (ret string) {
	return "No greaselib"
}

// The library currently only supports a single instance
// This variable tracks the singleton
var greaseInstance *GreaseLib = nil

func getGreaseLib() *GreaseLib {
}

type GreaseLibTargetFileOpts struct {
	//	uint32 _enabledFlags
	//	uint32_t _enabledFlags;
	Mode            uint32 // permissions for file (recommend default)
	Flags           uint32 // file flags (recommend default)
	Max_files       uint32 // max # of files to maintain (rotation)
	Max_file_size   uint32 // max size for any one file
	Max_total_size  uint64 // max total size to maintain in rotation
	Rotate_on_start bool
}

// analgous to greaseLib GreaseLibTargetOpts
type GreaseLibTargetOpts struct {
	Delim          *string
	TTY            *string
	File           *string
	OptsId         int                      // filled in automatically
	TargetCB       GreaseLibTargetCB        // used if this is target is a callback
	FileOpts       *GreaseLibTargetFileOpts // nil if not used
	Format_pre     *string
	Format_time    *string
	Format_level   *string
	Format_tag     *string
	Format_origin  *string
	Format_post    *string
	Format_pre_msg *string
	Name           *string // not used by greaseLib - but used for a string reference name
	// for the target ID
	flags    uint32
	NumBanks uint32
}

const GREASE_JSON_ESCAPE_STRINGS uint32 = C.GREASE_JSON_ESCAPE_STRINGS

func TargetOptsSetFlags(opts *GreaseLibTargetOpts, flag uint32) {
}

var nextOptsId uint32 = 0

//var mutexAddTargetMap = make(map[uint32]

type GreaseIdMap map[string]uint32

var DefaultLevelMap GreaseIdMap
var DefaultTagMap GreaseIdMap

// [string]:[target ID]
var TargetMap GreaseIdMap

// This interface is used for a 'callback target' in go. The
// greaseLibCallback(err,data) will be called with a 'data' string representing
// all log data since the last callback
type GreaseLibTargetCB func(err *GreaseError, data *TargetCallbackData)

// used temporarily when waiting for the callback when we call AddTarget
type addTargetCallbackData struct {
	addTargetCB GreaseLibAddTargetCB
	targetCB    GreaseLibTargetCB
}

var addTargetCallbackMap map[int]*addTargetCallbackData
var addTargetCallbackMapMutex *sync.Mutex

type TargetCallbackData struct {
	buf    *byte
	targId uint32
}

// after calling this function, the data, and any slice
// form GetBufferAsSlice() is invalid
func RetireCallbackData(data *TargetCallbackData) {
}

func (data *TargetCallbackData) GetBufferAsSlice() []byte {
}

// used to map callback targets (targets which have or are a callback)
var targetCallbackMap map[uint32]GreaseLibTargetCB
var targetCallbackMapMutex *sync.Mutex

type GreaseLevel uint32

const GREASE_ALL_LEVELS GreaseLevel = 0xFFFFFFFF //C.GREASE_ALL_LEVELS

func init() {
}


//export do_startGreaseLib_cb
func do_startGreaseLib_cb() {
}

// Start the library. The the cb.GreaseLib_start_cb() callback will be called
// upon start.
func StartGreaseLib(cb GreaseLibStartCB) {
}

// looks for a field in a struct with 'tag' and returns that
// field's reflect.Type
func findTypeByTag(tag string, in interface{}) reflect.Type {
	return nil
}

func SetSelfOriginLabel(label string) {
}

func AddLevelLabel(val uint32, label string) {
}

func AddTagLabel(val uint32, label string) {
}

func AddOriginLabel(val uint32, label string) {
}

// Assigns values to a struct based on StructTags of `greaseAssign` and `greaseType`
// Not that with string, this always assumes the structure it will fill will have a *string, not a string
func AssignFromStruct(opts interface{}, obj interface{}) { //, typ reflect.Type) {
}

func convertOptsToCGreaseLib(opts *GreaseLibTargetOpts) {
}

//export do_addTargetCB
func do_addTargetCB(err *C.GreaseLibError, info *C.GreaseLibStartedTargetInfo) {
}

//export do_modifyDefaultTargetCB
func do_modifyDefaultTargetCB(err *C.GreaseLibError, info *C.GreaseLibStartedTargetInfo) {
}

func NewGreaseLibTargetOpts() *GreaseLibTargetOpts {
}

//export do_commonTargetCB
func do_commonTargetCB(err *C.GreaseLibError, d *C.GreaseLibBuf, targetId C.uint32_t) {
}

func AddTarget(opts *GreaseLibTargetOpts, cb GreaseLibAddTargetCB) {
}

func ModifyDefaultTarget(opts *GreaseLibTargetOpts) int {
}

const GREASE_LIB_SET_FILEOPTS_MODE uint32 = 0x10000000
const GREASE_LIB_SET_FILEOPTS_FLAGS uint32 = 0x20000000
const GREASE_LIB_SET_FILEOPTS_MAXFILES uint32 = 0x40000000
const GREASE_LIB_SET_FILEOPTS_MAXFILESIZE uint32 = 0x80000000
const GREASE_LIB_SET_FILEOPTS_MAXTOTALSIZE uint32 = 0x01000000
const GREASE_LIB_SET_FILEOPTS_ROTATEONSTART uint32 = 0x02000000 // set if you want files to rotate on start
const GREASE_LIB_SET_FILEOPTS_ROTATE uint32 = 0x04000000        // set if you want files to rotate, if not set all other rotate options are skipped

func SetFileOpts(opts GreaseLibTargetOpts, flag uint32, val uint32) {
}

func SetupStandardLevels() int {
}
func SetupStandardTags() int {
}

type GreaseLibFilter struct {
}

const GREASE_LIB_SET_FILTER_ORIGIN uint32 = 0x1
const GREASE_LIB_SET_FILTER_TAG uint32 = 0x2
const GREASE_LIB_SET_FILTER_TARGET uint32 = 0x4
const GREASE_LIB_SET_FILTER_MASK uint32 = 0x8

func NewGreaseLibFilter() *GreaseLibFilter {
}

func SetFilterValue(filter *GreaseLibFilter, flag uint32, val uint32) {
}

func convertFilterToCGreaseLib(opts *GreaseLibFilter) {
}

func AddFilter(opts *GreaseLibFilter) int {
}
func DisableFilter(opts *GreaseLibFilter) int {
}
func EnableFilter(opts *GreaseLibFilter) int {
}

const GREASE_LIB_SINK_UNIXDGRAM uint32 = 0x1
const GREASE_LIB_SINK_PIPE uint32 = 0x2
const GREASE_LIB_SINK_SYSLOGDGRAM uint32 = 0x3
const GREASE_LIB_SINK_KLOG uint32 = 0x4
const GREASE_LIB_SINK_KLOG2 uint32 = 0x5

type GreaseLibSink struct {
	_binding C.GreaseLibSink
	id       uint32
}

func NewGreaseLibSink(sinkType uint32, path *string) *GreaseLibSink {
}

func AddSink(sink *GreaseLibSink) int {
}

type GreaseLibProcessClosedRedirectCallback func(err *GreaseError, stream_type int, pid int)

var closedRedirectCB GreaseLibProcessClosedRedirectCallback

func AssignChildClosedFDCallback(cb GreaseLibProcessClosedRedirectCallback) {
}

//export do_childClosedFDCallback
func do_childClosedFDCallback(err *C.GreaseLibError, stream_type C.int, fd C.int) {
}

func AddFDForStdout(fd int, originId uint32) {
}

func AddFDForStderr(fd int, originId uint32) {
}

func RemoveFDForStderr(fd int) {
}
func RemoveFDForStdout(fd int) {
}

func SetInternalTagName(name string) {

}

func SetInternalLogOrigin(originid uint32, name string) {
	//	internal_origin = C.uint32_t(originid)

}

func LogError(a ...interface{}) {
}

func LogErrorf(format string, a ...interface{}) {
}

func LogWarning(a ...interface{}) {
}

func LogWarningf(format string, a ...interface{}) {
}

func LogInfo(a ...interface{}) {
}

func LogInfof(format string, a ...interface{}) {
}

func LogDebug(a ...interface{}) {
}

func LogDebugf(format string, a ...interface{}) {
}

func LogSuccess(a ...interface{}) {
}

func LogSuccessf(format string, a ...interface{}) {
}

func LogError_noOrigin(a ...interface{}) {
}

func LogErrorf_noOrigin(format string, a ...interface{}) {
}

func LogWarning_noOrigin(a ...interface{}) {
}

func LogWarningf_noOrigin(format string, a ...interface{}) {
}

func LogInfo_noOrigin(a ...interface{}) {
}

func LogInfof_noOrigin(format string, a ...interface{}) {
}

func LogDebug_noOrigin(a ...interface{}) {
}

func LogDebugf_noOrigin(format string, a ...interface{}) {
}

func LogSuccess_noOrigin(a ...interface{}) {
}

func LogSuccessf_noOrigin(format string, a ...interface{}) {
}

func GetUnusedTagId() (goret uint32) {
}

func GetUnusedOriginId() (goret uint32) {
}
